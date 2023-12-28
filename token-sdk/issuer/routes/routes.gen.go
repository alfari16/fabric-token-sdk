// Package routes provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package routes

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Amount The amount to issue, transfer or redeem.
type Amount struct {
	// Code the code of the token
	Code string `json:"code"`

	// Value value in base units (usually cents)
	Value int64 `json:"value"`
}

// Counterparty The counterparty in a Transfer or Issuance transaction.
type Counterparty struct {
	Account string `json:"account"`

	// Node The node that holds the recipient account
	Node string `json:"node"`
}

// Error defines model for Error.
type Error struct {
	// Message High level error message
	Message string `json:"message"`

	// Payload Details about the error
	Payload string `json:"payload"`
}

// TransferRequest Instructions to issue or transfer tokens to an account
type TransferRequest struct {
	// Amount The amount to issue, transfer or redeem.
	Amount Amount `json:"amount"`

	// Counterparty The counterparty in a Transfer or Issuance transaction.
	Counterparty Counterparty `json:"counterparty"`

	// Message optional message that will be sent and stored with the transfer transaction
	Message *string `json:"message,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse = Error

// HealthSuccess defines model for HealthSuccess.
type HealthSuccess struct {
	// Message ok
	Message string `json:"message"`
}

// IssueSuccess defines model for IssueSuccess.
type IssueSuccess struct {
	Message string `json:"message"`

	// Payload Transaction id
	Payload string `json:"payload"`
}

// IssueJSONRequestBody defines body for Issue for application/json ContentType.
type IssueJSONRequestBody = TransferRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /healthz)
	Healthz(ctx echo.Context) error
	// Issue tokens to an account
	// (POST /issuer/issue)
	Issue(ctx echo.Context) error
	// Harvest tokens to an account
	// (POST /harvest)
	Harvest(ctx echo.Context) error
	// RequestKabayan tokens to an account
	// (POST /request-kabayan)
	RequestKabayan(ctx echo.Context) error

	// (GET /readyz)
	Readyz(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Healthz converts echo context to params.
func (w *ServerInterfaceWrapper) Healthz(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Healthz(ctx)
	return err
}

// Issue converts echo context to params.
func (w *ServerInterfaceWrapper) Issue(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Issue(ctx)
	return err
}

// Issue converts echo context to params.
func (w *ServerInterfaceWrapper) Harvest(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Harvest(ctx)
	return err
}

// Issue converts echo context to params.
func (w *ServerInterfaceWrapper) RequestKabayan(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RequestKabayan(ctx)
	return err
}

// Readyz converts echo context to params.
func (w *ServerInterfaceWrapper) Readyz(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Readyz(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/healthz", wrapper.Healthz)
	router.POST(baseURL+"/issuer/issue", wrapper.Issue)
	router.POST(baseURL+"/harvest", wrapper.Harvest)
	router.POST(baseURL+"/request-kabayan", wrapper.RequestKabayan)
	router.GET(baseURL+"/readyz", wrapper.Readyz)

}

type ErrorResponseJSONResponse Error

type HealthSuccessJSONResponse struct {
	// Message ok
	Message string `json:"message"`
}

type IssueSuccessJSONResponse struct {
	Message string `json:"message"`

	// Payload Transaction id
	Payload string `json:"payload"`
}

type HealthzRequestObject struct {
}

type HealthzResponseObject interface {
	VisitHealthzResponse(w http.ResponseWriter) error
}

type Healthz200JSONResponse struct{ HealthSuccessJSONResponse }

func (response Healthz200JSONResponse) VisitHealthzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type Healthz503JSONResponse struct{ ErrorResponseJSONResponse }

func (response Healthz503JSONResponse) VisitHealthzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(response)
}

type IssueRequestObject struct {
	Body *IssueJSONRequestBody
}

type IssueResponseObject interface {
	VisitIssueResponse(w http.ResponseWriter) error
}

type Issue200JSONResponse struct{ IssueSuccessJSONResponse }

func (response Issue200JSONResponse) VisitIssueResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type IssuedefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response IssuedefaultJSONResponse) VisitIssueResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type ReadyzRequestObject struct {
}

type ReadyzResponseObject interface {
	VisitReadyzResponse(w http.ResponseWriter) error
}

type Readyz200JSONResponse struct{ HealthSuccessJSONResponse }

func (response Readyz200JSONResponse) VisitReadyzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type Readyz503JSONResponse struct{ ErrorResponseJSONResponse }

func (response Readyz503JSONResponse) VisitReadyzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /healthz)
	Healthz(ctx context.Context, request HealthzRequestObject) (HealthzResponseObject, error)
	// Issue tokens to an account
	// (POST /issuer/issue)
	Issue(ctx context.Context, request IssueRequestObject) (IssueResponseObject, error)

	// (GET /readyz)
	Readyz(ctx context.Context, request ReadyzRequestObject) (ReadyzResponseObject, error)
}

type StrictHandlerFunc = runtime.StrictEchoHandlerFunc
type StrictMiddlewareFunc = runtime.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// Healthz operation middleware
func (sh *strictHandler) Healthz(ctx echo.Context) error {
	var request HealthzRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Healthz(ctx.Request().Context(), request.(HealthzRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Healthz")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(HealthzResponseObject); ok {
		return validResponse.VisitHealthzResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Issue operation middleware
func (sh *strictHandler) Issue(ctx echo.Context) error {
	var request IssueRequestObject

	var body IssueJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Issue(ctx.Request().Context(), request.(IssueRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Issue")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(IssueResponseObject); ok {
		return validResponse.VisitIssueResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) Harvest(ctx echo.Context) error {
	var request IssueRequestObject

	var body IssueJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body
	request.Body.Amount.Code = "IDR"

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Issue(ctx.Request().Context(), request.(IssueRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Issue")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(IssueResponseObject); ok {
		return validResponse.VisitIssueResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) RequestKabayan(ctx echo.Context) error {
	var request IssueRequestObject

	var body IssueJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body
	request.Body.Amount.Code = "KBY"

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Issue(ctx.Request().Context(), request.(IssueRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Issue")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(IssueResponseObject); ok {
		return validResponse.VisitIssueResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Readyz operation middleware
func (sh *strictHandler) Readyz(ctx echo.Context) error {
	var request ReadyzRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Readyz(ctx.Request().Context(), request.(ReadyzRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Readyz")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ReadyzResponseObject); ok {
		return validResponse.VisitReadyzResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWQW/jNhP9K8R832EXICIlaQtUt912gQS9FNkUKLrNYUyNLW4oUktSTt1A/70gKcu2",
	"LDvxNkCzJ1sSOXxv5s3jPIIwdWM0ae+geARLrjHaUXz4YK2xN/2b8EIY7Un78BebRkmBXhqdfXZGh3dO",
	"VFRj+Pd/S3Mo4H/ZJnqWvrosRoWu6ziU5ISVTQgCRTqOrRFAx+GKUPnqYysEOXcSAPoL60ZF0DU5hwuC",
	"Asx9CNpY05D1MnEcvj6O0Jh74OBXTdjovJV6AQGypS+ttFRC8WnYezcsNLPPJPwUuZ7EDr1r51r6GnYH",
	"KYzwcmhwpQyW+/RuLWqHIjwxWT6b6ibiKaSNZTQqbsd7PpHCu9q0ifgIZUUM4zfmDZMhX5z5AH1OlsWA",
	"JVF9Bny74sKUAdeH325+Bw5LVC1BcZ7ne8VPC8Ohc2yV3+zZReErYmEpM3MW/ntzT3o/ZcNRYxbxNZOa",
	"zdARa7X0jr1pXYtKrZgIzfEWOMyNrTFgkNr/8B1wqKWWdVtDkQ9HSe1pQXavPJHI+vz9ynD4KeSQbIPW",
	"r6bTLLZWBKzIbrfyHKSKWlBKftLNKOsoRCoioJIiwNGpDuZBkz3fb71hw4Ru9VCZMc7whfkKPauMKl0s",
	"iCUhG0nas3XMp/SsU8LWy6dSlnzqkJckPe+3RQEneMyVXFRM0ZIUG8d7fif/TB6lcgxnpvUxHTHWi7Q0",
	"h7UIbuhLS26iRa+187aNgnBDkwbJDG0a2yV+Q71VoJEYBgc4dnf0PtFxECNBH9u1I/6OH/H8+AfVug5J",
	"Zw9SKTYj5qLAdMmcN5ZK9iB9lfxgYLppjifTv0OAr/lPu6rUczOR+WSHoQ22TDEATK7YJ/6MvUdxTyWb",
	"rRiyUgY8s9ZTyRSVC7L8T91YcmSXUi9YY+USxYq1Ljz9QdawX7R5iEvZr9aYuQt976UPHQG38YhgPWRd",
	"gnV+lockm4Y0NhIKuDzLzy6jzHwVa51hW0pvbNaLwWWPsuziFUc2BILi05hsvwU4tFZBAZX3TZFlyghU",
	"lXG++DHP8wwbmS3PM+juOn7gmGyrSO7lz6ziyPJ3CLygqOeg8niRXwd3uOq/891R6yLPD6l4WJftjkMd",
	"h+/zy6d37U5xQU4eF4HuBpmDgN21dY12BQXckG+tduwiz5lMd17UhyAmHUsUYydlsdtt+omTiXETpKNS",
	"IemfnH9vytWLjZNjg5oYQ6Ytat7qct+VNi3qbUvd15RpZ6yLYPrp4vRKHdZmSvwRaZ7vSHO7uhHfIVde",
	"a6OPf/cfgAi6ijPD0LhPtGkYN6LvzczsCJiLbTB8HEWgVcbFMCXqI2Eu93p+F+xzzOwVIs7SpfGKge86",
	"VLzh3sxaq9/2MoJDzE5w/NdYmPXV/o2U5nZq7jO+CrPJVodbwnJ1+J68SZ+/4WsyEozshaDGM4FKuSdc",
	"/cSJg/87Q+avTEN9wsfx3vU3Ny5RKpwpikkdMqWxpq3U7eOZ3D9kqt/ePz9zd2xT9oBKkXebIPH1RIyP",
	"vSrS5NSP6FhKHQS62b3RWXfX/RMAAP//uoh3hJoTAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
