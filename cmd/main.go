package main

//import (
//	"log"
//	_ "github.com/lib/pq"
//)
//
//func main() {
//	cfg := config.Load(".")

//grpcConn, err := grpcPkg.New(cfg)
//if err != nil {
//	log.Fatalf("failed to get grpc connections: %v", err)
//}
//
//apiServer := api.New(&api.RouterOptions{
//	Cfg:        &cfg,
//	GrpcClient: grpcConn,
//})
//
//err = apiServer.Run(cfg.HttpPort)
//if err != nil {
//	log.Fatalf("failed to run server: %v", err)
//}
//}
