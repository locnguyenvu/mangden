package user

import (
    "context"
    "errors"

    pb "github.com/locnguyenvu/mdn/pkg/grpc"
)

func (ss *ServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserActionResponse, error) {

    _, err := ss.repository.Create(
        in.GetUserName(),
        in.GetPassword(),
        in.GetFirstName(),
        in.GetLastName(),
        int(in.GetYob()),
    )
    if err != nil {
        return nil, errors.New("Something went wrong")
    }
    return &pb.UserActionResponse{
        Success: true,
        Message: "",
    }, nil
}
