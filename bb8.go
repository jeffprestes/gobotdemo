package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

//BB-4a14
func main() {
	bleAdaptor := ble.NewClientAdaptor("BB-4A14")
	myBB8 := bb8.NewDriver(bleAdaptor)

	work := func() {

		myBB8.On("collision", func(data interface{}) {
			collisionInfo := data.(sphero.CollisionPacket)
			if collisionInfo.Speed > 0 {
				fmt.Printf("Collision! %+v\n", data)
				myBB8.SetRGB(255, 0, 0)
				myBB8.Roll(30, uint16(gobot.Rand(360)))
				time.Sleep(time.Second * 2)
			}

		})

		gobot.Every(3*time.Second, func() {
			myBB8.Roll(30, uint16(gobot.Rand(360)))
		})

		gobot.Every(2*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			myBB8.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("BB-8",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{myBB8},
		work,
	)

	robot.Start()
}
