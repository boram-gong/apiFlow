package server

import (
	"github.com/boram-gong/apiFlow/server/handlers"
	"github.com/boram-gong/service"
)

func Run(port, path string) {
	s := service.NewService(port, path)

	s.AddHTTPHandler("GET", "/lgi/adapter/json", handlers.JsonDecorator, DecodeTagJsonReq)
	s.AddHTTPHandler("GET", "/lgi/adapter/rule", handlers.ReadJsonRule, DecodeJsonRuleId)
	s.AddHTTPHandler("POST", "/lgi/adapter/rule", handlers.SaveRule, DecodePostJsonRule)
	s.AddHTTPHandler("PUT", "/lgi/adapter/rule", handlers.SaveRule, DecodePutJsonRule)
	s.AddHTTPHandler("DELETE", "/lgi/adapter/rule", handlers.DeleteRule, DecodeJsonRuleId)
	s.AddHTTPHandler("GET", "/lgi/responseAdapter/re", handlers.ReRule, DecodeNull)

	s.AddHTTPHandler("GET", "/lgi/apiFlow/db", handlers.GetDBClient, DecodeDbName)
	s.AddHTTPHandler("POST", "/lgi/apiFlow/db", handlers.MakeDBClient, DecodePostDbClient)
	s.AddHTTPHandler("PUT", "/lgi/apiFlow/db", handlers.MakeDBClient, DecodePutDbClient)
	s.AddHTTPHandler("DELETE", "/lgi/apiFlow/db", handlers.MakeDBClient, DecodeDeleteDbClient)

	s.AddHTTPHandler("GET", "/lgi/apiFlow/sql-server", handlers.GetApiServer, DecodeNull)
	s.AddHTTPHandler("POST", "/lgi/apiFlow/sql-server", handlers.MakeApiServer, DecodeServerApiReq)
	s.AddHTTPHandler("PUT", "/lgi/apiFlow/sql-server", handlers.ChangeApiServer, DecodeServerApiReq)
	s.AddHTTPHandler("DELETE", "/lgi/apiFlow/sql-server", handlers.DeleteApiServer, DecodeServerApiPathReq)

	s.Run()
}
