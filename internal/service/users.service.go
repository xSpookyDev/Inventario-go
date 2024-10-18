package service

import (
	"context"
	"errors"
	"fmt"
	"inventario-go/encryption"
	"inventario-go/internal/models"

	"github.com/rs/zerolog/log"
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("Error al verificar si el usuario existe: %w", err)
	}
	if u != nil {
		log.Warn().Msgf("Intento de registro fallido: el usuario con email %s ya existe", email)
		return fmt.Errorf("El usuario con el email ya existe: %w", err)
	}
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPass, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPass) != password {
		return nil, errors.New("Invalid Password")
	}
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
