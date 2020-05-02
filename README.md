


# Start
init
```bash
git clone https://github.com/MarkusSagen/MediaBox.git
cd MediaBox
```
#### Run server
```bash
cd server
docker build -t mediabox .
docker run -d -p 8000:8000 mediabox
```
or
```bash
cd server
go build server.go
./MediaBox
```

#### Run volume
```bash
mkdir -p ~/logs/mediabox
docker build -t mediabox-volume -f Dockerfile.volume .
docker run -d -p 8000:8000 -v ~/logs/mediabox:/app/logs mediabox-volume
tail -200f ~/logs/mediabox/app.log 	# Read server log 
```

#### Run frontend
```bash
cd frontend
npm install && npm start
```

### Test server
```bash
curl http://localhost:8000/api\?name\=Markus
curl http://localhost:8000/api?name=Markus
```



## Install GO dependencies

## Install Node dependencies


## Install Python dependencies
```
python3 -m pip install --upgrade pip
pip3 install virtualenv
virtualenv venv
source venv/bin/activate
python3 -m pip install --upgrade pip
pip3 install -r requirements.txt
```




## Testing application
```bash
curl http://localhost:8000?name=Markus
```

### Listing and deleting container/images
```bash
docker images ls 
docker images rm <id>
docker container ls
docker container stop <id>
```




## On Errors
#### package XXX is not in GOROOT
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

