# Goroku

Go client library for sending commands to roku devices and tvs<br>
Create a client:<br>
`ip := net.IPv4(192, 168, 1, 1)`<br>
`client := goroku.NewClientByIP(ip)`<br>
<br>
Issue commands:<br>
`err := client.SendCommand(goroku.CmdVolumeUp)`<br>
`deviceInfo, err := client.GetDeviceInfo()`<br>
