FROM golang:alpine AS builder

WORKDIR /go/src/HushTell
COPY . .
RUN go build -o server

FROM alpine
WORKDIR /HushTell
COPY --from=builder /go/src/HushTell/server /HushTell/server
CMD ["./server"]