package main

import (
	"context"
	"fmt"

	"github.com/vsior/udevadm"
)

func main() {
	ctx := context.Background()
	monitor, err := udevadm.NewMonitor(ctx)
	if err != nil {
		panic(err)
	}
	for dev := range monitor.Read() {
		fmt.Println(dev)
	}
}
