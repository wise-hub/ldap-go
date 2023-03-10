package main

import (
	"ldapgo/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/ldap-login", src.LDAPHandler())
	r.Run(":55555")
}
