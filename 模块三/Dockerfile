From golang:1.17 AS build
WORKDIR /httpserver/
COPY ..
ENV CGO_ENABLED=0
ENV GOLLLMODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

From busybox
COPY --from build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8360
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]