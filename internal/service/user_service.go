package service

import (
	"context"
	"time"
	db "user-api/db/sqlc"
	"user-api/internal/logger"
	"user-api/internal/models"
	"user-api/internal/repository"

	"go.uber.org/zap"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	// Adjust if birthday hasn't happened yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// 1. Create User
func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  parsedDob,
	})
	if err != nil {
		logger.Log.Error("Failed to create user in DB", zap.Error(err))
		return models.UserResponse{}, err
	}

	logger.Log.Info("User created successfully", zap.Int32("id", user.ID))

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

// 2. Get User
func (s *UserService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

// 3. List All Users
func (s *UserService) ListUsers(ctx context.Context, page int32, limit int32) ([]models.UserResponse, error) {
    if page < 1 {
        page = 1
    }
    offset := (page - 1) * limit

    arg := db.ListUsersParams{
        Limit:  limit,
        Offset: offset,
    }

    users, err := s.repo.ListUsers(ctx, arg)
    if err != nil {
        logger.Log.Error("Failed to list users", zap.Error(err))
        return nil, err
    }

    var response []models.UserResponse
    for _, user := range users {
        response = append(response, models.UserResponse{
            ID:   user.ID,
            Name: user.Name,
            Dob:  user.Dob.Format("2006-01-02"),
            Age:  calculateAge(user.Dob),
        })
    }

    return response, nil
}

// 4. Update User
func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.CreateUserRequest) (models.UserResponse, error) {
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)

	// UpdateUserParams requires ID, Name, and Dob
	arg := db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  parsedDob,
	}

	user, err := s.repo.UpdateUser(ctx, arg)
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Int32("id", id), zap.Error(err))
		return models.UserResponse{}, err
	}

	logger.Log.Info("User updated", zap.Int32("id", id))

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

// 5. Delete User 
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		logger.Log.Error("Failed to delete user", zap.Int32("id", id), zap.Error(err))
		return err
	}
	logger.Log.Info("User deleted", zap.Int32("id", id))
	return nil
}