section .text
global sum_to_n
sum_to_n:
	mov rax, 0 ; empty the return register
	mov rbx, rdi ; initialize a register to the value of the argument
.loop:
	add	rax, rbx ; add to running sum
	dec	rbx ; decrement counter
	jnp	.loop ; This relies on the fact that the parity flag is set to 0 if the decrement results in 0.
	ret
