// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwareSourceErrataRequest wrapper for the ListSoftwareSourceErrata operation
type ListSoftwareSourceErrataRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// A filter to return resources that may partially match the erratum name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only errata that match the given advisory types.
	AdvisoryType []AdvisoryTypesEnum `contributesTo:"query" name:"advisoryType" omitEmpty:"true" collectionFormat:"multi"`

	// The advisory severity.
	AdvisorySeverity []AdvisorySeverityEnum `contributesTo:"query" name:"advisorySeverity" omitEmpty:"true" collectionFormat:"multi"`

	// The issue date after which to list all errata, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeIssueDateStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIssueDateStart"`

	// The issue date before which to list all errata, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeIssueDateEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIssueDateEnd"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSoftwareSourceErrataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort errata by. Only one sort order may be provided. Default order for timeIssued is descending. Default order for name is ascending. If no value is specified timeIssued is default.
	SortBy ListSoftwareSourceErrataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourceErrataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourceErrataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareSourceErrataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourceErrataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareSourceErrataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.AdvisoryType {
		if _, ok := GetMappingAdvisoryTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisoryType: %s. Supported values are: %s.", val, strings.Join(GetAdvisoryTypesEnumStringValues(), ",")))
		}
	}

	for _, val := range request.AdvisorySeverity {
		if _, ok := GetMappingAdvisorySeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisorySeverity: %s. Supported values are: %s.", val, strings.Join(GetAdvisorySeverityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSoftwareSourceErrataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareSourceErrataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareSourceErrataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareSourceErrataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareSourceErrataResponse wrapper for the ListSoftwareSourceErrata operation
type ListSoftwareSourceErrataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ErratumCollection instances
	ErratumCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareSourceErrataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourceErrataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourceErrataSortOrderEnum Enum with underlying type: string
type ListSoftwareSourceErrataSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourceErrataSortOrderEnum
const (
	ListSoftwareSourceErrataSortOrderAsc  ListSoftwareSourceErrataSortOrderEnum = "ASC"
	ListSoftwareSourceErrataSortOrderDesc ListSoftwareSourceErrataSortOrderEnum = "DESC"
)

var mappingListSoftwareSourceErrataSortOrderEnum = map[string]ListSoftwareSourceErrataSortOrderEnum{
	"ASC":  ListSoftwareSourceErrataSortOrderAsc,
	"DESC": ListSoftwareSourceErrataSortOrderDesc,
}

var mappingListSoftwareSourceErrataSortOrderEnumLowerCase = map[string]ListSoftwareSourceErrataSortOrderEnum{
	"asc":  ListSoftwareSourceErrataSortOrderAsc,
	"desc": ListSoftwareSourceErrataSortOrderDesc,
}

// GetListSoftwareSourceErrataSortOrderEnumValues Enumerates the set of values for ListSoftwareSourceErrataSortOrderEnum
func GetListSoftwareSourceErrataSortOrderEnumValues() []ListSoftwareSourceErrataSortOrderEnum {
	values := make([]ListSoftwareSourceErrataSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourceErrataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourceErrataSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareSourceErrataSortOrderEnum
func GetListSoftwareSourceErrataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareSourceErrataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourceErrataSortOrderEnum(val string) (ListSoftwareSourceErrataSortOrderEnum, bool) {
	enum, ok := mappingListSoftwareSourceErrataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourceErrataSortByEnum Enum with underlying type: string
type ListSoftwareSourceErrataSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourceErrataSortByEnum
const (
	ListSoftwareSourceErrataSortByTimeissued ListSoftwareSourceErrataSortByEnum = "timeIssued"
	ListSoftwareSourceErrataSortByName       ListSoftwareSourceErrataSortByEnum = "name"
)

var mappingListSoftwareSourceErrataSortByEnum = map[string]ListSoftwareSourceErrataSortByEnum{
	"timeIssued": ListSoftwareSourceErrataSortByTimeissued,
	"name":       ListSoftwareSourceErrataSortByName,
}

var mappingListSoftwareSourceErrataSortByEnumLowerCase = map[string]ListSoftwareSourceErrataSortByEnum{
	"timeissued": ListSoftwareSourceErrataSortByTimeissued,
	"name":       ListSoftwareSourceErrataSortByName,
}

// GetListSoftwareSourceErrataSortByEnumValues Enumerates the set of values for ListSoftwareSourceErrataSortByEnum
func GetListSoftwareSourceErrataSortByEnumValues() []ListSoftwareSourceErrataSortByEnum {
	values := make([]ListSoftwareSourceErrataSortByEnum, 0)
	for _, v := range mappingListSoftwareSourceErrataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourceErrataSortByEnumStringValues Enumerates the set of values in String for ListSoftwareSourceErrataSortByEnum
func GetListSoftwareSourceErrataSortByEnumStringValues() []string {
	return []string{
		"timeIssued",
		"name",
	}
}

// GetMappingListSoftwareSourceErrataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourceErrataSortByEnum(val string) (ListSoftwareSourceErrataSortByEnum, bool) {
	enum, ok := mappingListSoftwareSourceErrataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
