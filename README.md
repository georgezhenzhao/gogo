# gogo

It has a few proto-functions.

- configuration by json
- main menu
- logging with different levels
- standalone http server
- senda HTTP Get request and save the response to a logging file
- support multiple threads

## build a exe file

    C:\..\gogo>make build

or

    C:\..\main>go build -o gogo.exe

## run the execute file

Please try different parameters to see what the results are.

### test mode

    C:\..\main>gogo.exe test

### start Http server

    C:\..\main>gogo.exe start

### Http Get request

    C:\..\main>gogo.exe get

### try different parameters or none parameters 

    C:\..\main>gogo.exe
    C:\..\main>gogo.exe params

### Check Pid

    C:\..\main>gogo.exe ps 123

## see configuration

../conf/conf.json

## Links

- [how to run golang multiple files in main package](https://stackoverflow.com/questions/28081486/golang-multiple-files-in-main-package)

- [What is golang good for](https://www.quora.com/What-is-golang-good-for)
