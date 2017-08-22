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

### 1. Test mode

    C:\..\main>gogo.exe test

### 2. Start Http server

    C:\..\main>gogo.exe start

### 3. Http Get request

    C:\..\main>gogo.exe get

### 4. Try different parameters or none parameters 

    C:\..\main>gogo.exe
    C:\..\main>gogo.exe params

### 5. Check Pid

    C:\..\main>gogo.exe ps 123

## see configuration

../conf/conf.json

## Links

- [how to run golang multiple files in main package](https://stackoverflow.com/questions/28081486/golang-multiple-files-in-main-package)

- [What is golang good for](https://www.quora.com/What-is-golang-good-for)
