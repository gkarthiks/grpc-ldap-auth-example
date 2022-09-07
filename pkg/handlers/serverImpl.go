package handlers

import (
	v1alpha1 "grpc-ldap-auth-example/proto"
)

type SimpleLDAPService struct {
	v1alpha1.UnimplementedSimpleLDAPServiceServer
}
