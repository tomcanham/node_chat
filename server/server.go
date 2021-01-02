package main

import (
	"log"
	"sync"
)

// Server controls the list of channels and maps clients to channels.
// It also processes client commands.
type Server struct {
	channels map[string]*Channel
	mu       sync.Mutex
}

var _serverInstance *Server
var once sync.Once

func getServer() *Server {
	once.Do(func() {
		_serverInstance = &Server{
			channels: make(map[string]*Channel),
		}
	})

	return _serverInstance
}

func (s *Server) getChannel(name string) *Channel {
	s.mu.Lock()
	defer s.mu.Unlock()

	var ch *Channel
	var ok bool

	if ch, ok = s.channels[name]; !ok {
		log.Printf("Creating new channel: %s\n", name)

		ch = newChannel(name)
		s.channels[name] = ch
		go ch.run()
	}

	return ch
}
