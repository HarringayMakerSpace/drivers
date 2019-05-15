package main

import (
	"machine"

	"github.com/tinygo-org/drivers/ubloxgps"
)

func main() {
	println("GPS Example")
	machine.I2C0.Configure(machine.I2CConfig{})
	gps := ubloxgps.New(machine.I2C0)
	parser := ubloxgps.Parser(gps)
	var fix ubloxgps.Fix
	for {
		parser.NextFix(&fix)
		if fix.Valid {
			print(fix.Time)
			print(", lat=", fix.Latitude)
			print(", long=", fix.Longitude)
			print(", altitude:=", fix.Altitude)
			print(", satellites=", fix.Satellites)
			println()
		} else {
			println("No fix")
		}
	}
}
