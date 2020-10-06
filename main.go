package main

import (
	"fmt"
	"image/color"
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})

	d := ssd1306.NewI2C(machine.I2C0)
	d.Configure(ssd1306.Config{
		Width: 128, Height: 32,
		Address: ssd1306.Address_128_32,
	})

	d.ClearDisplay()
	fmt.Println(`Start`) // To avoid hang up at 40sec.

	black := color.RGBA{1, 1, 1, 255}

	for i := 0; ; i++ {
		d.ClearBuffer()
		tinyfont.WriteLine(&d, &freemono.Regular9pt7b, 0, 12, strconv.FormatInt(int64(i), 10), black)
		d.Display()
		time.Sleep(1000 * time.Millisecond)
	}
}
