package sdk

import (
	"code.be.staff.com/staff/StaffGo/interface/example02/sdk/demoTwo"
	"code.be.staff.com/staff/StaffGo/interface/example02/sdk/demoOne"
	"code.be.staff.com/staff/StaffGo/interface/example02/sdk/model"
)


var (
	ys  = demoOne.NewOneManagerImpl()
	tt = demoTwo.NewTwoManagerImpl()

)

type API interface {
	GetList(number,page,count int) (model.Response,error)
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
