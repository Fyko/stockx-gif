FROM golang:1.16-alpine AS build

WORKDIR /go/src/github.com/fyko/stockx-gif-next/ 
COPY . .
RUN go mod download
RUN go build -o build/server cmd/server/main.go   

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/fyko/stockx-gif-next/app
COPY --from=build /go/src/github.com/fyko/stockx-gif-next/build/server .

CMD ["./server"]