# s3 upload micro service

This project is an upload and download micro service of Amazon S3 and KS3.

[![Build Status](https://api.travis-ci.org/lxlxw/s3-micro.svg?branch=master)](https://travis-ci.org/lxlxw/s3-micro)
[![GitHub release](https://img.shields.io/badge/releases-v1.0.1-brightgreen.svg)](https://github.com/lxlxw/s3-micro/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/lxlxw/s3-micro)](https://goreportcard.com/report/github.com/lxlxw/s3-micro)
[![GoDoc](https://godoc.org/github.com/lxlxw/s3-micro?status.svg)](https://godoc.org/github.com/lxlxw/s3-micro)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[ [English](https://github.com/lxlxw/micro-s3-grpc)
| [中文](https://github.com/lxlxw/micro-s3-grpc)
 ]

## Table of Contents
- [Installation](#installation)
- [Configuration](#configuration)
- [Build](#build)
    - [Run Grpc Server](#run-grpc-server)
    - [Run Http Server](#run-http-server)
- [Usage](#usage)
    - [RESTful API examples](#restful-api-examples)
    - [Grpc examples](#all-usage-examples)
- [Features](#features)
- [Swagger](#swagger)
- [Support and Feedback](#support-and-feedback)
- [Contact](#contact)
- [License](#license)


## Installation
Use go get to retrieve the project to add it to your GOPATH workspace, or project's Go module dependencies.

```bash
go get github.com/lxlxw/micro-s3-grpc
```

To update the project use go get -u to retrieve the latest version of the project.

```bash
go get -u github.com/lxlxw/micro-s3-grpc
```

## Configuration
```bash
cat ./conf/s3.toml
```
```bash
[S3]
accesskey = "AKLTRPycUPrRSDOP492EPQO6Bw"
secretkey = "xxxx"
region = ""
endpoint = "ks3testb.s3.amazonaws.com"
```

## Build

```bash
make build
```

### Run Grpc Server

```bash
make server
```

### Run Http Server

```bash
make http
```

## Usage

### RESTful API examples

```bash
curl -X POST -k http://localhost:8088/api/object/upload -d '{"bucketname": "test_bucket", "key":"test/test.txt", "filecontent":"xxxxxx"}'
```
You find more detailed api documentation at /doc.

### Grpc examples

create a service client, make a request, handle the error, and process the response.

```go
package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/lxlxw/micro-s3-grpc/rpc"
)

func main() {

	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// new client
	c := pb.NewStoreApiServiceClient(conn)

	// get method
	req := &pb.CreateBucketRequest{Store: "ks3", Bucketname: "test-bucket"}
	res, err := c.CreateBucket(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println(res.Msg)
}
```
## Features

* Grpc
* RESTful API
* Swagger UI
* Middleware
* App configurable
* Logging
* JWT Authorization

## Swagger

## Support and Feedback

If you find a bug, please submit the issue in Github directly.
[Micro-S3-Grpc Issues](https://github.com/lxlxw/micro-s3-grpc/issues)

## Contact

- Email：<x@xwlin.com>

## License

Micro-S3-Grpc is based on the MIT protocol.

<http://www.opensource.org/licenses/MIT>

