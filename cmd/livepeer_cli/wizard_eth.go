package main

import (
	"fmt"
	"net/url"
)

func (w *wizard) setMaxGasPrice() {
	fmt.Printf("Current maximum gas price: %v\n", w.maxGasPrice())
	fmt.Printf("Enter new maximum gas price in Wei (enter \"0\" for no maximum gas price)")
	amount := w.readBigInt()

	val := url.Values{
		"amount": {fmt.Sprintf("%v", amount.String())},
	}

	httpPostWithParams(fmt.Sprintf("http://%v:%v/setMaxGasPrice", w.host, w.httpPort), val)
}

func (w *wizard) signMessage() {
	fmt.Printf("Enter or paste the message to sign: \n")
	msg := w.readMultilineString()
	val := url.Values{
		"message": {msg},
	}
	result, ok := httpPostWithParams(fmt.Sprintf("http://%v:%v/signMessage", w.host, w.httpPort), val)
	if !ok {
		fmt.Printf("Error signing message: %v\n", result)
	}
	fmt.Println(fmt.Sprintf("\n\nSignature:\n0x%x", result))
}
