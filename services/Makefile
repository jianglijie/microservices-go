.PHONY: build run clean all

export GOOS= linux
export GOARCH=amd64

compile_date = $(shell date "+%Y%m%d")
repo=jianglj
imageName=demo
tag=`git rev-parse --short HEAD`
imageWholeName=${repo}/${imageName}:${compile_date}-${tag}
imageLastName = ${repo}/${imageName}:latest
port=50051
pes_parent_dir:=$(shell pwd)/$(lastword $(MAKEFILE_LIST))
cur_makefile_path:=$(shell dirname $(pes_parent_dir))
proto_dir = ${cur_makefile_path}/proto
bin_name = markets_service


build:
	protoc -I=${proto_dir} --go_out=plugins=grpc:${proto_dir} ${proto_dir}/hello.proto
	protoc -I=${proto_dir} --go_out=plugins=grpc:${proto_dir} ${proto_dir}/finance.proto
	go build -o ${cur_makefile_path}/${bin_name} ${cur_makefile_path}

	docker build -f ${cur_makefile_path}/Dockerfile -t ${imageWholeName} ${cur_makefile_path}
	docker push ${imageWholeName}
	docker tag ${imageWholeName} ${imageLastName}
	docker push ${imageLastName}
run:
	docker run -d --name ${imageName} -p ${port}:${port} ${imageWholeName}
clean:
	docker rm -f ${imageName}

all: build run
#
#build:
#	protoc -I. --go_out=plugins=grpc:. proto/hello.proto
#	# 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
#	GOOS=linux GOARCH=amd64 go build
#	# 根据当前目录下的 Dockerfile 生成名为的镜像
#	docker build -t services .
#run:
#	docker run -p 50051:50051 services