package session

type SessionStore interface {
	Get(string) string
	Set(string, string)
}

type SessionService struct {
	store SessionStore
}

func NewSessionService(s SessionStore) *SessionService {
	return &SessionService{
		store: s,
	}
}

func (s *SessionService) SetSession(user string, session string) {
	s.store.Set(user, session)
}

func (s *SessionService) GetSession(user string) string {
	return s.store.Get(user)
}
