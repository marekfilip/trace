package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface that describes an object
// capable of tracing event throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

// Trace method of tracer writes
// any count of message to buffer
func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// Trace method of nilTracer method
// does not write any message
func (t *nilTracer) Trace(a ...interface{}) {
}

// Creates new tracer that implements
// Tracer interface
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Creates new empty tracer that implements
// Tracer interface
func Off() Tracer {
	return &nilTracer{}
}
