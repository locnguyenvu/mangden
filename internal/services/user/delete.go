package user

import (
    "context"
    "errors"

    pb "mck.co/fuel/pkg/grpc"
)

func (ss *ServiceServer) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*pb.UserActionResponse, error) {
    u := ss.repository.Get(in.GetId())
    if u == nil {
        return &pb.UserActionResponse{
            Success: false,
            Message: "User not found",
        }, errors.New("User not found")
    }

    err := ss.repository.Delete(u.ID)
    if err != nil {
        return nil, errors.New("Something went wrong")
    }
    return &pb.UserActionResponse{
        Success: true,
        Message: "",
    }, nil
}
