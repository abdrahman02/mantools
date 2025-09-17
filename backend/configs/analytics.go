package configs

import (
	"context"
	"log"
	"os"

	analytics "google.golang.org/api/analyticsdata/v1beta"
	"google.golang.org/api/option"
)

var AnalyticsService *analytics.Service

func InitAnalyticsClient() {
	ctx := context.Background()

	service, err := analytics.NewService(ctx, option.WithCredentialsFile(os.Getenv("GCLOUD_CONSOLE_CREDENTIALS")))
	if err != nil { log.Fatal("Failed to create analytics service: ", err) }

	AnalyticsService = service
}