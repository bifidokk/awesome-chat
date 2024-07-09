package converter

import (
	"github.com/bifidokk/awesome-chat/auth/internal/model"
	desc "github.com/bifidokk/awesome-chat/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToCreateUserFromCreateRequest converts a CreateRequest from the gRPC layer to a CreateUser model for the business logic layer.
func ToCreateUserFromCreateRequest(req *desc.CreateRequest) *model.CreateUser {
	return &model.CreateUser{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		Role:            req.Role.String(),
		ConfirmPassword: req.PasswordConfirm,
	}
}

// ToGetUserResponseFromUser converts a User model from the business logic layer to a GetResponse for the gRPC layer.
func ToGetUserResponseFromUser(user *model.User) *desc.GetResponse {
	var createdAt *timestamppb.Timestamp
	var updatedAt *timestamppb.Timestamp

	createdAt = timestamppb.New(user.CreatedAt)

	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(desc.Role_value[user.Role]),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// ToUpdateUserFromUpdateRequest converts an UpdateRequest from the gRPC layer to an UpdateUser model for the business logic layer.
func ToUpdateUserFromUpdateRequest(req *desc.UpdateRequest) *model.UpdateUser {
	var name string
	var email string

	if req.Name != nil {
		name = req.Name.Value
	}

	if req.Email != nil {
		email = req.Email.Value
	}

	return &model.UpdateUser{
		ID:    req.Id,
		Name:  name,
		Email: email,
	}
}
