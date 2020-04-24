package dto

import (
	"github.com/jacobsa/go-serial/serial"
)

// SerialRequest is a DTO
type SerialRequest struct {
	PortName                string
	BaudRate                uint
	DataBits                uint
	StopBits                uint
	ParityMode              serial.ParityMode
	RTSCTSFlowControl       bool
	InterCharacterTimeout   uint
	MinimumReadSize         uint
	Rs485Enable             bool
	Rs485RtsHighDuringSend  bool
	Rs485RtsHighAfterSend   bool
	Rs485RxDuringTx         bool
	Rs485DelayRtsBeforeSend int
	Rs485DelayRtsAfterSend  int

	DataToSend []byte
}

// SendToSerial Open serial port and write data
func (sr *SerialRequest) SendToSerial() ResAPI {
	options := serial.OpenOptions{
		PortName:                sr.PortName,
		BaudRate:                sr.BaudRate,
		DataBits:                sr.DataBits,
		StopBits:                sr.StopBits,
		ParityMode:              sr.ParityMode,
		RTSCTSFlowControl:       sr.RTSCTSFlowControl,
		InterCharacterTimeout:   sr.InterCharacterTimeout,
		MinimumReadSize:         sr.MinimumReadSize,
		Rs485Enable:             sr.Rs485Enable,
		Rs485RtsHighDuringSend:  sr.Rs485RtsHighDuringSend,
		Rs485RtsHighAfterSend:   sr.Rs485RtsHighAfterSend,
		Rs485RxDuringTx:         sr.Rs485RxDuringTx,
		Rs485DelayRtsBeforeSend: sr.Rs485DelayRtsBeforeSend,
		Rs485DelayRtsAfterSend:  sr.Rs485DelayRtsAfterSend,
	}

	stream, err := serial.Open(options)
	if err != nil {
		return ResAPI{
			Status: 400,
			Result: err.Error(),
		}
	}
	defer stream.Close()

	_, err = stream.Write(sr.DataToSend)
	if err != nil {
		return ResAPI{
			Status: 400,
			Result: err.Error(),
		}
	}

	return ResAPI{
		Status: 200,
		Result: "OK",
	}
}
