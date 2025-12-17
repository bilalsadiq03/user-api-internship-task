package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/bilalsadiq03/user-api-internship-task/internal/repository"
	"github.com/bilalsadiq03/user-api-internship-task/internal/service"
)


type UserHandler struct {
	repo *repository.UserRepository
	validate *validator.Validate
	logger  *zap.Logger
}


func NewUserHandler(
	repo *repository.UserRepository,
	logger *zap.Logger,
) *UserHandler {
	return &UserHandler{
		repo:     repo,
		validate: validator.New(),
		logger:   logger,
	}
}


// Craete User
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid dob format")
	}

	user, err := h.repo.Create(c.Context(), req.Name, dob)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	h.logger.Info("user created", zap.Int32("id", user.ID))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	})
}



// GET User By ID
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.repo.GetByID(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  age,
	})
}


//Update User
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.ErrBadRequest
	} 

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil{
		return fiber.ErrBadRequest
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} 

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid dob format")
	}

	user, err := h.repo.UpdateByID(c.Context(), int32(id), req.Name, dob)
	if err != nil {
		
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	})
}

// Delete User By ID
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.repo.DeleteByID(c.Context(), int32(id)); err != nil {
		return fiber.ErrNotFound
	}

	return c.SendStatus(fiber.StatusNoContent)
}
