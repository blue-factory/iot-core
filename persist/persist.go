package persist

import (
	"errors"
	"sync"
)

// Persist ...
type Persist struct {
	mu sync.RWMutex
	s  map[string]string
	b  map[string]bool
}

// NewPersist ...
func NewPersist() *Persist {
	return &Persist{
		s: make(map[string]string),
		b: make(map[string]bool),
	}
}

// SetString ...
func (p *Persist) SetString(key string, value string) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	_, err := p.GetString(key)
	if err == nil {
		return errors.New(key + " persist string already exist")
	}

	p.s[key] = value
	return nil
}

// GetString ...
func (p *Persist) GetString(key string) (string, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	str, ok := p.s[key]
	if !ok {
		return "nil", errors.New(key + " persist string not found")
	}

	return str, nil
}

// SetBool ...
func (p *Persist) SetBool(key string, value bool) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	_, err := p.GetBool(key)
	if err == nil {
		return errors.New(key + " persist bool already exist")
	}

	p.b[key] = value
	return nil
}

// GetBool ...
func (p *Persist) GetBool(key string) (bool, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	b, ok := p.b[key]
	if !ok {
		return false, errors.New(key + " persist bool not found")
	}

	return b, nil
}
