package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT
type data struct {
	num1Unexported int `json:"Num_1_unexported"`           // 3
	num2Unexported int `json:"num_2_unexported"`           // 4
	num3Unexported int `json:"Num_3_unexported,omitempty"` // 0
	num4Unexported int `json:"num_4_unexported,omitempty"` // 0
	Num5Exported   int `json:"Num_5_exported"`             // 5
	Num6Exported   int `json:"num_6_exported"`             // 6
	Num7Exported   int `json:"Num_7_exported,omitempty"`   // 0
	Num8Exported   int `json:"num_8_exported,omitempty"`   // 0
	Num9Exported   int `json:"num_6_exported"`             // 7
	Num10Exported  int // 8
}

func main() {
	d := data{3, 4, 0, 0, 5, 6, 0, 0, 7, 8}
	b, err := json.Marshal(d)
	fmt.Println(b, err, string(b))

	var d2 data
	err = json.Unmarshal(b, &d2)
	fmt.Printf("%#v", d2)
}

// END OMIT
