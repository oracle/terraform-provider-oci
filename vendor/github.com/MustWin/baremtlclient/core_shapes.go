package baremtlsdk

import (
	"encoding/json"
	"net/url"
)

// ShapeList contains a list of shapes as well as optional OPCNextPage which
// can be used to pass as the Page field of CoreOptions in subsequent List calls.
// In conjunction with Limit is used in paginating result.
// OPCRequestID is used to identify the request for support issues.
type ShapeList struct {
	OPCNextPage  string
	OPCRequestID string
	Shapes       []string
}

// ListShapes retrieves a list of shapes. compartmentID is a required parameter.
// Additional optional parameters may be assigned and passed in as options.
// These include AvailabilityDomain, Limit and Page for pagination support, and
// an ImageID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listShapes
func (c *Client) ListShapes(compartmentID string, options ...CoreOptions) (shapes *ShapeList, e error) {
	query := url.Values{}
	query.Set(queryCompartmentID, compartmentID)
	c.setCoreOptions(query, options...)
	urlStr := buildCoreURL(resourceShapes, query)

	var resp *requestResponse
	if resp, e = c.api.getRequest(urlStr, nil); e != nil {
		return
	}

	shapes = &ShapeList{
		OPCNextPage:  resp.header.Get(headerOPCNextPage),
		OPCRequestID: resp.header.Get(headerOPCRequestID),
	}

	if e = json.Unmarshal(resp.body, &shapes.Shapes); e != nil {
		return
	}

	return
}
