package goroku

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
)

type Client struct {
	url        string
	httpClient *http.Client
	reader     *bytes.Reader
}

type DeviceInfo struct {
	Udn                         string  `xml:"udn"`
	SerialNumber                string  `xml:"serial-number"`
	DeviceID                    string  `xml:"device-id"`
	AdvertisingID               string  `xml:"advertising-id"`
	VendorName                  string  `xml:"vendor-name"`
	ModelName                   string  `xml:"model-name"`
	ModelNumber                 string  `xml:"model-number"`
	ModelRegion                 string  `xml:"model-region"`
	IsTV                        bool    `xml:"is-tv"`
	IsStick                     bool    `xml:"is-stick"`
	ScreenSize                  float64 `xml:"screen-size"`
	PanelID                     int     `xml:"panel-id"`
	MobileHasLiveTV             bool    `xml:"mobile-has-live-tv"`
	UIResolution                string  `xml:"ui-resolution"`
	TunerType                   string  `xml:"tuner-type"`
	SupportsEthernet            bool    `xml:"supports-ethernet"`
	WifiMAC                     string  `xml:"wifi-mac"`
	WifiDriver                  string  `xml:"wifi-driver"`
	HasWifi5GSupport            bool    `xml:"has-wifi-5G-support"`
	EthernetMAC                 string  `xml:"ethernet-mac"`
	NetworkType                 string  `xml:"network-type"`
	NetworkName                 string  `xml:"network-name"`
	FriendlyDeviceName          string  `xml:"friendly-device-name"`
	FriendlyModelName           string  `xml:"friendly-model-name"`
	DefaultDeviceName           string  `xml:"default-device-name"`
	UserDeviceName              string  `xml:"user-device-name"`
	UserDeviceLocation          string  `xml:"user-device-location"`
	BuildNumber                 string  `xml:"build-number"`
	SoftwareVersion             string  `xml:"software-version"`
	SoftwareBuild               int     `xml:"software-build"`
	LightningBaseBuildNumber    string  `xml:"lightning-base-build-number"`
	UIBuildNumber               string  `xml:"ui-build-number"`
	UISoftwareVersion           string  `xml:"ui-software-version"`
	UISoftwareBuild             int     `xml:"ui-software-build"`
	SecureDevice                bool    `xml:"secure-device"`
	Language                    string  `xml:"language"`
	Country                     string  `xml:"country"`
	Locale                      string  `xml:"locale"`
	TimeZoneAuto                bool    `xml:"time-zone-auto"`
	TimeZone                    string  `xml:"time-zone"`
	TimeZoneName                string  `xml:"time-zone-name"`
	TimeZoneTZ                  string  `xml:"time-zone-tz"`
	TimeZoneOffset              int     `xml:"time-zone-offset"`
	ClockFormat                 string  `xml:"clock-format"`
	Uptime                      int     `xml:"uptime"`
	PowerMode                   string  `xml:"power-mode"`
	SupportsSuspend             bool    `xml:"supports-suspend"`
	SupportsFindRemote          bool    `xml:"supports-find-remote"`
	SupportsAudioGuide          bool    `xml:"supports-audio-guide"`
	SupportsRVA                 bool    `xml:"supports-rva"`
	HasHandsFreeVoiceRemote     bool    `xml:"has-hands-free-voice-remote"`
	DeveloperEnabled            bool    `xml:"developer-enabled"`
	KeyedDeveloperID            string  `xml:"keyed-developer-id"`
	SearchEnabled               bool    `xml:"search-enabled"`
	SearchChannelsEnabled       bool    `xml:"search-channels-enabled"`
	VoiceSearchEnabled          bool    `xml:"voice-search-enabled"`
	SupportsPrivateListening    bool    `xml:"supports-private-listening"`
	SupportsPrivateListeningDTV bool    `xml:"supports-private-listening-dtv"`
	SupportsWarmStandby         bool    `xml:"supports-warm-standby"`
	HeadphonesConnected         bool    `xml:"headphones-connected"`
	SupportsAudioSettings       bool    `xml:"supports-audio-settings"`
	ExpertPQEnabled             float64 `xml:"expert-pq-enabled"`
	SupportsECSTextEdit         bool    `xml:"supports-ecs-textedit"`
	SupportsECSMicrophone       bool    `xml:"supports-ecs-microphone"`
	SupportsWakeOnWlan          bool    `xml:"supports-wake-on-wlan"`
	SupportsAirplay             bool    `xml:"supports-airplay"`
	HasPlayOnRoku               bool    `xml:"has-play-on-roku"`
	HasMobileScreensaver        bool    `xml:"has-mobile-screensaver"`
	SupportURL                  string  `xml:"support-url"`
	GrandCentralVersion         string  `xml:"grandcentral-version"`
	SupportsTRC                 bool    `xml:"supports-trc"`
	TRCVersion                  string  `xml:"trc-version"`
	TRCChannelVersion           string  `xml:"trc-channel-version"`
	AVSyncCalibrationEnabled    float64 `xml:"av-sync-calibration-enabled"`
	RoomCalibrationVersion      string  `xml:"room-calibration-version"`
}

const (
	CmdTogglePower = "keypress/power "
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

func (c *Client) GetDeviceInfo() (*DeviceInfo, error) {
	requestURL := fmt.Sprintf("%v/query/device-info", c.url)
	fmt.Printf("Sending command: %v\n", requestURL)

	request, err := http.NewRequest(http.MethodGet, requestURL, c.reader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println("Problem sending request to Roku Device:", err)
		return nil, err
	}
	defer func() {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not read the response body:", err)
		return nil, err
	}

	var di DeviceInfo
	err = xml.Unmarshal(responseBody, &di)
	if err != nil {
		fmt.Println("Problem unmarshalling response into DeviceInfo structure:", err)
		return nil, err
	}

	return &di, nil
}
