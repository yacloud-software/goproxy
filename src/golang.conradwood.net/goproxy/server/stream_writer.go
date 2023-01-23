package main

import (
	"bytes"
	hg "golang.conradwood.net/apis/h2gproxy"
)

type StreamWriter struct {
	rec receiver
	buf *bytes.Buffer
}
type receiver interface {
	Send(*hg.StreamDataResponse) error
}

func NewStreamWriter(rec receiver) *StreamWriter {
	res := &StreamWriter{rec: rec, buf: &bytes.Buffer{}}
	return res
}
func (sw *StreamWriter) Bytes() []byte {
	return sw.buf.Bytes()
}
func (sw *StreamWriter) Write(buf []byte) (int, error) {
	sw.buf.Write(buf)
	offset := 0
	repeat := true
	for repeat {
		size := 8192
		if offset+size > len(buf) {
			size = len(buf) - offset
			repeat = false
		}
		data := buf[offset : offset+size]
		err := sw.rec.Send(&hg.StreamDataResponse{Data: data})
		if err != nil {
			return 0, err
		}
		offset = offset + size
	}
	return len(buf), nil
}
