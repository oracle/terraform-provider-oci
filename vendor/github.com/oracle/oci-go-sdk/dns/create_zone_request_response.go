// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// CreateZoneRequest wrapper for the CreateZone operation
type CreateZoneRequest struct {

	// Details for creating a new zone.
	CreateZoneDetails `contributesTo:"body"`

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`
}

func (request CreateZoneRequest) String() string {
	return common.PointerString(request)
}

// CreateZoneResponse wrapper for the CreateZone operation
type CreateZoneResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Zone instance
	Zone `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the zone, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response CreateZoneResponse) String() string {
	return common.PointerString(response)
}
