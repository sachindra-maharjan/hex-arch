package main

import (
	"log"
	"os"

	gRPC "github.com/hex-arch/internal/adapters/framework/left/grpc"

	"github.com/hex-arch/internal/adapters/app/api"
	"github.com/hex-arch/internal/adapters/core/arithmetic"
	"github.com/hex-arch/internal/adapters/framework/right/db"
	"github.com/hex-arch/internal/ports"
)

func main() {

	var err error

	var dbaseAdapter ports.DBPort
	var arithAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	arithAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, arithAdapter)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}
