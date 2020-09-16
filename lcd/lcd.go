package lcd

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// LCD pin
	LCD_RST_PIN = rpio.Pin(27)
	LCD_DC_PIN  = rpio.Pin(25)
	LCD_CS_PIN  = rpio.Pin(8)
	LCD_BL_PIN  = rpio.Pin(24)
	// LCD size
	LCD_WIDTH  = 128
	LCD_HEIGHT = 128
	LCD_X      = 2
	LCD_Y      = 1
)

/* ## NOTES ##
# Conversion
py: GPIO.output(pin, value)
go: rpio.WritePin(pin, value)

*/

// ## CONFIG

func epd_digital_write(pin rpio.Pin, value rpio.State) {
	rpio.WritePin(pin, value)
}

func Driver_Delay_ms(xms int) {
	time.Sleep(time.Millisecond * time.Duration(xms))
}

func SPI_Write_Byte(data byte) {
	rpio.SpiTransmit(data)
}

func GPIOinit() {
	rpio.PinMode(LCD_RST_PIN, rpio.Output)
	rpio.PinMode(LCD_DC_PIN, rpio.Output)
	rpio.PinMode(LCD_CS_PIN, rpio.Output)
	rpio.PinMode(LCD_BL_PIN, rpio.Output)
	rpio.SpiSpeed(9000000)

}

// ## LCD
/*
func LCD_Reset() {
	rpio.PinMode(LCD_RST_PIN, High)
	Driver_Delay_ms(100)
	rpio.PinMode(LCD_RST_PIN, rpio.LOW)
	Driver_Delay_ms(100)
	rpio.PinMode(LCD_RST_PIN, rpio.HIGH)
	Driver_Delay_ms(100)
}
*/
