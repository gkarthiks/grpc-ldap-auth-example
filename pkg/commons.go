package pkg

var (
	LdapHost         = "ldap.server.name"
	LdapBindDN       = "ldap-bind-dn"
	LdapBindPassword = "ldap-search-password"
	LdapSearchDN     = "ldap - search - domain"
	RequestingUser   string
	IsLdapAuth       bool
	StdUser          = "testuser"
	StdPassword      = "secret-password-from-vault"
)
