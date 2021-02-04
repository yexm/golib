package timewheel

import (
	"testing"
	"time"
)

func BenchmarkTimingWheel(b *testing.B) {

	tw := NewTimingWheel(time.Second, 100, func(v interface{}) {})
	for i := 0; i < b.N; i++ {
		tw.Start()
		tw.Stop()
	}
}
