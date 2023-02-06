FROM golang:1.20.0
#working directory of the app in the container
WORKDIR /usr/src/app
#copy all files from current directory to working directory
COPY . .
#install all packages
RUN go mod tidy
#build the application
RUN go build 
#start the application
CMD ./go_serverless

