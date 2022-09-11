	.text
	.file	"mstore.c"
	.globl	multstore                       # -- Begin function multstore
	.p2align	4, 0x90
	.type	multstore,@function
multstore:                              # @multstore
	.cfi_startproc
# %bb.0:
	pushq	%rbx
	.cfi_def_cfa_offset 16
	.cfi_offset %rbx, -16
	movq	%rdx, %rbx
	callq	mult2
	movq	%rax, (%rbx)
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	multstore, .Lfunc_end0-multstore
	.cfi_endproc
                                        # -- End function
	.ident	"Apple clang version 13.1.6 (clang-1316.0.21.2.5)"
	.section	".note.GNU-stack","",@progbits
	.addrsig
