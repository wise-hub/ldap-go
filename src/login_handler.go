package src

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LDAPHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := &Login{}

		res := gin.H{
			"result": "Invalid username or password (0)",
		}

		if err := c.ShouldBind(&login); err != nil {
			// bad request
			c.JSON(403, gin.H{"result": err.Error()})
			return
		}

		if len(login.Username) > 30 {
			c.JSON(200, gin.H{"result": "Invalid username or password (1)"})
			return
		}

		userDn, err := LdapAuth(login)

		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"result": "Invalid username or password (2)"})
			return
		}

		res["result"] = "OK"
		res["user_dn"] = userDn

		c.JSON(200, res)
	}
}
