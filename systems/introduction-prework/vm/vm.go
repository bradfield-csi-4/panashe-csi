package vm

import "fmt"

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//

func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	// Keep looping, like a physical computer's clock
loop:
	for {

		op := memory[registers[0]] // fetch the opcode
		registers[0]++
		// decode and execute
		switch op {
		case Load:
			var reg, addr byte
			reg, registers[0] = memory[registers[0]], registers[0]+1
			addr, registers[0] = memory[registers[0]], registers[0]+1
			registers[reg] = memory[addr]
		case Store:
			var reg, addr byte
			reg, registers[0] = memory[registers[0]], registers[0]+1
			addr, registers[0] = memory[registers[0]], registers[0]+1
			memory[addr] = registers[reg]
		case Add:
			var reg1, reg2 byte
			reg1, registers[0] = memory[registers[0]], registers[0]+1
			reg2, registers[0] = memory[registers[0]], registers[0]+1
			registers[reg1] = registers[reg1] + registers[reg2]
		case Sub:
			var reg1, reg2 byte
			reg1, registers[0] = memory[registers[0]], registers[0]+1
			reg2, registers[0] = memory[registers[0]], registers[0]+1
			registers[reg1] = registers[reg1] - registers[reg2]
		case Addi:
			var reg, v byte
			reg, registers[0] = memory[registers[0]], registers[0]+1
			v, registers[0] = memory[registers[0]], registers[0]+1
			registers[reg] = registers[reg] + v
		case Subi:
			var reg, v byte
			reg, registers[0] = memory[registers[0]], registers[0]+1
			v, registers[0] = memory[registers[0]], registers[0]+1
			registers[reg] = registers[reg] - v
		case Jump:
			var loc byte
			loc, registers[0] = memory[registers[0]], registers[0]+1
			registers[0] = loc
		case Beqz:
			var reg, offset byte
			reg, registers[0] = memory[registers[0]], registers[0]+1
			offset, registers[0] = memory[registers[0]], registers[0]+1
			if registers[reg] == 0 {
				registers[0] += offset
			}
		case Halt:
			break loop
		default:
			panic(fmt.Sprintf("Unrecognized operation %x", op))
		}
	}
}
