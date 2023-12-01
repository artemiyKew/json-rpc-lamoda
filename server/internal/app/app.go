package app

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/artemiyKew/json-rpc-lamoda/config"
	"github.com/artemiyKew/json-rpc-lamoda/internal/delivery"
	"github.com/artemiyKew/json-rpc-lamoda/internal/repo"
	"github.com/artemiyKew/json-rpc-lamoda/internal/repo/postgres"
	"github.com/artemiyKew/json-rpc-lamoda/internal/service"
	"github.com/sirupsen/logrus"
)

func Run(configPath string) {
	ctx, cancel := context.WithCancel(context.Background())
	// Init config
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// Logger
	SetLogrus("")

	// Postgres
	logrus.Info("Initializing postgres...")
	db, err := postgres.NewDB(cfg.DataBaseURL)
	if err != nil {
		logrus.Fatal(err)
	}
	pg := postgres.New(db)
	defer pg.Close()

	// Repositories
	logrus.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	// Services dependencies
	logrus.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)

	// Routes
	logrus.Info("Initializing routes...")
	delivery := delivery.NewRoutes(ctx, services)

	// rpc server
	warehouseRoutes := delivery.WarehouseRoutes
	productRoutes := delivery.ProductRoutes
	// shippingRoutes := delivery.ShippingRoutes
	if err := rpc.Register(warehouseRoutes); err != nil {
		logrus.Fatal(err)
	}
	if err := rpc.Register(productRoutes); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Tcp listen port %s...", cfg.BindAddr)
	listener, err := net.Listen("tcp", cfg.BindAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		gracefulShutDown(conn, cancel)
	}

}

func gracefulShutDown(conn net.Conn, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errMsg := fmt.Sprintf("Received shutdown signal %v. Shutdown Done", sig)
	if err := conn.Close(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(errMsg)
	cancel()
}
