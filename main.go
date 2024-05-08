package main

import (
	"context"
	"database/sql"
	"dsvm/internal/config"
	"dsvm/internal/handler"
	"dsvm/internal/svc"
	"embed"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	xhttp "github.com/zeromicro/x/http"
	"go/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()
	var cfg config.Config
	conf.MustLoad(*configFile, &cfg)
	logx.MustSetup(cfg.LogxConf)
	ctx := svc.NewServiceContext(cfg)
	logger := logx.WithContext(context.Background())
	go startGoose(logger)
	// 启动 grpc 服务
	go startRpcServer(cfg.RpcServerConf, logger, ctx)
	// 启动 rest 服务
	startRestServer(cfg.RestServerConf, logger, ctx)
}

func startRestServer(cfg rest.RestConf, logger logx.Logger, ctx *svc.ServiceContext) {
	restServer, err := rest.NewServer(cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	httpx.SetErrorHandler(func(err error) (int, any) {
		return http.StatusOK,
			xhttp.BaseResponse[types.Nil]{
				Code: xhttp.BusinessCodeError,
				Msg:  err.Error(),
			}
	})
	handler.RegisterHandlers(restServer, ctx)
	defer restServer.Stop()
	logger.Infof("start rest server at %s:%d...", cfg.Host, cfg.Port)
	restServer.Start()
}

func startRpcServer(cfg zrpc.RpcServerConf, logger logx.Logger, ctx *svc.ServiceContext) {
	s := zrpc.MustNewServer(cfg, func(grpcServer *grpc.Server) {
		if cfg.Mode == service.DevMode || cfg.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logger.Infof("start grpc server at %s...", cfg.ListenOn)
	s.Start()
}

func startGoose(logger logx.Logger) {
	dsn := "root:root123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	goose.SetBaseFS(embedMigrations)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	GooseUpByDbName(logger, db, "mysql", "test")
	GooseUpByDbName(logger, db, "mysql", "test1")
	GooseUpByDbName(logger, db, "mysql", "test2")
}

func GooseUpByDbName(logger logx.Logger, db *sql.DB, dialect, dbName string) {
	_, err := db.Exec(fmt.Sprintf("use %s;", dbName))
	if err != nil {
		logger.Error(err)
	}
	if err := goose.SetDialect(dialect); err != nil {
		logger.Error(err)
		panic(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		logger.Error(err)
	}
}
