package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"universal-oauth/util"
)

const feedly_url = "https://sandbox7.feedly.com"
const clientID = "sandbox"
const clientSecret = "NdyYvHssp6H6c2iTiU6mMaBQQ409pMOy"

/* func Oauth(w http.ResponseWriter, r *http.Request) {
	oauthConfig := util.GetOAuthConfig()
	oauthStateString := "pseudo-random"
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
} */

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	oauthConfig := util.GetOAuthConfig()
	oauthStateString := "pseudo-random"
	url := oauthConfig.AuthCodeURL(oauthStateString)
	return &events.APIGatewayProxyResponse{
		StatusCode: 301,
		MultiValueHeaders: map[string][]string{
			"Location": []string{url},
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
