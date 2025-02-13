// Copyright 2022 Alim Zanibekov
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package teltonika

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestCodec12Encode(t *testing.T) {
	buffer, err := EncodePacketTCP(&Packet{
		CodecID:  Codec12,
		Data:     nil,
		Messages: []Message{{Type: TypeCommand, Text: "getinfo"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("000000000000000F0C010500000007676574696E666F0100004312")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodec13Encode(t *testing.T) {
	buffer, err := EncodePacketTCP(&Packet{
		CodecID:  Codec13,
		Data:     nil,
		Messages: []Message{{Type: TypeResponse, Text: "getinfo", Timestamp: 176276256}},
	})
	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("00000000000000130d01060000000b0a81c320676574696e666f0100001d6b")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodec14Encode(t *testing.T) {
	buffer, err := EncodePacketTCP(&Packet{
		CodecID:  Codec14,
		Data:     nil,
		Messages: []Message{{Type: TypeCommand, Text: "getver", Imei: "352093081452251"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("00000000000000160E01050000000E0352093081452251676574766572010000D2C1")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodec15Encode(t *testing.T) {
	buffer, err := EncodePacketTCP(&Packet{
		CodecID:  Codec15,
		Data:     nil,
		Messages: []Message{{Type: 0x0B, Text: "Hello!\n", Imei: "123456789123456", Timestamp: 1699440036}},
	})
	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("000000000000001b0f010b00000013654b65a4012345678912345648656c6c6f210a01000093d6")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodecs8EncodeTCP(t *testing.T) {
	buffer, err := EncodePacketTCP(&Packet{
		CodecID: Codec8,
		Data: []Data{{
			TimestampMs:    1720003694000,
			Lat:            42.373737,
			Lng:            42.373737,
			Altitude:       731,
			Angle:          262,
			EventID:        239,
			Speed:          0,
			Satellites:     5,
			Priority:       1,
			GenerationType: Unknown,
			Elements: []IOElement{
				{Id: 239, Value: []byte{1}},
				{Id: 240, Value: []byte{1}},
				{Id: 1, Value: []byte{0}},
				{Id: 66, Value: []byte{0x31, 0x4a}},
				{Id: 67, Value: []byte{0x10, 0x31}},
				{Id: 9, Value: []byte{0, 0x2b}},
			},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("000000000000003008010000019078358db0011941b81a1941b81a02db0106050000ef0603ef01f00101000342314a43103109002b0000010000c897")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodecs8EEncodeTCP(t *testing.T) {
	msg11, _ := hex.DecodeString("000000003544c87a")
	msg14, _ := hex.DecodeString("000000001dd7e06a")
	buffer, err := EncodePacketTCP(&Packet{
		CodecID: Codec8E,
		Data: []Data{{
			TimestampMs: 1560166592000,
			Lat:         0, Lng: 0, Altitude: 0, Angle: 0, EventID: 1, Speed: 0, Satellites: 0, Priority: 1,
			GenerationType: Unknown,
			Elements: []IOElement{
				{Id: 1, Value: []byte{1}},
				{Id: 17, Value: []byte{0, 0x1d}},
				{Id: 16, Value: []byte{0x01, 0x5e, 0x2c, 0x88}},
				{Id: 11, Value: msg11},
				{Id: 14, Value: msg14},
			},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("000000000000004A8E010000016B412CEE000100000000000000000000000000000000010005000100010100010011001D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A00000100002994")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodecs16EncodeTCP(t *testing.T) {
	buffer, err := EncodePacketTCP(&Packet{
		CodecID: Codec16,
		Data: []Data{{
			TimestampMs: 1562760414000,
			Lat:         0, Lng: 0, Altitude: 0, Angle: 0, EventID: 11, Speed: 0, Satellites: 0, Priority: 0,
			GenerationType: OnChange,
			Elements: []IOElement{
				{Id: 1, Value: []byte{0}},
				{Id: 3, Value: []byte{0}},
				{Id: 11, Value: []byte{0, 0x27}},
				{Id: 66, Value: []byte{0x56, 0x3a}},
			},
		}, {
			TimestampMs: 1562760415000,
			Lat:         0, Lng: 0, Altitude: 0, Angle: 0, EventID: 11, Speed: 0, Satellites: 0, Priority: 0,
			GenerationType: OnChange,
			Elements: []IOElement{
				{Id: 1, Value: []byte{0}},
				{Id: 3, Value: []byte{0}},
				{Id: 11, Value: []byte{0, 0x26}},
				{Id: 66, Value: []byte{0x56, 0x3a}},
			},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("000000000000005F10020000016BDBC7833000000000000000000000000000000000000B05040200010000030002000B00270042563A00000000016BDBC7871800000000000000000000000000000000000B05040200010000030002000B00260042563A00000200005FB3")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodecs8EncodeUDP(t *testing.T) {
	buffer, err := EncodePacketUDP("352093086403655", 51966, 5, &Packet{
		CodecID: Codec8,
		Data: []Data{{
			TimestampMs: 1560407006000,
			Lat:         0, Lng: 0, Altitude: 0, Angle: 0, EventID: 1, Speed: 0, Satellites: 0, Priority: 1,
			GenerationType: Unknown,
			Elements: []IOElement{
				{Id: 21, Value: []byte{3}},
				{Id: 1, Value: []byte{1}},
				{Id: 66, Value: []byte{0x5d, 0xbc}},
			},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("003DCAFE0105000F33353230393330383634303336353508010000016B4F815B30010000000000000000000000000000000103021503010101425DBC000001")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodecs8EEncodeUDP(t *testing.T) {
	msg16, _ := hex.DecodeString("015e2c88")
	msg11, _ := hex.DecodeString("000000003544c87a")
	msg14, _ := hex.DecodeString("000000001dd7e06a")
	buffer, err := EncodePacketUDP("352093086403655", 51966, 7, &Packet{
		CodecID: Codec8E,
		Data: []Data{{
			TimestampMs: 1560407121000,
			Lat:         0, Lng: 0, Altitude: 0, Angle: 0, EventID: 1, Speed: 0, Satellites: 0, Priority: 1,
			GenerationType: Unknown,
			Elements: []IOElement{
				{Id: 1, Value: []byte{1}},
				{Id: 17, Value: []byte{0, 0x9d}},
				{Id: 16, Value: msg16},
				{Id: 11, Value: msg11},
				{Id: 14, Value: msg14},
			},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("005FCAFE0107000F3335323039333038363430333635358E010000016B4F831C680100000000000000000000000000000000010005000100010100010011009D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A000001")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestCodecs16EncodeUDP(t *testing.T) {
	buffer, err := EncodePacketUDP("352094085231592", 51966, 7, &Packet{
		CodecID: Codec16,
		Data: []Data{{
			TimestampMs: 1447804801000,
			Lat:         0, Lng: 0, Altitude: 0, Angle: 0, EventID: 239, Speed: 0, Satellites: 0, Priority: 0,
			GenerationType: OnChange,
			Elements: []IOElement{
				{Id: 1, Value: []byte{0}},
				{Id: 3, Value: []byte{0}},
				{Id: 180, Value: []byte{0}},
				{Id: 239, Value: []byte{1}},
				{Id: 66, Value: []byte{0x11, 0x1a}},
			},
		}},
	})

	if err != nil {
		t.Fatal(err)
	}

	encoded := hex.EncodeToString(buffer)
	expected := strings.ToLower("0048CAFE0107000F33353230393430383532333135393210010000015117E40FE80000000000000000000000000000000000EF05050400010000030000B40000EF01010042111A000001")

	if encoded != expected {
		t.Error("encoded: ", encoded)
		t.Error("expected:", expected)
	}
}

func TestTCPCodecsDecodeWithoutErrors(t *testing.T) {
	cases := []string{
		"000000000000004308020000016B40D57B480100000000000000000000000000000001010101000000000000016B40D5C198010000000000000000000000000000000101010101000000020000252C",
		/* "000000000000002808010000016B40D9AD80010000000000000000000000000000000103021503010101425E100000010000F22A",
		"000000000000004308020000016B40D57B480100000000000000000000000000000001010101000000000000016B40D5C198010000000000000000000000000000000101010101000000020000252C",
		"000000000000005F10020000016BDBC7833000000000000000000000000000000000000B05040200010000030002000B00270042563A00000000016BDBC7871800000000000000000000000000000000000B05040200010000030002000B00260042563A00000200005FB3",
		"000000000000004A8E010000016B412CEE000100000000000000000000000000000000010005000100010100010011001D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A00000100002994",
		"00000000000000A98E020000017357633410000F0DC39B2095964A00AC00F80B00000000000B000500F00100150400C800004501007156000500B5000500B600040018000000430FE00044011B000100F10000601B000000000000017357633BE1000F0DC39B2095964A00AC00F80B000001810001000000000000000000010181002D11213102030405060708090A0B0C0D0E0F104545010ABC212102030405060708090A0B0C0D0E0F10020B010AAD020000BF30",
		"000000000000000F0C010500000007676574696E666F0100004312",
		"00000000000000130d01060000000b0a81c320676574696e666f0100001d6b",
		"00000000000000160E01050000000E0352093081452251676574766572010000D2C1",
		"000000000000001b0f010b00000013654b65a4012345678912345648656c6c6f210a01000093d6",
		"00000000000003c68e04000001909d6417c801062c583a25ce28b9004f004530000000ef0011000500ef0000f00100150300c800004501000a00b5000800b60004004234a90018000000430ef700440052001100900012fc4c0013fb77000f03e8000200f100005e89001000003f9a00000000000001909d606e5201062bc70c25ccd6d50096010032000000ef0011000500ef0100f00100150500c800004501000a00b5000800b60005004237780018000000430ed500440000001100da0012fd0e0013f9ef000f03e8000200f100005e89001000003a6400000000000001909d4f2c5001062bc70c25ccd6d5009601002c000000ef0011000500ef0000f00100150500c800004501000a00b5000800b60005004234b10018000000430f0300440041001100db0012fd0d0013f9e9000f0047000200f100005e89001000003a6400000000000001909d4e03d501062c0f5025cdd1170068008130004800f700030002013d0100f70500000000000000010101025801dffe02f95d01c2fdc4f97001b1fde2f96f0193fdd2f97001a1fe01f95f0171fdfff95001b1fdf3f97e01a1fe01f95f0200fdd5f97d01eefe12f95c01affe31f96d01d0fe33f9aa01cffe62f9a901b1fe43f9ba01a0fe51f99b0200fe05f99b01d0fe33f99b0170fe6ff99b0161fe6ff9ab0164fe22f9cd0192fe22f99d0183fe32f9cc0170fe6ff99b0142fe4ff99d0133fe6ef9bc0172fe30f99d01c2fe15f9cb019ffe70f99a0182fe11f97e0183fdf2f99f0152fe2ff98e0163fe21f9ae0181fe2ff96e0152fe2ff98e0162fe30f98e0191fe11f97e01c0fe12f97d019ffe20f94e0180fe2ff95e0192fde1f96001a1fdf1f95f01d1fdc4f96f01a2fdc2f95101bffdf0f9300182fde1f96101d1fdb3f94101a1fdd0f93201c2fda3f95101b1fdd1f94101b1fdd1f94101effde3f94e01b0fdf1f94001c0fde1f94001c1fdd2f9500171fdeef92201b0fde1f94001a1fdc1f93201a0fdc0f9130171fdcff92301dffdb2f9120190fdfff9300191fde0f9410191fe00f95f0172fdd0f94201cffde1f92001bdfdeef8e201affdd0f91201c0fdc2f93101a1fdb1f93301bffde0f92101f1fd74f92301a1fd80f8f601d1fd83f91401a0fdaff8f401a0fde0f92101affdc0f90301bffdd0f90201dffdb1f90201cefdbff8e301eefda2f8f3019ffdaef8c501f0fd94f932019ffddff90201b0fe01f94f01ddfdfff90001bbfe4cf8ee01befe30f94d01cdfe1ff90f01cffdf2f94f019ffe1ff93f01b1fdd2f9600191fe00f95f0184fdc3f9910192fe12f99d01a0fe10f95e0180fe4ff97c01a1fe31f98c01a0fe21f96e0193fdf2f98f0191fe21f98d0400006a96", */
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		log.Println("buf", buf)
		size, _, err := DecodeTCPFromSlice(buf)
		log.Println("size", size)
		if err != nil {
			t.Fatal(err)
		}
		if size != len(buf) {
			t.Error("[DecodeTCPFromSlice] payload not fully processed")
		}
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		reader := bytes.NewReader(buf)
		read, _, err := DecodeTCPFromReader(reader)
		if err != nil {
			t.Fatal(err)
		}
		if reader.Len() != 0 {
			t.Error("[DecodeTCPFromReader] payload not fully processed")
		}
		if hex.Dump(read) != hex.Dump(buf) {
			t.Error("[DecodeTCPFromReader] read bytes buffer invalid")
		}
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		reader := bytes.NewReader(buf)
		output := make([]byte, 1300)
		n, _, err := DecodeTCPFromReaderBuf(reader, output)
		if err != nil {
			t.Fatal(err)
		}
		if reader.Len() != 0 {
			t.Error("[DecodeTCPFromReaderBuf] payload not fully processed")
		}
		if hex.Dump(output[:n]) != hex.Dump(buf) {
			t.Error("[DecodeTCPFromReaderBuf] read bytes buffer invalid")
		}
	}
}

func TestTCPCodecsDecodeMustFail(t *testing.T) {
	cases := []string{
		"000001000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7C1",
		"000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000050000C7C1",
		"000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7C1",
		"000000000000002808010000016B40D9AD80010000000000000000000000000000000103021503010101425E100000010000F23A",
		"000000000000003604010000001B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000050000C7C1",
		"000000000000003611010000019B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000050000C7C1",
		"000000000000004308020000016B40D57B480100000000000000000000000000000001010101000000000000016B40D5C198010000000000000000000000000000000101010101000000020000254C",
		"000000000000005F10020000016BDBC7833000000000000000000000000000000000000B05040200010000030002000B00270042563A00000000016BDBC7871800000000000000000000000000000000000B05040200010000030002000B00260042563A00000200005F13",
		"000000000000004A8E010000016B412CEE000100000000000000000000000000000000010005000100010100010011001D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A00000100002991",
		"00000000000000A98E020000017357633410000F0DC39B2095964A00AC00F80B00000000000B000500F00100150400C800004501007156000500B5000500B600040018000000430FE00044011B000100F10000601B000000000000017357633BE1000F0DC39B2095964A00AC00F80B000001810001000000000000000000010181002D11213102030405060708090A0B0C0D0E0F104545010ABC212102030405060708090A0B0C0D0E0F10020B010AAD020000BF40",
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		_, _, err := DecodeTCPFromSlice(buf)
		if err == nil {
			t.Error("[DecodeTCPFromSlice] invalid payload processed successfully")
		}
	}
}

func TestUDPCodecsDecodeWithoutErrors(t *testing.T) {
	cases := []string{
		"005FCAFE0107000F3335323039333038363430333635358E010000016B4F831C680100000000000000000000000000000000010005000100010100010011009D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A000001",
		"003DCAFE0105000F33353230393330383634303336353508010000016B4F815B30010000000000000000000000000000000103021503010101425DBC000001",
		"0086CAFE0101000F3335323039333038353639383230368E0100000167EFA919800200000000000000000000000000000000FC0013000800EF0000F00000150500C80000450200010000710000FC00000900B5000000B600000042305600CD432A00CE6064001100090012FF22001303D1000F0000000200F1000059D90010000000000000000001",
		"0083CAFE0101000F3335323039333038353639383230368E0100000167F1AEEC00000A750E8F1D43443100F800B210000000000012000700EF0000F00000150500C800004501000100007142000900B5000600B6000500422FB300CD432A00CE60640011000700120007001303EC000F0000000200F1000059D90010000000000000000001",
		"01E4CAFE0126000F333532303934303839333937343634080400000163C803B420010A259E1A1D4A057D00DA0128130057421B0A4503F00150051503EF01510052005900BE00C1000AB50008B60005427025CD79D8CE605A5400005500007300005A0000C0000007C700000018F1000059D910002D32C85300000000570000000064000000F7BF000000000000000163C803AC50010A25A9D21D4A01B600DB0128130056421B0A4503F00150051503EF01510052005900BE00C1000AB50008B6000542702ECD79D8CE605A5400005500007300005A0000C0000007C700000017F1000059D910002D32B05300000000570000000064000000F7BF000000000000000163C803A868010A25B5581D49FE5400DB0127130057421B0A4503F00150051503EF01510052005900BE00C1000AB50008B60005427039CD79D8CE605A5400005500007300005A0000C0000007C700000017F1000059D910002D32995300000000570000000064000000F7BF000000000000000163C803A4B2010A25CC861D49F75C00DB0124130058421B0A4503F00150051503EF01510052005900BE00C1000AB50008B6000542703CCD79D8CE605A5400005500007300005A0000C0000007C700000018F1000059D910002D32695300000000570000000064000000F7BF000000000004",
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		size, _, err := DecodeUDPFromSlice(buf)
		if err != nil {
			t.Fatal(err)
		}
		if size != len(buf) {
			t.Error("[DecodeUDPFromSlice] payload not fully processed")
		}
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		reader := bytes.NewReader(buf)
		read, _, err := DecodeUDPFromReader(reader)
		if err != nil {
			t.Fatal(err)
		}
		if reader.Len() != 0 {
			t.Error("[DecodeUDPFromReader] payload not fully processed")
		}
		if hex.Dump(read) != hex.Dump(buf) {
			t.Error("[DecodeUDPFromReader] read bytes buffer invalid")
		}
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		reader := bytes.NewReader(buf)
		output := make([]byte, 1300)
		n, _, err := DecodeUDPFromReaderBuf(reader, output)
		if err != nil {
			t.Fatal(err)
		}
		if reader.Len() != 0 {
			t.Error("[DecodeUDPFromReaderBuf] payload not fully processed")
		}
		if hex.Dump(output[:n]) != hex.Dump(buf) {
			t.Error("[DecodeUDPFromReaderBuf] read bytes buffer invalid")
		}
	}
}

func TestUDPCodecsDecodeMustFail(t *testing.T) {
	cases := []string{
		"005FCAFE010700043335323039333038363430333635358E010000016B4F831C680100000000000000000000000000000000010005000100010100010011009D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A000001",
		"013DCAFE0105000F33353230393330383634303336353508010000016B4F815B30010000000000000000000000000000000103021503010101425DBC000001",
		"015BCAFE0101000F33353230393430383532333135393210070000015117E40FE80000000000000000000000000000000000EF05050400010000030000B40000EF01010042111A000001",
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		_, _, err := DecodeUDPFromSlice(buf)
		if err == nil {
			t.Error("[DecodeUDPFromSlice] invalid payload processed successfully")
		}
	}
}

func TestCodecsDecodeCommandResponseWithoutErrors(t *testing.T) {
	cases := []string{
		"00000000000000900C010600000088494E493A323031392F372F323220373A3232205254433A323031392F372F323220373A3533205253543A32204552523A312053523A302042523A302043463A302046473A3020464C3A302054553A302F302055543A3020534D533A30204E4F4750533A303A3330204750533A31205341543A302052533A332052463A36352053463A31204D443A30010000C78F",
		"00000000000000370C01060000002F4449313A31204449323A30204449333A302041494E313A302041494E323A313639323420444F313A3020444F323A3101000066E3",
		"00000000000000AB0E0106000000A303520930814522515665723A30332E31382E31345F3034204750533A41584E5F352E31305F333333332048773A464D42313230204D6F643A313520494D45493A33353230393330383134353232353120496E69743A323031382D31312D323220373A313320557074696D653A3137323334204D41433A363042444430303136323631205350433A312830292041584C3A30204F42443A3020424C3A312E362042543A340100007AAE",
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		size, _, err := DecodeTCPFromSlice(buf)
		if err != nil {
			t.Fatal(err)
		}
		if size != len(buf) {
			t.Error("[DecodeTCPFromSlice] payload not fully processed")
		}
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		reader := bytes.NewReader(buf)
		read, _, err := DecodeTCPFromReader(reader)
		if err != nil {
			t.Fatal(err)
		}
		if reader.Len() != 0 {
			t.Error("[DecodeTCPFromReader] payload not fully processed")
		}
		if hex.Dump(read) != hex.Dump(buf) {
			t.Error("[DecodeTCPFromReader] read bytes buffer invalid")
		}
	}
}

func TestPacketJSONMarshalUnmarshal(t *testing.T) {
	cases := []string{
		"00000000000000900C010600000088494E493A323031392F372F323220373A3232205254433A323031392F372F323220373A3533205253543A32204552523A312053523A302042523A302043463A302046473A3020464C3A302054553A302F302055543A3020534D533A30204E4F4750533A303A3330204750533A31205341543A302052533A332052463A36352053463A31204D443A30010000C78F",
		"00000000000000370C01060000002F4449313A31204449323A30204449333A302041494E313A302041494E323A313639323420444F313A3020444F323A3101000066E3",
		"000000000000005F10020000016BDBC7833000000000000000000000000000000000000B05040200010000030002000B00270042563A00000000016BDBC7871800000000000000000000000000000000000B05040200010000030002000B00260042563A00000200005FB3",
	}

	for _, s := range cases {
		buf, _ := hex.DecodeString(s)
		size, decoded, err := DecodeTCPFromSlice(buf)
		if err != nil {
			t.Fatal(err)
		}
		if size != len(buf) {
			t.Error("[DecodeTCPFromSlice] payload not fully processed")
		}

		res, err := json.Marshal(decoded.Packet)
		if err != nil {
			t.Fatal(err)
		}
		var packet Packet
		err = json.Unmarshal(res, &packet)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func tcpBenchCases(n int ) [][]byte {
	cases := []string{
		"000000000000003608010000016B40D8EA30010000000000000000000000000000000105021503010101425E0F01F10000601A014E0000000000000000010000C7CF",
		"000000000000002808010000016B40D9AD80010000000000000000000000000000000103021503010101425E100000010000F22A",
		"000000000000004308020000016B40D57B480100000000000000000000000000000001010101000000000000016B40D5C198010000000000000000000000000000000101010101000000020000252C",
		"000000000000005F10020000016BDBC7833000000000000000000000000000000000000B05040200010000030002000B00270042563A00000000016BDBC7871800000000000000000000000000000000000B05040200010000030002000B00260042563A00000200005FB3",
		"000000000000004A8E010000016B412CEE000100000000000000000000000000000000010005000100010100010011001D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A00000100002994",
		"00000000000000A98E020000017357633410000F0DC39B2095964A00AC00F80B00000000000B000500F00100150400C800004501007156000500B5000500B600040018000000430FE00044011B000100F10000601B000000000000017357633BE1000F0DC39B2095964A00AC00F80B000001810001000000000000000000010181002D11213102030405060708090A0B0C0D0E0F104545010ABC212102030405060708090A0B0C0D0E0F10020B010AAD020000BF30",
	}

	casesBin := make([][]byte, len(cases))

	for i := range cases {
		bs, _ := hex.DecodeString(cases[i])
		casesBin[i] = bs

	}

	benchCases := make([][]byte, n)

	for i := 0; i < n; i++ {
		benchCases[i] = casesBin[rand.Intn(len(cases))]
	}

	return benchCases
}

func benchmarkTCPDecodeSlice(b *testing.B, config ...*DecodeConfig) {
	benchCases := tcpBenchCases(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := DecodeTCPFromSlice(benchCases[i], config...)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkTCPDecodeReader(b *testing.B, config ...*DecodeConfig) {
	benchCases := tcpBenchCases(b.N)
	readers := make([]io.Reader, len(benchCases))
	for i, benchCase := range benchCases {
		readers[i] = bytes.NewReader(benchCase)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := DecodeTCPFromReader(readers[i], config...)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func udpBenchCases(n int) [][]byte {
	cases := []string{
		"005FCAFE0107000F3335323039333038363430333635358E010000016B4F831C680100000000000000000000000000000000010005000100010100010011009D00010010015E2C880002000B000000003544C87A000E000000001DD7E06A000001",
		"003DCAFE0105000F33353230393330383634303336353508010000016B4F815B30010000000000000000000000000000000103021503010101425DBC000001",
		"0086CAFE0101000F3335323039333038353639383230368E0100000167EFA919800200000000000000000000000000000000FC0013000800EF0000F00000150500C80000450200010000710000FC00000900B5000000B600000042305600CD432A00CE6064001100090012FF22001303D1000F0000000200F1000059D90010000000000000000001",
		"0083CAFE0101000F3335323039333038353639383230368E0100000167F1AEEC00000A750E8F1D43443100F800B210000000000012000700EF0000F00000150500C800004501000100007142000900B5000600B6000500422FB300CD432A00CE60640011000700120007001303EC000F0000000200F1000059D90010000000000000000001",
		"01E4CAFE0126000F333532303934303839333937343634080400000163C803B420010A259E1A1D4A057D00DA0128130057421B0A4503F00150051503EF01510052005900BE00C1000AB50008B60005427025CD79D8CE605A5400005500007300005A0000C0000007C700000018F1000059D910002D32C85300000000570000000064000000F7BF000000000000000163C803AC50010A25A9D21D4A01B600DB0128130056421B0A4503F00150051503EF01510052005900BE00C1000AB50008B6000542702ECD79D8CE605A5400005500007300005A0000C0000007C700000017F1000059D910002D32B05300000000570000000064000000F7BF000000000000000163C803A868010A25B5581D49FE5400DB0127130057421B0A4503F00150051503EF01510052005900BE00C1000AB50008B60005427039CD79D8CE605A5400005500007300005A0000C0000007C700000017F1000059D910002D32995300000000570000000064000000F7BF000000000000000163C803A4B2010A25CC861D49F75C00DB0124130058421B0A4503F00150051503EF01510052005900BE00C1000AB50008B6000542703CCD79D8CE605A5400005500007300005A0000C0000007C700000018F1000059D910002D32695300000000570000000064000000F7BF000000000004",
	}

	casesBin := make([][]byte, len(cases))

	for i := range cases {
		bs, _ := hex.DecodeString(cases[i])
		casesBin[i] = bs
	}

	benchCases := make([][]byte, n)

	for i := 0; i < n; i++ {
		benchCases[i] = casesBin[rand.Intn(len(cases))]
	}

	return benchCases
}

func benchmarkUDPDecodeSlice(b *testing.B, config ...*DecodeConfig) {
	benchCases := udpBenchCases(b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := DecodeUDPFromSlice(benchCases[i], config...)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkUDPDecodeReader(b *testing.B, config ...*DecodeConfig) {
	benchCases := udpBenchCases(b.N)
	readers := make([]io.Reader, len(benchCases))
	for i, benchCase := range benchCases {
		readers[i] = bytes.NewReader(benchCase)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := DecodeUDPFromReader(readers[i], config...)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTCPDecode(b *testing.B) {
	benchmarkTCPDecodeSlice(b)
}

func BenchmarkTCPDecodeReader(b *testing.B) {
	benchmarkTCPDecodeReader(b)
}

func BenchmarkUDPDecodeSlice(b *testing.B) {
	benchmarkUDPDecodeSlice(b)
}

func BenchmarkUDPDecodeReader(b *testing.B) {
	benchmarkUDPDecodeReader(b)
}

func BenchmarkTCPDecodeAllocElementsOnReadBuffer(b *testing.B) {
	benchmarkTCPDecodeSlice(b, &DecodeConfig{OnReadBuffer})
}

func BenchmarkTCPDecodeReaderAllocElementsOnReadBuffer(b *testing.B) {
	benchmarkTCPDecodeReader(b, &DecodeConfig{OnReadBuffer})
}

func BenchmarkUDPDecodeSliceAllocElementsOnReadBuffer(b *testing.B) {
	benchmarkUDPDecodeSlice(b, &DecodeConfig{OnReadBuffer})
}

func BenchmarkUDPDecodeReaderAllocElementsOnReadBuffer(b *testing.B) {
	benchmarkUDPDecodeReader(b, &DecodeConfig{OnReadBuffer})
}

func BenchmarkEncodeTCP(b *testing.B) {
	cases := []*Packet{
		{
			CodecID:  Codec12,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getinfo"}},
		},
		{
			CodecID:  Codec12,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getio"}},
		},
		{
			CodecID:  Codec12,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getver"}},
		},
		{
			CodecID:  Codec13,
			Data:     nil,
			Messages: []Message{{Type: TypeResponse, Text: "getver", Timestamp: uint32(time.Now().Unix())}},
		},
		{
			CodecID:  Codec13,
			Data:     nil,
			Messages: []Message{{Type: TypeResponse, Text: "getinfo", Timestamp: uint32(time.Now().Unix())}},
		},
		{
			CodecID:  Codec14,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getver", Imei: "352093081452251"}},
		},
		{
			CodecID:  Codec14,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getio", Imei: "352093081452251"}},
		},
	}

	benchCase := make([]*Packet, b.N)

	for i := 0; i < b.N; i++ {
		benchCase[i] = cases[rand.Intn(len(cases))]
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := EncodePacketTCP(benchCase[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeUDP(b *testing.B) {
	cases := []*Packet{
		{
			CodecID:  Codec12,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getinfo"}},
		},
		{
			CodecID:  Codec12,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getio"}},
		},
		{
			CodecID:  Codec12,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getver"}},
		},
		{
			CodecID:  Codec13,
			Data:     nil,
			Messages: []Message{{Type: TypeResponse, Text: "getver", Timestamp: uint32(time.Now().Unix())}},
		},
		{
			CodecID:  Codec13,
			Data:     nil,
			Messages: []Message{{Type: TypeResponse, Text: "getinfo", Timestamp: uint32(time.Now().Unix())}},
		},
		{
			CodecID:  Codec14,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getver", Imei: "352093081452251"}},
		},
		{
			CodecID:  Codec14,
			Data:     nil,
			Messages: []Message{{Type: TypeCommand, Text: "getio", Imei: "352093081452251"}},
		},
	}

	benchCase := make([]*Packet, b.N)

	for i := 0; i < b.N; i++ {
		benchCase[i] = cases[rand.Intn(len(cases))]
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := EncodePacketUDP("352093081452251", 0, 0, benchCase[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCrc16IBMGenerateLookupTable(b *testing.B) {
	// save results here to avoid result table allocation on the stack
	// and possibly some optimizations.
	results := make([][]uint16, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		results[i%256] = genCrc16IBMLookupTable()
	}
}

func BenchmarkCrc16IBMWithLookupTable(b *testing.B) {
	benchCases := make([][]byte, 256)
	for i := range benchCases {
		benchCases[i] = make([]byte, 1024)
		rand.Read(benchCases[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Crc16IBM(benchCases[i%256])
	}
}

func BenchmarkCrc16IBMWithoutLookupTable(b *testing.B) {
	benchCases := make([][]byte, 256)
	for i := range benchCases {
		benchCases[i] = make([]byte, 1024)
		rand.Read(benchCases[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		crc16IBMNoTable(benchCases[i%256])
	}
}

func crc16IBMNoTable(data []byte) uint16 {
	crc := uint16(0)
	size := len(data)
	for i := 0; i < size; i++ {
		crc ^= uint16(data[i])
		for j := 0; j < 8; j++ {
			if (crc & 0x0001) == 1 {
				crc = (crc >> 1) ^ 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	return crc
}
