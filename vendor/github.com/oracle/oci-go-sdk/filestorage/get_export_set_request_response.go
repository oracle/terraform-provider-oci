// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetExportSetRequest wrapper for the GetExportSet operation
type GetExportSetRequest struct {

	// The OCID of the export set.
	ExportSetId *string `mandatory:"true" contributesTo:"path" name:"exportSetId"`
}

func (request GetExportSetRequest) String() string {
	return common.PointerString(request)
}

// GetExportSetResponse wrapper for the GetExportSet operation
type GetExportSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExportSet instance
	ExportSet `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetExportSetResponse) String() string {
	return common.PointerString(response)
}
