package config

import "net"

type NetworkConfig struct {
	Port     int
	IP       net.IP
	Protocol string
}

type SomeIPHeader struct {
	ServiceID        uint16
	MethodID         uint16
	Lenth            uint32
	ClientID         uint16
	SessionID        uint16
	ProtocolVersion  uint8
	InterfaceVersion uint8
	MessageType      uint8
	ReturnCode       uint8
}

// SOMEIP Message
type SOMEIPMessage struct {
	Header  SomeIPHeader
	Payload []byte
}

func NewNetworkConfig(port int, ip net.IP, protocol string) *NetworkConfig {
	return &NetworkConfig{
		Port:     30490,
		IP:       net.ParseIP("127.0.0.1"),
		Protocol: "udp",
	}
}
