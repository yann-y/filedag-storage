package main

import (
	"encoding/json"
	"errors"
	"github.com/filedag-project/filedag-storage/dag/config"
	"github.com/filedag-project/filedag-storage/dag/pool/poolservice"
	"github.com/filedag-project/filedag-storage/dag/pool/server"
	"github.com/filedag-project/filedag-storage/dag/proto"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
)

const (
	EnvRootUser     = "DAGPOOL_ROOT_USER"
	EnvRootPassword = "DAGPOOL_ROOT_PASSWORD"
)

var log = logging.Logger("pool-main")
var startCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a dag pool process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "listen",
			Usage: "set server listen",
			Value: ":50001",
		},
		&cli.StringFlag{
			Name:  "datadir",
			Usage: "directory to store data in",
			Value: "./dp-data",
		},
		&cli.StringFlag{
			Name:  "config",
			Usage: "set config path",
			Value: "./conf/node_config.json",
		},
		&cli.StringFlag{
			Name:    "root-user",
			Usage:   "set root user",
			EnvVars: []string{EnvRootUser},
			Value:   "dagpool",
		},
		&cli.StringFlag{
			Name:    "root-password",
			Usage:   "set root password",
			EnvVars: []string{EnvRootPassword},
			Value:   "dagpool",
		},
	},
	Action: func(cctx *cli.Context) error {
		cfg, err := loadPoolConfig(cctx)
		if err != nil {
			return err
		}
		startDagPoolServer(cfg)
		return nil
	},
}

func startDagPoolServer(cfg config.PoolConfig) {
	log.Infof("dagpool start...")
	log.Infof("listen %s", cfg.Listen)
	// listen port
	lis, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	// new server
	s := grpc.NewServer()
	service, err := poolservice.NewDagPoolService(cfg)
	if err != nil {
		log.Errorf("NewDagPoolService err:%v", err)
		return
	}
	defer service.Close()

	proto.RegisterDagPoolServer(s, &server.DagPoolService{DagPool: service})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Errorf("failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	s.GracefulStop()

	log.Info("Server exit")
}

func loadPoolConfig(cctx *cli.Context) (config.PoolConfig, error) {
	var cfg config.PoolConfig
	cfg.Listen = cctx.String("listen")
	datadir := cctx.String("datadir")
	if err := os.MkdirAll(datadir, 0777); err != nil {
		return config.PoolConfig{}, err
	}
	cfg.LeveldbPath = path.Join(datadir, "leveldb")
	cfg.RootUser = cctx.String("root-user")
	if cfg.RootUser == "" {
		return config.PoolConfig{}, errors.New("root param is invalid")
	}
	cfg.RootPassword = cctx.String("root-password")
	nodeConfigPath := cctx.String("config")

	var nodeConfigs []config.DagNodeConfig
	for _, path := range strings.Split(nodeConfigPath, ",") {
		var nc config.DagNodeConfig
		file, err := ioutil.ReadFile(path)
		if err != nil {
			log.Errorf("ReadFile err:%v", err)
			return config.PoolConfig{}, err
		}
		err = json.Unmarshal(file, &nc)
		if err != nil {
			log.Errorf("Unmarshal err:%v", err)
			return config.PoolConfig{}, err
		}
		nodeConfigs = append(nodeConfigs, nc)
	}
	cfg.DagNodeConfig = nodeConfigs
	return cfg, nil
}
