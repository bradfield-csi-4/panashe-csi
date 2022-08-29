  /*

Two different ways to loop over an array of arrays.

Spotted at:
http://stackoverflow.com/questions/9936132/why-does-the-order-of-the-loops-affect-performance-when-iterating-over-a-2d-arra

*/

void option_one() {
  int i, j;
  static int x[4000][4000];
  for (i = 0; i < 4000; i++) {
    for (j = 0; j < 4000; j++) {
      x[i][j] = i + j;
    }
  }
}

void option_two() {
  int i, j;
  static int x[4000][4000];
  for (i = 0; i < 4000; i++) {
    for (j = 0; j < 4000; j++) {
      x[j][i] = i + j;
    }
  }
}

int main() {
  option_one();
  //option_two();
  return 0;
}


/*
 *
Option one stats with cachegrind
==35997== I   refs:      144,186,702
==35997== I1  misses:          1,004
==35997== LLi misses:            995
==35997== I1  miss rate:        0.00%
==35997== LLi miss rate:        0.00%
==35997==
==35997== D   refs:       16,052,126  (39,216 rd   + 16,012,910 wr)
==35997== D1  misses:      1,002,019  ( 1,459 rd   +  1,000,560 wr)
==35997== LLd misses:      1,001,860  ( 1,316 rd   +  1,000,544 wr)
==35997== D1  miss rate:         6.2% (   3.7%     +        6.2%  )
==35997== LLd miss rate:         6.2% (   3.4%     +        6.2%  )
==35997==
==35997== LL refs:         1,003,023  ( 2,463 rd   +  1,000,560 wr)
==35997== LL misses:       1,002,855  ( 2,311 rd   +  1,000,544 wr)
==35997== LL miss rate:          0.6% (   0.0%     +        6.2%  )

Option two:
==35950== I   refs:      144,186,702
==35950== I1  misses:          1,004
==35950== LLi misses:            995
==35950== I1  miss rate:        0.00%
==35950== LLi miss rate:        0.00%
==35950==
==35950== D   refs:       16,052,126  (39,216 rd   + 16,012,910 wr)
==35950== D1  misses:     16,002,018  ( 1,459 rd   + 16,000,559 wr)
==35950== LLd misses:      1,005,859  ( 1,316 rd   +  1,004,543 wr)
==35950== D1  miss rate:        99.7% (   3.7%     +       99.9%  )
==35950== LLd miss rate:         6.3% (   3.4%     +        6.3%  )
==35950==
==35950== LL refs:        16,003,022  ( 2,463 rd   + 16,000,559 wr)
==35950== LL misses:       1,006,854  ( 2,311 rd   +  1,004,543 wr)
==35950== LL miss rate:          0.6% (   0.0%     +        6.3%  )

The big difference here is that option a way higher D1 cache miss rate - 99.7% vs 6.2%
*/
