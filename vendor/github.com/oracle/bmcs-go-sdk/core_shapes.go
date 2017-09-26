// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type Shape struct {
	Name string `json:"shape"`
}

// ListShapes contains a list of shapes as well as optional OPCNextPage which
// can be used to pass as the Page field of CoreOptions in subsequent List calls.
// In conjunction with Limit is used in paginating result.
// OPCRequestID is used to identify the request for support issues.
type ListShapes struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Shapes []Shape
}

func (l *ListShapes) GetList() interface{} {
	return &l.Shapes
}

// ListShapes retrieves a list of shapes. compartmentID is a required parameter.
// Additional optional parameters may be assigned and passed in as options.
// These include AvailabilityDomain, Limit and Page for pagination support, and
// an ImageID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Shape/ListShapes
func (c *Client) ListShapes(compartmentID string, opts *ListShapesOptions) (shapes *ListShapes, e error) {
	details := &requestDetails{
		name:     resourceShapes,
		optional: opts,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	shapes = &ListShapes{}
	e = resp.unmarshal(shapes)
	return
}
