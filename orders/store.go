package main

import "context"

type store struct {
}

func NewsStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context) error {
	return nil
}
