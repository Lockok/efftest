package service

import (
	"context"
	"fmt"
)

func (s *subscriptionService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidSubscriptionID
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete subscription: %w", err)
	}

	return nil
}