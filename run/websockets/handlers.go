package main

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
	"log"
	"net/http"
)

type WSHandler struct {
	m *melody.Melody

	service *Service
}

func NewWebsocketHandler(host string) *WSHandler {
	service := &Service{
		sessions: map[uuid.UUID]*melody.Session{},
		host:     host,
	}

	m := melody.New()
	m.HandleConnect(service.handleConnect)
	m.HandleMessage(service.handleMessage)
	return &WSHandler{
		m: m,
		service: service,
	}
}

func (re *WSHandler) handleInitConsumer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	session := re.service.NewConsumer()
	re.m.HandleRequestWithKeys(w, r, map[string]interface{}{
		"session_type": session.Type,
		"session_id":   session.ID,
	})
}

func (re *WSHandler) handleInitPublisher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sessionID, err := uuid.Parse(chi.URLParam(r, "session_id"))
	if err != nil {
		log.Printf("failed to parse sessionID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := re.service.NewPublisher(sessionID)
	if err != nil {
		log.Printf("failed to create new publisher: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	re.m.HandleRequestWithKeys(w, r, map[string]interface{}{
		"session_type":     session.Type,
		"session_id":       session.ID,
		"consumer_session": session.Session,
	})
}
