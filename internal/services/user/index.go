package user

import (
    "context"
    "errors"

    pb "github.com/locnguyenvu/mdn/pkg/grpc"
)

func (ss *ServiceServer) List(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {

    users, err := ss.repository.ListLatest()
    if err != nil {
        return nil, errors.New("Something went wrong")
    }
    var userlist []*pb.GetUserResponse
    for _, elem := range users {
        userlist = append(userlist, &pb.GetUserResponse{
            Id: elem.ID,
            UserName: elem.Username,
            FirstName: elem.FirstName,
            LastName: elem.LastName,
            Yob: int64(elem.Yob),
        })
    }
    return &pb.ListUserResponse{Data: userlist}, nil
}
