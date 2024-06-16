package bytecounter

import (
	"bufio"
	"io"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	var start int = 0
	for {
		advance, token, _ := bufio.ScanWords(p[start:], true)
		if len(token) == 0 {
			break
		}
		start += advance
		*c += 1
	}
	return len(p), nil
}

type CounterWriter struct {
	counter int64
	writer  io.Writer
}

func (cw *CounterWriter) Write(p []byte) (int, error) {
	cw.counter += int64(len(p))
	return cw.writer.Write(p)
}

// // newWriter is a Writer Wrapper, return original Writer
// // and a Counter which record bytes have written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{0, w}
	return &cw, &cw.counter
}
