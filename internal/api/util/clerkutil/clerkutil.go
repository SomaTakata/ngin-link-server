package clerkutil

import (
	"fmt"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func NewClerkClient() clerk.Client {
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	client, _ := clerk.NewClient(clerkSecretKey)
	return client
}

func GetClerkID(ctx *gin.Context, client clerk.Client) (string, error) {
	sessionToken := ctx.GetHeader("Authorization")
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

	// verify the session
	sessClaims, err := client.VerifyToken(sessionToken)
	if err != nil {
		return "", err
	}

	// get the user, and say welcome!
	user, err := client.Users().Read(sessClaims.Claims.Subject)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", sessClaims.Claims)
	return sessClaims.Claims.Subject, nil
}
