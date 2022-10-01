#ifndef Py_OPCODE_H
#define Py_OPCODE_H
#ifdef __cplusplus
extern "C" {
#endif


/* Instruction opcodes for compiled code */

#define STOP_CODE	0x00
#define POP_TOP		0x01
#define ROT_TWO		0x02
#define ROT_THREE	0x03
#define DUP_TOP		0x04
#define ROT_FOUR	0x05
#define NOP		0x09

#define UNARY_POSITIVE	0x0a
#define UNARY_NEGATIVE	0x0b
#define UNARY_NOT	0x0c
#define UNARY_CONVERT	0x0d

#define UNARY_INVERT	0x0f

#define BINARY_POWER	0x13

#define BINARY_MULTIPLY	0x14
#define BINARY_DIVIDE	0x15
#define BINARY_MODULO	0x16
#define BINARY_ADD	0x17
#define BINARY_SUBTRACT	0x18
#define BINARY_SUBSCR	0x19
#define BINARY_FLOOR_DIVIDE 0x1a
#define BINARY_TRUE_DIVIDE 0x1b
#define INPLACE_FLOOR_DIVIDE 0x1c
#define INPLACE_TRUE_DIVIDE 0x1d

#define SLICE		0x1e
/* Also uses 0x1f-0x21 */
#define SLICE_0x01		0x1f
#define SLICE_0x02		0x20
#define SLICE_0x03		0x21

#define STORE_SLICE	0x28
/* Also uses 0x29-0x2b */
#define STORE_SLICE_0x01	0x29
#define STORE_SLICE_0x02	0x2a
#define STORE_SLICE_0x03	0x2b

#define DELETE_SLICE	0x32
/* Also uses 0x33-0x35 */
#define DELETE_SLICE_0x01	0x33
#define DELETE_SLICE_0x02	0x34
#define DELETE_SLICE_0x03	0x35

#define STORE_MAP	0x36
#define INPLACE_ADD	0x37
#define INPLACE_SUBTRACT	0x38
#define INPLACE_MULTIPLY	0x39
#define INPLACE_DIVIDE	0x3a
#define INPLACE_MODULO	0x3b
#define STORE_SUBSCR	0x3c
#define DELETE_SUBSCR	0x3d

#define BINARY_LSHIFT	0x3e
#define BINARY_RSHIFT	0x3f
#define BINARY_AND	0x40
#define BINARY_XOR	0x41
#define BINARY_OR	0x42
#define INPLACE_POWER	0x43
#define GET_ITER	0x44

#define PRINT_EXPR	0x46
#define PRINT_ITEM	0x47
#define PRINT_NEWLINE	0x48
#define PRINT_ITEM_TO   0x49
#define PRINT_NEWLINE_TO 0x4a
#define INPLACE_LSHIFT	0x4b
#define INPLACE_RSHIFT	0x4c
#define INPLACE_AND	0x4d
#define INPLACE_XOR	0x4e
#define INPLACE_OR	0x4f
#define BREAK_LOOP	0x50
#define WITH_CLEANUP    0x51
#define LOAD_LOCALS	0x52
#define RETURN_VALUE	0x53
#define IMPORT_STAR	0x54
#define EXEC_STMT	0x55
#define YIELD_VALUE	0x56
#define POP_BLOCK	0x57
#define END_FINALLY	0x58
#define BUILD_CLASS	0x59

#define HAVE_ARGUMENT	0x5a	/* Opcodes from here have an argument: */

#define STORE_NAME	0x5a	/* Index in name list */
#define DELETE_NAME	0x5b	/* "" */
#define UNPACK_SEQUENCE	0x5c	/* Number of sequence items */
#define FOR_ITER	0x5d
#define LIST_APPEND	0x5e

#define STORE_ATTR	0x5f	/* Index in name list */
#define DELETE_ATTR	0x60	/* "" */
#define STORE_GLOBAL	0x61	/* "" */
#define DELETE_GLOBAL	0x62	/* "" */
#define DUP_TOPX	0x63	/* number of items to duplicate */
#define LOAD_CONST	0x64	/* Index in const list */
#define LOAD_NAME	0x65	/* Index in name list */
#define BUILD_TUPLE	0x66	/* Number of tuple items */
#define BUILD_LIST	0x67	/* Number of list items */
#define BUILD_SET	0x68     /* Number of set items */
#define BUILD_MAP	0x69	/* Always zero for now */
#define LOAD_ATTR	0x6a	/* Index in name list */
#define COMPARE_OP	0x6b	/* Comparison operator */
#define IMPORT_NAME	0x6c	/* Index in name list */
#define IMPORT_FROM	0x6d	/* Index in name list */
#define JUMP_FORWARD	0x6e	/* Number of bytes to skip */

#define JUMP_IF_FALSE_OR_POP 0x6f /* Target byte offset from beginning
                                    of code */
#define JUMP_IF_TRUE_OR_POP 0x70	/* "" */
#define JUMP_ABSOLUTE	0x71	/* "" */
#define POP_JUMP_IF_FALSE 0x72	/* "" */
#define POP_JUMP_IF_TRUE 0x73	/* "" */

#define LOAD_GLOBAL	0x74	/* Index in name list */

#define CONTINUE_LOOP	0x77	/* Start of loop (absolute) */
#define SETUP_LOOP	0x78	/* Target address (relative) */
#define SETUP_EXCEPT	0x79	/* "" */
#define SETUP_FINALLY	0x7a	/* "" */

#define LOAD_FAST	0x7c	/* Local variable number */
#define STORE_FAST	0x7d	/* Local variable number */
#define DELETE_FAST	0x7e	/* Local variable number */

#define RAISE_VARARGS	0x82	/* Number of raise arguments (0x01, 0x02 or 0x03) */
/* CALL_FUNCTION_XXX opcodes defined below depend on this definition */
#define CALL_FUNCTION	0x83	/* #args + (#kwargs<<0x08) */
#define MAKE_FUNCTION	0x84	/* #defaults */
#define BUILD_SLICE 	0x85	/* Number of items */

#define MAKE_CLOSURE    0x86     /* #free vars */
#define LOAD_CLOSURE    0x87     /* Load free variable from closure */
#define LOAD_DEREF      0x88     /* Load and dereference from closure cell */ 
#define STORE_DEREF     0x89     /* Store into cell */ 

/* The next 0x03 opcodes must be contiguous and satisfy
   (CALL_FUNCTION_VAR - CALL_FUNCTION) & 0x03 == 0x01  */
#define CALL_FUNCTION_VAR          0x8c	/* #args + (#kwargs<<0x08) */
#define CALL_FUNCTION_KW           0x8d	/* #args + (#kwargs<<0x08) */
#define CALL_FUNCTION_VAR_KW       0x8e	/* #args + (#kwargs<<0x08) */

#define SETUP_WITH 0x8f

/* Support for opargs more than 0x10 bits long */
#define EXTENDED_ARG  0x91

#define SET_ADD         0x92
#define MAP_ADD         0x93


enum cmp_op {PyCmp_LT=Py_LT, PyCmp_LE=Py_LE, PyCmp_EQ=Py_EQ, PyCmp_NE=Py_NE, PyCmp_GT=Py_GT, PyCmp_GE=Py_GE,
	     PyCmp_IN, PyCmp_NOT_IN, PyCmp_IS, PyCmp_IS_NOT, PyCmp_EXC_MATCH, PyCmp_BAD};

#define HAS_ARG(op) ((op) >= HAVE_ARGUMENT)

#ifdef __cplusplus
}
#endif
#endif /* !Py_OPCODE_H */
