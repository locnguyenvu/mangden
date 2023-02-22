package user

import (
    "context"
    "errors"

    pb "github.com/locnguyenvu/mdn/pkg/grpc"
)

func (ss *ServiceServer) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserActionResponse, error) {
    u := ss.repository.Get(in.GetId())
    if u == nil {
        return &pb.UserActionResponse{
            Success: false,
            Message: "User not found",
        }, errors.New("User not found")
    }

    u.FirstName = in.GetFirstName()
    u.LastName = in.GetLastName()
    u.Yob = int(in.GetYob())

    err := ss.repository.Save(u)
    if err != nil {
        return nil, errors.New("Something went wrong")
    }
    return &pb.UserActionResponse{
        Success: true,
        Message: "",
    }, nil
}
