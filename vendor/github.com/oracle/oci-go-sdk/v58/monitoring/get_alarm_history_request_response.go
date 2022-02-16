// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// GetAlarmHistoryRequest wrapper for the GetAlarmHistory operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/GetAlarmHistory.go.html to see an example of how to use GetAlarmHistoryRequest.
type GetAlarmHistoryRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an alarm.
	AlarmId *string `mandatory:"true" contributesTo:"path" name:"alarmId"`

	// Customer part of the request identifier token. If you need to contact Oracle about a particular
	// request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The type of history entries to retrieve. State history (STATE_HISTORY) or state transition history (STATE_TRANSITION_HISTORY).
	// If not specified, entries of both types are retrieved.
	// Example: `STATE_HISTORY`
	AlarmHistorytype GetAlarmHistoryAlarmHistorytypeEnum `mandatory:"false" contributesTo:"query" name:"alarmHistorytype" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Default: 1000
	// Example: 500
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only alarm history entries with timestamps occurring on or after the specified date and time. Format defined by RFC3339.
	// Example: `2019-01-01T01:00:00.789Z`
	TimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampGreaterThanOrEqualTo"`

	// A filter to return only alarm history entries with timestamps occurring before the specified date and time. Format defined by RFC3339.
	// Example: `2019-01-02T01:00:00.789Z`
	TimestampLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timestampLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAlarmHistoryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAlarmHistoryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAlarmHistoryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAlarmHistoryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAlarmHistoryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAlarmHistoryAlarmHistorytypeEnum(string(request.AlarmHistorytype)); !ok && request.AlarmHistorytype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlarmHistorytype: %s. Supported values are: %s.", request.AlarmHistorytype, strings.Join(GetGetAlarmHistoryAlarmHistorytypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAlarmHistoryResponse wrapper for the GetAlarmHistory operation
type GetAlarmHistoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlarmHistoryCollection instances
	AlarmHistoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetAlarmHistoryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAlarmHistoryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAlarmHistoryAlarmHistorytypeEnum Enum with underlying type: string
type GetAlarmHistoryAlarmHistorytypeEnum string

// Set of constants representing the allowable values for GetAlarmHistoryAlarmHistorytypeEnum
const (
	GetAlarmHistoryAlarmHistorytypeHistory           GetAlarmHistoryAlarmHistorytypeEnum = "STATE_HISTORY"
	GetAlarmHistoryAlarmHistorytypeTransitionHistory GetAlarmHistoryAlarmHistorytypeEnum = "STATE_TRANSITION_HISTORY"
)

var mappingGetAlarmHistoryAlarmHistorytypeEnum = map[string]GetAlarmHistoryAlarmHistorytypeEnum{
	"STATE_HISTORY":            GetAlarmHistoryAlarmHistorytypeHistory,
	"STATE_TRANSITION_HISTORY": GetAlarmHistoryAlarmHistorytypeTransitionHistory,
}

// GetGetAlarmHistoryAlarmHistorytypeEnumValues Enumerates the set of values for GetAlarmHistoryAlarmHistorytypeEnum
func GetGetAlarmHistoryAlarmHistorytypeEnumValues() []GetAlarmHistoryAlarmHistorytypeEnum {
	values := make([]GetAlarmHistoryAlarmHistorytypeEnum, 0)
	for _, v := range mappingGetAlarmHistoryAlarmHistorytypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAlarmHistoryAlarmHistorytypeEnumStringValues Enumerates the set of values in String for GetAlarmHistoryAlarmHistorytypeEnum
func GetGetAlarmHistoryAlarmHistorytypeEnumStringValues() []string {
	return []string{
		"STATE_HISTORY",
		"STATE_TRANSITION_HISTORY",
	}
}

// GetMappingGetAlarmHistoryAlarmHistorytypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAlarmHistoryAlarmHistorytypeEnum(val string) (GetAlarmHistoryAlarmHistorytypeEnum, bool) {
	mappingGetAlarmHistoryAlarmHistorytypeEnumIgnoreCase := make(map[string]GetAlarmHistoryAlarmHistorytypeEnum)
	for k, v := range mappingGetAlarmHistoryAlarmHistorytypeEnum {
		mappingGetAlarmHistoryAlarmHistorytypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetAlarmHistoryAlarmHistorytypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
