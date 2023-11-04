##  general commands for testing:
```
you can structure the file like 

server.go
server_test.go

(or)

server.go
test
    server_test.go


in both the cases the package for _test.go files should be same
```


##  general commands for testing:
```
go test .  					                 #run all test in the current pkg
go test shiva/cartService                    #run all tests in the given pkg 
go test shiva/billingService
go test -run TestNotify .                    #run only TestNotify which is present in the current pkh
go test -run TestNotify shiva/billingService #run only TestNotify which is present in the given pkg
go test ./...                                #run all the test in current and recursive pkg inside

#you cannot run a (go test service_test.go), you have to run tests for the entire pkg

include -v flag for verbosity
```
##  code coverage
```

cover and cover profile works only if calc.go and calc_test.go are in same folder
go test ./... --cover #cover shows the test coverage percentage
go test ./... -coverprofile=coverage.output #generates coverage report to a file 
go tool cover  -html=coverage.output #to see the report in a html format
```

##  checking for race conditions
```
go build --race main.go
./main.go (will give warning if race condition)

for test files
go test ./... -race (to check if any race condition errors)

```

##  testing main - not IMP
```
func TestMain(m *testing.M) {
	fmt.Println("Hello World")
	ret := m.Run()
	fmt.Println("Tests have executed")
	os.Exit(ret)
}
go test file.go -v
```

## using tags to run specific tests all at once

```
add these lines above package declaration
// +build unit
// +build integration

use the bellow command to run specific tests mentioned in tags flag
go test ./... --tags=unit -v
go test ./... --tags=integration -v
```
