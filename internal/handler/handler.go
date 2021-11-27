package handler

import (
	"errors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/tingpo/pupgobackend/internal/graph"
	"github.com/tingpo/pupgobackend/internal/graph/generated"
	"gorm.io/gorm"
	"net/http"
)

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

func GraphQLHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
func SignInHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signinPayload SingInPayload
		// Deserialize payload
		err := c.ShouldBindJSON(&signinPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// verify google Id token
		switch signinPayload.Type {
		case "google":
			_, err = verifyGoogleToken(signinPayload.Token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		case "facebook":
			err = verifyFBToken(signinPayload.Token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		// create or signin
		tmpUser, err := sqlCnter.FindUserByEmail(signinPayload.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tmpUser = sqlCnter.CreateUser(signinPayload.Account, signinPayload.Email)
		}

		accessToken, err := CreateJWT(signinPayload.Account, signinPayload.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Write to DB
		userToken, err := sqlCnter.FindTokenByID(tmpUser.Id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = sqlCnter.CreateUserToken(tmpUser.Id, accessToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
		} else {
			userToken.Token = accessToken
			err = sqlCnter.UpdateTokenByID(tmpUser.Id, userToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, userToken)

	}
}
func SignUpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signinPayload SingInPayload
		// Deserialize payload
		err := c.ShouldBindJSON(&signinPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// verify google Id token
		switch signinPayload.Type {
		case "google":
			_, err = verifyGoogleToken(signinPayload.Token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		case "facebook":
			err = verifyFBToken(signinPayload.Token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		// create or signin
		tmpUser, err := sqlCnter.FindUserByEmail(signinPayload.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tmpUser = sqlCnter.CreateUser(signinPayload.Account, signinPayload.Email)
		}

		accessToken, err := CreateJWT(signinPayload.Account, signinPayload.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Write to DB
		userToken, err := sqlCnter.FindTokenByID(tmpUser.Id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = sqlCnter.CreateUserToken(tmpUser.Id, accessToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
		} else {
			userToken.Token = accessToken
			err = sqlCnter.UpdateTokenByID(tmpUser.Id, userToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, userToken)

	}
}
