package copy

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var a int
	a = 1
	fmt.Println("start")
	time.Sleep(time.Duration(a) * time.Second)
	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled", time.Now())
			return
		case <-time.After(time.Second * 2):
			fmt.Println("Invoked", time.Now())
			return
		}
	}
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("Cancelled", time.Now())
	//		return
	//	default:
	//		time.Sleep(time.Second * 2)
	//		fmt.Println("Invoked", time.Now())
	//	}
	//}
}
