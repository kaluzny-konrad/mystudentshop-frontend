package main

import "google.golang.org/grpc"

type frontendServer struct {
	productCatalogSvcAddr string
	productCatalogSvcConn *grpc.ClientConn
}
