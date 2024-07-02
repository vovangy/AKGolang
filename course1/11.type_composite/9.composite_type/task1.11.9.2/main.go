package main

import "fmt"

type TV interface {
	switchOFF()
	GetModel()
	GetStatus()
	switchOn()
}

type LgTV struct {
	status bool
	model  string
}

type SamsungTV struct {
	status bool
	model  string
}

func (tv SamsungTV) GetStatus() bool {
	return tv.status
}

func (tv SamsungTV) GetModel() string {
	return tv.model
}

func (tv *SamsungTV) switchOFF() bool {
	tv.status = false

	return true
}

func (tv *SamsungTV) switchOn() bool {
	tv.status = true

	return true
}

func (tv LgTV) GetStatus() bool {
	return tv.status
}

func (tv LgTV) GetModel() string {
	return tv.model
}

func (tv *LgTV) switchOFF() bool {
	tv.status = false

	return true
}

func (tv *LgTV) switchOn() bool {
	tv.status = true

	return true
}

func (tv LgTV) LGHub() string {
	return "LGHub"
}

func (tv SamsungTV) SamsungHub() string {
	return "SamsungHub"
}

func main() {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tv.GetStatus())
	fmt.Println(tv.GetModel())
	fmt.Println(tv.SamsungHub())
	fmt.Println(tv.GetStatus())
	fmt.Println(tv.switchOn())
	fmt.Println(tv.GetStatus())
	fmt.Println(tv.switchOFF())
	fmt.Println(tv.GetStatus())

}
