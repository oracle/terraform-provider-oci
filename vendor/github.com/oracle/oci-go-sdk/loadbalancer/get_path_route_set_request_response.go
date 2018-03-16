// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetPathRouteSetRequest wrapper for the GetPathRouteSet operation
type GetPathRouteSetRequest struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the specified load balancer.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the path route set to retrieve.
	// Example: `path-route-set-001`
	PathRouteSetName *string `mandatory:"true" contributesTo:"path" name:"pathRouteSetName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetPathRouteSetRequest) String() string {
	return common.PointerString(request)
}

// GetPathRouteSetResponse wrapper for the GetPathRouteSet operation
type GetPathRouteSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PathRouteSet instance
	PathRouteSet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPathRouteSetResponse) String() string {
	return common.PointerString(response)
}
