package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/sphero"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor("BB-4A14")
	myBB8 := bb8.NewDriver(bleAdaptor)
	mqttAdaptor := mqtt.NewAdaptor("tcp://iot.eclipse.org:1883", "raspberry-bb8")

	work := func() {

		myBB8.On("collision", func(data interface{}) {
			collisionInfo := data.(sphero.CollisionPacket)
			if collisionInfo.Speed > 0 {
				fmt.Printf("Collision! %+v\n", collisionInfo)
				myBB8.SetRGB(255, 0, 0)
				myBB8.Roll(40, uint16(180))
				time.Sleep(time.Second * 2)
			} else {
				r := uint8(gobot.Rand(255))
				g := uint8(gobot.Rand(255))
				b := uint8(gobot.Rand(255))
				myBB8.SetRGB(r, g, b)
				myBB8.Roll(40, uint16(gobot.Rand(360)))
			}

		})

		myBB8.On("sensordata", func(data interface{}) {
			botinfo := data.(sphero.DataStreamingPacket)
			fmt.Printf("[Lendo sensores]! %+v\n", botinfo)
			if botinfo.VeloX == 0 && botinfo.VeloY == 0 {
				fmt.Printf("\r\nTa parado! %+v\n", data)

			}

		})

		mqttAdaptor.On("/jeffprestes/bb8", func(msg mqtt.Message) {
			fmt.Printf("[main][work][mqttadption][on] %+v\r\n", msg)
			msgText := string(msg.Payload())
			switch msgText {
			case "anda":
				r := uint8(gobot.Rand(255))
				g := uint8(gobot.Rand(255))
				b := uint8(gobot.Rand(255))
				myBB8.SetRGB(r, g, b)
				myBB8.Roll(40, uint16(gobot.Rand(360)))

			case "ficavermelho":
				myBB8.SetRGB(255, 0, 0)

			case "ficafrio":
				myBB8.SetRGB(0, 0, 200)
			case "parar":
				myBB8.Stop()
			}

		})

	}

	robot := gobot.NewRobot("BB-8",
		[]gobot.Connection{bleAdaptor, mqttAdaptor},
		[]gobot.Device{myBB8},
		work,
	)

	robot.Start()
	bleAdaptor.Finalize()
}
