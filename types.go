package goroku

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

type MediaPlayer struct {
	State    string `xml:"state,attr"`
	Error    string `xml:"error,attr"`
	Plugin   Plugin `xml:"plugin"`
	Format   Format `xml:"format"`
	Position string `xml:"position"`
}

type Plugin struct {
	ID        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Bandwidth string `xml:"bandwidth,attr"`
}

type Format struct {
	Audio    string `xml:"audio,attr"`
	Video    string `xml:"video,attr"`
	Captions string `xml:"captions,attr"`
	DRM      string `xml:"drm,attr"`
}
