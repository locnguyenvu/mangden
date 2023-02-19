package user

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mck.co/fuel/internal/user"
	pb "mck.co/fuel/pkg/grpc"
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
