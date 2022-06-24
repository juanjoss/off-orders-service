package docs

import (
	"github.com/juanjoss/off-orders-service/ports"
)

// swagger:route GET /randomProductFromUserSSD get-requests endpoint
// Returns a random product from a random user's device.
// responses:
//   200: randProdFromUserSsdResponse
//	 default: randProdFromUserSsdErrorResponse

// Get random product from user SSD request success.
// swagger:response randProdFromUserSsdResponse
type randProdFromUserSsdResponseWrapper struct {
	// in:body
	Response ports.GetRandomProductFromUserSsdResponse
}

// Get random product from user SSD request error.
// swagger:response randProdFromUserSsdErrorResponse
type randProdFromUserSsdErrorResponseWrapper struct {
	// in:body
	Msg string
}
