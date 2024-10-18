package server

import (
	"context"

	"github.com/frankie-mur/habit_tracker/api"
)

func (s *Server) CreateHabit(
	_ context.Context,
	request *api.CreateHabitRequest,
) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("CreateHabit request received: $s", request)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Name:            request.Name,
			WeeklyFrequency: request.WeeklyFrequency,
		},
	}, nil

}
