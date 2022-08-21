	.file	"unoptimized-pagecount3.c"
	.text
	.globl	pagecount
	.type	pagecount, @function
pagecount:
.LFB11:
	.cfi_startproc
	movl	%esi, %ecx
	shrq	%cl, %rdi
	movabsq	$-6148914691236517205, %rdx
	movq	%rdi, %rax
	mulq	%rdx
	movq	%rdx, %rax
	shrq	%rax
	ret
	.cfi_endproc
.LFE11:
	.size	pagecount, .-pagecount
	.section	.rodata.str1.8,"aMS",@progbits,1
	.align 8
.LC3:
	.string	"%.2fs to run %d tests (%.2fns per test)\n"
	.text
	.globl	main
	.type	main, @function
main:
.LFB12:
	.cfi_startproc
	pushq	%r15
	.cfi_def_cfa_offset 16
	.cfi_offset 15, -16
	pushq	%r14
	.cfi_def_cfa_offset 24
	.cfi_offset 14, -24
	pushq	%r13
	.cfi_def_cfa_offset 32
	.cfi_offset 13, -32
	pushq	%r12
	.cfi_def_cfa_offset 40
	.cfi_offset 12, -40
	pushq	%rbp
	.cfi_def_cfa_offset 48
	.cfi_offset 6, -48
	pushq	%rbx
	.cfi_def_cfa_offset 56
	.cfi_offset 3, -56
	subq	$88, %rsp
	.cfi_def_cfa_offset 144
	movabsq	$4294967296, %rax
	movq	%rax, 48(%rsp)
	movabsq	$1099511627776, %rax
	movq	%rax, 56(%rsp)
	movabsq	$4503599627370496, %rax
	movq	%rax, 64(%rsp)
	movq	$12, 16(%rsp)
	movq	$16, 24(%rsp)
	movq	$32, 32(%rsp)
	call	clock
	movq	%rax, 8(%rsp)
	movl	$0, %ebp
	movl	$0, %edx
	jmp	.L3
.L4:
	movslq	%edx, %rax
	imulq	$1431655766, %rax, %rax
	shrq	$32, %rax
	movl	%edx, %ecx
	sarl	$31, %ecx
	subl	%ecx, %eax
	leal	(%rax,%rax,2), %ecx
	movl	%edx, %eax
	subl	%ecx, %eax
	cltq
	movq	16(%rsp,%rax,8), %rcx
	movl	%ecx, %esi
	addl	48(%rsp,%rax,8), %esi
	leal	1(%rsi,%rbp), %ebp
	addl	$1, %edx
.L3:
	cmpl	$9999999, %edx
	jle	.L4
	call	clock
	movq	%rax, %r14
	call	clock
	movq	%rax, %r15
	movl	$0, %ebx
	jmp	.L5
.L6:
	movslq	%ebx, %rax
	imulq	$1431655766, %rax, %rax
	shrq	$32, %rax
	movl	%ebx, %edx
	sarl	$31, %edx
	subl	%edx, %eax
	leal	(%rax,%rax,2), %edx
	movl	%ebx, %eax
	subl	%edx, %eax
	cltq
	movq	48(%rsp,%rax,8), %r13
	movq	16(%rsp,%rax,8), %r12
	movq	%r12, %rsi
	movq	%r13, %rdi
	call	pagecount
	addq	%r13, %rax
	addq	%r12, %rax
	addl	%eax, %ebp
	addl	$1, %ebx
.L5:
	cmpl	$9999999, %ebx
	jle	.L6
	call	clock
	subq	%r15, %rax
	subq	8(%rsp), %r14
	subq	%r14, %rax
	pxor	%xmm0, %xmm0
	cvtsi2sdq	%rax, %xmm0
	divsd	.LC0(%rip), %xmm0
	movapd	%xmm0, %xmm1
	mulsd	.LC1(%rip), %xmm1
	divsd	.LC2(%rip), %xmm1
	movl	$10000000, %esi
	movl	$.LC3, %edi
	movl	$2, %eax
	call	printf
	movl	%ebp, %eax
	addq	$88, %rsp
	.cfi_def_cfa_offset 56
	popq	%rbx
	.cfi_def_cfa_offset 48
	popq	%rbp
	.cfi_def_cfa_offset 40
	popq	%r12
	.cfi_def_cfa_offset 32
	popq	%r13
	.cfi_def_cfa_offset 24
	popq	%r14
	.cfi_def_cfa_offset 16
	popq	%r15
	.cfi_def_cfa_offset 8
	ret
	.cfi_endproc
.LFE12:
	.size	main, .-main
	.section	.rodata.cst8,"aM",@progbits,8
	.align 8
.LC0:
	.long	0
	.long	1093567616
	.align 8
.LC1:
	.long	0
	.long	1104006501
	.align 8
.LC2:
	.long	0
	.long	1097011920
	.ident	"GCC: (GNU) 11.3.0"
	.section	.note.GNU-stack,"",@progbits
