// Copyright 2024 Scott Long
//
// SPDX-License-Identifier: MIT

package server

import (
	"context"
	"os"
	"syscall"

	"github.com/smxlong/kit/logger"
	"github.com/smxlong/kit/signalcontext"
	"github.com/smxlong/vaultic/version"
)

// Server is the Vaultic server.
type Server struct {
}

// New returns a new Vaultic server.
func New() (*Server, error) {
	return &Server{}, nil
}

// Run runs the server.
func (s *Server) Run() error {
	l, err := logger.New()
	if err != nil {
		return err
	}
	defer func() {
		_ = l.Sync()
	}()

	l.Infow("welcome to vaultic", "version", version.Version)

	ctx, cancel := signalcontext.WithSignals(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	<-ctx.Done()
	return nil
}
