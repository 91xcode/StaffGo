package sdk

import (
	"code.be.staff.com/staff/StaffGo/interface/example01/sdk/demoTwo"
	"code.be.staff.com/staff/StaffGo/interface/example01/sdk/demoOne"
)


var (
	ys  = demoOne.NewOne()
	tt = demoTwo.NewTwo()

)

type API interface {
	Get(id int, name string) error
}

func GetAPI(cpID int) API {
	switch {
	case cpID >= 2000:
		return tt
	case cpID >= 1000:
		return ys
	default:
		return nil
	}
}
