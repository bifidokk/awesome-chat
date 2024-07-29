package chat

import (
	"context"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.chatRepository.Delete(ctx, id)

	return err
}
