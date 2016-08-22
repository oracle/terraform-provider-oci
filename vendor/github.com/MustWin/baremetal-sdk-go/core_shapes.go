package baremetal

type Shape struct {
	Name string `json:"shape"`
}

// ShapeList contains a list of shapes as well as optional OPCNextPage which
// can be used to pass as the Page field of CoreOptions in subsequent List calls.
// In conjunction with Limit is used in paginating result.
// OPCRequestID is used to identify the request for support issues.
type ShapeList struct {
	ResourceContainer
	Shapes []Shape
}

func (l *ShapeList) GetList() interface{} {
	return &l.Shapes
}

// ListShapes retrieves a list of shapes. compartmentID is a required parameter.
// Additional optional parameters may be assigned and passed in as options.
// These include AvailabilityDomain, Limit and Page for pagination support, and
// an ImageID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listShapes
func (c *Client) ListShapes(compartmentID string, options ...Options) (shapes *ShapeList, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceShapes,
		ocid:    compartmentID,
		options: options,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	shapes = &ShapeList{}
	e = resp.unmarshal(shapes)
	return
}
