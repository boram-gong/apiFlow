package svc

import (
	"context"
	"encoding/json"
	"net/http"

	svc_http "github.com/boram-gong/apiFlow/service/svc/http"

	"github.com/gin-gonic/gin"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshal, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshal.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if head, ok := err.(svc_http.Headerer); ok {
		for k := range head.Headers() {
			w.Header().Set(k, head.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(svc_http.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func MakeHTTPHandler(engine *gin.Engine, endpoints Endpoints) {
	serverOptions := []svc_http.ServerOption{
		svc_http.ServerBefore(headersToContext),
		svc_http.ServerErrorEncoder(errorEncoder),
		svc_http.ServerErrorHandler(svc_http.NewNopErrorHandler()),
		svc_http.ServerAfter(svc_http.SetContentType(contentType)),
	}

	// json-adapter
	engine.Handle("GET", "/lgi/adapter/json", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.JsonDecoratorEndpoint,
			svc_http.WrapS(c, DecodeTagJsonReq),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("GET", "/lgi/adapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ReadRuleEndpoint,
			svc_http.WrapS(c, DecodeJsonRuleId),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("POST", "/lgi/adapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.SaveRuleEndpoint,
			svc_http.WrapS(c, DecodePostJsonRule),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("PUT", "/lgi/adapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.SaveRuleEndpoint,
			svc_http.WrapS(c, DecodePutJsonRule),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("DELETE", "/lgi/adapter/rule", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.DeleteRuleEndpoint,
			svc_http.WrapS(c, DecodeJsonRuleId),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})
	engine.Handle("GET", "/lgi/responseAdapter/re", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ReRuleEndpoint,
			svc_http.WrapS(c, DecodeNull),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	// operation-client
	engine.Handle("GET", "/lgi/apiFlow/operation", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.GetDbClientEndpoint,
			svc_http.WrapS(c, DecodeDbName),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("POST", "/lgi/apiFlow/operation", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ChangeClientEndpoint,
			svc_http.WrapS(c, DecodePostDbClient),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("PUT", "/lgi/apiFlow/operation", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ChangeClientEndpoint,
			svc_http.WrapS(c, DecodePutDbClient),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("DELETE", "/lgi/apiFlow/operation", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ChangeClientEndpoint,
			svc_http.WrapS(c, DecodeDeleteDbClient),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	// sql-server
	engine.Handle("GET", "/lgi/apiFlow/sql-server", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.GetApiServerEndpoint,
			svc_http.WrapS(c, DecodeNull),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("POST", "/lgi/apiFlow/sql-server", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.MakeApiServerEndpoint,
			svc_http.WrapS(c, DecodeServerApiReq),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("PUT", "/lgi/apiFlow/sql-server", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.ChangeApiServerEndpoint,
			svc_http.WrapS(c, DecodeServerApiReq),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("DELETE", "/lgi/apiFlow/sql-server", func(c *gin.Context) {
		svc_http.NewServer(
			endpoints.DeleteApiServerEndpoint,
			svc_http.WrapS(c, DecodeServerApiPathReq),
			EncodeHTTPGenericResponse,
			serverOptions...,
		).ServeHTTP(c.Writer, c.Request)
	})

}
