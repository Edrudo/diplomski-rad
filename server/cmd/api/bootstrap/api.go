package bootstrap

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/quic-go/quicvarint"
	"go.uber.org/zap"

	"http3-server-poc/cmd/api/config"
	"http3-server-poc/internal/api/controller"
	controllermappers "http3-server-poc/internal/api/controller/mappers"
	"http3-server-poc/internal/api/router"
	"http3-server-poc/internal/domain/services"
	"http3-server-poc/internal/infrastructure/filesystem"
	"http3-server-poc/internal/infrastructure/inmemorycache"
	"http3-server-poc/internal/infrastructure/mysql"
	"http3-server-poc/internal/tlsconfig"
)

func newController(
	partsStoringService controller.PartsStoringService,
	logger *zap.Logger,
) *controller.Controller {
	serverRequestMapper := controllermappers.NewServerRequestMapper()
	return controller.NewController(partsStoringService, logger, serverRequestMapper)
}

func newPartsProcessingEngine(
	dataHashChan chan string,
	partsRepository services.PartsRepository,
	getdataRepository services.GeodataRepository,
	imageStore services.ImageStore,
	jsonStore services.JsonStore,
) *services.PartsProcessingEngine {
	return services.NewPartsProcessingEngine(
		dataHashChan,
		partsRepository,
		getdataRepository,
		imageStore,
		jsonStore,
	)
}

func newHttp3Server(handler http.Handler) http3.Server {
	return http3.Server{
		Addr: fmt.Sprintf(
			"%v:%v",
			config.Cfg.ServerConfig.Http3ServerAddress,
			config.Cfg.ServerConfig.Http3ServerPort,
		),
		TLSConfig: tlsconfig.GetTLSConfig(),
		QuicConfig: &quic.Config{
			HandshakeIdleTimeout: time.Millisecond * time.Duration(config.Cfg.QuicConfig.HandshakeIdleTimeoutMs),
			MaxIdleTimeout:       time.Millisecond * time.Duration(config.Cfg.QuicConfig.MaxIdleTimeoutMs),
			RequireAddressValidation: func(addr net.Addr) bool {
				return false // for now, should whitelist our clients address
			},
			InitialStreamReceiveWindow:     quicvarint.Max,
			MaxStreamReceiveWindow:         quicvarint.Max,
			InitialConnectionReceiveWindow: quicvarint.Max,
			MaxConnectionReceiveWindow:     quicvarint.Max,
			KeepAlivePeriod:                time.Millisecond * time.Duration(config.Cfg.QuicConfig.KeepAlivePeriod),
			DisablePathMTUDiscovery:        false,
		},
		Handler: handler,
	}
}

func Api(logger *zap.Logger) http3.Server {
	defer logger.Sync() // flushes buffer, if any

	mysqlConnection := newMysqlConnection(logger)
	gedataRepository := mysql.NewGeodataRepository(mysqlConnection, logger)
	partsRepository := inmemorycache.NewPartsRepository()
	imageStore := filesystem.NewImageStore()
	jsonStore := filesystem.NewJsonStore()
	dataHashChan := make(chan string)
	partsProcessingEngine := newPartsProcessingEngine(
		dataHashChan,
		partsRepository,
		gedataRepository,
		imageStore,
		jsonStore,
	)
	go partsProcessingEngine.StartProcessing()

	partsStoringService := services.NewPartsStoringService(partsRepository, dataHashChan)

	controller := newController(partsStoringService, logger)

	handler, err := router.GenerateRoutingHandler(controller)
	if err != nil {
		panic(err)
	}

	return newHttp3Server(handler)
}
