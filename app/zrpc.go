// Code generated by jzero. DO NOT EDIT.

package app

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jzero-io/jzero/app/internal/config"
	"github.com/jzero-io/jzero/app/internal/svc"

	credentialsvr "github.com/jzero-io/jzero/app/internal/server/credential"
	machinesvr "github.com/jzero-io/jzero/app/internal/server/machine"

	"github.com/jzero-io/jzero/app/internal/pb/credentialpb"
	"github.com/jzero-io/jzero/app/internal/pb/machinepb"
)

func getZrpcServer(c config.Config, ctx *svc.ServiceContext) *zrpc.RpcServer {
	s := zrpc.MustNewServer(c.Zrpc, func(grpcServer *grpc.Server) {
	    
		credentialpb.RegisterCredentialServer(grpcServer, credentialsvr.NewCredentialServer(ctx))
		machinepb.RegisterMachineServer(grpcServer, machinesvr.NewMachineServer(ctx))

		if c.Zrpc.Mode == service.DevMode || c.Zrpc.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	return s
}
