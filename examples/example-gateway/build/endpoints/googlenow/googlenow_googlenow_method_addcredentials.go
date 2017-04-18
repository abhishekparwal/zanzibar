// Code generated by zanzibar
// @generated

package googlenow

import (
	"context"
	"net/http"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/googlenow"
)

// AddCredentialsHandler is the handler for "/googlenow/add-credentials"
type AddCredentialsHandler struct {
	Clients *clients.Clients
}

// NewAddCredentialsEndpoint creates a handler
func NewAddCredentialsEndpoint(
	gateway *zanzibar.Gateway,
) *AddCredentialsHandler {
	return &AddCredentialsHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/googlenow/add-credentials".
func (handler *AddCredentialsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}
	var requestBody AddCredentialsHTTPRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := AddCredentialsEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	respHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	} // TODO(sindelar): implement check headers on response

	res.WriteJSONBytes(202, respHeaders, nil)
}

// AddCredentialsEndpoint calls thrift client GoogleNow.AddCredentials
type AddCredentialsEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w AddCredentialsEndpoint) Handle(
	ctx context.Context,
	// TODO(sindelar): Switch to zanzibar.Headers when tchannel
	// generation is implemented.
	headers http.Header,
	r *AddCredentialsHTTPRequest,
) (map[string]string, error) {
	clientRequest := convertToAddCredentialsClientRequest(r)

	clientHeaders := map[string]string{}
	for k, v := range map[string]string{"X-Uuid": "X-Uuid", "X-Token": "X-Token"} {
		clientHeaders[v] = headers.Get(k)
	}

	respHeaders, err := w.Clients.GoogleNow.AddCredentials(
		ctx, clientHeaders, clientRequest,
	)
	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		// TODO(sindelar): Consider returning partial headers in error case.
		return nil, err
	}

	// Filter and map response headers from client to server response.
	endRespHead := map[string]string{}
	for k, v := range map[string]string{"X-Token": "X-Token", "X-Uuid": "X-Uuid"} {
		endRespHead[v] = respHeaders[k]
	}

	return endRespHead, nil
}

func convertToAddCredentialsClientRequest(body *AddCredentialsHTTPRequest) *googlenowClient.AddCredentialsHTTPRequest {
	clientRequest := &googlenowClient.AddCredentialsHTTPRequest{}

	clientRequest.AuthCode = string(body.AuthCode)

	return clientRequest
}
