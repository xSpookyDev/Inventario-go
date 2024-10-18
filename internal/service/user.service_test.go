package service

import (
	"context"
	"inventario-go/encryption"
	"inventario-go/internal/entity"
	"inventario-go/internal/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(&entity.User{Email: "test@exists.com", Password: encryptedPassword}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	s = New(repo)
	code := m.Run()
	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{{
		Name:          "RegisterUser_Success",
		Email:         "test@test.com",
		UserName:      "test",
		Password:      "validPassword",
		ExpectedError: nil,
	},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			s := New(repo)

			err := s.RegisterUser(ctx, tc.Email, tc.Name, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected errror %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUSer_InvalidPassword",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidPassword,
		},
	}
	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, errr := s.LoginUser(ctx, tc.Email, tc.Password)
			if errr != tc.ExpectedError {
				t.Errorf("Expected Error %v, got %v", tc.ExpectedError, errr)
			}
		})
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserId        int
		RoleId        int
		ExpectedError error
	}{
		{
			Name:          "AddUserRole_Success",
			UserId:        1,
			RoleId:        1,
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.AddUserRole(ctx, tc.UserId, tc.RoleId)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
