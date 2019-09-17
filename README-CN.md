[English](README.md) | 简体中文

# Rpc Client for Go
[![Latest Stable Version](https://badge.fury.io/gh/aliyun%2Frpc-client.svg)](https://badge.fury.io/gh/aliyun%2Frpc-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/aliyun/rpc-client)](https://goreportcard.com/report/github.com/aliyun/rpc-client)
[![codecov](https://codecov.io/gh/aliyun/rpc-client/branch/master/graph/badge.svg)](https://codecov.io/gh/aliyun/rpc-client)
[![Travis Build Status](https://travis-ci.org/aliyun/rpc-client.svg?branch=master)](https://travis-ci.org/aliyun/rpc-client)
[![Appveyor Build Status](https://ci.appveyor.com/api/projects/status/6sxnwbriw1gwehx8/branch/master?svg=true)](https://ci.appveyor.com/project/aliyun/rpc-client)

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

Alibaba Cloud Rpc Client for Go 是帮助 GO 开发者管理客户端的工具。
                   
本文将介绍如何获取和使用 Alibaba Cloud Rpc Client for Go。

## 要求
- 请确保你的系统安装了不低于 1.10.x 版本的 Go 环境。

## 安装
使用 `go get` 下载安装

```sh
$ go get -u github.com/aliyun/rpc-client
```

如果你使用 `dep` 来管理你的依赖包，你可以使用以下命令:

```sh
$ dep ensure -add  github.com/aliyun/rpc-client
```

##快速使用
在您开始之前，您需要注册阿里云帐户并获取您的[凭证](https://usercenter.console.aliyun.com/#/manage/ak)。