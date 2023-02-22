
package user

import (
    "context"
    "errors"

    pb "github.com/locnguyenvu/mdn/pkg/grpc"
)

func (ss *ServiceServer) Get(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {

    u := ss.repository.Get(in.GetId())
    if u == nil {
        return nil, errors.New("Something went wrong")
    }
    return &pb.GetUserResponse{
            Id: u.ID,
            UserName: u.Username,
            FirstName: u.FirstName,
            LastName: u.LastName,
            Yob: int64(u.Yob),
        }, nil
}
