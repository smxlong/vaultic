package server

// Server is the Vaultic server.
type Server struct {
}

// New returns a new Vaultic server.
func New() (*Server, error) {
	return &Server{}, nil
}

// Run runs the server.
func (s *Server) Run() error {
	return nil
}
