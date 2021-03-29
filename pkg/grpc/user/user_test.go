package user

import (
	context "context"
	"errors"
	"reflect"
	"testing"

	"github.com/aryannr97/data-server/pkg/repository"
	mocks "github.com/aryannr97/data-server/pkg/repository/mocks"
	testifymock "github.com/stretchr/testify/mock"
)

func TestServer_AddUser(t *testing.T) {
	userMockRepo := new(mocks.MockUserRepository)

	userMockRepo.On("AddUser", testifymock.Anything, testifymock.Anything).Return("606157f940032d116cbc7aba", nil).Once()
	userMockRepo.On("AddUser", testifymock.Anything, testifymock.Anything).Return("", errors.New("Add error")).Once()

	type args struct {
		ctx     context.Context
		request *UserAddRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *UserAddResponse
		wantErr bool
	}{
		{
			"Positive case",
			&Server{UserStore: userMockRepo},
			args{
				ctx:     context.Background(),
				request: &UserAddRequest{},
			},
			&UserAddResponse{Id: "606157f940032d116cbc7aba"},
			false,
		},
		{
			"Negative case",
			&Server{UserStore: userMockRepo},
			args{
				ctx:     context.Background(),
				request: &UserAddRequest{},
			},
			&UserAddResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AddUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUser(t *testing.T) {
	userMockRepo := new(mocks.MockUserRepository)
	userDoc := repository.UserDocument{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@mail.com",
		Phone:     "4567893412",
		Country:   "US",
	}

	userMockRepo.On("GetUser", testifymock.Anything, testifymock.Anything).Return(userDoc, nil).Once()
	userMockRepo.On("GetUser", testifymock.Anything, testifymock.Anything).Return(repository.UserDocument{}, errors.New("Get error")).Once()

	type args struct {
		ctx     context.Context
		request *UserGetReuqest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *UserGetResponse
		wantErr bool
	}{
		{
			"Positive case",
			&Server{UserStore: userMockRepo},
			args{
				ctx:     context.Background(),
				request: &UserGetReuqest{Id: "606157f940032d116cbc7aba"},
			},
			&UserGetResponse{
				Id:        "606157f940032d116cbc7aba",
				Firstname: "John",
				Lastname:  "Doe",
				Email:     "john@mail.com",
				Phone:     "4567893412",
				Country:   "US",
			},
			false,
		},
		{
			"Negative case",
			&Server{UserStore: userMockRepo},
			args{
				ctx:     context.Background(),
				request: &UserGetReuqest{Id: "606157f940032d116cbc7aba"},
			},
			&UserGetResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateUser(t *testing.T) {
	userMockRepo := new(mocks.MockUserRepository)

	userMockRepo.On("UpdateUser", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(nil).Once()
	userMockRepo.On("UpdateUser", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(errors.New("Update error")).Once()

	type args struct {
		ctx     context.Context
		request *UserUpdateRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *Empty
		wantErr bool
	}{
		{
			"Positive case",
			&Server{UserStore: userMockRepo},
			args{
				ctx: context.Background(),
				request: &UserUpdateRequest{
					Id:        "606157f940032d116cbc7aba",
					Firstname: "John",
					Lastname:  "Doe",
					Email:     "john@mail.com",
					Phone:     "4567893412",
					Country:   "US",
				},
			},
			&Empty{},
			false,
		},
		{
			"Negative case",
			&Server{UserStore: userMockRepo},
			args{
				ctx: context.Background(),
				request: &UserUpdateRequest{
					Id:        "606157f940032d116cbc7aba",
					Firstname: "John",
					Lastname:  "Doe",
					Email:     "john@mail.com",
					Phone:     "4567893412",
					Country:   "US",
				},
			},
			&Empty{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteUser(t *testing.T) {
	userMockRepo := new(mocks.MockUserRepository)

	userMockRepo.On("DeleteUser", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(nil).Once()
	userMockRepo.On("DeleteUser", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(errors.New("Delete error")).Once()

	type args struct {
		ctx     context.Context
		request *UserDeleteRequest
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *Empty
		wantErr bool
	}{
		{
			"Positive case",
			&Server{UserStore: userMockRepo},
			args{
				ctx:     context.Background(),
				request: &UserDeleteRequest{Id: "606157f940032d116cbc7aba"},
			},
			&Empty{},
			false,
		},
		{
			"Negative case",
			&Server{UserStore: userMockRepo},
			args{
				ctx:     context.Background(),
				request: &UserDeleteRequest{Id: "606157f940032d116cbc7aba"},
			},
			&Empty{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
