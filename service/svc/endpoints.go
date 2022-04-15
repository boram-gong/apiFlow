package svc

import (
	"github.com/boram-gong/apiFlow/service/svc/endpoint"
)

type Endpoints struct {
	// adapter
	JsonDecoratorEndpoint endpoint.Endpoint
	SaveRuleEndpoint      endpoint.Endpoint
	ReRuleEndpoint        endpoint.Endpoint
	DeleteRuleEndpoint    endpoint.Endpoint
	ReadRuleEndpoint      endpoint.Endpoint

	// operation-client
	GetDbClientEndpoint  endpoint.Endpoint
	ChangeClientEndpoint endpoint.Endpoint

	// api_server-flow
	GetApiServerEndpoint    endpoint.Endpoint
	MakeApiServerEndpoint   endpoint.Endpoint
	ChangeApiServerEndpoint endpoint.Endpoint
	DeleteApiServerEndpoint endpoint.Endpoint
}
