package handlers

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1alpha1 "grpc-ldap-auth-example/proto"
	"os"
)

func EstablishServerConnection() (ldapServiceClient v1alpha1.SimpleLDAPServiceClient) {
	logrus.Infoln("establishing a connection to the ldap server over gRPC...")
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", "localhost", "6000"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("error occurred while establishing the gRPC connection with onboarding service: %v", err)
		os.Exit(1)
	}
	ldapServiceClient = v1alpha1.NewSimpleLDAPServiceClient(conn)
	logrus.Infoln("connection established")
	return
}
