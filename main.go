package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	LCD_RST_PIN = rpio.Pin(27)
	LCD_DC_PIN  = rpio.Pin(25)
	LCD_CS_PIN  = rpio.Pin(8)
	LCD_BL_PIN  = rpio.Pin(24)
)

func main() {
	// Clean up incase it exited improperly
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println(sig)
			fmt.Println("Performing safe exit")
			rpio.Close()
			os.Exit(2)

		}
	}()
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Start pressing buttons on your interface!")
	for {

		time.Sleep(150000000)

		// key UP
		if rpio.Pin(6).Read() == 0 {
			fmt.Println("UP")
		}

		// key DOWN
		if rpio.Pin(19).Read() == 0 {
			fmt.Println("DOWN")
		}

		//key LEFT
		if rpio.Pin(5).Read() == 0 {
			fmt.Println("LEFT")
		}

		//key RIGHT
		if rpio.Pin(26).Read() == 0 {
			fmt.Println("RIGHT")
		}

		//key CLICK
		if rpio.Pin(13).Read() == 0 {
			fmt.Println("CLICK")
		}

		//key 1
		if rpio.Pin(21).Read() == 0 {
			fmt.Println("Key1")
		}

		//key 2
		if rpio.Pin(20).Read() == 0 {
			fmt.Println("Key2")
		}

		//key 3
		if rpio.Pin(16).Read() == 0 {
			fmt.Println("Key3")
		}

		// kill
		if rpio.Pin(13).Read() == 0 && rpio.Pin(20).Read() == 0 {
			rpio.Close()
			os.Exit(5)
		}
	}
	rpio.Close()
}
