package main

import (
	"github.com/jonas-meng/emulator/global"
	"os"
	"os/signal"
	"fmt"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	system := global.Init(3)
	go system.Run()
	fmt.Printf("Signal: %v\n", <-c)
}
