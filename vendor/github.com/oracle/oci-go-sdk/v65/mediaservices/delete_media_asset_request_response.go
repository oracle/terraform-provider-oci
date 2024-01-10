// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DeleteMediaAssetRequest wrapper for the DeleteMediaAsset operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteMediaAsset.go.html to see an example of how to use DeleteMediaAssetRequest.
type DeleteMediaAssetRequest struct {

	// Unique MediaAsset identifier
	MediaAssetId *string `mandatory:"true" contributesTo:"path" name:"mediaAssetId"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// DeleteMode decides whether to delete all the immediate children or all assets with the asset's ID as their masterMediaAssetId.
	DeleteMode DeleteMediaAssetDeleteModeEnum `mandatory:"false" contributesTo:"query" name:"deleteMode" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteMediaAssetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteMediaAssetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteMediaAssetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteMediaAssetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteMediaAssetRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeleteMediaAssetDeleteModeEnum(string(request.DeleteMode)); !ok && request.DeleteMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeleteMode: %s. Supported values are: %s.", request.DeleteMode, strings.Join(GetDeleteMediaAssetDeleteModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteMediaAssetResponse wrapper for the DeleteMediaAsset operation
type DeleteMediaAssetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteMediaAssetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteMediaAssetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteMediaAssetDeleteModeEnum Enum with underlying type: string
type DeleteMediaAssetDeleteModeEnum string

// Set of constants representing the allowable values for DeleteMediaAssetDeleteModeEnum
const (
	DeleteMediaAssetDeleteModeChildren    DeleteMediaAssetDeleteModeEnum = "DELETE_CHILDREN"
	DeleteMediaAssetDeleteModeDerivations DeleteMediaAssetDeleteModeEnum = "DELETE_DERIVATIONS"
)

var mappingDeleteMediaAssetDeleteModeEnum = map[string]DeleteMediaAssetDeleteModeEnum{
	"DELETE_CHILDREN":    DeleteMediaAssetDeleteModeChildren,
	"DELETE_DERIVATIONS": DeleteMediaAssetDeleteModeDerivations,
}

var mappingDeleteMediaAssetDeleteModeEnumLowerCase = map[string]DeleteMediaAssetDeleteModeEnum{
	"delete_children":    DeleteMediaAssetDeleteModeChildren,
	"delete_derivations": DeleteMediaAssetDeleteModeDerivations,
}

// GetDeleteMediaAssetDeleteModeEnumValues Enumerates the set of values for DeleteMediaAssetDeleteModeEnum
func GetDeleteMediaAssetDeleteModeEnumValues() []DeleteMediaAssetDeleteModeEnum {
	values := make([]DeleteMediaAssetDeleteModeEnum, 0)
	for _, v := range mappingDeleteMediaAssetDeleteModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeleteMediaAssetDeleteModeEnumStringValues Enumerates the set of values in String for DeleteMediaAssetDeleteModeEnum
func GetDeleteMediaAssetDeleteModeEnumStringValues() []string {
	return []string{
		"DELETE_CHILDREN",
		"DELETE_DERIVATIONS",
	}
}

// GetMappingDeleteMediaAssetDeleteModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeleteMediaAssetDeleteModeEnum(val string) (DeleteMediaAssetDeleteModeEnum, bool) {
	enum, ok := mappingDeleteMediaAssetDeleteModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
