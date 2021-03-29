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

// GetUser retrives user details for given id
func (s *Server) GetUser(ctx context.Context, request *UserGetReuqest) (*UserGetResponse, error) {
	res, err := s.UserStore.GetUser(request.Id)
	if err != nil {
		return nil, err
	}

	resp := &UserGetResponse{
		Id:        request.Id,
		Firstname: res.FirstName,
		Lastname:  res.LastName,
		Email:     res.Email,
		Phone:     res.Phone,
		Country:   res.Country,
	}

	return resp, nil
}

// UpdateUser updates user details for given id
func (s *Server) UpdateUser(ctx context.Context, request *UserUpdateRequest) (*Empty, error) {
	log.Println("Updating user data")

	userDoc := repository.UserEntity{
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Email:     request.Email,
		Phone:     request.Phone,
		Country:   request.Country,
	}

	return &Empty{}, s.UserStore.UpdateUser(request.Id, userDoc)
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {}
