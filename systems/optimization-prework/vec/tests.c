#include "vendor/unity.h"
<<<<<<< HEAD
#include <time.h>

#include "vec.h"

#define TEST_LOOPS 100

=======

#include "vec.h"

>>>>>>> be55fec (Add systems assembly and optimization prework)
extern data_t dotproduct(vec_ptr, vec_ptr);

void setUp(void) {
}

void tearDown(void) {
}

void test_empty(void) {
  vec_ptr u = new_vec(0);
  vec_ptr v = new_vec(0);

  TEST_ASSERT_EQUAL(0, dotproduct(u, v));

  free_vec(u);
  free_vec(v);
}

void test_basic(void) {
  vec_ptr u = new_vec(3);
  vec_ptr v = new_vec(3);

  set_vec_element(u, 0, 1);
  set_vec_element(u, 1, 2);
  set_vec_element(u, 2, 3);
  set_vec_element(v, 0, 4);
  set_vec_element(v, 1, 5);
  set_vec_element(v, 2, 6);

  TEST_ASSERT_EQUAL(32, dotproduct(u, v));

  free_vec(u);
  free_vec(v);
}

void test_longer(void) {
  long n = 1000000;
  vec_ptr u = new_vec(n);
  vec_ptr v = new_vec(n);

  for (long i = 0; i < n; i++) {
    set_vec_element(u, i, i + 1);
    set_vec_element(v, i, i + 1);
  }

  long expected = (2 * n * n * n + 3 * n * n + n) / 6;
  TEST_ASSERT_EQUAL(expected, dotproduct(u, v));

  free_vec(u);
  free_vec(v);
}

int main(void) {
    UNITY_BEGIN();

    RUN_TEST(test_empty);
    RUN_TEST(test_basic);
    RUN_TEST(test_longer);

<<<<<<< HEAD
    int ignore, i = 0;
    uint64_t memory_size, p_size_pow;
    uint64_t msizes[] = {1L << 32, 1L << 40, 1L << 52};
    uint64_t psize_pows[] = {12, 16, 32};
    clock_t baseline_start = clock();
    for (i = 0; i < TEST_LOOPS; i++) {
      memory_size = msizes[i % 3];
      p_size_pow = psize_pows[i % 3];
      ignore += 1 + memory_size +
                p_size_pow; // so that this loop isn't just optimized away
    }
    clock_t baseline_end = clock();

    clock_t start = clock();
    for (i = 0; i < TEST_LOOPS; i++) {
	    //test_empty();
	    //test_basic();
	    test_longer();
    }
    clock_t end_clock = clock();
    clock_t clocks_elapsed = end_clock - start - (baseline_end - baseline_start);
    double time_elapsed = clocks_elapsed * 1.0 / CLOCKS_PER_SEC;

    printf("%.2fs to run %d `longer` tests (%.2fns per test)\n", time_elapsed, TEST_LOOPS, time_elapsed * 1e9 / TEST_LOOPS);
    return UNITY_END();
}

/*
 * Initial data (very consistent, +- 10ms):
 * 0.88s to run 100 `longer` tests (8811300.00ns per test)
 *
 * Move length check before the loop
 * 0.83s to run 100 `longer` tests (8317900.00ns per test)
 *
 * Access internal data array directly instead of function call
 * 0.70s to run 100 `longer` tests (6953300.00ns per test)
 *
 * 2 element loop unroll
 * 0.68s to run 100 `longer` tests (6843270.00ns per test)
 */
=======
    return UNITY_END();
}
>>>>>>> be55fec (Add systems assembly and optimization prework)
