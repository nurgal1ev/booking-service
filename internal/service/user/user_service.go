package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Role      string
	UpdatedAt *time.Time
}

type UserService struct {
	userRepo  *user.UserRepo
	jwtSecret string
}

func NewUserService(userRepo *user.UserRepo, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(u.Password) == "" {
		return errors.New("password is required")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if u.Role == "" {
		u.Role = "guest"
	}
	if u.Role != "guest" && u.Role != "admin" {
		return errors.New("role must be 'guest' or 'admin'")
	}
	return nil
}

func (s *UserService) Create(ctx context.Context, u *User) (*models.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	userEmail, err := s.userRepo.FindByEmail(ctx, u.Email)
	if err != nil {
		return nil, err
	}

	if userEmail != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user = models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Password:  string(hashedPassword),
		Role:      u.Role,
	}

	err = s.userRepo.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) FindByID(ctx context.Context, id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}

	return user, err
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("invalid email")
	}

	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}

	return user, err
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJWT(user.ID, user.Email, s.jwtSecret)
	return token, nil
}
