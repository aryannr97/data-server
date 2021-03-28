package user

import (
	context "context"
	"log"

	"github.com/aryannr97/data-server/pkg/repository"
)

// Server handles user operations
type Server struct {
	UserStore repository.IUserRepository
}

// AddUser processes user add operation
func (s *Server) AddUser(ctx context.Context, request *UserAddRequest) (*UserAddResponse, error) {
	log.Println("Adding new user")

	userDoc := repository.UserEntity{
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Email:     request.Email,
		Phone:     request.Phone,
		Country:   request.Country,
	}
	id, err := s.UserStore.AddUser(userDoc)
	if err != nil {
		return &UserAddResponse{}, err
	}

	return &UserAddResponse{Id: id}, nil
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {}
