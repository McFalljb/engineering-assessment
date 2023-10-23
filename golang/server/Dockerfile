# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.21-alpine AS go-build

WORKDIR /app

RUN apk add --no-cache gcc musl-dev linux-headers

COPY ./golang/go.mod ./golang/go.sum ./
RUN go mod download

COPY ./golang ./

# compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o foodTruck ./server/main.go

EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD ["./foodTruck"]
