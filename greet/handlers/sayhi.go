package handlers

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	v1alpha1 "grpc-ldap-auth-example/proto"
)

func SubmitRequestToGreet(greetToName, authToken string) {
	client := EstablishServerConnection()
	metaDataPairs := metadata.Pairs("authorization", fmt.Sprintf("Basic %s", authToken))
	ctx := metadata.NewOutgoingContext(context.Background(), metaDataPairs)
	sayHiResponse, err := client.SayHi(ctx, &v1alpha1.SayHiRequest{MyName: greetToName})

	if err != nil {
		logrus.Errorln(err)
	} else {
		logrus.Println(sayHiResponse.GetGreetingResponse())
	}

}
