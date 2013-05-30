// Copyright 2012, Google, Inc. All rights reserved.
// Copyright 2009-2011 Andreas Krennmair. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"code.google.com/p/gopacket"
	"fmt"
	"net"
	"reflect"
	"strings"
	"testing"
)

var testSimpleTCPPacket []byte = []byte{
	0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
	0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, 0x06,
	0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5, 0xf7,
	0x00, 0x50, 0xc5, 0x7e, 0x0e, 0x48, 0x49, 0x07, 0x42, 0x32, 0x80, 0x18,
	0x00, 0x73, 0xab, 0xb1, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x03, 0x77,
	0x37, 0x9c, 0x42, 0x77, 0x5e, 0x3a, 0x47, 0x45, 0x54, 0x20, 0x2f, 0x20,
	0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31, 0x0d, 0x0a, 0x48, 0x6f,
	0x73, 0x74, 0x3a, 0x20, 0x77, 0x77, 0x77, 0x2e, 0x66, 0x69, 0x73, 0x68,
	0x2e, 0x63, 0x6f, 0x6d, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x6b, 0x65, 0x65, 0x70, 0x2d, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x0d, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4d, 0x6f, 0x7a, 0x69, 0x6c, 0x6c,
	0x61, 0x2f, 0x35, 0x2e, 0x30, 0x20, 0x28, 0x58, 0x31, 0x31, 0x3b, 0x20,
	0x4c, 0x69, 0x6e, 0x75, 0x78, 0x20, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34,
	0x29, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x57, 0x65, 0x62, 0x4b, 0x69,
	0x74, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32, 0x20, 0x28, 0x4b, 0x48, 0x54,
	0x4d, 0x4c, 0x2c, 0x20, 0x6c, 0x69, 0x6b, 0x65, 0x20, 0x47, 0x65, 0x63,
	0x6b, 0x6f, 0x29, 0x20, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2f, 0x31,
	0x35, 0x2e, 0x30, 0x2e, 0x38, 0x37, 0x34, 0x2e, 0x31, 0x32, 0x31, 0x20,
	0x53, 0x61, 0x66, 0x61, 0x72, 0x69, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32,
	0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x3a, 0x20, 0x74, 0x65,
	0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x68, 0x74, 0x6d,
	0x6c, 0x2b, 0x78, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x6d, 0x6c, 0x3b, 0x71, 0x3d,
	0x30, 0x2e, 0x39, 0x2c, 0x2a, 0x2f, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e,
	0x38, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x67, 0x7a, 0x69, 0x70,
	0x2c, 0x64, 0x65, 0x66, 0x6c, 0x61, 0x74, 0x65, 0x2c, 0x73, 0x64, 0x63,
	0x68, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x65, 0x6e, 0x2d, 0x55,
	0x53, 0x2c, 0x65, 0x6e, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x38, 0x0d, 0x0a,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x43, 0x68, 0x61, 0x72, 0x73,
	0x65, 0x74, 0x3a, 0x20, 0x49, 0x53, 0x4f, 0x2d, 0x38, 0x38, 0x35, 0x39,
	0x2d, 0x31, 0x2c, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x3b, 0x71, 0x3d, 0x30,
	0x2e, 0x37, 0x2c, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x33, 0x0d, 0x0a,
	0x0d, 0x0a,
}

// Benchmarks for actual gopacket code

func BenchmarkLayerClassSliceContains(b *testing.B) {
	lc := gopacket.NewLayerClassSlice([]gopacket.LayerType{LayerTypeTCP, LayerTypeEthernet})
	for i := 0; i < b.N; i++ {
		_ = lc.Contains(LayerTypeTCP)
	}
}

func BenchmarkLayerClassMapContains(b *testing.B) {
	lc := gopacket.NewLayerClassMap([]gopacket.LayerType{LayerTypeTCP, LayerTypeEthernet})
	for i := 0; i < b.N; i++ {
		_ = lc.Contains(LayerTypeTCP)
	}
}

func BenchmarkLazyNoCopyEthLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true}).Layer(LayerTypeEthernet)
	}
}

func BenchmarkLazyNoCopyIPLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true}).Layer(LayerTypeIPv4)
	}
}

func BenchmarkLazyNoCopyTCPLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true}).Layer(LayerTypeTCP)
	}
}

func BenchmarkLazyNoCopyAllLayers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true}).Layers()
	}
}

func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	}
}

func BenchmarkLazy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Lazy)
	}
}

func BenchmarkNoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.NoCopy)
	}
}

func BenchmarkLazyNoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
	}
}

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &TCP{}
	}
}

func BenchmarkFlow(b *testing.B) {
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
	net := p.NetworkLayer()
	for i := 0; i < b.N; i++ {
		net.NetworkFlow()
	}
}

func BenchmarkEndpoints(b *testing.B) {
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
	flow := p.NetworkLayer().NetworkFlow()
	for i := 0; i < b.N; i++ {
		flow.Endpoints()
	}
}

func BenchmarkTCPLayerFromDecodedPacket(b *testing.B) {
	b.StopTimer()
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Layer(LayerTypeTCP)
	}
}

func BenchmarkTCPLayerClassFromDecodedPacket(b *testing.B) {
	b.StopTimer()
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	lc := gopacket.NewLayerClass([]gopacket.LayerType{LayerTypeTCP})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = p.LayerClass(lc)
	}
}

func BenchmarkTCPTransportLayerFromDecodedPacket(b *testing.B) {
	b.StopTimer()
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = p.TransportLayer()
	}
}

func testDecoder([]byte, gopacket.PacketBuilder) error {
	return nil
}

func BenchmarkDecodeFuncCallOverheadDirectCall(b *testing.B) {
	var data []byte
	var pb gopacket.PacketBuilder
	for i := 0; i < b.N; i++ {
		_ = testDecoder(data, pb)
	}
}

func BenchmarkDecodeFuncCallOverheadDecoderCall(b *testing.B) {
	d := gopacket.DecodeFunc(testDecoder)
	var data []byte
	var pb gopacket.PacketBuilder
	for i := 0; i < b.N; i++ {
		_ = d.Decode(data, pb)
	}
}

func BenchmarkDecodeFuncCallOverheadArrayCall(b *testing.B) {
	EthernetTypeMetadata[1] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(testDecoder)}
	d := EthernetType(1)
	var data []byte
	var pb gopacket.PacketBuilder
	for i := 0; i < b.N; i++ {
		_ = d.Decode(data, pb)
	}
}

