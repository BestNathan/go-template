FROM bestnathan/golang:1.16.3-alpine3.13-builder as build

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io

WORKDIR /app
COPY ./go.mod ./go.sum ./

RUN go mod download -x

COPY . .

RUN sh ./build.sh

FROM harbor.sensoro.com/library/alpine:3
WORKDIR /app
COPY --from=build /app/build build/
COPY --from=build /app/configs configs/
CMD [ "./build/lins-push" ]