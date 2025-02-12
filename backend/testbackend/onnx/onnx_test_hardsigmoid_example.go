package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("HardSigmoid", "TestHardsigmoidExample", NewTestHardsigmoidExample)
}

// NewTestHardsigmoidExample version: 3.
func NewTestHardsigmoidExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "HardSigmoid",
		Title:  "TestHardsigmoidExample",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x72, 0xa, 0x34, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xb, 0x48, 0x61, 0x72, 0x64, 0x53, 0x69, 0x67, 0x6d, 0x6f, 0x69, 0x64, 0x2a, 0xf, 0xa, 0x5, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x15, 0x0, 0x0, 0x0, 0x3f, 0xa0, 0x1, 0x1, 0x2a, 0xe, 0xa, 0x4, 0x62, 0x65, 0x74, 0x61, 0x15, 0x9a, 0x99, 0x19, 0x3f, 0xa0, 0x1, 0x1, 0x12, 0x18, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x73, 0x69, 0x67, 0x6d, 0x6f, 0x69, 0x64, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "HardSigmoid",
		     Attributes: ([]*pb.AttributeProto) (len=2 cap=2) {
		    (*pb.AttributeProto)(0xc00023aa00)(name:"alpha" type:FLOAT f:0.5 ),
		    (*pb.AttributeProto)(0xc00023ab00)(name:"beta" type:FLOAT f:0.6 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-1, 0, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{0.100000024, 0.6, 1}),
			),
		},
	}
}
