package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	config := config()

	client := mqtt.NewClient(config)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	publish(client)

	client.Disconnect(250)
	fmt.Println("Disconnected")
}

func config() *mqtt.ClientOptions {
	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://localhost:1883")
	options.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected")
	}
	options.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Printf("Connection lost")
	}

	return options
}

func publish(client mqtt.Client) {
	num := 50
	for i := 0; i < num; i++ {
		value := fmt.Sprintf("%d", i)

		token := client.Publish("progress", 1, false, value)
		token.Wait()

		time.Sleep(time.Millisecond * 100)
	}
}
