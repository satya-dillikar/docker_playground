https://tutorialedge.net/golang/go-docker-tutorial/

cd .../golang-docker
export GO111MODULE=off
export GOBIN=~/go/bin
export GOPATH=~/go
go build -o main .
go build -o main ./src
docker build --tag golang-docker .
docker images
docker run -d -p 8080:8081 golang-docker


docker ps
curl http://localhost:8080/hi
curl http://localhost:8080/


docker build -t dsatya6/golang-docker:0.1.0 .