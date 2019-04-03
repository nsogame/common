package packets

import (
	"bytes"
	"io"
	"log"

	"github.com/Mempler/osubinary"
	"github.com/pkg/errors"
)

func Parse(body []byte) (packets []Packet, err error) {
	var id uint16
	var length uint32
	var actualLength int

	reader := bytes.NewReader(body)
	for {
		id, err = osubinary.RUInt16(reader)
		if err == io.EOF {
			break
		} else if err != nil {
			err = errors.Wrap(err, "packets.Parse: error reading packet type")
			return
		}

		// skip a byte
		_, err = osubinary.RInt8(reader)
		if err != nil {
			err = errors.Wrap(err, "packets.Parse: error reading skipbyte")
			return
		}

		length, err = osubinary.RUInt32(reader)
		if err != nil {
			err = errors.Wrap(err, "packets.Parse: error reading length")
			return
		}

		// actually read the data
		data := make([]byte, length)
		actualLength, err = reader.Read(data)
		if actualLength < int(length) {
			err = errors.Wrap(err, "packets.Parse: packet is too short")
			return
		}

		var packet Packet
		switch id {
		default:
			log.Println("Unknown packet number", id)
			continue
		}
		packets = append(packets, packet)
	}
	return
}
