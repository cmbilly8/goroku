package goroku_test

import (
	"net"
	"testing"

	"github.com/cmbilly8/goroku"
)

func TestClient_VolumeUp(t *testing.T) {
	ip := net.IPv4(192, 168, 40, 100)
	client := goroku.NewClientByIP(ip)
	err := client.SendCommand(goroku.CmdVolumeUp)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
