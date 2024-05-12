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
	"http3-server-poc/internal/tlsconfig"
)

func newController(
	partsStoringService controller.PartsStoringService,
	logger *zap.Logger,
) *controller.Controller {
	serverRequestMapper := controllermappers.NewServerRequestMapper()
	return controller.NewController(partsStoringService, logger, serverRequestMapper)
}

func newPartProcessingEngine(
	partsHashChan chan string,
	partsRepository services.PartsRepository,
	dataStore services.DataStore,
) *services.PartProcessingEngine {
	return services.NewPartProcessingEngine(partsHashChan, partsRepository, dataStore)
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

	partsRepository := inmemorycache.NewPartsRepository()
	dataStore := filesystem.NewDataStore()
	hashChan := make(chan string)
	partProcessingEngine := newPartProcessingEngine(hashChan, partsRepository, dataStore)
	go partProcessingEngine.StartProcessing()

	partStoringService := services.NewPartStoringService(partsRepository, hashChan)

	controller := newController(partStoringService, logger)

	handler, err := router.GenerateRoutingHandler(controller)
	if err != nil {
		panic(err)
	}

	return newHttp3Server(handler)
}
