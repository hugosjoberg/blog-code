package session

import "github.com/hugosjoberg/interface-struct-bad-example/cache"

type SessionService struct {
	cache *cache.Cache
}

func NewSessionService(c *cache.Cache) *SessionService {
	return &SessionService{
		cache: c,
	}
}

func (s *SessionService) SetSession(user string, session string) {
	s.cache.Set(user, session)
}

func (s *SessionService) GetSession(user string) string {
	return s.cache.Get(user)
}
