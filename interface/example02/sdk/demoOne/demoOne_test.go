package demoOne

import (
	"testing"
)



func TestMain(m *testing.M) {
	m.Run()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("One", testOne)
}

func testOne(t *testing.T)  {

	bd := NewOneManagerImpl()
	_,err:=bd.GetList(5000,1,10)
	if err != nil {
		// %v 是以默认方式打印此值
		t.Errorf("Error : %v", err)
	}

	//t.Logf("list: %v", list)

}


