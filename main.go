package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc-ldap-auth-example/pkg/handlers"
	"grpc-ldap-auth-example/pkg/interceptors"
	v1alpha1 "grpc-ldap-auth-example/proto"
	"net"
)

var (
	grpcPort = "6000"
)

func main() {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptors.BasicAuthInterceptor,
			),
		),
	)

	logrus.Debugln("registering the CNP Onboard server")
	v1alpha1.RegisterSimpleLDAPServiceServer(grpcServer, &handlers.SimpleLDAPService{})

	// Start gRPC server
	logrus.Debugf("starting the CNP gRPC server on port: %s\n", grpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	logrus.Infof("Starting gRPC server on %s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Errorf("error serving gRPC: %v", err)
		return
	}

}
