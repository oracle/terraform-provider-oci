// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListStagesRequest wrapper for the ListStages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListStages.go.html to see an example of how to use ListStagesRequest.
type ListStagesRequest struct {

	// Unique Remediation Run identifier path parameter.
	RemediationRunId *string `mandatory:"true" contributesTo:"path" name:"remediationRunId"`

	// A filter to return only Stages that match the specified type.
	Type ListStagesTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return only Stages that match the specified status.
	Status RemediationRunStageStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListStagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field used to sort Stages. Only one sort order is allowed.
	// Default order for status is the following sequence: **CREATED, IN_PROGRESS, SUCCEEDED, FAILED, CANCELING, and CANCELED**.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _timeFinished_ is **descending**.
	// Default order for _timeStarted_ is **descending**.
	// Default order for _type_ is the following sequence: **DETECT, RECOMMEND, VERIFY, and APPLY**.
	SortBy ListStagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListStagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListStagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListStagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListStagesTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListStagesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRemediationRunStageStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetRemediationRunStageStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListStagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListStagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListStagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListStagesResponse wrapper for the ListStages operation
type ListStagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RemediationRunStageCollection instances
	RemediationRunStageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListStagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListStagesTypeEnum Enum with underlying type: string
type ListStagesTypeEnum string

// Set of constants representing the allowable values for ListStagesTypeEnum
const (
	ListStagesTypeDetect    ListStagesTypeEnum = "DETECT"
	ListStagesTypeRecommend ListStagesTypeEnum = "RECOMMEND"
	ListStagesTypeVerify    ListStagesTypeEnum = "VERIFY"
	ListStagesTypeApply     ListStagesTypeEnum = "APPLY"
)

var mappingListStagesTypeEnum = map[string]ListStagesTypeEnum{
	"DETECT":    ListStagesTypeDetect,
	"RECOMMEND": ListStagesTypeRecommend,
	"VERIFY":    ListStagesTypeVerify,
	"APPLY":     ListStagesTypeApply,
}

var mappingListStagesTypeEnumLowerCase = map[string]ListStagesTypeEnum{
	"detect":    ListStagesTypeDetect,
	"recommend": ListStagesTypeRecommend,
	"verify":    ListStagesTypeVerify,
	"apply":     ListStagesTypeApply,
}

// GetListStagesTypeEnumValues Enumerates the set of values for ListStagesTypeEnum
func GetListStagesTypeEnumValues() []ListStagesTypeEnum {
	values := make([]ListStagesTypeEnum, 0)
	for _, v := range mappingListStagesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListStagesTypeEnumStringValues Enumerates the set of values in String for ListStagesTypeEnum
func GetListStagesTypeEnumStringValues() []string {
	return []string{
		"DETECT",
		"RECOMMEND",
		"VERIFY",
		"APPLY",
	}
}

// GetMappingListStagesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStagesTypeEnum(val string) (ListStagesTypeEnum, bool) {
	enum, ok := mappingListStagesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStagesSortOrderEnum Enum with underlying type: string
type ListStagesSortOrderEnum string

// Set of constants representing the allowable values for ListStagesSortOrderEnum
const (
	ListStagesSortOrderAsc  ListStagesSortOrderEnum = "ASC"
	ListStagesSortOrderDesc ListStagesSortOrderEnum = "DESC"
)

var mappingListStagesSortOrderEnum = map[string]ListStagesSortOrderEnum{
	"ASC":  ListStagesSortOrderAsc,
	"DESC": ListStagesSortOrderDesc,
}

var mappingListStagesSortOrderEnumLowerCase = map[string]ListStagesSortOrderEnum{
	"asc":  ListStagesSortOrderAsc,
	"desc": ListStagesSortOrderDesc,
}

// GetListStagesSortOrderEnumValues Enumerates the set of values for ListStagesSortOrderEnum
func GetListStagesSortOrderEnumValues() []ListStagesSortOrderEnum {
	values := make([]ListStagesSortOrderEnum, 0)
	for _, v := range mappingListStagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListStagesSortOrderEnumStringValues Enumerates the set of values in String for ListStagesSortOrderEnum
func GetListStagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListStagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStagesSortOrderEnum(val string) (ListStagesSortOrderEnum, bool) {
	enum, ok := mappingListStagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListStagesSortByEnum Enum with underlying type: string
type ListStagesSortByEnum string

// Set of constants representing the allowable values for ListStagesSortByEnum
const (
	ListStagesSortByStatus       ListStagesSortByEnum = "status"
	ListStagesSortByTimecreated  ListStagesSortByEnum = "timeCreated"
	ListStagesSortByTimefinished ListStagesSortByEnum = "timeFinished"
	ListStagesSortByTimestarted  ListStagesSortByEnum = "timeStarted"
	ListStagesSortByType         ListStagesSortByEnum = "type"
)

var mappingListStagesSortByEnum = map[string]ListStagesSortByEnum{
	"status":       ListStagesSortByStatus,
	"timeCreated":  ListStagesSortByTimecreated,
	"timeFinished": ListStagesSortByTimefinished,
	"timeStarted":  ListStagesSortByTimestarted,
	"type":         ListStagesSortByType,
}

var mappingListStagesSortByEnumLowerCase = map[string]ListStagesSortByEnum{
	"status":       ListStagesSortByStatus,
	"timecreated":  ListStagesSortByTimecreated,
	"timefinished": ListStagesSortByTimefinished,
	"timestarted":  ListStagesSortByTimestarted,
	"type":         ListStagesSortByType,
}

// GetListStagesSortByEnumValues Enumerates the set of values for ListStagesSortByEnum
func GetListStagesSortByEnumValues() []ListStagesSortByEnum {
	values := make([]ListStagesSortByEnum, 0)
	for _, v := range mappingListStagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListStagesSortByEnumStringValues Enumerates the set of values in String for ListStagesSortByEnum
func GetListStagesSortByEnumStringValues() []string {
	return []string{
		"status",
		"timeCreated",
		"timeFinished",
		"timeStarted",
		"type",
	}
}

// GetMappingListStagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListStagesSortByEnum(val string) (ListStagesSortByEnum, bool) {
	enum, ok := mappingListStagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
