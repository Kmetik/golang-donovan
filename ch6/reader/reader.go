package reader

import "io"

type limitReader struct {
	r     io.Reader
	limit int
}

// Read([]byte) (int, error)

type stringReader struct {
	s string
}

func (sr stringReader) Read(p []byte) (int, error) {
	n := copy(p, sr.s)
	sr.s = sr.s[n:]
	if len(sr.s) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func (lr *limitReader) Read(p []byte) (int, error) {
	n, err := lr.r.Read(p[:lr.limit])
	if err != nil {
		return 0, err
	}
	return n, nil
}

func LimitReader(r io.Reader, limit int) io.Reader {
	nr := limitReader{r: r, limit: limit}
	return nr.r
}

func NewReader(s string) io.Reader {

	return stringReader{s: s}
}
