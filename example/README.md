# AWS Util Example

A trivial example application to showcase the [awsutil](https://github.com/go-nacelle/awsutil) library.

## Overview

This example application uses DynamoDB to provide a simple string get/set API over HTTP. The **main** function boots [nacelle](https://github.com/go-nacelle/nacelle) with a initializer that creates a DynamoDB client and a server initializer for the process provided by the [httpbase](https://github.com/go-nacelle) library. The client created by the former is injected into the later.

## Building and Running

Simply run `docker-compose up`. This will compile the example application via a multi-stage build and start a container for the API as well as a container for [localstack](https://github.com/localstack/localstack). The provided docker compose environment sets the environment for the application to direct all DynamoDB requests to localstack.

After localstack has started, you must create the DynamoDB table. All requests made to the API before this is done will result in an internal server error. Simply run teh following bash script.

```bash
./setup.sh
```

## Usage

```bash
$ curl -i http://localhost:5000/example-key
HTTP/1.1 404 Not Found
Date: Mon, 24 Jun 2019 19:15:30 GMT
Content-Length: 0
```

```bash
$ curl -i http://localhost:5000/example-key -X POST -d 'payload'
HTTP/1.1 200 OK
Date: Mon, 24 Jun 2019 19:15:35 GMT
Content-Length: 0
```

```bash
$ curl -i http://localhost:5000/example-key
HTTP/1.1 200 OK
Date: Mon, 24 Jun 2019 19:15:40 GMT
Content-Length: 7
Content-Type: text/plain; charset=utf-8

payload
```
