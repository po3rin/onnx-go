package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ConvTranspose", "TestConvtransposeDilations", NewTestConvtransposeDilations)
}

// NewTestConvtransposeDilations version: 4.
func NewTestConvtransposeDilations() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ConvTranspose",
		Title:  "TestConvtransposeDilations",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa3, 0x1, 0xa, 0x2c, 0xa, 0x1, 0x58, 0xa, 0x1, 0x57, 0x12, 0x1, 0x59, 0x22, 0xd, 0x43, 0x6f, 0x6e, 0x76, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x2a, 0x12, 0xa, 0x9, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x40, 0x2, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x12, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5a, 0x1b, 0xa, 0x1, 0x58, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x1b, 0xa, 0x1, 0x57, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1b, 0xa, 0x1, 0x59, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"X", "W"},
		     Output:    []string{"Y"},
		     Name:      "",
		     OpType:    "ConvTranspose",
		     Attributes: ([]*pb.AttributeProto) (len=1 cap=1) {
		    (*pb.AttributeProto)(0xc000118d00)(name:"dilations" type:INTS ints:2 ints:2 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 3, 3),
				tensor.WithBacking([]float32{3, 8, 1, 9, 5, 7, 3, 2, 6}),
			),

			tensor.New(
				tensor.WithShape(1, 1, 2, 2),
				tensor.WithBacking([]float32{7, 2, 1, 9}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 5, 5),
				tensor.WithBacking([]float32{21, 56, 13, 16, 2, 63, 35, 67, 10, 14, 24, 22, 76, 76, 21, 9, 5, 88, 45, 63, 3, 2, 33, 18, 54}),
			),
		},
	}
}
