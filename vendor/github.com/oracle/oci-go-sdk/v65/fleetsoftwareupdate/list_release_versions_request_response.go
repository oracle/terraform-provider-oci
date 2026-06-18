// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReleaseVersionsRequest wrapper for the ListReleaseVersions operation
type ListReleaseVersionsRequest struct {

	// Release version category to return.
	VersionType ListReleaseVersionsVersionTypeEnum `mandatory:"true" contributesTo:"query" name:"versionType" omitEmpty:"true"`

	// A filter to return only release versions that belong to the specified major version family.
	MajorFamily *string `mandatory:"false" contributesTo:"query" name:"majorFamily"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReleaseVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReleaseVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReleaseVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReleaseVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReleaseVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListReleaseVersionsVersionTypeEnum(string(request.VersionType)); !ok && request.VersionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionType: %s. Supported values are: %s.", request.VersionType, strings.Join(GetListReleaseVersionsVersionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReleaseVersionsResponse wrapper for the ListReleaseVersions operation
type ListReleaseVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReleaseVersionCollection instances
	ReleaseVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReleaseVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReleaseVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReleaseVersionsVersionTypeEnum Enum with underlying type: string
type ListReleaseVersionsVersionTypeEnum string

// Set of constants representing the allowable values for ListReleaseVersionsVersionTypeEnum
const (
	ListReleaseVersionsVersionTypeExadataReleaseVersion ListReleaseVersionsVersionTypeEnum = "EXADATA_RELEASE_VERSION"
)

var mappingListReleaseVersionsVersionTypeEnum = map[string]ListReleaseVersionsVersionTypeEnum{
	"EXADATA_RELEASE_VERSION": ListReleaseVersionsVersionTypeExadataReleaseVersion,
}

var mappingListReleaseVersionsVersionTypeEnumLowerCase = map[string]ListReleaseVersionsVersionTypeEnum{
	"exadata_release_version": ListReleaseVersionsVersionTypeExadataReleaseVersion,
}

// GetListReleaseVersionsVersionTypeEnumValues Enumerates the set of values for ListReleaseVersionsVersionTypeEnum
func GetListReleaseVersionsVersionTypeEnumValues() []ListReleaseVersionsVersionTypeEnum {
	values := make([]ListReleaseVersionsVersionTypeEnum, 0)
	for _, v := range mappingListReleaseVersionsVersionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListReleaseVersionsVersionTypeEnumStringValues Enumerates the set of values in String for ListReleaseVersionsVersionTypeEnum
func GetListReleaseVersionsVersionTypeEnumStringValues() []string {
	return []string{
		"EXADATA_RELEASE_VERSION",
	}
}

// GetMappingListReleaseVersionsVersionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReleaseVersionsVersionTypeEnum(val string) (ListReleaseVersionsVersionTypeEnum, bool) {
	enum, ok := mappingListReleaseVersionsVersionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
