package main

import (
	"slack-application/database"
	"slack-application/http_server"
	"slack-application/log"
	"sync"
)

var (
	logger = log.GetNewLogger()
	wg     = &sync.WaitGroup{}
)

func main() {

	logger.Info("slack application started!")

	// setup db connection
	database.SetUpDatabase()

	// run GRPC http_server
	wg.Add(1)
	go http_server.StartServer(wg)
	//grpcServer, lis := runGrpcServer(*port)
	//
	//err := grpcServer.Serve(lis)
	//if err != nil {
	//	logger.Fatal("Unable to listen on service", zap.Int("port", 3000), zap.Error(err))
	//}
	wg.Wait()
}

//func runGrpcServer(port int) (*grpc.Server, net.Listener) {
//	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
//	if err != nil {
//		logger.Fatal("Unable to listen on port", zap.Int("port", port), zap.Error(err))
//	}
//
//	grpcServer := grpc.NewServer()
//
//	return grpcServer, lis
//}
