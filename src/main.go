package main

import (
	"fmt"
	"golang-wasm-stun/packages/src/dom"
	"syscall/js"

	"github.com/pion/stun"
)


func clickTest(_ js.Value, _ []js.Value) interface{} {
	dom.SetValue("result", "innerText", "Clicked")
	return nil
}

func createStunRequest(_ js.Value, _ []js.Value) interface{} {
	// c, err := stun.Dial("udp", "stun.l.google.com:19302")
	c, err := stun.Dial("udp", "stun.webrtc.ecl.ntt.com:3478")
	if err != nil {
		dom.SetValue("result", "innerText", fmt.Sprintln("STUN dial faild:", err.Error()))
        dom.AddClass("result", "error")
	}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN server, waiting for response message.
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			dom.SetValue("result", "innerText", err.Error())
            dom.AddClass("result", "error")
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			dom.SetValue("result", "innerText", fmt.Sprintln("xorAddr FAILED!!:", err.Error()))
            dom.AddClass("result", "error")
		}
        dom.SetValue("result", "innerText", fmt.Sprintln("your IP is", xorAddr.IP))
	}); err != nil {
		dom.SetValue("result", "innerText", fmt.Sprintln("ERROR in c.Do()!!:", err.Error()))
        dom.AddClass("result", "error")
	}
    return nil
}

func main() {
	dom.Hide("loading")
	dom.AddEventListener("run-button", "click", createStunRequest)

	ch := make(chan struct{})
	<-ch
	fmt.Println("Program has started!")
}
