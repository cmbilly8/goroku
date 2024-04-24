# Goroku

This is a poorly thought out go client library to control roku TVs and devices<br>
Create a client:
`ip := net.IPv4(192, 168, 1, 1)`
`client := goroku.NewClientByIP(ip)`
<br>
Issue commands:
`err := client.SendCommand(goroku.CmdVolumeUp)`
`deviceInfo, err := client.GetDeviceInfo()`
