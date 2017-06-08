FROM golang:1.8


#### protobuf protoc
RUN apt-get update -y
RUN apt-get install -y autoconf automake libtool curl make g++ unzip

RUN mkdir /tmp/protobuf
WORKDIR /tmp/protobuf

RUN git clone https://github.com/google/protobuf -b 3.0.x --depth 1  .

RUN ./autogen.sh

RUN ./configure
RUN make
RUN make check
RUN make install
RUN ldconfig

RUN protoc --version


RUN go get -u github.com/micro/protobuf/proto github.com/micro/protobuf/protoc-gen-go
