// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/clients/bar"
)

// HandleArgNotStructRequest handles "/bar/arg-not-struct-path".
func HandleArgNotStructRequest(
	ctx context.Context,
	inc *zanzibar.IncomingMessage,
	g *zanzibar.Gateway,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}

	// Handle request body.
	rawBody, ok := inc.ReadAll()
	if !ok {
		return
	}
	var body ArgNotStructHTTPRequest
	if ok := inc.UnmarshalBody(&body, rawBody); !ok {
		return
	}
	clientRequest := convertToArgNotStructClientRequest(&body)
	clientResp, err := clients.Bar.ArgNotStruct(ctx, clientRequest, h)
	if err != nil {
		g.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		inc.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	// Handle client respnse.
	if !inc.IsOKResponse(clientResp.StatusCode, []int{200}) {
		g.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	inc.WriteJSONBytes(clientResp.StatusCode, nil)
}

func convertToArgNotStructClientRequest(body *ArgNotStructHTTPRequest) *barClient.ArgNotStructHTTPRequest {
	clientRequest := barClient.ArgNotStructHTTPRequest{}

	clientRequest.Request = body.Request

	return &clientRequest
}
