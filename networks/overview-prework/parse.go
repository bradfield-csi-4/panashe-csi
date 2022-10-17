package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"time"
)

const pcapMagicNum = 0xa1b2c3d4

func main() {
	var file *os.File
	var err error

	filename := flag.String("filename", "", "Pass filename")
	flag.Parse()
	if *filename == "" {
		panic("Empty filename")
	}
	file, err = os.Open(*filename)
	if err != nil {
		panic(err)
	}
	fileBytes := make([]byte, 80_000)
	_, err = file.Read(fileBytes)
	if err != nil {
		panic(err)
	}
	parsePcap(fileBytes)

}

type FrameField struct {
	len  int
	data []byte
}

func NewFrameField(dataLen int) FrameField {
	return FrameField{len: dataLen, data: make([]byte, dataLen)}
}

func parsePcap(buffer []byte) {
	fieldNames := []string{
		"magicNum", "majorVer", "minorVer", "TZOffset", "tsAcc",
		"snapshotLen", "llHeaderType",
	}
	pcapFields := map[string]FrameField{
		fieldNames[0]: NewFrameField(4),
		fieldNames[1]: NewFrameField(2),
		fieldNames[2]: NewFrameField(2),
		fieldNames[3]: NewFrameField(4),
		fieldNames[4]: NewFrameField(4),
		fieldNames[5]: NewFrameField(4),
		fieldNames[6]: NewFrameField(4),
	}

	i := 0
	for _, name := range fieldNames {
		field := pcapFields[name]
		copy(field.data, buffer[i:i+field.len])
		i += field.len
	}

	leMagicNum := binary.LittleEndian.Uint32(pcapFields[fieldNames[0]].data)
	beMagicNum := binary.BigEndian.Uint32(pcapFields[fieldNames[0]].data)
	// NOTE: I'm not handling the slight variation for seconds and nanoseconds in timestamp
	if leMagicNum == pcapMagicNum {
		parseFileHeaderLE(fieldNames[1:], pcapFields)
		parsePacketsLE(buffer[24:])
	} else if beMagicNum == pcapMagicNum {
		parseBigEndian(fieldNames[1:], pcapFields)
	} else {
		panic(fmt.Sprintf("Unrecognized magic number.\nBE: %x\nLE: %x", beMagicNum, leMagicNum))
	}
}

func parse4ByteBufferLE(buffer []byte) uint32 {
	return binary.LittleEndian.Uint32(buffer)
}

func parse4ByteFieldLE(field FrameField) uint32 {
	return parse4ByteBufferLE(field.data)
}

func parse2ByteBufferLE(buffer []byte) uint16 {
	return binary.LittleEndian.Uint16(buffer)
}

func parse2ByteFieldLE(field FrameField) uint16 {
	return parse2ByteBufferLE(field.data)
}

func parseFileHeaderLE(names []string, fields map[string]FrameField) {
	var tzOffset, tsAcc FrameField
	var ok bool
	if tzOffset, ok = fields["TZOffset"]; !ok {
		panic("No TZOffset")
	}
	if tsAcc, ok = fields["tsAcc"]; !ok {
		panic("No tsAcc")
	}
	if parse4ByteFieldLE(tzOffset) != 0 || parse4ByteFieldLE(tsAcc) != 0 {
		panic("Unexpected time headers")
	}
	fmt.Println("File info:")
	fmt.Printf(
		"Version: %d.%d\n",
		parse2ByteFieldLE(fields["majorVer"]),
		parse2ByteFieldLE(fields["minorVer"]),
	)
	fmt.Printf(
		"Snapshot length: %d\n",
		parse4ByteFieldLE(fields["snapshotLen"]),
	)
	var llHeaderType string
	// From https://www.tcpdump.org/linktypes.html
	switch parse4ByteFieldLE(fields["llHeaderType"]) {
	case 1:
		llHeaderType = "LINKTYPE_ETHERNET"
	default:
		llHeaderType = "unknown"
	}
	fmt.Printf(
		"Link layer header type: %s\n",
		llHeaderType,
	)
	fmt.Println("===========================\n")
}

func parsePacketsLE(buffer []byte) {
	offset := 0
	counter := 0
	var oldEtherType uint16
	for offset < len(buffer) {
		if offset+16 > len(buffer) {
			fmt.Printf("Incomplete header\n")
			offset += 16
			return
		}
		tsSecs := parse4ByteBufferLE(buffer[:offset+4])
		// NOTE: this can be nanoseconds, depending on the file magic number
		tsMicro := parse4ByteBufferLE(buffer[offset+4 : offset+8])
		capturedPacketLen := parse4ByteBufferLE(buffer[offset+8 : offset+12])
		untruncatedPacketLen := parse4ByteBufferLE(buffer[offset+12 : offset+16])
		offset += 16 // Advance past packet header

		macDest := buffer[offset:offset+6]
		macSource := buffer[offset+6:offset+12] 
		ethertypeVal := parse2ByteBufferLE(buffer[offset+12:offset+14])
		var ethertype string
		if ethertypeVal == 0x08 { // Reference https://en.wikipedia.org/wiki/EtherType#Values
			ethertype = "IPv4"
		} else {
			ethertype = "unknown"
		}

		offset += 14 // Advance past ethernet headers
		ipv4Version := buffer[offset] & 0b1111

		offset += int(capturedPacketLen-14)
		if capturedPacketLen != untruncatedPacketLen {
			panic(
				fmt.Sprintf(
					"Found truncated packet. Captured: %d, original %d",
					capturedPacketLen,
					untruncatedPacketLen,
				),
			)
		}
		if capturedPacketLen == 0 {
			break
		}
		if counter > 1 && ethertypeVal != oldEtherType {
			panic(fmt.Sprintf("found two inconsistent ethertypes %x and %x. Offset %d", oldEtherType, ethertypeVal, offset))
		}
		oldEtherType = ethertypeVal
		counter++
		fmt.Printf("Packet no. %d\n", counter)
		fmt.Printf("timestamp: %v\n", time.Unix(int64(tsSecs), int64(tsMicro*1000)))
		fmt.Printf("Captured packet len: %d\n", capturedPacketLen)
		fmt.Printf("Untruncated packet len: %d\n", untruncatedPacketLen)
		fmt.Printf("MAC destination: %x\n", macDest)
		fmt.Printf("MAC source: %x\n", macSource)
		fmt.Printf("Ethertype: %s\n", ethertype)
		fmt.Printf("IPv4 version: %x\n", ipv4Version)
		fmt.Println()
	}
}

func parseBigEndian(names []string, fieldNames map[string]FrameField) {
	// TODO:
}
