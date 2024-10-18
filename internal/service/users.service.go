package service

import (
	"context"
	"errors"
	"fmt"
	"inventario-go/encryption"
	"inventario-go/internal/models"

	"github.com/rs/zerolog/log"
)

var (
	ErrUserAlreadyExists = errors.New("El usuario ya existe")
	ErrInvalidPassword   = errors.New("Contraseña inválida")
	ErrRoleAlreadyAdded  = errors.New("El error ya pertenece al usuario")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb)
	log.Info().Msgf("Usuario con email %s se registro perfectamente", email)
	return s.repo.SaveUser(ctx, email, name, pass)

}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error().Err(err).Msgf("Error al buscar usuario con email %s", email)
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		log.Error().Err(err).Msgf("Error al decodificar la contraseña del usuario con email %s", email)
		return nil, fmt.Errorf("error al decodificar la contraseña: %w", err)
	}

	decryptedPass, err := encryption.Decrypt(bb)
	if err != nil {
		log.Error().Err(err).Msgf("Error al desencriptar la contraseña del usuario con email %s", email)
		return nil, fmt.Errorf("error al desencriptar la contraseña: %w", err)
	}

	if string(decryptedPass) != password {
		log.Warn().Msgf("Intento de login fallido: contraseña incorrecta para el usuario con email %s", email)
		return nil, ErrInvalidPassword
	}

	log.Info().Msgf("Usuario con email %s inició sesión exitosamente", email)
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userId, roleId int) error {
	roles, err := s.repo.GetUserRoles(ctx, userId)
	if err != nil {
		return err
	}
	for _, r := range roles {
		if r.RoleID == roleId {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userId, roleId)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return nil
	}
	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true
			break
		}
	}
	if !roleFound {
		return nil
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
