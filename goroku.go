package goroku

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
)

type Client struct {
	url string
}

const (
	cmdTogglePower = "keypress/power"
	cmdVolumeUp    = "keypress/VolumeUp"
)

func NewClientByIP(device_ip net.IP) *Client {
	url := fmt.Sprintf("http://%v:8060", device_ip.String())
	fmt.Println(url)
	return &Client{
		url: url,
	}
}

func (c *Client) TogglePower() error {
	err := sendCommand(cmdTogglePower, c.url)
	return err
}

func (c *Client) VolumeUp() error {
	err := sendCommand(cmdVolumeUp, c.url)
	return err
}

func getEmptyReader() *bytes.Reader {
	jsonBody := []byte("")
	return bytes.NewReader(jsonBody)
}

func sendCommand(cmd string, url string) error {
	requestURL := fmt.Sprintf("%v/%v", url, cmd)
	_, err := http.NewRequest(http.MethodPost, requestURL, getEmptyReader())
	return err
}
