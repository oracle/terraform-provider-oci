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

// GetProjectRequest wrapper for the GetProject operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/GetProject.go.html to see an example of how to use GetProjectRequest.
type GetProjectRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The project key.
	ProjectKey *string `mandatory:"true" contributesTo:"path" name:"projectKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This parameter allows users to specify which view of the object to return. CHILD_COUNT_STATISTICS - This option is used to get statistics on immediate children of the object by their type.
	Projection []GetProjectProjectionEnum `contributesTo:"query" name:"projection" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetProjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetProjectRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetProjectRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetProjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetProjectRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Projection {
		if _, ok := GetMappingGetProjectProjectionEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Projection: %s. Supported values are: %s.", val, strings.Join(GetGetProjectProjectionEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetProjectResponse wrapper for the GetProject operation
type GetProjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Project instance
	Project `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetProjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetProjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetProjectProjectionEnum Enum with underlying type: string
type GetProjectProjectionEnum string

// Set of constants representing the allowable values for GetProjectProjectionEnum
const (
	GetProjectProjectionChildCountStatistics GetProjectProjectionEnum = "CHILD_COUNT_STATISTICS"
)

var mappingGetProjectProjectionEnum = map[string]GetProjectProjectionEnum{
	"CHILD_COUNT_STATISTICS": GetProjectProjectionChildCountStatistics,
}

// GetGetProjectProjectionEnumValues Enumerates the set of values for GetProjectProjectionEnum
func GetGetProjectProjectionEnumValues() []GetProjectProjectionEnum {
	values := make([]GetProjectProjectionEnum, 0)
	for _, v := range mappingGetProjectProjectionEnum {
		values = append(values, v)
	}
	return values
}

// GetGetProjectProjectionEnumStringValues Enumerates the set of values in String for GetProjectProjectionEnum
func GetGetProjectProjectionEnumStringValues() []string {
	return []string{
		"CHILD_COUNT_STATISTICS",
	}
}

// GetMappingGetProjectProjectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetProjectProjectionEnum(val string) (GetProjectProjectionEnum, bool) {
	mappingGetProjectProjectionEnumIgnoreCase := make(map[string]GetProjectProjectionEnum)
	for k, v := range mappingGetProjectProjectionEnum {
		mappingGetProjectProjectionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetProjectProjectionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
