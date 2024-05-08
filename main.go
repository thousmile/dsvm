package main

import (
	"context"
	"dsvm/internal/handler"
	"dsvm/internal/svc"
	"flag"
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

	"dsvm/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()
	var cfg config.Config
	conf.MustLoad(*configFile, &cfg)
	logx.MustSetup(cfg.LogxConf)
	ctx := svc.NewServiceContext(cfg)
	logger := logx.WithContext(context.Background())
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
