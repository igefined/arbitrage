package bundle

import "context"

func (s *service) Clear(ctx context.Context) error {
	return s.bundle.Clear(ctx)
}
