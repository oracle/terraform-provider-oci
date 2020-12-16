// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package objectstorage

import (
	"context"
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// CallWithDetails executes the http request, the given context using details specified in the paremeters, this function
// provides a way to override some settings present in the client
func (client ObjectStorageClient) CallWithDetails(ctx context.Context, request *http.Request, details common.ClientCallDetails) (response *http.Response, err error) {

	checkZeroLengthRequestBody(request)

	return client.BaseClient.CallWithDetails(ctx, request, details)
}

//If content length is zero, to avoid sending transfer-coding: chunked header, need to explicitly set the body to nil/Nobody.
func checkZeroLengthRequestBody(request *http.Request) {
	if request.Header != nil && request.Body != nil && request.Body != http.NoBody &&
		common.ParseContentLength(request.Header.Get("Content-Length")) == 0 {
		request.Body = http.NoBody
	}
	return
}
