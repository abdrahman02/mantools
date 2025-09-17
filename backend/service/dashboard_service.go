package service

import (
	"backend/repository"
	"context"
)

type DashboardService interface {
	GetDashboardData(ctx context.Context) (map[string]interface{}, error)
}

type dashboardService struct {
	repo repository.AnalyticsRepository
	propertyID string
}

func NewDashboardService(r repository.AnalyticsRepository, propertyID string) DashboardService {
	return &dashboardService{repo: r, propertyID: propertyID}
}

func (s *dashboardService) GetDashboardData(ctx context.Context) (map[string]interface{}, error) {
	visitors, err := s.repo.GetActiveUsers(ctx, s.propertyID)
	if err != nil {	return nil, err	}

	popularPage, err := s.repo.GetPopularPage(ctx, s.propertyID)
	if err != nil { return nil, err	}

	usersByCountry, err := s.repo.GetUsersByCountry(ctx, s.propertyID)
	if err != nil { return nil, err }

	result := map[string]interface{}{
		"visitors": visitors,
		"popularPage": popularPage,
		"usersByCountry": usersByCountry,
	}

	return result, nil
}