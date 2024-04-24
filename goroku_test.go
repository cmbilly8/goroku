package goroku_test

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/cmbilly8/goroku"
)

// func TestClient_VolumeUp(t *testing.T) {
// 	ip := net.IPv4(192, 168, 40, 100)
// 	client := goroku.NewClientByIP(ip)
// 	err := client.SendCommand(goroku.CmdVolumeUp)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	}
// }

func TestClient_GetDeviceInfo(t *testing.T) {
	ip := net.IPv4(192, 168, 40, 100)
	client := goroku.NewClientByIP(ip)
	deviceInfo, err := client.GetDeviceInfo()
	if err != nil {
		fmt.Print("error getting device info")
		log.Fatal()
	}
	printStructFields(deviceInfo)
}

func printStructFields(di *goroku.DeviceInfo) {
	val := reflect.ValueOf(di).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()

		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
