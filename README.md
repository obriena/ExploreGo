# Development setup for Docker
## Docker
This readme file is based on this article:
[Simple Go Setup](https://levelup.gitconnected.com/setup-simple-go-development-environment-with-docker-b8b9c0d4e0a8)

Download lastest docker image of go  
```
docker pull golang
```

Run the image with options:
```
docker run --rm -it --name go-restful -v $PWD:/go/src/github.com/the-evengers/go-restful golang
```  

-v or --volume option allows me to mount a directory from my local machine $PWD (working directory) to a directory within the container /go/src/github.com/the-evengers/go-restful.

### Step up
Goal: Run a docker image that will compile and execute our code on our hard drive with out having to restart it.  Our code will be a server so when Docker executes the code our server will be running serving requests.  We also want all of the changes to be loaded automatically.

There is a file called Local.Dockerfile - this should be built locally
```
docker build -f Local.Dockerfile -t flyingspheres/devgoservice:1.0 .
```

## Development Setup
Init the module:
```
go mod init flyingspheres.com/test
touch test.go
```



