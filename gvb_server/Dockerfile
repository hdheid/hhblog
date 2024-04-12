FROM golang:1.20 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

# Download Go modules
ADD go.mod .
ADD go.sum .
COPY . .
RUN go mod download

#COPY go.mod go.sum ./
#RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
#COPY *.go ./

# Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
RUN GOOS=linux go build -ldflags="-s -w" -o /main main.go


FROM alpine:3.12

WORKDIR /app

COPY --from=builder /main /app/main
RUN chmod +x /main
ENV TZ=Asia/Shanghai
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/main"]