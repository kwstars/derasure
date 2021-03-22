FROM golang:1.15 as build

ENV APPNAME=derasure

ADD . /go/src/${APPNAME}

WORKDIR /go/src/${APPNAME}

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && make build-in-docker


FROM alpine:latest

ENV APPNAME=derasure

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk --update add ca-certificates \
    && apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    #&& apk del tzdata \
    && rm -rf /var/cache/apk/*

WORKDIR /opt/myapp/

COPY --from=build /go/src/${APPNAME}/dist/${APPNAME} app
COPY ./configs/${APPNAME}/config.yaml .
COPY ./public ./public
COPY ./templates ./templates
#COPY ./configs/fonts ./configs/fonts

#VOLUME /var/log/uname
#VOLUME /etc/uname
#ENV CADDYPATH=/var/lib/caddy

EXPOSE 8080

CMD ["/opt/myapp/app -f config.yaml"]
