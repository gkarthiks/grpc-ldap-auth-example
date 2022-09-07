package handlers

import (
	"context"
	"fmt"
	v1alpha1 "grpc-ldap-auth-example/proto"
)

func (s *SimpleLDAPService) SayHi(ctx context.Context, request *v1alpha1.SayHiRequest) (*v1alpha1.SayHiResponse, error) {
	requestString := request.GetMyName()
	var greetTo string
	if len(requestString) > 0 {
		greetTo = requestString
	} else {
		greetTo = "stranger"
	}
	return &v1alpha1.SayHiResponse{GreetingResponse: fmt.Sprintf("Hello there %s !", greetTo)}, nil
}
