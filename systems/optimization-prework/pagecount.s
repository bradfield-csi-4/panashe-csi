<<<<<<< HEAD
	.file	"pagecount.c"
	.text
	.globl	pagecount
	.type	pagecount, @function
pagecount:
.LFB11:
	.cfi_startproc
	movq	%rdi, %rax
	movl	%esi, %ecx
	shrq	%cl, %rax
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
=======
	.arch armv8-a
	.text
	.align	2
	.globl _pagecount
_pagecount:
LFB1:
	udiv	x0, x0, x1
	ret
LFE1:
	.section __TEXT,__text_startup,regular,pure_instructions
	.align	2
	.globl _main
_main:
LFB2:
	stp	x29, x30, [sp, -112]!
LCFI0:
	mov	x29, sp
	stp	x19, x20, [sp, 16]
	stp	x21, x22, [sp, 32]
LCFI1:
	adrp	x0, ___stack_chk_guard
	ldr	x0, [x0, ___stack_chk_guard];momd
	ldr	x1, [x0]
	str	x1, [sp, 104]
	mov	x1, 0
	adrp	x0, lC0@GOTPAGE
	ldr	x0, [x0, lC0@GOTPAGEOFF];momd
	ldp	x2, x3, [x0]
	stp	x2, x3, [sp, 56]
	ldr	x0, [x0, 16]
	str	x0, [sp, 72]
	adrp	x0, lC1@GOTPAGE
	ldr	x0, [x0, lC1@GOTPAGEOFF];momd
	ldp	x2, x3, [x0]
	stp	x2, x3, [sp, 80]
	ldr	x0, [x0, 16]
	str	x0, [sp, 96]
	bl	_clock
	mov	w19, 0
	mov	w1, 0
	b	L3
L4:
	mov	w0, 21846
	movk	w0, 0x5555, lsl 16
	smull	x0, w1, w0
	lsr	x0, x0, 32
	sub	w0, w0, w1, asr 31
	add	w0, w0, w0, lsl 1
	sub	w0, w1, w0
	sxtw	x0, w0
	add	x2, sp, 56
	ldr	x2, [x2, x0, lsl 3]
	add	x3, sp, 80
	ldr	x0, [x3, x0, lsl 3]
	add	w0, w2, w0
	add	w0, w19, w0
	add	w19, w0, 1
	add	w1, w1, 1
L3:
	mov	w0, 38527
	movk	w0, 0x98, lsl 16
	cmp	w1, w0
	ble	L4
	bl	_clock
	bl	_clock
	mov	w20, 0
	b	L5
L6:
	mov	w0, 21846
	movk	w0, 0x5555, lsl 16
	smull	x0, w20, w0
	lsr	x0, x0, 32
	sub	w0, w0, w20, asr 31
	add	w0, w0, w0, lsl 1
	sub	w0, w20, w0
	sxtw	x0, w0
	add	x1, sp, 56
	ldr	x22, [x1, x0, lsl 3]
	add	x1, sp, 80
	ldr	x21, [x1, x0, lsl 3]
	mov	x1, x21
	mov	x0, x22
	bl	_pagecount
	add	x1, x0, x22
	add	x1, x1, x21
	add	w19, w19, w1
	add	w20, w20, 1
L5:
	mov	w0, 38527
	movk	w0, 0x98, lsl 16
	cmp	w20, w0
	ble	L6
	bl	_clock
	adrp	x0, ___stack_chk_guard
	ldr	x0, [x0, ___stack_chk_guard];momd
	ldr	x2, [sp, 104]
	ldr	x1, [x0]
	subs	x2, x2, x1
	mov	x1, 0
	bne	L9
	mov	w0, w19
	ldp	x19, x20, [sp, 16]
	ldp	x21, x22, [sp, 32]
	ldp	x29, x30, [sp], 112
LCFI2:
	ret
L9:
LCFI3:
	bl	___stack_chk_fail
LFE2:
	.const
	.align	3
lC0:
	.xword	4294967296
	.xword	1099511627776
	.xword	4503599627370496
	.align	3
lC1:
	.xword	4096
	.xword	65536
	.xword	4294967296
	.section __TEXT,__text_startup,regular,pure_instructions
	.section __TEXT,__eh_frame,coalesced,no_toc+strip_static_syms+live_support
EH_frame1:
	.set L$set$0,LECIE1-LSCIE1
	.long L$set$0
LSCIE1:
	.long	0
	.byte	0x1
	.ascii "zR\0"
	.byte	0x1
	.byte	0x78
	.byte	0x1e
	.byte	0x1
	.byte	0x10
	.byte	0xc
	.byte	0x1f
	.byte	0
	.align	3
LECIE1:
LSFDE1:
	.set L$set$1,LEFDE1-LASFDE1
	.long L$set$1
LASFDE1:
	.long	LASFDE1-EH_frame1
	.quad	LFB1-.
	.set L$set$2,LFE1-LFB1
	.quad L$set$2
	.byte	0
	.align	3
LEFDE1:
LSFDE3:
	.set L$set$3,LEFDE3-LASFDE3
	.long L$set$3
LASFDE3:
	.long	LASFDE3-EH_frame1
	.quad	LFB2-.
	.set L$set$4,LFE2-LFB2
	.quad L$set$4
	.byte	0
	.byte	0x4
	.set L$set$5,LCFI0-LFB2
	.long L$set$5
	.byte	0xe
	.byte	0x70
	.byte	0x9d
	.byte	0xe
	.byte	0x9e
	.byte	0xd
	.byte	0x4
	.set L$set$6,LCFI1-LCFI0
	.long L$set$6
	.byte	0x93
	.byte	0xc
	.byte	0x94
	.byte	0xb
	.byte	0x95
	.byte	0xa
	.byte	0x96
	.byte	0x9
	.byte	0x4
	.set L$set$7,LCFI2-LCFI1
	.long L$set$7
	.byte	0xa
	.byte	0xde
	.byte	0xdd
	.byte	0xd5
	.byte	0xd6
	.byte	0xd3
	.byte	0xd4
	.byte	0xe
	.byte	0
	.byte	0x4
	.set L$set$8,LCFI3-LCFI2
	.long L$set$8
	.byte	0xb
	.align	3
LEFDE3:
	.ident	"GCC: (GNU) 11.2.0"
	.subsections_via_symbols
>>>>>>> be55fec (Add systems assembly and optimization prework)
