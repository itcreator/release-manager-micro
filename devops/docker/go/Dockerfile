FROM golang:1.7


#### protobuf protoc
#todo: move protobuf to separate container

RUN apt-get update -y
RUN apt-get install -y autoconf automake libtool curl make g++ unzip

RUN mkdir /src
WORKDIR /src

RUN git clone https://github.com/google/protobuf -b 3.0.x --depth 1  .

RUN ./autogen.sh

RUN ./configure
RUN make
RUN make check
RUN make install
RUN ldconfig

RUN protoc --version


RUN go get -u github.com/micro/protobuf/proto github.com/micro/protobuf/protoc-gen-go


### consul agent

RUN cd /tmp \
    && wget https://releases.hashicorp.com/consul/0.7.0/consul_0.7.0_linux_amd64.zip \
    && unzip consul_0.7.0_linux_amd64.zip \
    && chmod +x consul \
    && mv consul /usr/bin/consul

#install Glide package manager
RUN go get github.com/Masterminds/glide

EXPOSE 8080

