FROM golang:1.16.0-alpine as builder
LABEL maintainer="Ben Selby <benmatselby@gmail.com>"

ENV APPNAME hagen
ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go
ENV GO111MODULE on

COPY . /go/src/github.com/benmatselby/${APPNAME}

RUN apk update && \
	apk add --no-cache --virtual .build-deps \
	ca-certificates \
	gcc \
	libc-dev \
	libgcc \
	git \
	curl \
	make

RUN cd /go/src/github.com/benmatselby/${APPNAME} && \
	go install github.com/golang/mock/mockgen && \
	go install github.com/securego/gosec/cmd/gosec && \
	go get -u golang.org/x/lint/golint && \
	make static-all  && \
	mv ${APPNAME} /usr/bin/${APPNAME}  && \
	apk del .build-deps  && \
	rm -rf /go

FROM scratch

COPY --from=builder /usr/bin/${APPNAME} /usr/bin/${APPNAME}
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENV HOME /root

ENTRYPOINT [ "hagen" ]
CMD [ "--help" ]
