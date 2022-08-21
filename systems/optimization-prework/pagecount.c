#include <stdint.h>
#include <stdio.h>
#include <time.h>

#define TEST_LOOPS 10000000

uint64_t pagecount(uint64_t memory_size, uint64_t p_size_pow) {
  return memory_size >> p_size_pow;
}

int main (int argc, char** argv) {
  clock_t baseline_start, baseline_end, test_start, test_end;
  uint64_t memory_size, p_size_pow;
  double clocks_elapsed, time_elapsed;
  int i, ignore = 0;

  uint64_t msizes[] = {1L << 32, 1L << 40, 1L << 52};
  uint64_t psize_pows[] = {12, 16, 32};

  baseline_start = clock();
  for (i = 0; i < TEST_LOOPS; i++) {
    memory_size = msizes[i % 3];
    p_size_pow = psize_pows[i % 3];
    ignore += 1 + memory_size +
              p_size_pow; // so that this loop isn't just optimized away
  }
  baseline_end = clock();

  test_start = clock();
  for (i = 0; i < TEST_LOOPS; i++) {
    memory_size = msizes[i % 3];
    p_size_pow = psize_pows[i % 3];
    ignore += pagecount(memory_size, p_size_pow) + memory_size + p_size_pow;
  }
  test_end = clock();

  clocks_elapsed = test_end - test_start - (baseline_end - baseline_start);
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;

  printf("%.2fs to run %d tests (%.2fns per test)\n", time_elapsed, TEST_LOOPS,
         time_elapsed * 1e9 / TEST_LOOPS);
  return ignore;
}

/*
 *
 * As a first step, read the code for this function then consider:

    Which instructions would you expect your compiler to generate for this function?
    Expected a division

    What does it in fact generate?
    There are 4 instructions total:
	movq	%rdi, %rax ; memory size
	movl	$0, %edx
	divq	%rsi ; page size
	ret

    The first two are necessary just to set up the registers for division



    If you change the optimization level, is the function substantially different?
    At -O3, I get
	movq	%rdi, %rax
	xorl	%edx, %edx
	divq	%rsi
	ret

    As discussed by oz, the xor might be slightly cheaper than moving 0, but it's not substantially different.

    Use godbolt.org to explore a few different compilers. Do any of them generate substantially different instructions?
    No difference between gcc, icx and clang here.


    By using Agner Fog’s instruction tables or reviewing CS:APP chapter 5.7, can you determine which of the generated instructions may be slow?
    From Fogg's table, for Tiger lake (11th gen intel):
    Instruction           Latency            Reciprocal throughput
    mov[r, r]             1                  1/4
    mov[r, i]             _                  1/4
    div[r64]              15                 10

   Division's a whole order of magnitude slower!

   We rewrote this to use bitshifting instead of explicit division.
	movq	%rdi, %rax
	movl	%esi, %ecx
	shrq	%cl, %rax
	ret
   
   Instruction            Latency            Reciprocal throughput
   shr[r, cl]             1                  1


  Next, let’s improve performance!

    Noting that a page size is always a power of 2, and that the size of memory will always be cleanly divisible by the page size, can you think of a performance optimization we could employ? You are welcome to change the function signature and test runner code.
    Bit shift for division

    How much of an improvement would you expect to see?
    Before, we had a total reciprocal throughput of 10.5. We now have one of 1.5. So this should be about 7 times faster now.

    Go ahead and make the improvement, and measure the speed up. Did this match your expectations?
    Post improvement: 0.8ns per test
    Pre improvement: ~1ns per test

    So not as dramatic as I though :laughing emoji:

    Consider, what is stopping the compiler from making the same optimization that you did?
    The compiler doesn't "know" that it will always have a power of 2 in the denominator, or what power of 2 it is.
 

    Using a profiling tool like perf on Linux or Instruments on macOS, see if you can measure the number of CPU cycles used for each version. Does this correspond to our measure of CPU time?

Percent│
       │
       │
       │   Disassembly of section .text:
       │
       │   0000000000401136 <pagecount>:
       │   pagecount():
       │     mov %rdi,%rax
       │     mov $0x0,%edx
 83.78 │     div %rsi
 16.22 │   ← ret

    The unoptimized version spends 83.78% of the CPU cycles were spent in div (and 31 out of a total of 114 sampled cycles).

Percent│
       │
       │
       │   Disassembly of section .text:
       │
       │   0000000000401136 <pagecount>:
       │   pagecount():
       │     mov %rdi,%rax
       │     mov %esi,%ecx
100.00 │     shr %cl,%rax
       │   ← ret

The optimized version spends ~100% of CPU time on the shift right op. However, this is only 10 samples of 66 total, ~15%

$ perf stat ./optimized-pagecount
0.00s to run 10000000 tests (0.05ns per test)

 Performance counter stats for './optimized-pagecount':

             19.85 msec task-clock:u              #    0.979 CPUs utilized          
                 0      context-switches:u        #    0.000 /sec                   
                 0      cpu-migrations:u          #    0.000 /sec                   
                55      page-faults:u             #    2.770 K/sec                  
        79,671,761      cpu_core/cycles:u/        #    4.013 G/sec


$ perf stat ./unoptimized-pagecount
0.01s to run 10000000 tests (1.40ns per test)

 Performance counter stats for './unoptimized-pagecount':

             32.15 msec task-clock:u              #    0.987 CPUs utilized          
                 0      context-switches:u        #    0.000 /sec                   
                 0      cpu-migrations:u          #    0.000 /sec                   
                55      page-faults:u             #    1.711 K/sec                  
       134,564,054      cpu_core/cycles:u/        #    4.185 G/sec

So the optimized version uses nearly half as many cpu cycles by that measure
*/
