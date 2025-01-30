package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// Base URLs for store links
)

// Base URLs for store links
const (
	androidStoreBaseURL = "https://play.google.com/store/apps/details?id=com.example.app&referrer="
	iosStoreBaseURL     = "https://apps.apple.com/app/id123456789?ref="
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	fmt.Println("request", request)

	userAgent := request.Headers["User-Agent"]
	fmt.Println("agent", userAgent)
	eventID := strings.TrimPrefix(request.Path, "/event")
	fmt.Println("eventID", eventID)
	if eventID == "" {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Missing event ID",
		}, nil
	}

	if strings.Contains(strings.ToLower(userAgent), "android") {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusFound,
			Headers: map[string]string{
				"Location": fmt.Sprintf("%s%s", androidStoreBaseURL, eventID),
			},
		}, nil
	} else if strings.Contains(strings.ToLower(userAgent), "iphone") || strings.Contains(strings.ToLower(userAgent), "ipad") {

		// http.Redirect(w, r, fmt.Sprintf("%s%s", androidStoreBaseURL, eventID), http.StatusFound)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusFound,
			Headers: map[string]string{
				"Location": fmt.Sprintf("%s%s", iosStoreBaseURL, eventID),
			},
		}, nil
	}

	// Default response for unknown devices
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       fmt.Sprintf("Device not supported %s", userAgent),
	}, nil
}

// eventHandler redirects users based on their device type with eventID

func main() {
	lambda.Start(handler)

}
