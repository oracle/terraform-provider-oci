// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetFileSystemRequest wrapper for the GetFileSystem operation
type GetFileSystemRequest struct {

	// The OCID of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`
}

func (request GetFileSystemRequest) String() string {
	return common.PointerString(request)
}

// GetFileSystemResponse wrapper for the GetFileSystem operation
type GetFileSystemResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The FileSystem instance
	FileSystem `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFileSystemResponse) String() string {
	return common.PointerString(response)
}
