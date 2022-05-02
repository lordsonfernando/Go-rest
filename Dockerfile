##We specify the base image 
## added builder...this image contains executables
FROM golang:1.12.0-alpine3.9 AS builder
## We create an /app directory within our
## image that will hold our application source
## files
## Add git to the image to download git dependencies inside the image
RUN apk update && apk add --no-cache git

RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app
## get gorilla/mux dependency from git
## run go build to compile the binary
## executable of our Go program
RUN go get github.com/gorilla/mux
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable

## Executes build from builder image ...this reduces the size of docker image from 372 Mb to 13.5 Mb
FROM alpine:latest
COPY --from=builder /app .

CMD ["./main"]