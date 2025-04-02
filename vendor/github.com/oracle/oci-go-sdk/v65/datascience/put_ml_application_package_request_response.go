// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// PutMlApplicationPackageRequest wrapper for the PutMlApplicationPackage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/PutMlApplicationPackage.go.html to see an example of how to use PutMlApplicationPackageRequest.
type PutMlApplicationPackageRequest struct {

	// unique MlApplicationImplementation identifier
	MlApplicationImplementationId *string `mandatory:"true" contributesTo:"path" name:"mlApplicationImplementationId"`

	// ML Application Package to upload
	PutMlApplicationPackage io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The content length of the body.
	ContentLength *int64 `mandatory:"false" contributesTo:"header" name:"content-length"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// This header allows you to specify a filename during upload. This file name is used to dispose of the file contents
	// while downloading the file. If this optional field is not populated in the request, then the OCID of the model is used for the file
	// name when downloading.
	// Example: `{"Content-Disposition": "attachment"
	//            "filename"="model.tar.gz"
	//            "Content-Length": "2347"
	//            "Content-Type": "application/gzip"}`
	ContentDisposition *string `mandatory:"false" contributesTo:"header" name:"content-disposition"`

	// List of arguments (Json map - argument name to argument value) for ML Application package (available arguments are in ML Application package descriptor). E.g. {"vcnId": "ocid1.vcn.oc1.iad.abcd...", "logId":"ocid1.log.oc1.iad.abcd..."}
	OpcMlAppPackageArgs *string `mandatory:"false" contributesTo:"header" name:"opc-ml-app-package-args"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutMlApplicationPackageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutMlApplicationPackageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutMlApplicationPackageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.PutMlApplicationPackage)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutMlApplicationPackageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PutMlApplicationPackageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutMlApplicationPackageResponse wrapper for the PutMlApplicationPackage operation
type PutMlApplicationPackageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID (https://docs.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the work request. Use GetWorkRequest (https://docs.oracle.com/iaas/api/#/en/workrequests/20160918/WorkRequest/GetWorkRequest)
	// with this ID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response PutMlApplicationPackageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutMlApplicationPackageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
