package trace
import (
	"bytes"
	"testing"
)

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("nilです")
	} else {
		tracer.Trace("こんにちは、trace")
		if buf.String() != "こんにちは、trace¥n" {
			t.Errorf("'%s'という誤った文字列です", buf.String())
		}
	}
}