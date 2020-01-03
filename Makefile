.PHONY: build  start push
 
# 版本号
VERSION_TAG = 1.0.0
MILESTONE_TAG = Sep.2019
NAME = sunpengwei/go_common
 
# 源码最后一次提交的commit id，使用commit id可以将docker镜像与提交的代码绑定在一起
COMMIT_ID := $(shell git rev-parse HEAD)
 
# 镜像构建的时间戳
BUILD_TS := $(shell date +'%Y%m%d%H%M%S')
 
# 源码分支
BRANCH_TAG := $(shell git rev-parse --abbrev-ref HEAD)
 
# 镜像版本信息
VERSION := $(VERSION_TAG)-build-$(BRANCH_TAG)-$(BUILD_TS)
 
# 镜像构建信息
BUILD := $(VERSION_TAG)-$(MILESTONE_TAG)-build-$(BUILD_TS)-$(BRANCH_TAG)-$(COMMIT_ID)


build: build-version
 
# 默认读取当前目录下的Dockerfile。 --build-arg：给Dockerfile添加参数，BUILD=$(BUILD)中间不能有空格
build-version:
	docker build --build-arg BUILD=$(BUILD) -t $(NAME):$(VERSION) . >/dev/null
 
# docker image 打tag
tag-latest:
	docker tag $(NAME):$(VERSION) $(NAME):latest >/dev/null
 
# push 到镜像仓库
push: build-version tag-latest
	docker push $(NAME):$(VERSION); docker push $(NAME):latest
 
# 运行镜像
start:
	docker run -it --rm $(NAME):$(VERSION) /bin/bash