func BenchmarkFmtVerboseString(b *testing.B) {
	b.StopTimer()
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%#v", p)
	}
}

func BenchmarkPacketString(b *testing.B) {
	b.StopTimer()
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.String()
	}
}

func BenchmarkPacketDumpString(b *testing.B) {
	b.StopTimer()
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.Default)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.String()
	}
}

// TestFlowMapKey makes sure a flow and an endpoint can be used as map keys.
func TestFlowMapKey(t *testing.T) {
	_ = map[gopacket.Flow]bool{}
	_ = map[gopacket.Endpoint]bool{}
	_ = map[[2]gopacket.Flow]bool{}
}

func TestDecodeSimpleTCPPacket(t *testing.T) {
	equal := func(desc, want string, got fmt.Stringer) {
		if want != got.String() {
			t.Errorf("%s: got %q want %q", desc, got.String(), want)
		}
	}
	p := gopacket.NewPacket(testSimpleTCPPacket, LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
	if eth := p.LinkLayer(); eth == nil {
		t.Error("No ethernet layer found")
	} else {
		equal("Eth Src", "bc:30:5b:e8:d3:49", eth.LinkFlow().Src())
		equal("Eth Dst", "00:00:0c:9f:f0:20", eth.LinkFlow().Dst())
	}
	if net := p.NetworkLayer(); net == nil {
		t.Error("No net layer found")
	} else if ip, ok := net.(*IPv4); !ok {
		t.Error("Net layer is not IP layer")
	} else {
		equal("IP Src", "172.17.81.73", net.NetworkFlow().Src())
		equal("IP Dst", "173.222.254.225", net.NetworkFlow().Dst())
		want := &IPv4{
			BaseLayer:  BaseLayer{testSimpleTCPPacket[14:34], testSimpleTCPPacket[34:]},
			Version:    4,
			IHL:        5,
			TOS:        0,
			Length:     420,
			Id:         14815,
			Flags:      0x02,
			FragOffset: 0,
			TTL:        64,
			Protocol:   6,
			Checksum:   0x555A,
			SrcIP:      []byte{172, 17, 81, 73},
			DstIP:      []byte{173, 222, 254, 225},
		}
		if !reflect.DeepEqual(ip, want) {
			t.Errorf("IP layer mismatch, \ngot  %#v\nwant %#v\n", ip, want)
		}
	}
	if trans := p.TransportLayer(); trans == nil {
		t.Error("No transport layer found")
	} else if tcp, ok := trans.(*TCP); !ok {
		t.Error("Transport layer is not TCP layer")
	} else {
		equal("TCP Src", "50679", trans.TransportFlow().Src())
		equal("TCP Dst", "80", trans.TransportFlow().Dst())
		want := &TCP{
			BaseLayer:  BaseLayer{testSimpleTCPPacket[34:66], testSimpleTCPPacket[66:]},
			SrcPort:    50679,
			DstPort:    80,
			Seq:        0xc57e0e48,
			Ack:        0x49074232,
			DataOffset: 8,
			ACK:        true,
			PSH:        true,
			Window:     0x73,
			Checksum:   0xabb1,
			Urgent:     0,
			sPort:      []byte{0xc5, 0xf7},
			dPort:      []byte{0x0, 0x50},
			Options: []TCPOption{
				TCPOption{
					OptionType:   0x1,
					OptionLength: 0x1,
				},
				TCPOption{
					OptionType:   0x1,
					OptionLength: 0x1,
				},
				TCPOption{
					OptionType:   0x8,
					OptionLength: 0xa,
					OptionData:   []byte{0x3, 0x77, 0x37, 0x9c, 0x42, 0x77, 0x5e, 0x3a},
				},
			},
		}
		if !reflect.DeepEqual(tcp, want) {
			t.Errorf("TCP layer mismatch\ngot  %#v\nwant %#v", tcp, want)
		}
	}
	if payload, ok := p.Layer(gopacket.LayerTypePayload).(*gopacket.Payload); payload == nil || !ok {
		t.Error("No payload layer found")
	} else {
		if string(payload.Payload()) != "GET / HTTP/1.1\r\nHost: www.fish.com\r\nConnection: keep-alive\r\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.121 Safari/535.2\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip,deflate,sdch\r\nAccept-Language: en-US,en;q=0.8\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.3\r\n\r\n" {
			t.Error("--- Payload STRING ---\n", string(payload.Payload()), "\n--- Payload BYTES ---\n", payload.Payload())
		}
	}
}

// Makes sure packet payload doesn't display the 6 trailing null of this packet
// as part of the payload.  They're actually the ethernet trailer.
func TestDecodeSmallTCPPacketHasEmptyPayload(t *testing.T) {
	p := gopacket.NewPacket(
		[]byte{
			0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49, 0xb8, 0xac, 0x6f, 0x92, 0xd5, 0xbf,
			0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
			0x3f, 0x9f, 0xac, 0x11, 0x51, 0xc5, 0xac, 0x11, 0x51, 0x49, 0x00, 0x63,
			0x9a, 0xef, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xc1, 0x27, 0x83, 0x50, 0x14,
			0x00, 0x00, 0xc3, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, LinkTypeEthernet, gopacket.Default)

	if payload := p.Layer(gopacket.LayerTypePayload); payload != nil {
		t.Error("Payload found for empty TCP packet")
	}
}

func TestDecodeVLANPacket(t *testing.T) {
	p := gopacket.NewPacket(
		[]byte{
			0x00, 0x10, 0xdb, 0xff, 0x10, 0x00, 0x00, 0x15, 0x2c, 0x9d, 0xcc, 0x00,
			0x81, 0x00, 0x01, 0xf7, 0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x29, 0x8d,
			0x40, 0x00, 0x7d, 0x06, 0x83, 0xa0, 0xac, 0x1b, 0xca, 0x8e, 0x45, 0x16,
			0x94, 0xe2, 0xd4, 0x0a, 0x00, 0x50, 0xdf, 0xab, 0x9c, 0xc6, 0xcd, 0x1e,
			0xe5, 0xd1, 0x50, 0x10, 0x01, 0x00, 0x5a, 0x74, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}, LinkTypeEthernet, gopacket.Default)
	if err := p.ErrorLayer(); err != nil {
		t.Error("Error while parsing vlan packet:", err)
	}
	if vlan := p.Layer(LayerTypeDot1Q); vlan == nil {
		t.Error("Didn't detect vlan")
	} else if _, ok := vlan.(*Dot1Q); !ok {
		t.Error("LayerTypeDot1Q layer is not a Dot1Q object")
	}
	for i, l := range p.Layers() {
		t.Logf("Layer %d: %#v", i, l)
	}
	want := []gopacket.LayerType{LayerTypeEthernet, LayerTypeDot1Q, LayerTypeIPv4, LayerTypeTCP}
	checkLayers(p, want, t)
}

func TestDecodeSCTPPackets(t *testing.T) {
	sctpPackets := [][]byte{
		[]byte{ // INIT
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x1f, 0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x08, 0x00, 0x45, 0x02,
			0x00, 0x44, 0x00, 0x00, 0x40, 0x00, 0x40, 0x84, 0xc4, 0x22, 0xac, 0x1d, 0x14, 0x0f, 0xac, 0x19,
			0x09, 0xcc, 0x27, 0x0f, 0x22, 0xb8, 0x00, 0x00, 0x00, 0x00, 0x19, 0x6b, 0x0b, 0x40, 0x01, 0x00,
			0x00, 0x24, 0xb6, 0x96, 0xb0, 0x9e, 0x00, 0x01, 0xc0, 0x00, 0x00, 0x0a, 0xff, 0xff, 0xdb, 0x85,
			0x60, 0x23, 0x00, 0x0c, 0x00, 0x06, 0x00, 0x05, 0x00, 0x00, 0x80, 0x00, 0x00, 0x04, 0xc0, 0x00,
			0x00, 0x04,
		}, []byte{ // INIT ACK
			0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x00, 0x1f, 0xca, 0xb3, 0x76, 0x40, 0x08, 0x00, 0x45, 0x20,
			0x01, 0x24, 0x00, 0x00, 0x40, 0x00, 0x36, 0x84, 0xcd, 0x24, 0xac, 0x19, 0x09, 0xcc, 0xac, 0x1d,
			0x14, 0x0f, 0x22, 0xb8, 0x27, 0x0f, 0xb6, 0x96, 0xb0, 0x9e, 0x4b, 0xab, 0x40, 0x9a, 0x02, 0x00,
			0x01, 0x04, 0x32, 0x80, 0xfb, 0x42, 0x00, 0x00, 0xf4, 0x00, 0x00, 0x0a, 0x00, 0x0a, 0x85, 0x98,
			0xb1, 0x26, 0x00, 0x07, 0x00, 0xe8, 0xd3, 0x08, 0xce, 0xe2, 0x52, 0x95, 0xcc, 0x09, 0xa1, 0x4c,
			0x6f, 0xa7, 0x9e, 0xba, 0x03, 0xa1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0xfb, 0x80, 0x32, 0x9e, 0xb0,
			0x96, 0xb6, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x46, 0xc2, 0x50, 0x00, 0x00,
			0x00, 0x00, 0x5e, 0x25, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x0a, 0x00, 0x26, 0xb1,
			0x98, 0x85, 0x02, 0x00, 0x27, 0x0f, 0xac, 0x1d, 0x14, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x22,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0x6a, 0x72, 0x5c, 0x1c, 0x3c, 0xaa,
			0x7a, 0xcd, 0xd3, 0x8f, 0x52, 0x78, 0x7c, 0x77, 0xfd, 0x46, 0xbd, 0x72, 0x82, 0xc1, 0x1f, 0x70,
			0x44, 0xcc, 0xc7, 0x9b, 0x9b, 0x7b, 0x13, 0x54, 0x3f, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x24, 0xb6, 0x96,
			0xb0, 0x9e, 0x00, 0x01, 0xc0, 0x00, 0x00, 0x0a, 0xff, 0xff, 0xdb, 0x85, 0x60, 0x23, 0x00, 0x0c,
			0x00, 0x06, 0x00, 0x05, 0x00, 0x00, 0x80, 0x00, 0x00, 0x04, 0xc0, 0x00, 0x00, 0x04, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x04, 0xc0, 0x00,
			0x00, 0x04,
		}, []byte{ // COOKIE ECHO, DATA
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x1f, 0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x08, 0x00, 0x45, 0x02,
			0x01, 0x20, 0x00, 0x00, 0x40, 0x00, 0x40, 0x84, 0xc3, 0x46, 0xac, 0x1d, 0x14, 0x0f, 0xac, 0x19,
			0x09, 0xcc, 0x27, 0x0f, 0x22, 0xb8, 0x32, 0x80, 0xfb, 0x42, 0x01, 0xf9, 0xf3, 0xa9, 0x0a, 0x00,
			0x00, 0xe8, 0xd3, 0x08, 0xce, 0xe2, 0x52, 0x95, 0xcc, 0x09, 0xa1, 0x4c, 0x6f, 0xa7, 0x9e, 0xba,
			0x03, 0xa1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x42, 0xfb, 0x80, 0x32, 0x9e, 0xb0, 0x96, 0xb6, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x46, 0xc2, 0x50, 0x00, 0x00, 0x00, 0x00, 0x5e, 0x25,
			0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x0a, 0x00, 0x26, 0xb1, 0x98, 0x85, 0x02, 0x00,
			0x27, 0x0f, 0xac, 0x1d, 0x14, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x22, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0x6a, 0x72, 0x5c, 0x1c, 0x3c, 0xaa, 0x7a, 0xcd, 0xd3, 0x8f,
			0x52, 0x78, 0x7c, 0x77, 0xfd, 0x46, 0xbd, 0x72, 0x82, 0xc1, 0x1f, 0x70, 0x44, 0xcc, 0xc7, 0x9b,
			0x9b, 0x7b, 0x13, 0x54, 0x3f, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x24, 0xb6, 0x96, 0xb0, 0x9e, 0x00, 0x01,
			0xc0, 0x00, 0x00, 0x0a, 0xff, 0xff, 0xdb, 0x85, 0x60, 0x23, 0x00, 0x0c, 0x00, 0x06, 0x00, 0x05,
			0x00, 0x00, 0x80, 0x00, 0x00, 0x04, 0xc0, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x16, 0xdb, 0x85, 0x60, 0x23, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x6f, 0x6f, 0x21, 0x0a, 0x00, 0x00, 0x00,
		}, []byte{ // COOKIE ACK, SACK
			0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x00, 0x1f, 0xca, 0xb3, 0x76, 0x40, 0x08, 0x00, 0x45, 0x20,
			0x00, 0x34, 0x00, 0x00, 0x40, 0x00, 0x36, 0x84, 0xce, 0x14, 0xac, 0x19, 0x09, 0xcc, 0xac, 0x1d,
			0x14, 0x0f, 0x22, 0xb8, 0x27, 0x0f, 0xb6, 0x96, 0xb0, 0x9e, 0xed, 0x64, 0x30, 0x98, 0x0b, 0x00,
			0x00, 0x04, 0x03, 0x00, 0x00, 0x10, 0xdb, 0x85, 0x60, 0x23, 0x00, 0x00, 0xf3, 0xfa, 0x00, 0x00,
			0x00, 0x00,
		}, []byte{ // DATA
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x1f, 0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x08, 0x00, 0x45, 0x02,
			0x00, 0x3c, 0x00, 0x00, 0x40, 0x00, 0x40, 0x84, 0xc4, 0x2a, 0xac, 0x1d, 0x14, 0x0f, 0xac, 0x19,
			0x09, 0xcc, 0x27, 0x0f, 0x22, 0xb8, 0x32, 0x80, 0xfb, 0x42, 0xa1, 0xe3, 0xb2, 0x31, 0x00, 0x03,
			0x00, 0x19, 0xdb, 0x85, 0x60, 0x24, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x62, 0x69,
			0x7a, 0x7a, 0x6c, 0x65, 0x21, 0x0a, 0x00, 0x00, 0x00, 0x00,
		}, []byte{ // SACK
			0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x00, 0x1f, 0xca, 0xb3, 0x76, 0x40, 0x08, 0x00, 0x45, 0x20,
			0x00, 0x30, 0x00, 0x00, 0x40, 0x00, 0x36, 0x84, 0xce, 0x18, 0xac, 0x19, 0x09, 0xcc, 0xac, 0x1d,
			0x14, 0x0f, 0x22, 0xb8, 0x27, 0x0f, 0xb6, 0x96, 0xb0, 0x9e, 0xfa, 0x49, 0x94, 0x3a, 0x03, 0x00,
			0x00, 0x10, 0xdb, 0x85, 0x60, 0x24, 0x00, 0x00, 0xf4, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, []byte{ // SHUTDOWN
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x1f, 0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x08, 0x00, 0x45, 0x02,
			0x00, 0x28, 0x00, 0x00, 0x40, 0x00, 0x40, 0x84, 0xc4, 0x3e, 0xac, 0x1d, 0x14, 0x0f, 0xac, 0x19,
			0x09, 0xcc, 0x27, 0x0f, 0x22, 0xb8, 0x32, 0x80, 0xfb, 0x42, 0x3f, 0x29, 0x59, 0x23, 0x07, 0x00,
			0x00, 0x08, 0x85, 0x98, 0xb1, 0x25,
		}, []byte{ // SHUTDOWN ACK
			0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x00, 0x1f, 0xca, 0xb3, 0x76, 0x40, 0x08, 0x00, 0x45, 0x20,
			0x00, 0x24, 0x00, 0x00, 0x40, 0x00, 0x36, 0x84, 0xce, 0x24, 0xac, 0x19, 0x09, 0xcc, 0xac, 0x1d,
			0x14, 0x0f, 0x22, 0xb8, 0x27, 0x0f, 0xb6, 0x96, 0xb0, 0x9e, 0xb2, 0xc8, 0x99, 0x24, 0x08, 0x00,
			0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, []byte{ // SHUTDOWN COMPLETE
			0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x1f, 0x24, 0xbe, 0x05, 0x27, 0x0b, 0x17, 0x08, 0x00, 0x45, 0x02,
			0x00, 0x24, 0x00, 0x00, 0x40, 0x00, 0x40, 0x84, 0xc4, 0x42, 0xac, 0x1d, 0x14, 0x0f, 0xac, 0x19,
			0x09, 0xcc, 0x27, 0x0f, 0x22, 0xb8, 0x32, 0x80, 0xfb, 0x42, 0xa8, 0xd1, 0x86, 0x85, 0x0e, 0x00,
			0x00, 0x04,
		}}
	wantLayers := [][]gopacket.LayerType{
		[]gopacket.LayerType{LayerTypeSCTPInit},
		[]gopacket.LayerType{LayerTypeSCTPInitAck},
		[]gopacket.LayerType{LayerTypeSCTPCookieEcho, LayerTypeSCTPData},
		[]gopacket.LayerType{LayerTypeSCTPCookieAck, LayerTypeSCTPSack},
		[]gopacket.LayerType{LayerTypeSCTPData},
		[]gopacket.LayerType{LayerTypeSCTPSack},
		[]gopacket.LayerType{LayerTypeSCTPShutdown},
		[]gopacket.LayerType{LayerTypeSCTPShutdownAck},
		[]gopacket.LayerType{LayerTypeSCTPShutdownComplete},
	}
	for i, data := range sctpPackets {
		p := gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
		for _, typ := range wantLayers[i] {
			if p.Layer(typ) == nil {
				t.Errorf("Packet %d missing layer type %v, got:", i, typ)
				for _, layer := range p.Layers() {
					t.Errorf("\t%v", layer.LayerType())
				}
				if p.ErrorLayer() != nil {
					t.Error("\tPacket layer error:", p.ErrorLayer().Error())
				}
			}
		}
	}
}

func TestDecodeCiscoDiscovery(t *testing.T) {
	// http://wiki.wireshark.org/SampleCaptures?action=AttachFile&do=get&target=cdp_v2.pcap
	data := []byte{
		0x01, 0x00, 0x0c, 0xcc, 0xcc, 0xcc, 0x00, 0x0b, 0xbe, 0x18, 0x9a, 0x41, 0x01, 0xc3, 0xaa, 0xaa,
		0x03, 0x00, 0x00, 0x0c, 0x20, 0x00, 0x02, 0xb4, 0x09, 0xa0, 0x00, 0x01, 0x00, 0x0c, 0x6d, 0x79,
		0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x00, 0x02, 0x00, 0x11, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
		0xcc, 0x00, 0x04, 0xc0, 0xa8, 0x00, 0xfd, 0x00, 0x03, 0x00, 0x13, 0x46, 0x61, 0x73, 0x74, 0x45,
		0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x30, 0x2f, 0x31, 0x00, 0x04, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x28, 0x00, 0x05, 0x01, 0x14, 0x43, 0x69, 0x73, 0x63, 0x6f, 0x20, 0x49, 0x6e, 0x74, 0x65,
		0x72, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x20, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
		0x6e, 0x67, 0x20, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x20, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61,
		0x72, 0x65, 0x20, 0x0a, 0x49, 0x4f, 0x53, 0x20, 0x28, 0x74, 0x6d, 0x29, 0x20, 0x43, 0x32, 0x39,
		0x35, 0x30, 0x20, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x20, 0x28, 0x43, 0x32, 0x39,
		0x35, 0x30, 0x2d, 0x49, 0x36, 0x4b, 0x32, 0x4c, 0x32, 0x51, 0x34, 0x2d, 0x4d, 0x29, 0x2c, 0x20,
		0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x31, 0x32, 0x2e, 0x31, 0x28, 0x32, 0x32, 0x29,
		0x45, 0x41, 0x31, 0x34, 0x2c, 0x20, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x20, 0x53, 0x4f,
		0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x20, 0x28, 0x66, 0x63, 0x31, 0x29, 0x0a, 0x54, 0x65, 0x63,
		0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x20,
		0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f,
		0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
		0x0a, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x28, 0x63, 0x29, 0x20, 0x31,
		0x39, 0x38, 0x36, 0x2d, 0x32, 0x30, 0x31, 0x30, 0x20, 0x62, 0x79, 0x20, 0x63, 0x69, 0x73, 0x63,
		0x6f, 0x20, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x0a,
		0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x54, 0x75, 0x65, 0x20, 0x32, 0x36, 0x2d,
		0x4f, 0x63, 0x74, 0x2d, 0x31, 0x30, 0x20, 0x31, 0x30, 0x3a, 0x33, 0x35, 0x20, 0x62, 0x79, 0x20,
		0x6e, 0x62, 0x75, 0x72, 0x72, 0x61, 0x00, 0x06, 0x00, 0x15, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x20,
		0x57, 0x53, 0x2d, 0x43, 0x32, 0x39, 0x35, 0x30, 0x2d, 0x31, 0x32, 0x00, 0x08, 0x00, 0x24, 0x00,
		0x00, 0x0c, 0x01, 0x12, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x01, 0x02, 0x20, 0xff,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0xbe, 0x18, 0x9a, 0x40, 0xff, 0x00, 0x00, 0x00,
		0x09, 0x00, 0x0c, 0x4d, 0x59, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x00, 0x0a, 0x00, 0x06, 0x00,
		0x01, 0x00, 0x0b, 0x00, 0x05, 0x01, 0x00, 0x12, 0x00, 0x05, 0x00, 0x00, 0x13, 0x00, 0x05, 0x00,
		0x00, 0x16, 0x00, 0x11, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0xcc, 0x00, 0x04, 0xc0, 0xa8, 0x00,
		0xfd,
	}
	p := gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	wantLayers := []gopacket.LayerType{LayerTypeEthernet, LayerTypeLLC, LayerTypeSNAP, LayerTypeCiscoDiscovery, LayerTypeCiscoDiscoveryInfo}
	checkLayers(p, wantLayers, t)

	want := &CiscoDiscoveryInfo{
		CDPHello: CDPHello{
			OUI:              []byte{0, 0, 12},
			ProtocolID:       274,
			ClusterMaster:    []byte{0, 0, 0, 0},
			Unknown1:         []byte{255, 255, 255, 255},
			Version:          1,
			SubVersion:       2,
			Status:           32,
			Unknown2:         255,
			ClusterCommander: net.HardwareAddr{0, 0, 0, 0, 0, 0},
			SwitchMAC:        net.HardwareAddr{0, 0x0b, 0xbe, 0x18, 0x9a, 0x40},
			Unknown3:         255,
			ManagementVLAN:   0,
		},
		DeviceID:      "myswitch",
		Addresses:     []net.IP{net.IPv4(192, 168, 0, 253)},
		PortID:        "FastEthernet0/1",
		Capabilities:  CDPCapabilities{false, false, false, true, false, true, false, false, false},
		Version:       "Cisco Internetwork Operating System Software \nIOS (tm) C2950 Software (C2950-I6K2L2Q4-M), Version 12.1(22)EA14, RELEASE SOFTWARE (fc1)\nTechnical Support: http://www.cisco.com/techsupport\nCopyright (c) 1986-2010 by cisco Systems, Inc.\nCompiled Tue 26-Oct-10 10:35 by nburra",
		Platform:      "cisco WS-C2950-12",
		VTPDomain:     "MYDOMAIN",
		NativeVLAN:    1,
		FullDuplex:    true,
		MgmtAddresses: []net.IP{net.IPv4(192, 168, 0, 253)},
		BaseLayer:     BaseLayer{Contents: data[26:]},
	}
	cdpL := p.Layer(LayerTypeCiscoDiscoveryInfo)
	info, _ := cdpL.(*CiscoDiscoveryInfo)
	if !reflect.DeepEqual(info, want) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info, want)
	}
}

func TestDecodeLinkLayerDiscovery(t *testing.T) {
	// http://wiki.wireshark.org/SampleCaptures?action=AttachFile&do=get&target=lldp.detailed.pcap
	data := []byte{
		0x01, 0x80, 0xc2, 0x00, 0x00, 0x0e, 0x00, 0x01, 0x30, 0xf9, 0xad, 0xa0,
		0x88, 0xcc, 0x02, 0x07, 0x04, 0x00, 0x01, 0x30, 0xf9, 0xad, 0xa0, 0x04,
		0x04, 0x05, 0x31, 0x2f, 0x31, 0x06, 0x02, 0x00, 0x78, 0x08, 0x17, 0x53,
		0x75, 0x6d, 0x6d, 0x69, 0x74, 0x33, 0x30, 0x30, 0x2d, 0x34, 0x38, 0x2d,
		0x50, 0x6f, 0x72, 0x74, 0x20, 0x31, 0x30, 0x30, 0x31, 0x00, 0x0a, 0x0d,
		0x53, 0x75, 0x6d, 0x6d, 0x69, 0x74, 0x33, 0x30, 0x30, 0x2d, 0x34, 0x38,
		0x00, 0x0c, 0x4c, 0x53, 0x75, 0x6d, 0x6d, 0x69, 0x74, 0x33, 0x30, 0x30,
		0x2d, 0x34, 0x38, 0x20, 0x2d, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
		0x6e, 0x20, 0x37, 0x2e, 0x34, 0x65, 0x2e, 0x31, 0x20, 0x28, 0x42, 0x75,
		0x69, 0x6c, 0x64, 0x20, 0x35, 0x29, 0x20, 0x62, 0x79, 0x20, 0x52, 0x65,
		0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
		0x20, 0x30, 0x35, 0x2f, 0x32, 0x37, 0x2f, 0x30, 0x35, 0x20, 0x30, 0x34,
		0x3a, 0x35, 0x33, 0x3a, 0x31, 0x31, 0x00, 0x0e, 0x04, 0x00, 0x14, 0x00,
		0x14, 0x10, 0x0e, 0x07, 0x06, 0x00, 0x01, 0x30, 0xf9, 0xad, 0xa0, 0x02,
		0x00, 0x00, 0x03, 0xe9, 0x00, 0xfe, 0x07, 0x00, 0x12, 0x0f, 0x02, 0x07,
		0x01, 0x00, 0xfe, 0x09, 0x00, 0x12, 0x0f, 0x01, 0x03, 0x6c, 0x00, 0x00,
		0x10, 0xfe, 0x09, 0x00, 0x12, 0x0f, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00,
		0xfe, 0x06, 0x00, 0x12, 0x0f, 0x04, 0x05, 0xf2, 0xfe, 0x06, 0x00, 0x80,
		0xc2, 0x01, 0x01, 0xe8, 0xfe, 0x07, 0x00, 0x80, 0xc2, 0x02, 0x01, 0x00,
		0x00, 0xfe, 0x17, 0x00, 0x80, 0xc2, 0x03, 0x01, 0xe8, 0x10, 0x76, 0x32,
		0x2d, 0x30, 0x34, 0x38, 0x38, 0x2d, 0x30, 0x33, 0x2d, 0x30, 0x35, 0x30,
		0x35, 0x00, 0xfe, 0x05, 0x00, 0x80, 0xc2, 0x04, 0x00, 0x00, 0x00,
	}

	p := gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	wantLayers := []gopacket.LayerType{LayerTypeEthernet, LayerTypeLinkLayerDiscovery, LayerTypeLinkLayerDiscoveryInfo}
	checkLayers(p, wantLayers, t)
	lldpL := p.Layer(LayerTypeLinkLayerDiscovery)
	lldp := lldpL.(*LinkLayerDiscovery)
	want := &LinkLayerDiscovery{
		ChassisID: LLDPChassisID{LLDPChassisIDSubTypeMACAddr, []byte{0x00, 0x01, 0x30, 0xf9, 0xad, 0xa0}},
		PortID:    LLDPPortID{LLDPPortIDSubtypeIfaceName, []byte("1/1")},
		TTL:       120,
		BaseLayer: BaseLayer{Contents: data[14:]},
	}
	lldp.Values = nil // test these in next stage
	if !reflect.DeepEqual(lldp, want) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", lldp, want)
	}

	infoL := p.Layer(LayerTypeLinkLayerDiscoveryInfo)
	info := infoL.(*LinkLayerDiscoveryInfo)
	wantinfo := &LinkLayerDiscoveryInfo{
		PortDescription: "Summit300-48-Port 1001\x00",
		SysName:         "Summit300-48\x00",
		SysDescription:  "Summit300-48 - Version 7.4e.1 (Build 5) by Release_Master 05/27/05 04:53:11\x00",
		SysCapabilities: LLDPSysCapabilities{
			SystemCap:  LLDPCapabilities{Bridge: true, Router: true},
			EnabledCap: LLDPCapabilities{Bridge: true, Router: true},
		},
		MgmtAddress: LLDPMgmtAddress{IANAAddressFamily802, []byte{0x00, 0x01, 0x30, 0xf9, 0xad, 0xa0}, LLDPInterfaceSubtypeifIndex, 1001, ""},
		OrgTLVs: []LLDPOrgSpecificTLV{
			LLDPOrgSpecificTLV{OUI: 0x120f, SubType: 0x2, Info: []uint8{0x7, 0x1, 0x0}},
			LLDPOrgSpecificTLV{OUI: 0x120f, SubType: 0x1, Info: []uint8{0x3, 0x6c, 0x0, 0x0, 0x10}},
			LLDPOrgSpecificTLV{OUI: 0x120f, SubType: 0x3, Info: []uint8{0x1, 0x0, 0x0, 0x0, 0x0}},
			LLDPOrgSpecificTLV{OUI: 0x120f, SubType: 0x4, Info: []uint8{0x5, 0xf2}},
			LLDPOrgSpecificTLV{OUI: 0x80c2, SubType: 0x1, Info: []uint8{0x1, 0xe8}},
			LLDPOrgSpecificTLV{OUI: 0x80c2, SubType: 0x2, Info: []uint8{0x1, 0x0, 0x0}},
			LLDPOrgSpecificTLV{OUI: 0x80c2, SubType: 0x3, Info: []uint8{0x1, 0xe8, 0x10, 0x76, 0x32, 0x2d, 0x30, 0x34, 0x38, 0x38, 0x2d, 0x30, 0x33, 0x2d, 0x30, 0x35, 0x30, 0x35, 0x0}},
			LLDPOrgSpecificTLV{OUI: 0x80c2, SubType: 0x4, Info: []uint8{0x0}},
		},
		Unknown: nil,
	}
	if !reflect.DeepEqual(info, wantinfo) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info, wantinfo)
	}
	info8021, err := info.Decode8021()
	if err != nil {
		t.Errorf("8021 Values decode error: %v", err)
	}
	want8021 := LLDPInfo8021{
		PVID:               488,
		PPVIDs:             []PortProtocolVLANID{PortProtocolVLANID{false, false, 0}},
		VLANNames:          []VLANName{VLANName{488, "v2-0488-03-0505\x00"}},
		ProtocolIdentities: nil,
		VIDUsageDigest:     0,
		ManagementVID:      0,
		LinkAggregation:    LLDPLinkAggregation{false, false, 0},
	}
	if !reflect.DeepEqual(info8021, want8021) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info8021, want8021)
	}
	info8023, err := info.Decode8023()
	if err != nil {
		t.Errorf("8023 Values decode error: %v", err)
	}
	want8023 := LLDPInfo8023{
		LinkAggregation:    LLDPLinkAggregation{true, false, 0},
		MACPHYConfigStatus: LLDPMACPHYConfigStatus{true, true, 0x6c00, 0x0010},
		PowerViaMDI:        LLDPPowerViaMDI8023{true, true, true, false, 1, 0, 0, 0, 0, 0, 0},
		MTU:                1522,
	}

	if !reflect.DeepEqual(info8023, want8023) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info8023, want8023)
	}

	// http://wiki.wireshark.org/SampleCaptures?action=AttachFile&do=get&target=lldpmed_civicloc.pcap
	data = []byte{
		0x01, 0x80, 0xc2, 0x00, 0x00, 0x0e, 0x00, 0x13, 0x21, 0x57, 0xca, 0x7f,
		0x88, 0xcc, 0x02, 0x07, 0x04, 0x00, 0x13, 0x21, 0x57, 0xca, 0x40, 0x04,
		0x02, 0x07, 0x31, 0x06, 0x02, 0x00, 0x78, 0x08, 0x01, 0x31, 0x0a, 0x1a,
		0x50, 0x72, 0x6f, 0x43, 0x75, 0x72, 0x76, 0x65, 0x20, 0x53, 0x77, 0x69,
		0x74, 0x63, 0x68, 0x20, 0x32, 0x36, 0x30, 0x30, 0x2d, 0x38, 0x2d, 0x50,
		0x57, 0x52, 0x0c, 0x5f, 0x50, 0x72, 0x6f, 0x43, 0x75, 0x72, 0x76, 0x65,
		0x20, 0x4a, 0x38, 0x37, 0x36, 0x32, 0x41, 0x20, 0x53, 0x77, 0x69, 0x74,
		0x63, 0x68, 0x20, 0x32, 0x36, 0x30, 0x30, 0x2d, 0x38, 0x2d, 0x50, 0x57,
		0x52, 0x2c, 0x20, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20,
		0x48, 0x2e, 0x30, 0x38, 0x2e, 0x38, 0x39, 0x2c, 0x20, 0x52, 0x4f, 0x4d,
		0x20, 0x48, 0x2e, 0x30, 0x38, 0x2e, 0x35, 0x58, 0x20, 0x28, 0x2f, 0x73,
		0x77, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
		0x2f, 0x66, 0x69, 0x73, 0x68, 0x28, 0x74, 0x73, 0x5f, 0x30, 0x38, 0x5f,
		0x35, 0x29, 0x29, 0x0e, 0x04, 0x00, 0x14, 0x00, 0x04, 0x10, 0x0c, 0x05,
		0x01, 0x0f, 0xff, 0x7a, 0x94, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe,
		0x09, 0x00, 0x12, 0x0f, 0x01, 0x03, 0x6c, 0x00, 0x00, 0x10, 0xfe, 0x07,
		0x00, 0x12, 0xbb, 0x01, 0x00, 0x0f, 0x04, 0xfe, 0x08, 0x00, 0x12, 0xbb,
		0x02, 0x01, 0x40, 0x65, 0xae, 0xfe, 0x2e, 0x00, 0x12, 0xbb, 0x03, 0x02,
		0x28, 0x02, 0x55, 0x53, 0x01, 0x02, 0x43, 0x41, 0x03, 0x09, 0x52, 0x6f,
		0x73, 0x65, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x06, 0x09, 0x46, 0x6f, 0x6f,
		0x74, 0x68, 0x69, 0x6c, 0x6c, 0x73, 0x13, 0x04, 0x38, 0x30, 0x30, 0x30,
		0x1a, 0x03, 0x52, 0x33, 0x4c, 0xfe, 0x07, 0x00, 0x12, 0xbb, 0x04, 0x03,
		0x00, 0x41, 0x00, 0x00,
	}

	p = gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	wantLayers = []gopacket.LayerType{LayerTypeEthernet, LayerTypeLinkLayerDiscovery, LayerTypeLinkLayerDiscoveryInfo}
	checkLayers(p, wantLayers, t)
	lldpL = p.Layer(LayerTypeLinkLayerDiscovery)
	lldp = lldpL.(*LinkLayerDiscovery)
	want = &LinkLayerDiscovery{
		ChassisID: LLDPChassisID{LLDPChassisIDSubTypeMACAddr, []byte{0x00, 0x13, 0x21, 0x57, 0xca, 0x40}},
		PortID:    LLDPPortID{LLDPPortIDSubtypeLocal, []byte("1")},
		TTL:       120,
		BaseLayer: BaseLayer{Contents: data[14:]},
	}
	lldp.Values = nil // test these in next stage
	if !reflect.DeepEqual(lldp, want) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", lldp, want)
	}

	infoL = p.Layer(LayerTypeLinkLayerDiscoveryInfo)
	info = infoL.(*LinkLayerDiscoveryInfo)
	wantinfo = &LinkLayerDiscoveryInfo{
		PortDescription: "1",
		SysName:         "ProCurve Switch 2600-8-PWR",
		SysDescription:  "ProCurve J8762A Switch 2600-8-PWR, revision H.08.89, ROM H.08.5X (/sw/code/build/fish(ts_08_5))",
		SysCapabilities: LLDPSysCapabilities{
			SystemCap:  LLDPCapabilities{Bridge: true, Router: true},
			EnabledCap: LLDPCapabilities{Bridge: true},
		},
		MgmtAddress: LLDPMgmtAddress{IANAAddressFamilyIPV4, []byte{0x0f, 0xff, 0x7a, 0x94}, LLDPInterfaceSubtypeifIndex, 0, ""},
		OrgTLVs: []LLDPOrgSpecificTLV{
			LLDPOrgSpecificTLV{OUI: 0x120f, SubType: 0x1, Info: []uint8{0x3, 0x6c, 0x0, 0x0, 0x10}},
			LLDPOrgSpecificTLV{OUI: 0x12bb, SubType: 0x1, Info: []uint8{0x0, 0xf, 0x4}},
			LLDPOrgSpecificTLV{OUI: 0x12bb, SubType: 0x2, Info: []uint8{0x1, 0x40, 0x65, 0xae}},
			LLDPOrgSpecificTLV{OUI: 0x12bb, SubType: 0x3, Info: []uint8{0x2, 0x28, 0x2, 0x55, 0x53, 0x1, 0x2, 0x43, 0x41, 0x3, 0x9, 0x52, 0x6f, 0x73, 0x65, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x6, 0x9, 0x46, 0x6f, 0x6f, 0x74, 0x68, 0x69, 0x6c, 0x6c, 0x73, 0x13, 0x4, 0x38, 0x30, 0x30, 0x30, 0x1a, 0x3, 0x52, 0x33, 0x4c}},
			LLDPOrgSpecificTLV{OUI: 0x12bb, SubType: 0x4, Info: []uint8{0x3, 0x0, 0x41}},
		},
		Unknown: nil,
	}
	if !reflect.DeepEqual(info, wantinfo) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info, wantinfo)
	}
	info8023, err = info.Decode8023()
	if err != nil {
		t.Errorf("8023 Values decode error: %v", err)
	}
	want8023 = LLDPInfo8023{
		MACPHYConfigStatus: LLDPMACPHYConfigStatus{true, true, 0x6c00, 0x0010},
	}

	if !reflect.DeepEqual(info8023, want8023) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info8023, want8023)
	}

	infoMedia, err := info.DecodeMedia()
	if err != nil {
		t.Errorf("8023 Values decode error: %v", err)
	}
	wantMedia := LLDPInfoMedia{
		MediaCapabilities: LLDPMediaCapabilities{true, true, true, true, false, false, LLDPMediaClassNetwork},
		NetworkPolicy:     LLDPNetworkPolicy{LLDPAppTypeVoice, true, true, 50, 6, 46},
		Location: LLDPLocation{Format: LLDPLocationFormatAddress, Address: LLDPLocationAddress{
			What:        LLDPLocationAddressWhatClient,
			CountryCode: "US",
			AddressLines: []LLDPLocationAddressLine{
				LLDPLocationAddressLine{LLDPLocationAddressTypeNational, "CA"},
				LLDPLocationAddressLine{LLDPLocationAddressTypeCity, "Roseville"},
				LLDPLocationAddressLine{LLDPLocationAddressTypeStreet, "Foothills"},
				LLDPLocationAddressLine{LLDPLocationAddressTypeHouseNum, "8000"},
				LLDPLocationAddressLine{LLDPLocationAddressTypeUnit, "R3L"},
			},
		}},
		PowerViaMDI: LLDPPowerViaMDI{0, 0, LLDPPowerPriorityLow, 6500},
	}

	if !reflect.DeepEqual(infoMedia, wantMedia) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", infoMedia, wantMedia)
	}

}

