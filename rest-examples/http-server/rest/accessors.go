package rest

func (s *Server) Running() bool {
	return s.running
}

// DEFAULTS

func RestHostDefault() string {
	return restHostDefault
}

func RestPortDefault() int {
	return restPortDefault
}
