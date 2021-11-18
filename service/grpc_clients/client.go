package grpc_client

import (
	"bitbucket.org/udevs/book_service/config"
)

type ServiceManager interface {
	// AccountService() billing_service.AccountServiceClient
}

type GrpcClients struct {
	// accountService billing_service.AccountServiceClient
}

func NewGrpcClients(conf *config.Config) (*GrpcClients, error) {
	// connBillingService, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", conf.BillingServiceHost, conf.BillingServicePort),
	// 	grpc.WithInsecure())
	// if err != nil {
	// 	return nil, err
	// }

	return &GrpcClients{
		// accountService: billing_service.NewAccountServiceClient(connBillingService),
	}, nil
}

// func (g *GrpcClients) AccountService() billing_service.AccountServiceClient {
// 	return g.accountService
// }
