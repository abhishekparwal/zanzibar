/*
 * CODE GENERATED AUTOMATICALLY
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package bar

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/clients/bar"
	"github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"
)

// HandleNormalRequest handles "/bar-path".
func HandleNormalRequest(
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
	var body NormalHTTPRequest
	if ok := inc.UnmarshalBody(&body, rawBody); !ok {
		return
	}
	clientRequest := convertToNormalClientRequest(&body)
	clientResp, err := clients.Bar.Normal(ctx, clientRequest, h)
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
	b, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		inc.SendError(500, errors.Wrap(err, "could not read client response body:"))
		return
	}
	var clientRespBody bar.BarResponse
	if err := clientRespBody.UnmarshalJSON(b); err != nil {
		inc.SendError(500, errors.Wrap(err, "could not unmarshal client response body:"))
		return
	}
	response := convertNormalClientResponse(&clientRespBody)
	inc.WriteJSON(clientResp.StatusCode, response)
}

func convertToNormalClientRequest(body *NormalHTTPRequest) *barClient.NormalHTTPRequest {
	// TODO: Add request fields mapping here.
	return &barClient.NormalHTTPRequest{}
}
func convertNormalClientResponse(body *bar.BarResponse) *bar.BarResponse {
	// TODO: Add response fields mapping here.
	return &bar.BarResponse{}
}