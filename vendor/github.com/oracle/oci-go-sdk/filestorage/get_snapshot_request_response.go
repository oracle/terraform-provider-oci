// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetSnapshotRequest wrapper for the GetSnapshot operation
type GetSnapshotRequest struct {

	// The OCID of the snapshot.
	SnapshotId *string `mandatory:"true" contributesTo:"path" name:"snapshotId"`
}

func (request GetSnapshotRequest) String() string {
	return common.PointerString(request)
}

// GetSnapshotResponse wrapper for the GetSnapshot operation
type GetSnapshotResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Snapshot instance
	Snapshot `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSnapshotResponse) String() string {
	return common.PointerString(response)
}
