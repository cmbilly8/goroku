package goroku

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
)

type Client struct {
	url        string
	httpClient *http.Client
	reader     *bytes.Reader
}

const (
	CmdTogglePower = "keypress/power"
	CmdVolumeUp    = "keypress/VolumeUp"
	CmdVolumeDown  = "keypress/VolumeDown"
	CmdVolumeMute  = "keypress/VolumeMute"
	CmdHome        = "keypress/Home"
	CmdReverse     = "keypress/Rev"
	CmdForward     = "keypress/Fwd"
	CmdPlay        = "keypress/Play"
	CmdSelect      = "keypress/Select"
	CmdLeft        = "keypress/Left"
	CmdRight       = "keypress/Right"
	CmdDown        = "keypress/Down"
	CmdUp          = "keypress/Up"
	CmdBack        = "keypress/Back"
	CmdSkipBack    = "keypress/InstantReplay"
	CmdInfo        = "keypress/Info"
	CmdBackspace   = "keypress/Backspace"
	CmdSearch      = "keypress/Search"
	CmdEnter       = "keypress/Enter"
)

func NewClientByIP(device_ip net.IP) *Client {
	url := fmt.Sprintf("http://%v:8060", device_ip.String())
	reader := getEmptyReader()
	return &Client{
		url:        url,
		httpClient: &http.Client{},
		reader:     reader,
	}
}

func (c *Client) GetUrl() string {
	return c.url
}

func getEmptyReader() *bytes.Reader {
	jsonBody := []byte("")
	return bytes.NewReader(jsonBody)
}

func (c *Client) SendCommand(Cmd string) error {
	requestURL := fmt.Sprintf("%v/%v", c.url, Cmd)
	fmt.Printf("Sending command: %v", requestURL)
	req, _ := http.NewRequest(http.MethodPost, requestURL, c.reader)
	_, err := c.httpClient.Do(req)
	return err
}