func TestDecodeNortelDiscovery(t *testing.T) {
	// http://www.thetechfirm.com/packets/nortel_btdp/btdp_nai.enc
	data := []byte{
		0x01, 0x00, 0x81, 0x00, 0x01, 0x00, 0x00, 0x04, 0x38, 0xe0, 0xcc, 0xde,
		0x00, 0x13, 0xaa, 0xaa, 0x03, 0x00, 0x00, 0x81, 0x01, 0xa2, 0xac, 0x13,
		0x58, 0x03, 0x00, 0x04, 0x15, 0x30, 0x0c, 0x02, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x38, 0xe0, 0xcc, 0xde, 0x80, 0x6a, 0x00, 0x01, 0x14, 0x00,
		0x02, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	p := gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	wantLayers := []gopacket.LayerType{LayerTypeEthernet, LayerTypeLLC, LayerTypeSNAP, LayerTypeNortelDiscovery}
	checkLayers(p, wantLayers, t)

	want := &NortelDiscovery{
		IPAddress: []byte{172, 19, 88, 3},
		SegmentID: []byte{0x00, 0x04, 0x15},
		Chassis:   NDPChassisBayStack450101001000Switches,
		Backplane: NDPBackplaneEthernetFastEthernetGigabitEthernet,
		State:     NDPStateHeartbeat,
		NumLinks:  0,
	}
	ndpL := p.Layer(LayerTypeNortelDiscovery)
	info, _ := ndpL.(*NortelDiscovery)
	if !reflect.DeepEqual(info, want) {
		t.Errorf("Values mismatch, \ngot  %#v\nwant %#v\n", info, want)
	}
}

func TestDecodeIPv6Jumbogram(t *testing.T) {
	// Haven't found any of these in the wild or on example pcaps online, so had
	// to generate one myself via scapy.  Unfortunately, scapy can only
	// str(packet) for packets with length < 65536, due to limitations in python's
	// struct library, so I generated the header with:
	// Ether() / IPv6(src='::1', dst='::2') / IPv6ExtHdrHopByHop(options=[Jumbo(jumboplen=70000)]) / TCP(sport=8888, dport=80)
	// then added the payload manually ("payload" * 9996).  The checksums here are
	// not correct, but we don't check, so who cares ;)
	dataStr := "\x00\x1f\xca\xb3v@$\xbe\x05'\x0b\x17\x86\xdd`\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x06\x00\xc2\x04\x00\x01\x11p\"\xb8\x00P\x00\x00\x00\x00\x00\x00\x00\x00P\x02 \x00l\xd8\x00\x00"
	payload := strings.Repeat("payload", 9996)
	data := []byte(dataStr + payload)
	p := gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeIPv6HopByHop, LayerTypeTCP, gopacket.LayerTypePayload}, t)
	if p.ApplicationLayer() == nil {
		t.Error("Packet has no application layer")
	} else if string(p.ApplicationLayer().Payload()) != payload {
		t.Errorf("Jumbogram payload wrong")
	}
	// Check truncated for jumbograms
	data = data[:len(data)-1]
	p = gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv6, LayerTypeIPv6HopByHop, LayerTypeTCP, gopacket.LayerTypePayload}, t)
	if !p.Metadata().Truncated {
		t.Error("Jumbogram should be truncated")
	}
}

