package docs

import (
	"github.com/juanjoss/off-orders-service/ports"
)

// swagger:route GET / get-requests endpoint
// Returns all products.
// responses:
//   200: allProductsResponse
//	 default: allProductsErrorResponse

// Get all products request success.
// swagger:response allProductsResponse
type allProductsResponseWrapper struct {
	// in:body
	Products ports.GetAllProductsResponse
}

// Get all products request error.
// swagger:response allProductsErrorResponse
type allProductsErrorResponseWrapper struct {
	// in:body
	Msg string
}
