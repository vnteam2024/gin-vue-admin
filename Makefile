SHELL = /bin/bash

#SCRIPT_DIR         = $(shell pwd)/etc/script
#Please select golang version
BUILD_IMAGE_SERVER  = golang:1.18
#Please select node version
BUILD_IMAGE_WEB     = node:16
#project name
PROJECT_NAME        = github.com/flipped-aurora/gin-vue-admin/server
#Configuration file directory
CONFIG_FILE         = config.yaml
#Mirror warehouse namespace
IMAGE_NAME          = gva
#Mirror address
REPOSITORY          = registry.cn-hangzhou.aliyuncs.com/${IMAGE_NAME}

ifeq ($(TAGS_OPT),)
TAGS_OPT            = latest
else
endif

ifeq ($(PLUGIN),)
PLUGIN            = email
else
endif

#Container environment front-end and back-end packaged together
build: build-web build-server
	docker run --name build-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-local

#Container environment packaging front end
build-web:
	docker run --name build-web-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_WEB} make build-web-local

#Container environment packaging backend
build-server:
	docker run --name build-server-local --rm -v $(shell pwd):/go/src/${PROJECT_NAME} -w /go/src/${PROJECT_NAME} ${BUILD_IMAGE_SERVER} make build-server-local

#Build web image
build-image-web:
	@cd web/ && docker build -t ${REPOSITORY}/web:${TAGS_OPT} .

#Build server image
build-image-server:
	@cd server/ && docker build -t ${REPOSITORY}/server:${TAGS_OPT} .

#Local environment packaging front-end and back-end
build-local:
	if [ -d "build" ];then rm -rf build; else echo "OK!"; fi \
	&& if [ -f "/.dockerenv" ];then echo "OK!"; else  make build-web-local && make build-server-local; fi \
	&& mkdir build && cp -r web/dist build/ && cp server/server build/ && cp -r server/resource build/resource 

#Local environment packaging front end
build-web-local:
	@cd web/ && if [ -d "dist" ];then rm -rf dist; else echo "OK!"; fi \
	&& yarn config set registry http://mirrors.cloud.tencent.com/npm/ && yarn install && yarn build

#Local environment packaging backend
build-server-local:
	@cd server/ && if [ -f "server" ];then rm -rf server; else echo "OK!"; fi \
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy \
	&& go build -ldflags "-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n') -X main.Version=${TAGS_OPT}" -v

#Package front-end and back-end two-in-one images
image: build 
	docker build -t ${REPOSITORY}/gin-vue-admin:${TAGS_OPT} -f deploy/docker/Dockerfile .

#Early taster version
images: build build-image-web build-image-server
	docker build -t ${REPOSITORY}/all:${TAGS_OPT} -f deploy/docker/Dockerfile .
	
#Plug-in quick packaging: make plugin PLUGIN="This is the plug-in folder name, the default is email"
plugin:
	if [ -d ".plugin" ];then rm -rf .plugin ; else echo "OK!"; fi && mkdir -p .plugin/${PLUGIN}/{server/plugin,web/plugin} \
	&& if [ -d "server/plugin/${PLUGIN}" ];then cp -r server/plugin/${PLUGIN} .plugin/${PLUGIN}/server/plugin/ ; else echo "OK!"; fi \
	&& if [ -d "web/src/plugin/${PLUGIN}" ];then cp -r web/src/plugin/${PLUGIN} .plugin/${PLUGIN}/web/plugin/ ; else echo "OK!"; fi \
	&& cd .plugin && zip -r ${PLUGIN}.zip ${PLUGIN} && mv ${PLUGIN}.zip ../ && cd ..
