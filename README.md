# Goroku

This is a poorly thought out go client library to control roku TVs and devices<br>
Create a client:
`ip := net.IPv4(192, 168, 1, 1)`<br>
`client := goroku.NewClientByIP(ip)`<br>
<br>
Issue commands:
`err := client.SendCommand(goroku.CmdVolumeUp)`<br>
`deviceInfo, err := client.GetDeviceInfo()`<br>
