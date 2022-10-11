package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(context.Context, string) (int, error)
	AverageDestination(context.Context, string) (float32, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float32, error) {

	//TO DO...

	return 1.1, nil
}
