package defaults

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Example struct {
	ProductId int `json:"ProductId" default:"5"`
	Price     int `json:"Price" default:"5"`
}

type Example_2 struct {
	ProductId *int `json:"ProductId" default:"5"`
	Price     *int `json:"Price" default:"5"`
}

func TestTMP(t *testing.T) {
	jsonData := `{"ProductId":0}`
	example := &Example{}
	json.Unmarshal([]byte(jsonData), &example)
	fmt.Printf("%+v", example) //&{ProductId:5 Price:0}

	SetDefaults(example)
	fmt.Printf("%+v", example) //&{ProductId:5 Price:5}
}

func TestPtr(t *testing.T) {
	jsonData := `{"ProductId":0}`
	example := &Example_2{}
	json.Unmarshal([]byte(jsonData), &example)
	SetDefaults(example)
	fmt.Printf("%v %v", *example.ProductId, *example.Price) // 0 5
}
