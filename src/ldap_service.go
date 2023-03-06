package src

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func LdapAuth(login *Login) (string, error) {
	// Set up the LDAP connection
	// ldap.forumsys.com:389
	// einstein : password
	l, err := ldap.Dial("tcp", "ldap.forumsys.com:389")
	if err != nil {
		return "", err
	}
	defer l.Close()
	// "dc=example,dc=com""
	// Search for the user's DN
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		//fmt.Sprintf("(cn=%s)", login.Username),
		fmt.Sprintf("(uid=%s)", login.Username),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return "", err
	}

	if len(sr.Entries) != 1 {
		return "", fmt.Errorf("User not found or too many entries returned: %v", len(sr.Entries))
	}

	// Bind as the user
	userDN := sr.Entries[0].DN
	err = l.Bind(userDN, login.Password)

	if err != nil {
		return "", err
	}

	//fmt.Println("auth success")

	return userDN, nil
}
