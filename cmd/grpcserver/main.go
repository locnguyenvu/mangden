package main


import (
    "os"
    "fmt"
    "net"

    "github.com/sirupsen/logrus"
    "go.uber.org/dig"
	"google.golang.org/grpc"

    "github.com/locnguyenvu/mdn/internal/services/user"
    "github.com/locnguyenvu/mdn/pkg/database/mysql"
    "github.com/locnguyenvu/mdn/pkg/logging"
    pb "github.com/locnguyenvu/mdn/pkg/grpc"
)

var (
    serverAddress = "0.0.0.0:50051"
)

func main() {
    c := dig.New()

    if val, exists := os.LookupEnv("GRPC_ADDR"); exists {
        serverAddress = val
    }

    constructors := []interface{}{
        logging.NewFromEnv,
        mysql.NewGormFromEnv,
        user.NewService,
    }

    for _, constructor := range constructors {
        if err := c.Provide(constructor); err != nil {
            fmt.Println("Failed to bootstrap", err.Error())
            os.Exit(1)
        }
    }

    c.Invoke(func(userService *user.ServiceServer, logger logrus.FieldLogger) {

        lis, err := net.Listen("tcp", serverAddress)
        if err != nil {
            logger.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterUserServiceServer(s, userService)
        logger.Infof("server listening at %v", lis.Addr())
        if err := s.Serve(lis); err != nil {
            logger.Fatalf("failed to serve: %v", err)
        }

    })
}


