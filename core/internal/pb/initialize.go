package pb

import (
	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.TaggerClient

func Initialize() error {

	conn, err := grpc.NewClient(config.Cfg.PB_TARGET, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client = proto.NewTaggerClient(conn)

	return nil
}
