```


Go 的Test学习

cd xx/dbops

go test -v

➜  dbops git:(master) ✗ go test -v
=== RUN   TestUserWorkFlow
=== RUN   TestUserWorkFlow/Add
=== RUN   TestUserWorkFlow/Get
=== RUN   TestUserWorkFlow/Del
=== RUN   TestUserWorkFlow/Reget
--- PASS: TestUserWorkFlow (0.01s)
    --- PASS: TestUserWorkFlow/Add (0.00s)
    --- PASS: TestUserWorkFlow/Get (0.00s)
        api_test.go:40: GetUser pwd: 123
    --- PASS: TestUserWorkFlow/Del (0.00s)
    --- PASS: TestUserWorkFlow/Reget (0.01s)
PASS
ok  	code.be.staff.com/staff/StaffGo/test/dbops	0.660s


```
