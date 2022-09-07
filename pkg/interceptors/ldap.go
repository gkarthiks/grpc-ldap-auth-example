package interceptors

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ldap.v2"
	"grpc-ldap-auth-example/pkg"
	"time"
)

func isAuthenticated(userName, password string) (bool, error) {
	ldapCon, err := ldap.Dial("tcp", pkg.LdapHost)
	if err != nil {
		errorObj := errors.New(fmt.Sprintf("Error occurred while dialing to LDAP server. Please check on LDAP host config %s. error: %v", pkg.LdapHost, err))
		return false, errorObj
	}
	defer ldapCon.Close()
	err = ldapCon.Bind(pkg.LdapBindDN, pkg.LdapBindPassword)
	if err != nil {
		errorObj := errors.New(fmt.Sprintf("Error occurred while binding to LDAP server %s. Please check on LDAP bind DN and password configs %s and %s. error: %v", pkg.LdapHost, pkg.LdapBindDN, pkg.LdapBindPassword, err))
		return false, errorObj
	}
	searchRequest := ldap.NewSearchRequest(
		pkg.LdapSearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", userName),
		[]string{"dn"}, nil,
	)

	searchResult, err := ldapCon.Search(searchRequest)
	if err != nil {
		errObj := errors.New(fmt.Sprintf("Error occurred while searching for user %s on LDAP server %s. error: %v", userName, pkg.LdapHost, err))
		return false, errObj
	}
	// if no item or more than 1 item is returned, this is not allowed in user authentication
	if len(searchResult.Entries) != 1 {
		logrus.Warnf("User %s does not exist or too many entries returned on LDAP host %s", userName, pkg.LdapHost)
	}

	userDn := searchResult.Entries[0].DN
	err = ldapCon.Bind(userDn, password)
	if err != nil {
		errorObj := errors.New(fmt.Sprintf("LDAP authentication failure for user %s due to incorrect password provided from user %s on LDAP server. error: %v", userName, pkg.LdapHost, err))
		return false, errorObj
	}
	logrus.Infof("User %s has been successfully authenticated at %s.", userName, time.Now().Format(time.RFC3339))
	return true, nil
}
