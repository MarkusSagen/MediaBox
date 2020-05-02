



## On Errors
#### packgage ....
Make sure on Mac looks like when running "go env" as: 
```bash
GOROOT=/usr/local/go
ROPATH=/Users/SagenOS/go
```

#### Error XXXXX
Make sure a github repo is created and put it under your $GOPATH, like:
```bash
cd $GOPATH/github.com/<Git-Username>
git clone https://github.com/<Git-Username>/<Repo>.git
cd <Repo>
go mod init github.com/<Git-Username>/<Repo>
```

#### Error GOPATH contains "go.mod" or "go.sum"
Remove go.mod and go.sum and make sure that 
```bash
$GOROOT
$GOPATH
```
Are different paths

