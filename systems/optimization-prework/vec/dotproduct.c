#include "vec.h"


data_t dotproduct(vec_ptr u, vec_ptr v) {
<<<<<<< HEAD
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
=======
   data_t sum = 0, u_val, v_val;

   for (long i = 0; i < vec_length(u); i++) { // we can assume both vectors are same length
        get_vec_element(u, i, &u_val);
        get_vec_element(v, i, &v_val);
        sum += u_val * v_val;
   }   
>>>>>>> be55fec (Add systems assembly and optimization prework)
   return sum;
}
