package repository

import (
	"context"
	"strconv"
	"strings"

	analytics "google.golang.org/api/analyticsdata/v1beta"
)

type AnalyticsRepository interface {
	GetActiveUsers(ctx context.Context, propertyID string) ([]map[string]string, error)
	GetPopularPage(ctx context.Context, propertyID string) ([]map[string]interface{}, error)
	GetUsersByCountry(ctx context.Context, propertyID string) (map[string]string, error)
}

type analyticsRepository struct {
	service *analytics.Service
}

func NewAnalyticsRepository(s *analytics.Service) AnalyticsRepository {
	return &analyticsRepository{service: s}
}

func (r *analyticsRepository) GetActiveUsers(ctx context.Context, propertyID string) ([]map[string]string, error) {
	req := &analytics.RunReportRequest{
		Dimensions: []*analytics.Dimension{{Name: "date"}},
		Metrics: []*analytics.Metric{{Name: "activeUsers"}},
		DateRanges: []*analytics.DateRange{{StartDate: "7daysAgo", EndDate: "today"}},
	}
	resp, err := r.service.Properties.RunReport("properties/" + propertyID, req).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	var result []map[string]string
	for _, row := range resp.Rows {
		result = append(result, map[string]string{
			"date": row.DimensionValues[0].Value,
			"users": row.MetricValues[0].Value,
		})
	}

	return result, nil
}

func (r *analyticsRepository) GetPopularPage(ctx context.Context, propertyID string) ([]map[string]interface{}, error) {
	req := &analytics.RunReportRequest{
		Dimensions: []*analytics.Dimension{{Name: "pagePath"}},
		Metrics: []*analytics.Metric{{Name: "screenPageViews"}},
		DateRanges: []*analytics.DateRange{{StartDate: "30daysAgo", EndDate: "today"}},
		OrderBys: []*analytics.OrderBy{
			{Desc: true, Metric: &analytics.MetricOrderBy{MetricName: "screenPageViews"}},
		},
		DimensionFilter: &analytics.FilterExpression{
			Filter: &analytics.Filter{
				FieldName: "pagePath",
				InListFilter: &analytics.InListFilter{
					Values: []string{
						"/text-formatter",
						"/text-case-converter",
						"/pdf-tools",
						"/images-compressor",
						"/images-converter",
						"/qr-generator",
						"/api-request-tester",
						"/jwt-decoder",
						"/hash-generator",
					},
				},
			},
		},
	}

	resp, err := r.service.Properties.RunReport("properties/" + propertyID, req).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, row := range resp.Rows {
		views, _ := strconv.Atoi(row.MetricValues[0].Value)
		result = append(result, map[string]interface{}{
			"pagePath": row.DimensionValues[0].Value,
			"pageName": strings.Title(strings.ReplaceAll(row.DimensionValues[0].Value[1:], "-", " ")),
			"views": views,
		})
	}

	return result, nil
}

func (r *analyticsRepository) GetUsersByCountry(ctx context.Context, propertyID string) (map[string]string, error) {
	req := &analytics.RunReportRequest{
		Dimensions: []*analytics.Dimension{{Name: "country"}},
		Metrics: []*analytics.Metric{{Name: "activeUsers"}},
		DateRanges: []*analytics.DateRange{{StartDate: "30daysAgo", EndDate: "today"}},
		OrderBys: []*analytics.OrderBy{
			{Desc: true, Metric: &analytics.MetricOrderBy{MetricName: "activeUsers"}},
		},
		Limit: 5,
	}

	resp, err := r.service.Properties.RunReport("properties/" + propertyID, req).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, row := range resp.Rows {
		result[row.DimensionValues[0].Value] = row.MetricValues[0].Value
	}

	return result, nil
}