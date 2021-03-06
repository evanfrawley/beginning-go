package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMessageEncodeDecode(t *testing.T) {
	var cases = []struct {
		message        message
		expectedResult []byte
	}{
		{ // When all the fields are empty
			message:        message{},
			expectedResult: []byte{0x78, 0x9c, 0xaa, 0x56, 0x2a, 0xa9, 0x2c, 0x48, 0x55, 0xb2, 0x32, 0xd0, 0x51, 0x4a, 0xca, 0x4f, 0xa9, 0x54, 0xb2, 0x52, 0x52, 0xd2, 0x51, 0x2a, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x2d, 0x56, 0xb2, 0xca, 0x2b, 0xcd, 0xc9, 0xa9, 0xe5, 0x2, 0x4, 0x0, 0x0, 0xff, 0xff, 0xe8, 0xed, 0xc, 0x47},
		},
		{ // When the message is a chat
			message: message{
				Type: messageTypeChat,
				Body: "Hello World",
			},
			expectedResult: []byte{0x78, 0x9c, 0xaa, 0x56, 0x2a, 0xa9, 0x2c, 0x48, 0x55, 0xb2, 0x32, 0xd4, 0x51, 0x4a, 0xca, 0x4f, 0xa9, 0x54, 0xb2, 0x52, 0xf2, 0x48, 0xcd, 0xc9, 0xc9, 0x57, 0x8, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0xd2, 0x51, 0x2a, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x2d, 0x56, 0xb2, 0xca, 0x2b, 0xcd, 0xc9, 0xa9, 0xe5, 0x2, 0x4, 0x0, 0x0, 0xff, 0xff, 0x8e, 0xb7, 0x10, 0x64},
		},
		{ // When the message is a username update
			message: message{
				Type: messageTypeUsernames,
				Usernames: map[NodeAddress]string{
					NodeAddress("192.168.0.10:9999"): "Server-1",
					NodeAddress("172.16.0.17:9999"):  "Server-2",
				},
			},
			expectedResult: []byte{0x78, 0x9c, 0x5c, 0xc9, 0x3d, 0xe, 0x85, 0x20, 0xc, 0x7, 0xf0, 0xfd, 0x1d, 0xe3, 0x3f, 0xf3, 0x8, 0x65, 0x10, 0xdb, 0x6b, 0x78, 0x2, 0x8d, 0x1d, 0xfd, 0x48, 0x51, 0x13, 0x42, 0xb8, 0xbb, 0x61, 0x75, 0xfe, 0x55, 0x5c, 0xe5, 0x54, 0x48, 0x74, 0x58, 0x8e, 0xb5, 0x40, 0x0, 0x87, 0x3b, 0xab, 0xed, 0xf3, 0xa6, 0x19, 0x52, 0x41, 0x29, 0x7a, 0x1a, 0x7c, 0xf0, 0x94, 0x84, 0x99, 0x19, 0x82, 0x49, 0xed, 0x51, 0xfb, 0x47, 0x38, 0x10, 0x77, 0x1d, 0x3b, 0x87, 0xf, 0x13, 0x5a, 0xfb, 0xbd, 0x1, 0x0, 0x0, 0xff, 0xff, 0x9, 0x36, 0x19, 0x96},
		},
		{ // When the message is a username request
			message: message{
				Type: messageTypeUsernameReq,
				Body: "192.168.0.32:9876",
			},
			expectedResult: []byte{0x78, 0x9c, 0xaa, 0x56, 0x2a, 0xa9, 0x2c, 0x48, 0x55, 0xb2, 0x32, 0xd6, 0x51, 0x4a, 0xca, 0x4f, 0xa9, 0x54, 0xb2, 0x52, 0x32, 0xb4, 0x34, 0xd2, 0x33, 0x34, 0xb3, 0xd0, 0x33, 0xd0, 0x33, 0x36, 0xb2, 0xb2, 0xb4, 0x30, 0x37, 0x53, 0xd2, 0x51, 0x2a, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x2d, 0x56, 0xb2, 0xca, 0x2b, 0xcd, 0xc9, 0xa9, 0xe5, 0x2, 0x4, 0x0, 0x0, 0xff, 0xff, 0xa8, 0xb6, 0xf, 0xbc},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := c.message.Encode()
			if !reflect.DeepEqual(result, c.expectedResult) {
				t.Fatalf("Encode - Expected %#v but got %#v", c.expectedResult, result)
			}

			var newMsg message
			err := newMsg.Decode(result)
			CheckNoError(t, err)
			if !reflect.DeepEqual(newMsg, c.message) {
				t.Fatalf("Decode - Expected %#v but got %#v", c.message, newMsg)
			}
		})
	}
}
