FROM golang:alpine AS build-env

WORKDIR $HOME/app

ENV CGO_ENABLED 0

RUN apk add --no-cache git
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN export PATH="$PATH:$GOPATH/bin"

COPY app ./app
COPY conf ./conf
COPY messages ./messages
COPY public ./public
COPY tests ./tests
COPY go.mod .

WORKDIR $HOME/

RUN revel build -m prod -a app -t prod_build

FROM alpine
WORKDIR /
COPY --from=build-env /prod_build /
ENTRYPOINT ./run.sh