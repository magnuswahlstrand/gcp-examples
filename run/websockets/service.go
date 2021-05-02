package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
	"log"
	"sync"
)

type Service struct {
	mu       sync.Mutex
	sessions map[uuid.UUID]*melody.Session
	host     string
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func writeJSON(sess *melody.Session, typ string, v interface{}) {
	b, err := json.MarshalIndent(Message{
		Type:    typ,
		Payload: v,
	},"","  ")
	if err != nil {
		log.Println("error: failed to marshal message", err)
		return
	}
	log.Println(string(b))

	sess.Write(b)
}

func writeJSONToConsumer(sess *melody.Session, typ string, v interface{}) {
	// Forward message
	consumerSess, ok := sess.MustGet("consumer_session").(*melody.Session)
	if !ok {
		sess.CloseWithMsg([]byte("unexpected session id"))
		return
	}
	writeJSON(consumerSess, typ, v)
}

func (s *Service) handleMessage(sess *melody.Session, msg []byte) {
	switch v := sess.MustGet("session_type"); v {
	case sessionTypePublisher:

		//payload := make(map[string]interface{})
		//if err := json.Unmarshal(msg, &v); err != nil {
		//	log.Println("error: failed to unmarshal message", err)
		//	return
		//}
		//log.Println("m",string(msg))
		//log.Println("p",payload)
		writeJSONToConsumer(sess, "update", json.RawMessage(msg))
	default:
		log.Printf("unexpected session type %v\n", v)
		sess.CloseWithMsg([]byte("unexpected connection"))
	}
}

func (s *Service) handleConnect(sess *melody.Session) {
	switch v := sess.MustGet("session_type"); v {
	case sessionTypeConsumer:
		s.handleConsumerConnect(sess)
	case sessionTypePublisher:
		writeJSONToConsumer(sess, "publisher_connected", map[string]interface{}{})
	default:
		log.Printf("unexpected session type %v\n", v)
		sess.CloseWithMsg([]byte("unexpected connection"))
		return
	}

	log.Printf("connected")
}

func (s *Service) handleConsumerConnect(sess *melody.Session) {
	sessionID, ok := sess.MustGet("session_id").(uuid.UUID)
	if !ok {
		sess.CloseWithMsg([]byte("unexpected session id"))
		return
	}

	s.mu.Lock()
	s.sessions[sessionID] = sess
	s.mu.Unlock()

	writeJSON(sess, "connected", map[string]string{
		"url": s.host + "/publisher/init/" + sessionID.String(),
	})
}

func (s *Service) handlePublisherConnect(sess *melody.Session) {
	sessionID, ok := sess.MustGet("session_id").(uuid.UUID)
	if !ok {
		sess.CloseWithMsg([]byte("unexpected session id"))
		return
	}

	s.mu.Lock()
	s.sessions[sessionID] = sess
	s.mu.Unlock()

	writeJSON(sess, "connected", map[string]string{
		"url": s.host + "/publisher/init/" + sessionID.String(),
	})
}

type SessionMetadata struct {
	Type string
	ID   uuid.UUID
	*melody.Session
}

func (s *Service) NewConsumer() *SessionMetadata {
	return &SessionMetadata{
		Type: sessionTypeConsumer,
		ID:   uuid.New(),
	}
}

func (s *Service) NewPublisher(sessionID uuid.UUID) (*SessionMetadata, error) {
	s.mu.Lock()
	sess, found := s.sessions[sessionID]
	s.mu.Unlock()

	if !found {
		return nil, fmt.Errorf("invalid session id: %s", sessionID)
	}

	return &SessionMetadata{
		Type:    sessionTypePublisher,
		ID:      sessionID,
		Session: sess,
	}, nil
}
