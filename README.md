English | [简体中文](README-CN.md)

# Alibaba Cloud RPC Client for Go
[![Latest Stable Version](https://badge.fury.io/gh/aliyun%2Frpc-client-go.svg)](https://badge.fury.io/gh/aliyun%2Frpc-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/aliyun/rpc-client-go)](https://goreportcard.com/report/github.com/aliyun/rpc-client-go)
[![codecov](https://codecov.io/gh/aliyun/rpc-client-go/branch/master/graph/badge.svg)](https://codecov.io/gh/aliyun/rpc-client-go)
[![Travis Build Status](https://travis-ci.org/aliyun/rpc-client-go.svg?branch=master)](https://travis-ci.org/aliyun/rpc-client-go)
[![Appveyor Build Status](https://ci.appveyor.com/api/projects/status/6sxnwbriw1gwehx8/branch/master?svg=true)](https://ci.appveyor.com/project/aliyun/rpc-client-go)

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

Alibaba Cloud RPC Client for Go a tool for Go developers to manage client.

This document introduces how to obtain and use Alibaba Cloud RPC Client for Go.

## Requirements
- It's necessary for you to make sure your system have installed a Go environment which is new than 1.10.x.

## Installation
Use `go get` to install:

```sh
$ go get -u github.com/aliyun/rpc-client-go
```

If you use `dep` to manage your dependence, you can use the following command:

```sh
$ dep ensure -add  github.com/aliyun/rpc-client-go
```

## Quick Examples
Before you begin, you need to sign up for an Alibaba Cloud account and retrieve your [Credentials](https://usercenter.console.aliyun.com/#/manage/ak).