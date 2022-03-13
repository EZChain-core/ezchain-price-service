FROM golang:1.14.2 AS build_base
LABEL maintainer="quandc <quandc@vccloud.vn>"

RUN apt-get update && apt-get install -y git pkg-config

# stage 2
from build_base AS build_go

ENV GO111MODULE=on

WORKDIR $GOPATH/src/git.paas.vn/iam/gray_titanic
COPY go.mod .
COPY go.sum .
RUN go mod download
#RUN go mod vendor
# # RUN CGO_ENABLED=0 GOOS=linux go get

# stage 3
FROM build_go AS server_builder

ENV GO111MODULE=on

COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -gcflags="-N -l" -o /bin/gray_titanic ./cmd/*
EXPOSE 8080
CMD ["/bin/gray_titanic"]

# Stage 4
FROM golang:1.14.2 AS gray_titanic

ENV TZ 'Asia/Ho_Chi_Minh'
RUN echo $TZ > /etc/timezone && \
    apt-get update && apt-get install -y tzdata && \
    rm /etc/localtime && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    apt-get clean

COPY --from=server_builder /bin/gray_titanic /bin/gray_titanic
# copy locales translate to go path to load message
COPY --from=server_builder $GOPATH/src/git.paas.vn/iam/gray_titanic/locales /go/locales

CMD ["/bin/gray_titanic"]



