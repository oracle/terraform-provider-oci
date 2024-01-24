// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModelsRequest wrapper for the ListModels operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListModels.go.html to see an example of how to use ListModelsRequest.
type ListModelsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire vendor given.
	Vendor *string `mandatory:"false" contributesTo:"query" name:"vendor"`

	// A filter to return only resources their capability matches the given capability.
	Capability []ModelCapabilityEnum `contributesTo:"query" name:"capability" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ModelLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the model.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated` is descending.
	SortBy ListModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Capability {
		if _, ok := GetMappingModelCapabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Capability: %s. Supported values are: %s.", val, strings.Join(GetModelCapabilityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingModelLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelsResponse wrapper for the ListModels operation
type ListModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ModelCollection instances
	ModelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelsSortOrderEnum Enum with underlying type: string
type ListModelsSortOrderEnum string

// Set of constants representing the allowable values for ListModelsSortOrderEnum
const (
	ListModelsSortOrderAsc  ListModelsSortOrderEnum = "ASC"
	ListModelsSortOrderDesc ListModelsSortOrderEnum = "DESC"
)

var mappingListModelsSortOrderEnum = map[string]ListModelsSortOrderEnum{
	"ASC":  ListModelsSortOrderAsc,
	"DESC": ListModelsSortOrderDesc,
}

var mappingListModelsSortOrderEnumLowerCase = map[string]ListModelsSortOrderEnum{
	"asc":  ListModelsSortOrderAsc,
	"desc": ListModelsSortOrderDesc,
}

// GetListModelsSortOrderEnumValues Enumerates the set of values for ListModelsSortOrderEnum
func GetListModelsSortOrderEnumValues() []ListModelsSortOrderEnum {
	values := make([]ListModelsSortOrderEnum, 0)
	for _, v := range mappingListModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelsSortOrderEnumStringValues Enumerates the set of values in String for ListModelsSortOrderEnum
func GetListModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelsSortOrderEnum(val string) (ListModelsSortOrderEnum, bool) {
	enum, ok := mappingListModelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelsSortByEnum Enum with underlying type: string
type ListModelsSortByEnum string

// Set of constants representing the allowable values for ListModelsSortByEnum
const (
	ListModelsSortByDisplayname    ListModelsSortByEnum = "displayName"
	ListModelsSortByLifecyclestate ListModelsSortByEnum = "lifecycleState"
	ListModelsSortByTimecreated    ListModelsSortByEnum = "timeCreated"
)

var mappingListModelsSortByEnum = map[string]ListModelsSortByEnum{
	"displayName":    ListModelsSortByDisplayname,
	"lifecycleState": ListModelsSortByLifecyclestate,
	"timeCreated":    ListModelsSortByTimecreated,
}

var mappingListModelsSortByEnumLowerCase = map[string]ListModelsSortByEnum{
	"displayname":    ListModelsSortByDisplayname,
	"lifecyclestate": ListModelsSortByLifecyclestate,
	"timecreated":    ListModelsSortByTimecreated,
}

// GetListModelsSortByEnumValues Enumerates the set of values for ListModelsSortByEnum
func GetListModelsSortByEnumValues() []ListModelsSortByEnum {
	values := make([]ListModelsSortByEnum, 0)
	for _, v := range mappingListModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelsSortByEnumStringValues Enumerates the set of values in String for ListModelsSortByEnum
func GetListModelsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"lifecycleState",
		"timeCreated",
	}
}

// GetMappingListModelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelsSortByEnum(val string) (ListModelsSortByEnum, bool) {
	enum, ok := mappingListModelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
