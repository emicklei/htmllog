package htmllog

// Copyright 2016 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
)

const (
	endOfBuffer = "EOB"
)

type limitedBuffer struct {
	*bytes.Buffer
	// keep track of how many bytes have been written
	written int
	// maximum number of bytes that can be written
	max int
}

// Write will attempt to write all bytes of p. If the write limit is reached then panic.
func (b *limitedBuffer) Write(p []byte) (n int, err error) {
	if b.written+len(p) > b.max {
		b.Buffer.Write(p[:b.max-b.written])
		// limit reached; no need to update written field
		panic(endOfBuffer)
	}
	n, err = b.Buffer.Write(p)
	b.written += n
	return
}

// sprintf is recoverable call to fmt.Fprintf
func (b *limitedBuffer) sprintf(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err == endOfBuffer {
			return
		}
	}()
	fmt.Fprintf(b, format, args...)
}

// LimitedSprintf returns the result of fmt.Sprintf limited to a number of bytes.
// Use this function to protect against printing recursive structures.
func LimitedSprintf(limit int, format string, args ...interface{}) string {
	w := &limitedBuffer{
		Buffer:  new(bytes.Buffer),
		written: 0,
		max:     limit,
	}
	w.sprintf(format, args...)
	return w.Buffer.String()
}
