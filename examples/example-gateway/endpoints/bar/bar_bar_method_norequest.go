// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"
)

// HandleNoRequestRequest handles "/bar/no-request-path".
func HandleNoRequestRequest(
	ctx context.Context,
	inc *zanzibar.IncomingMessage,
	g *zanzibar.Gateway,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}

	// Handle request body.
	clientResp, err := clients.Bar.NoRequest(ctx, h)
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
	response := convertNoRequestClientResponse(&clientRespBody)
	inc.WriteJSON(clientResp.StatusCode, response)
}

func convertNoRequestClientResponse(body *bar.BarResponse) *bar.BarResponse {
	downstreamResponse := bar.BarResponse{}

	return &downstreamResponse
}
