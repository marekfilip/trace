package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	given := "Hello from trace package."
	expected := "Hello from trace package.\n"

	if tracer == nil {
		t.Error("Return should not be nil")
	} else {
		tracer.Trace(given)
		if buf.String() != expected {
			t.Errorf("Trace buffer error\nExpected: %s\nActual: %s\n", expected, buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer = Off()
	silentTracer.Trace("Sample message")
}
