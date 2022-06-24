package docs

import "github.com/juanjoss/off-orders-service/ports"

// swagger:route GET /random get-requests endpoint
// Returns a random product.
// responses:
//   200: randomProductResponse
//	 default: randomProductErrorResponse

// Get random product request success.
// swagger:response randomProductResponse
type randomProductResponseWrapper struct {
	// in:body
	Product ports.GetRandomProductResponse
}

// Get random product request error.
// swagger:response randomProductErrorResponse
type randomProductErrorResponseWrapper struct {
	// in:body
	Msg string
}
