// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// GetFolderRequest wrapper for the GetFolder operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetFolder.go.html to see an example of how to use GetFolderRequest.
type GetFolderRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The folder key.
	FolderKey *string `mandatory:"true" contributesTo:"path" name:"folderKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This parameter allows users to specify which view of the object to return. CHILD_COUNT_STATISTICS - This option is used to get statistics on immediate children of the object by their type.
	Projection []GetFolderProjectionEnum `contributesTo:"query" name:"projection" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFolderRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFolderRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetFolderRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFolderRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetFolderRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Projection {
		if _, ok := GetMappingGetFolderProjectionEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Projection: %s. Supported values are: %s.", val, strings.Join(GetGetFolderProjectionEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetFolderResponse wrapper for the GetFolder operation
type GetFolderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Folder instance
	Folder `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFolderResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFolderResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetFolderProjectionEnum Enum with underlying type: string
type GetFolderProjectionEnum string

// Set of constants representing the allowable values for GetFolderProjectionEnum
const (
	GetFolderProjectionChildCountStatistics GetFolderProjectionEnum = "CHILD_COUNT_STATISTICS"
)

var mappingGetFolderProjectionEnum = map[string]GetFolderProjectionEnum{
	"CHILD_COUNT_STATISTICS": GetFolderProjectionChildCountStatistics,
}

// GetGetFolderProjectionEnumValues Enumerates the set of values for GetFolderProjectionEnum
func GetGetFolderProjectionEnumValues() []GetFolderProjectionEnum {
	values := make([]GetFolderProjectionEnum, 0)
	for _, v := range mappingGetFolderProjectionEnum {
		values = append(values, v)
	}
	return values
}

// GetGetFolderProjectionEnumStringValues Enumerates the set of values in String for GetFolderProjectionEnum
func GetGetFolderProjectionEnumStringValues() []string {
	return []string{
		"CHILD_COUNT_STATISTICS",
	}
}

// GetMappingGetFolderProjectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetFolderProjectionEnum(val string) (GetFolderProjectionEnum, bool) {
	mappingGetFolderProjectionEnumIgnoreCase := make(map[string]GetFolderProjectionEnum)
	for k, v := range mappingGetFolderProjectionEnum {
		mappingGetFolderProjectionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetFolderProjectionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
