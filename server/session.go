package server

type Sessions struct {
	id       int
	sessions map[int]*Session
}

type Session struct {
	Sender, Receiver string
	Messages         []Message
}

type Message struct {
	Sender string
	Val    string
}

func (s *Sessions) NewSession(Sender, Receiver string) int {
	// TODO: Concurrent access to s.id and s.sessions. Needs to be Write locked
	// TODO: s.id might overflow
	sessionID := s.id
	s.id++
	s.sessions[sessionID] = &Session{
		Receiver: Receiver,
		Sender:   Sender,
	}
	return sessionID
}

func (s *Sessions) List(user string) []int {
	ids := []int{}
	// TODO: Concurrent access to s.sessions. Needs to be Read locked
	for id, session := range s.sessions {
		if session.Sender == user || session.Receiver == user {
			ids = append(ids, id)
		}
	}
	return ids
}

func (s *Sessions) Get(id int) *Session {
	return s.sessions[id]
}

func (s *Session) AddMessage(m Message) {
	s.Messages = append(s.Messages, m)
}

func (s *Session) GetMessages() []Message {
	return s.Messages
}
