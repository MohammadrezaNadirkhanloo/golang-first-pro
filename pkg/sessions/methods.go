package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, key, value string) {
	session := sessions.Default(c)

	session.Set(key, value)
	session.Save()

}

func Flash(c *gin.Context, key string) string{
	session := sessions.Default(c)

	respons:=  session.Get(key)
	session.Save()

	session.Delete(key)
	session.Save()

	if respons !=nil{
		return respons.(string)
	}

	return ""
}