func TestDecodeUDPPacketTooSmall(t *testing.T) {
	data := []byte{
		0x00, 0x15, 0x2c, 0x9d, 0xcc, 0x00, 0x00, 0x10, 0xdb, 0xff, 0x10, 0x00, 0x81, 0x00, 0x01, 0xf7,
		0x08, 0x00, 0x45, 0x60, 0x00, 0x3c, 0x0f, 0xa9, 0x00, 0x00, 0x6e, 0x11, 0x01, 0x0a, 0x47, 0xe6,
		0xee, 0x2e, 0xac, 0x16, 0x59, 0x73, 0x00, 0x50, 0x00, 0x50, 0x00, 0x28, 0x4d, 0xad, 0x00, 0x67,
		0x00, 0x01, 0x00, 0x72, 0xd5, 0xc7, 0xf1, 0x07, 0x00, 0x00, 0x01, 0x01, 0x00, 0x0d, 0x00, 0x00,
		0x00, 0x14, 0x00, 0x00, 0x19, 0xba,
	}
	p := gopacket.NewPacket(data, LinkTypeEthernet, gopacket.Default)
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeDot1Q, LayerTypeIPv4, LayerTypeUDP, gopacket.LayerTypePayload}, t)
	if !p.Metadata().Truncated {
		t.Error("UDP short packet should be truncated")
	}
}
