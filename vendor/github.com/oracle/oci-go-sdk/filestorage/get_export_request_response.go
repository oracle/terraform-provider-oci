// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetExportRequest wrapper for the GetExport operation
type GetExportRequest struct {

	// The OCID of the export.
	ExportId *string `mandatory:"true" contributesTo:"path" name:"exportId"`
}

func (request GetExportRequest) String() string {
	return common.PointerString(request)
}

// GetExportResponse wrapper for the GetExport operation
type GetExportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Export instance
	Export `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetExportResponse) String() string {
	return common.PointerString(response)
}
