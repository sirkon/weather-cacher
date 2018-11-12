package main

//go:generate protoc --go_out=plugins=grpc:internal schema/forecast.proto schema/rpc.proto
