FROM golang:1.19-alpine AS build

WORKDIR /go/src/github.com/fyko/stockx-gif/ 
COPY . .
RUN go mod download
RUN go build -o dist/stockx-gif main.go   

FROM alpine:3.18.3  
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/fyko/stockx-gif-next/app
COPY --from=build /go/src/github.com/fyko/stockx-gif/dist/stockx-gif .

ENTRYPOINT [ "./stockx-gif" ]