package rest

func (s *Server) Running() bool {
	return s.running
}

// DEFAULTS

func RestServerHostDefault() string {
	return restServerHostDefault
}

func RestServerPortDefault() int {
	return restServerPortDefault
}

func RestHostDefault() string {
	return restHostDefault
}

func RestPortDefault() int {
	return restPortDefault
}
