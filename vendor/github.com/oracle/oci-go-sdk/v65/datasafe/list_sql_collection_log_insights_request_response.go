// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSqlCollectionLogInsightsRequest wrapper for the ListSqlCollectionLogInsights operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlCollectionLogInsights.go.html to see an example of how to use ListSqlCollectionLogInsightsRequest.
type ListSqlCollectionLogInsightsRequest struct {

	// An optional filter to return the stats of the SQL collection logs collected after the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeStarted"`

	// An optional filter to return the stats of the SQL collection logs collected before the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeEnded"`

	// The OCID of the SQL collection resource.
	SqlCollectionId *string `mandatory:"true" contributesTo:"path" name:"sqlCollectionId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The group by parameter to summarize SQL collection log insights aggregation.
	GroupBy ListSqlCollectionLogInsightsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlCollectionLogInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlCollectionLogInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlCollectionLogInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlCollectionLogInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlCollectionLogInsightsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlCollectionLogInsightsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListSqlCollectionLogInsightsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlCollectionLogInsightsResponse wrapper for the ListSqlCollectionLogInsights operation
type ListSqlCollectionLogInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlCollectionLogInsightsCollection instances
	SqlCollectionLogInsightsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlCollectionLogInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlCollectionLogInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlCollectionLogInsightsGroupByEnum Enum with underlying type: string
type ListSqlCollectionLogInsightsGroupByEnum string

// Set of constants representing the allowable values for ListSqlCollectionLogInsightsGroupByEnum
const (
	ListSqlCollectionLogInsightsGroupByClientip         ListSqlCollectionLogInsightsGroupByEnum = "clientIp"
	ListSqlCollectionLogInsightsGroupByClientprogram    ListSqlCollectionLogInsightsGroupByEnum = "clientProgram"
	ListSqlCollectionLogInsightsGroupByClientosusername ListSqlCollectionLogInsightsGroupByEnum = "clientOsUserName"
)

var mappingListSqlCollectionLogInsightsGroupByEnum = map[string]ListSqlCollectionLogInsightsGroupByEnum{
	"clientIp":         ListSqlCollectionLogInsightsGroupByClientip,
	"clientProgram":    ListSqlCollectionLogInsightsGroupByClientprogram,
	"clientOsUserName": ListSqlCollectionLogInsightsGroupByClientosusername,
}

var mappingListSqlCollectionLogInsightsGroupByEnumLowerCase = map[string]ListSqlCollectionLogInsightsGroupByEnum{
	"clientip":         ListSqlCollectionLogInsightsGroupByClientip,
	"clientprogram":    ListSqlCollectionLogInsightsGroupByClientprogram,
	"clientosusername": ListSqlCollectionLogInsightsGroupByClientosusername,
}

// GetListSqlCollectionLogInsightsGroupByEnumValues Enumerates the set of values for ListSqlCollectionLogInsightsGroupByEnum
func GetListSqlCollectionLogInsightsGroupByEnumValues() []ListSqlCollectionLogInsightsGroupByEnum {
	values := make([]ListSqlCollectionLogInsightsGroupByEnum, 0)
	for _, v := range mappingListSqlCollectionLogInsightsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionLogInsightsGroupByEnumStringValues Enumerates the set of values in String for ListSqlCollectionLogInsightsGroupByEnum
func GetListSqlCollectionLogInsightsGroupByEnumStringValues() []string {
	return []string{
		"clientIp",
		"clientProgram",
		"clientOsUserName",
	}
}

// GetMappingListSqlCollectionLogInsightsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionLogInsightsGroupByEnum(val string) (ListSqlCollectionLogInsightsGroupByEnum, bool) {
	enum, ok := mappingListSqlCollectionLogInsightsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
