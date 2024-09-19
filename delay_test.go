package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewDelay(t *testing.T) {
	Init()
	d := NewDelay(2 * time.Second)
	err := d.Push("a", "b")
	s := d.cli.Get(context.Background(), "a")
	fmt.Println(s.String())
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(d.Pop())
	d.Close()
	return
}
