// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operationsinsights

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
	"strings"
)

// ListAssociatedResourcesRequest wrapper for the ListAssociatedResources operation
type ListAssociatedResourcesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Operation Insights private endpoint.
	OperationsInsightsPrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"operationsInsightsPrivateEndpointId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// The field to query associated resource by resource type.
	ResourceType ListAssociatedResourcesResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceType" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociatedResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort associated resource using a specific Operation Insights Private Endpoint.
	SortBy ListAssociatedResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociatedResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociatedResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociatedResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociatedResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssociatedResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := mappingLifecycleStateEnum[string(val)]; !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := mappingListAssociatedResourcesResourceTypeEnum[string(request.ResourceType)]; !ok && request.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", request.ResourceType, strings.Join(GetListAssociatedResourcesResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := mappingListAssociatedResourcesSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssociatedResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListAssociatedResourcesSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssociatedResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssociatedResourcesResponse wrapper for the ListAssociatedResources operation
type ListAssociatedResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociatedResourceCollection instances
	AssociatedResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssociatedResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociatedResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociatedResourcesResourceTypeEnum Enum with underlying type: string
type ListAssociatedResourcesResourceTypeEnum string

// Set of constants representing the allowable values for ListAssociatedResourcesResourceTypeEnum
const (
	ListAssociatedResourcesResourceTypeDatabaseinsight ListAssociatedResourcesResourceTypeEnum = "databaseInsight"
	ListAssociatedResourcesResourceTypeHostinsight     ListAssociatedResourcesResourceTypeEnum = "hostInsight"
)

var mappingListAssociatedResourcesResourceTypeEnum = map[string]ListAssociatedResourcesResourceTypeEnum{
	"databaseInsight": ListAssociatedResourcesResourceTypeDatabaseinsight,
	"hostInsight":     ListAssociatedResourcesResourceTypeHostinsight,
}

// GetListAssociatedResourcesResourceTypeEnumValues Enumerates the set of values for ListAssociatedResourcesResourceTypeEnum
func GetListAssociatedResourcesResourceTypeEnumValues() []ListAssociatedResourcesResourceTypeEnum {
	values := make([]ListAssociatedResourcesResourceTypeEnum, 0)
	for _, v := range mappingListAssociatedResourcesResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedResourcesResourceTypeEnumStringValues Enumerates the set of values in String for ListAssociatedResourcesResourceTypeEnum
func GetListAssociatedResourcesResourceTypeEnumStringValues() []string {
	return []string{
		"databaseInsight",
		"hostInsight",
	}
}

// ListAssociatedResourcesSortOrderEnum Enum with underlying type: string
type ListAssociatedResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListAssociatedResourcesSortOrderEnum
const (
	ListAssociatedResourcesSortOrderAsc  ListAssociatedResourcesSortOrderEnum = "ASC"
	ListAssociatedResourcesSortOrderDesc ListAssociatedResourcesSortOrderEnum = "DESC"
)

var mappingListAssociatedResourcesSortOrderEnum = map[string]ListAssociatedResourcesSortOrderEnum{
	"ASC":  ListAssociatedResourcesSortOrderAsc,
	"DESC": ListAssociatedResourcesSortOrderDesc,
}

// GetListAssociatedResourcesSortOrderEnumValues Enumerates the set of values for ListAssociatedResourcesSortOrderEnum
func GetListAssociatedResourcesSortOrderEnumValues() []ListAssociatedResourcesSortOrderEnum {
	values := make([]ListAssociatedResourcesSortOrderEnum, 0)
	for _, v := range mappingListAssociatedResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListAssociatedResourcesSortOrderEnum
func GetListAssociatedResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListAssociatedResourcesSortByEnum Enum with underlying type: string
type ListAssociatedResourcesSortByEnum string

// Set of constants representing the allowable values for ListAssociatedResourcesSortByEnum
const (
	ListAssociatedResourcesSortByTimeregistered ListAssociatedResourcesSortByEnum = "timeRegistered"
	ListAssociatedResourcesSortById             ListAssociatedResourcesSortByEnum = "id"
	ListAssociatedResourcesSortByName           ListAssociatedResourcesSortByEnum = "name"
)

var mappingListAssociatedResourcesSortByEnum = map[string]ListAssociatedResourcesSortByEnum{
	"timeRegistered": ListAssociatedResourcesSortByTimeregistered,
	"id":             ListAssociatedResourcesSortById,
	"name":           ListAssociatedResourcesSortByName,
}

// GetListAssociatedResourcesSortByEnumValues Enumerates the set of values for ListAssociatedResourcesSortByEnum
func GetListAssociatedResourcesSortByEnumValues() []ListAssociatedResourcesSortByEnum {
	values := make([]ListAssociatedResourcesSortByEnum, 0)
	for _, v := range mappingListAssociatedResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedResourcesSortByEnumStringValues Enumerates the set of values in String for ListAssociatedResourcesSortByEnum
func GetListAssociatedResourcesSortByEnumStringValues() []string {
	return []string{
		"timeRegistered",
		"id",
		"name",
	}
}
