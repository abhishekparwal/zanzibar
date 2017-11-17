// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarArgWithQueryParamsHandler is the handler for "/bar/argWithQueryParams"
type BarArgWithQueryParamsHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarArgWithQueryParamsHandler creates a handler
func NewBarArgWithQueryParamsHandler(
	g *zanzibar.Gateway,
	deps *module.Dependencies,
) *BarArgWithQueryParamsHandler {
	handler := &BarArgWithQueryParamsHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "argWithQueryParams",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarArgWithQueryParamsHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/bar/argWithQueryParams",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/argWithQueryParams".
func (h *BarArgWithQueryParamsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBarBar.Bar_ArgWithQueryParams_Args

	nameOk := req.CheckQueryValue("name")
	if !nameOk {
		return
	}
	nameQuery, ok := req.GetQueryValue("name")
	if !ok {
		return
	}
	requestBody.Name = nameQuery

	userUUIDOk := req.HasQueryValue("userUUID")
	if userUUIDOk {
		userUUIDQuery, ok := req.GetQueryValue("userUUID")
		if !ok {
			return
		}
		requestBody.UserUUID = ptr.String(userUUIDQuery)
	}

	workflow := ArgWithQueryParamsEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		default:
			req.Logger.Warn("Workflow for endpoint returned error", zap.Error(errValue))
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// ArgWithQueryParamsEndpoint calls thrift client Bar.ArgWithQueryParams
type ArgWithQueryParamsEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w ArgWithQueryParamsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithQueryParams_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithQueryParamsClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.ArgWithQueryParams(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request", zap.Error(errValue))
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertArgWithQueryParamsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithQueryParamsClientRequest(in *endpointsBarBar.Bar_ArgWithQueryParams_Args) *clientsBarBar.Bar_ArgWithQueryParams_Args {
	out := &clientsBarBar.Bar_ArgWithQueryParams_Args{}

	out.Name = string(in.Name)
	out.UserUUID = (*string)(in.UserUUID)

	return out
}

func convertArgWithQueryParamsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key5, value6 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key5] = int32(value6)
	}
	out.BinaryField = []byte(in.BinaryField)

	return out
}
