// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetMountTargetRequest wrapper for the GetMountTarget operation
type GetMountTargetRequest struct {

	// The OCID of the mount target.
	MountTargetId *string `mandatory:"true" contributesTo:"path" name:"mountTargetId"`
}

func (request GetMountTargetRequest) String() string {
	return common.PointerString(request)
}

// GetMountTargetResponse wrapper for the GetMountTarget operation
type GetMountTargetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MountTarget instance
	MountTarget `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMountTargetResponse) String() string {
	return common.PointerString(response)
}
