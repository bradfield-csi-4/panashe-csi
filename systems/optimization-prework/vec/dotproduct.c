#include "vec.h"


data_t dotproduct(vec_ptr u, vec_ptr v) {
   data_t sum = 0, u_val_1, u_val_2, v_val_1, v_val_2;
   long i;

   long len = vec_length(u);
   for (i = 0; i < len-1; i+=2) { // we can assume both vectors are same length
	u_val_1 = u->data[i];
	u_val_2 = u->data[i+1];
	v_val_1 = v->data[i];
	v_val_2 = v->data[i+1];
        sum += u_val_1 * v_val_1 + u_val_2 * v_val_2;
   }   

   if (i < len) {
	   sum += u->data[i] * v->data[i];
   }
   return sum;
}
