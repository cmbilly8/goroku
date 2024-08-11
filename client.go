package goroku

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/koron/go-ssdp"
	"io"
	"net"
	"net/http"
)

const (
	CmdTogglePower   = "keypress/power"
	CmdVolumeUp      = "keypress/VolumeUp"
	CmdVolumeDown    = "keypress/VolumeDown"
	CmdVolumeMute    = "keypress/VolumeMute"
	CmdHome          = "keypress/Home"
	CmdReverse       = "keypress/Rev"
	CmdForward       = "keypress/Fwd"
	CmdPlay          = "keypress/Play"
	CmdSelect        = "keypress/Select"
	CmdLeft          = "keypress/Left"
	CmdRight         = "keypress/Right"
	CmdDown          = "keypress/Down"
	CmdUp            = "keypress/Up"
	CmdBack          = "keypress/Back"
	CmdSkipBack      = "keypress/InstantReplay"
	CmdInfo          = "keypress/Info"
	CmdBackspace     = "keypress/Backspace"
	CmdSearch        = "keypress/Search"
	CmdEnter         = "keypress/Enter"
	queryDeviceInfo  = "query/device-info"
	queryMediaPlayer = "query/media-player"
)

type Client struct {
	url        string
	httpClient *http.Client
	reader     *bytes.Reader
}

func NewClientByIP(device_ip net.IP) *Client {
	url := fmt.Sprintf("http://%v:8060", device_ip.String())
	reader := makeEmptyReader()
	return &Client{
		url:        url,
		httpClient: &http.Client{},
		reader:     reader,
	}
}

func (c *Client) Url() string {
	return c.url
}

func makeEmptyReader() *bytes.Reader {
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

func (c *Client) DeviceInfo() (*DeviceInfo, error) {
	var di DeviceInfo
	err := c.populateStructFromGetRequest(queryDeviceInfo, &di)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &di, nil
}

func (c *Client) MediaPlayer() (*MediaPlayer, error) {
	var mp MediaPlayer
	err := c.populateStructFromGetRequest(queryMediaPlayer, &mp)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &mp, nil
}

func (c *Client) populateStructFromGetRequest(endpoint string, v interface{}) error {
	requestURL := fmt.Sprintf("%v/%v", c.url, endpoint)
	fmt.Printf("Sending command: %v\n", requestURL)

	request, err := http.NewRequest(http.MethodGet, requestURL, c.reader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println("Problem sending request:", err)
		return err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not read the response body:", err)
		return err
	}

	err = xml.Unmarshal(responseBody, v)
	if err != nil {
		fmt.Println("Problem unmarshalling response:", err)
		return err
	}

	return nil
}

func Discover() (ssdp.Service, error) {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		addrs, _ := i.Addrs()
		fmt.Printf("Interface: %v, Addresses: %v\n", i.Name, addrs)
	}

	services, err := ssdp.Search("roku:ecp", 20, "239.255.255.250:1900")
	if err != nil {
		fmt.Println("Error during SSDP search:", err)
		return ssdp.Service{}, err
	}

	if len(services) == 0 {
		fmt.Println("No Roku services found on the network.")
		return ssdp.Service{}, fmt.Errorf("no services found")
	}

	fmt.Printf("Found %d Roku services\n", len(services))
	return services[0], nil
}
