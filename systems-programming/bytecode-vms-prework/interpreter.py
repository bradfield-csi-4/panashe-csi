import dis  # wait so some things survived the fire?
import marshal
import sys
import __builtin__

assert sys.version_info[:2] == (2, 7)


def parse_pyc(f):
    """
    Given a Python 2.7 .pyc file, read the key information and unmarshall and
    return the code object.
    """
    magic_number = f.read(4)
    assert magic_number.encode('hex') == '03f30d0a'
    f.read(4)  # next 4 bytes is the timestamp
    return marshal.load(f)

OP = {
    'NOP'           : 0x09,
    'BINARY_DIVIDE' : 0x15,
    'BINARY_MODULO' : 0x16,
    'BINARY_ADD'    : 0x17,
    'INPLACE_ADD'   : 0x37,
    'GET_ITER'      : 0x44,
    'PRINT_ITEM'    : 0x47,
    'PRINT_NEWLINE' : 0x48,
    'POP_BLOCK'     : 0x57,
    # Below have arguments
    'STORE_NAME'    : 0x5a,
    'FOR_ITER'      : 0x5d,
    'LOAD_CONST'    : 0x64,
    'LOAD_NAME'     : 0x65,
    'JUMP_ABSOLUTE' : 0x71,
    'RETURN_VALUE'  : 0x53,
    'SETUP_LOOP'    : 0x78,
    'CALL_FUNCTION' : 0x83,
    'MAKE_FUNCTION' : 0x84,


}

HAVE_ARGUMENT = 0x5a  # opcodes from here on have an argument

def interpret_bytecode(bytecode, i, values, local_vars, block_stack):
    while i < len(bytecode):
        opcode = ord(bytecode[i])
        i += 1
        if opcode >= HAVE_ARGUMENT:
            oparg = ord(bytecode[i]) + (ord(bytecode[i+1]) << 8)
            i += 2

        if opcode == OP['NOP']:
            continue
        elif opcode == OP['BINARY_DIVIDE']:
            tos = values.pop()
            tos1 = values.pop()
            values.append(tos1 / tos)
        elif opcode == OP['BINARY_MODULO']:
            tos = values.pop()
            tos1 = values.pop()
            values.append(tos1 % tos)
        elif opcode == OP['BINARY_ADD']:
            tos = values.pop()
            tos1 = values.pop()
            values.append(tos1 + tos)
        elif opcode == OP['LOAD_CONST']:
            values.append(code.co_consts[oparg])
        elif opcode == OP['LOAD_NAME']:
            name = code.co_names[oparg]
            if name in local_vars:
                values.append(local_vars[name])
            else:
                builtin = getattr(__builtin__, name)
                values.append(builtin)
        elif opcode == OP['INPLACE_ADD']:
            tos = values.pop()
            tos1 = values.pop()
            tos1 += tos
            values.append(tos1)
        elif opcode == OP['GET_ITER']:
            values.append(iter(values.pop()))
        elif opcode == OP['PRINT_ITEM']:
            print values.pop()
        elif opcode == OP['PRINT_NEWLINE']:
            print
        elif opcode == OP['POP_BLOCK']:
            start, end = block_stack.pop()
            interpret_bytecode(bytecode[start:end], i, values, local_vars, block_stack)
        elif opcode == OP['STORE_NAME']:
            name = code.co_names[oparg]
            local_vars[name] = values.pop()
        elif opcode == OP['FOR_ITER']:
            try:
                values.append(next(values[-1]))
            except StopIteration:
                values.pop()
                i += oparg
        elif opcode == OP['JUMP_ABSOLUTE']:
            i = oparg
        elif opcode == OP['RETURN_VALUE']:
            return values.pop()
        elif opcode == OP['SETUP_LOOP']:
            block_stack.append((i, i+oparg))
        elif opcode == OP['CALL_FUNCTION']:
            n_pos = oparg & 0xff
            n_key = oparg >> 8
            kwargs = {}
            while n_key > 0:
                n_key -= 1
                kw_val = values.pop()
                kw_name = values.pop()
                kwargs[kw_name] = kw_val

            args = []
            while n_pos > 0:
                n_pos -= 1
                args.append(values.pop())

            f = values.pop()
            values.append(f(*args, **kwargs))
        elif opcode == OP['MAKE_FUNCTION']:
            def fun():
                pass
            co = values.pop()
            fun.func_name = co.co_name
            fun.func_code = co
            fun.func_globals[fun.func_name] = fun
            defaults = []
            while (oparg > 0):
                defaults.append(values.pop())
                oparg-=1
            fun.func_defaults = tuple(reversed(defaults))
            values.append(fun)
        else:
            print 'Unknown opcode {:#x}'.format(opcode)
    
def interpret(code):
    """
    Given a code object, interpret (evaluate) the code.
    """
    interpret_bytecode(code.co_code, 0, [], {}, [])


if __name__ == '__main__':
    """
    Unmarshall the code object from the .pyc file reference as a command
    line argument, and intrepret it.

    Usage: python interpreter.py 1.pyc
    """
    f = open(sys.argv[1], 'rb')
    code = parse_pyc(f)
    # dis.dis(code)  # helpful for debugging! but kinda cheating
    print('Interpreting {}...\n'.format(sys.argv[1]))
    ret = interpret(code)
    print('\nFinished interpreting, and returned {}'.format(ret))
