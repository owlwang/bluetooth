package main

import (
	"time"

	"github.com/aykevl/go-bluetooth"
)

// flags + local name
var advPayload = []byte("\x02\x01\x06" + "\x07\x09TinyGo")

func main() {
	println("starting")
	adapter := bluetooth.DefaultAdapter
	must("enable SD", adapter.Enable())
	adv := adapter.NewAdvertisement()
	options := &bluetooth.AdvertiseOptions{
		Interval: bluetooth.NewAdvertiseInterval(100),
	}
	must("config adv", adv.Configure(advPayload, nil, options))
	must("start adv", adv.Start())

	var heartRateMeasurement bluetooth.Characteristic
	must("add service", adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.New16BitUUID(0x180D), // Heart Rate
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &heartRateMeasurement,
				UUID:   bluetooth.New16BitUUID(0x2A37), // Heart Rate Measurement
				Value:  []byte{0, 75},                  // 75bpm
			},
		},
	}))

	println("sleeping")
	for {
		// Sleep forever.
		time.Sleep(time.Hour)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
