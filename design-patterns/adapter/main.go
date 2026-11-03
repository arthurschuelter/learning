package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningInto(com Computer) {
	fmt.Println("Client inserts lightning port into computer")
	com.InsertLightning()
}

func main() {
	fmt.Println("Adapter")

	client := &Client{}
	mac := &Mac{}
	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowsMachine: windows,
	}

	client.InsertLightningInto(mac)
	client.InsertLightningInto(windowsAdapter)
}
