package main

import "fmt"

func main() {
	// f, err := NewFeeType().ValueOf("adult")
	f, err := FeeType.ValueOf("adult")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%s: %d円", f.Label(), f.Yen())
}

type Fee interface {
	Yen() int
	Label() string
}

type adultFee struct{}

func (adultFee) Yen() int {
	return 1500
}

func (adultFee) Label() string {
	return "大人"
}

type childFee struct{}

func (childFee) Yen() int {
	return 1000
}

func (childFee) Label() string {
	return "子供"
}

var FeeType = struct {
	Adult   Fee
	Child   Fee
	ValueOf func(s string) (fee Fee, err error)
}{
	Adult: adultFee{},
	Child: childFee{},
	ValueOf: func(s string) (fee Fee, err error) {
		switch s {
		case "adult":
			fee = adultFee{}
		case "child":
			fee = childFee{}
		default:
			err = fmt.Errorf("未定義の料金タイプ %s", s)
		}
		return
	},
}

func NewFeeType() struct {
	Adult   Fee
	Child   Fee
	ValueOf func(s string) (fee Fee, err error)
} {
	return struct {
		Adult   Fee
		Child   Fee
		ValueOf func(s string) (fee Fee, err error)
	}{
		Adult: adultFee{},
		Child: childFee{},
		ValueOf: func(s string) (fee Fee, err error) {
			switch s {
			case "adult":
				fee = adultFee{}
			case "child":
				fee = childFee{}
			default:
				err = fmt.Errorf("未定義の料金タイプ %s", s)
			}
			return
		},
	}
}
