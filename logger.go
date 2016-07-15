package htmllog

// Copyright 2016 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bytes"
	"html"
	"io"
	"os"
	"sync"
	"time"
)

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Fatal = "fatal"
)

// DefaultMaxMessageSize specifies the maxium length that a message can be.
var DefaultMaxMessageSize = 200

// EventFunc is the signature for a custom definition how to write a log event.
type EventFunc func(hyper *Logger, logLevel, htmlEscapedMessage string)

// Logger provides a thread-safe access to the Html file for writing log events.
type Logger struct {
	protect        *sync.Mutex
	name           string
	out            io.Writer
	htmlStyle      string
	scrollToBottom string
	eventFunc      EventFunc
	messageLimit   int
}

// New creates a new logger and opens a new (or existing) file.
func New(name string) (*Logger, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return NewOn(f)
}

// NewOn creates a new logger that uses a writer.
func NewOn(w io.Writer) (*Logger, error) {
	l := &Logger{name: "internal-writer",
		out:     w,
		protect: new(sync.Mutex),
	}
	l.Configure(DefaultScrollToBottom, DefaultStyle, DefaultMaxMessageSize, LogEvent)
	l.setup()
	return l, nil
}

// Configure is to override default behavior of the logger.
func (l *Logger) Configure(scroll string, style string, newLimit int, logFunc EventFunc) {
	l.scrollToBottom = scroll
	l.htmlStyle = style
	l.messageLimit = newLimit
	l.eventFunc = logFunc
}

// LogEvent is the default implementation for an EventFunc and will write
// a timestamp, the log level, the (limited) message and a Html linebreak.
// This function is called in a protected region.
func LogEvent(hyper *Logger, logLevel, htmlEscapedMessage string) {
	hyper.Timestamp()
	hyper.Level(logLevel)
	hyper.Div(logLevel+"msg", func() {
		hyper.Raw(htmlEscapedMessage)
	})
	hyper.Br()
}

// setup writes the start of a new Html document
func (l *Logger) setup() {
	io.WriteString(l.out, "<html><head>")
	io.WriteString(l.out, l.scrollToBottom)
	io.WriteString(l.out, "<style>")
	io.WriteString(l.out, l.htmlStyle)
	io.WriteString(l.out, "</style>")
	io.WriteString(l.out, "</head><body onLoad=\"javascript:loaded();\">")
}

// Event is the generic function for writing a log event
func (l *Logger) Event(level string, format string, args ...interface{}) {
	l.protect.Lock()
	defer l.protect.Unlock()
	l.eventFunc(l, level, html.EscapeString(LimitedSprintf(l.messageLimit, format, args...)))
}

// Level writes the log level. Must be called in a protected region.
func (l *Logger) Level(level string) {
	l.Div(level, func() { l.Raw("[" + l.Nbsp(level, 5) + "]") })
}

func (l *Logger) Reset() error {
	l.protect.Lock()
	defer l.protect.Unlock()
	if l.out == nil {
		return nil
	}
	f, ok := l.out.(*os.File)
	if !ok {
		return nil
	}
	if err := f.Close(); err != nil {
		return err
	}
	f, err := os.Create(l.name)
	if err != nil {
		return err
	}
	l.out = f
	l.setup()
	return nil
}

// Debugf writes a message with Debug level
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Event(Debug, format, args...)
}

// Infof writes a message with Info level
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Event(Info, format, args...)
}

// Warnf writes a message with Warn level
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Event(Warn, format, args...)
}

// Errorf writes a message with Error level
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Event(Error, format, args...)
}

// Fatalf writes a message with Fatal level
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Event(Fatal, format, args...)
}

// Raw writes the message as is. Must be called in a protected region.
func (l *Logger) Raw(htmlEscapedMessage string) {
	io.WriteString(l.out, htmlEscapedMessage)
}

// Br writes a Html linebreak. Must be called in a protected region.
func (l *Logger) Br() {
	io.WriteString(l.out, "<br>")
}

// Timestamp writes a formatted time (now). Must be called in a protected region.
func (l *Logger) Timestamp() {
	l.Div("time", func() { io.WriteString(l.out, time.Now().Format("01-02 15:04:05")) })
}

// Div writes the begin and end div tags around the block call. Must be called in a protected region.
func (l *Logger) Div(cls string, block func()) {
	io.WriteString(l.out, "<div class=\"")
	io.WriteString(l.out, cls)
	io.WriteString(l.out, "\">")
	block()
	io.WriteString(l.out, "</div>")
}

// Nbsp returns a new string of a fixed size, padded with Html non-breaking spaces.
func (l *Logger) Nbsp(s string, size int) string {
	var b bytes.Buffer
	b.WriteString(s)
	for i := len(s); i < size; i++ {
		b.WriteString("&nbsp;")
	}
	return b.String()
}
