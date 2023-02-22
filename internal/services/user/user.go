package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"github.com/locnguyenvu/mdn/internal/user"
	pb "github.com/locnguyenvu/mdn/pkg/grpc"
)


type ServiceServer struct {
    pb.UnimplementedUserServiceServer
    
    repository *user.Repository
}


func NewService(db *gorm.DB, logger logrus.FieldLogger) *ServiceServer {
    return &ServiceServer{
        repository: user.NewRepository(db),
    }
}
