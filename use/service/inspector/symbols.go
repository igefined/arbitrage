package inspector

import "context"

func (s *service) Symbols(ctx context.Context) []string {
	return s.symbols.Symbols(ctx)
}
