package main

import (
	"fmt"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func mqttSubscribe(timeout int64) {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%v:%v", env.MqttHost, env.MqttPort))
	opts.SetUsername(env.MqttUser)
	opts.SetPassword(env.MqttPassword)

	lastupdate := time.Time{}

	go func() {
		for {
			//если последнее время получения данных больше таймаута
			//повторная подписка на событие
			if time.Now().Sub(lastupdate).Milliseconds() > timeout {
				fmt.Println("New mqtt client")
				c := mqtt.NewClient(opts)
				if token := c.Connect(); token.Wait() && token.Error() != nil {
					lastupdate = time.Now()
					fmt.Println("Error get temperature")
				}

				msgRcvd := func(client mqtt.Client, message mqtt.Message) {
					t, _ := strconv.ParseFloat(string(message.Payload()), 32)
					temperature := float32(t)
					//fmt.Println("Temperature: " + fmt.Sprintf("%f", temperature))
					addTemp(temperature)

					lastupdate = time.Now()
				}

				if token := c.Subscribe("outdoor/temperature", 0, msgRcvd); token.Wait() && token.Error() != nil {
					fmt.Println(token.Error())
				}
			}
			<-time.After(time.Millisecond * time.Duration(timeout))
		}
	}()
}
