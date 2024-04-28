package goroku_test

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/cmbilly8/goroku"
)

// func TestClient_PowerOn(t *testing.T) {
// 	ip := net.IPv4(192, 168, 40, 100)
// 	client := goroku.NewClientByIP(ip)
// 	err := client.SendCommand(goroku.CmdTogglePower)
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	}
// }

func TestClient_DeviceInfo(t *testing.T) {
	fmt.Println("Start get device info test")
	ip := net.IPv4(192, 168, 40, 100)
	client := goroku.NewClientByIP(ip)
	deviceInfo, err := client.DeviceInfo()
	if err != nil {
		fmt.Print("error getting device info")
		log.Fatal()
	}
	printStructFields(deviceInfo)
	fmt.Println("End get device info test\n\t\n\t")
}

func TestClient_MediaPlayer(t *testing.T) {
	fmt.Println("Start get media player test")
	ip := net.IPv4(192, 168, 40, 100)
	client := goroku.NewClientByIP(ip)
	mediaPlayer, err := client.MediaPlayer()
	if err != nil {
		fmt.Print("error getting device info")
		log.Fatal()
	}
	printStructFields(mediaPlayer)
	fmt.Println("End get media player test\n\t\n\t")
}

func printStructFields(v interface{}) {
	fmt.Print("Response:\n")
	val := reflect.ValueOf(v).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()

		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
