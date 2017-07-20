package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	fmt.Println("Iniciando a configuração do Raspberry...")
	r := raspi.NewAdaptor()
	motor := gpio.NewRelayDriver(r, "11")
	fmt.Println("Vamos começar...")
	time.Sleep(5 * time.Second)
	err := motor.On()
	if err != nil {
		fmt.Println("[main] Erro ao ligar o relay: ", err.Error())
		return
	}
	fmt.Println("Motor funcionando...")
	time.Sleep(3 * time.Second)
	err = motor.Off()
	if err != nil {
		fmt.Println("[main] Erro ao desligar o relay: ", err.Error())
		return
	}
	fmt.Println("Relay desligado.")
	r.Finalize()
	fmt.Println("GPIO liberada.")
	return
}
