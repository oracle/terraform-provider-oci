// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstancesRequest wrapper for the ListManagedInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstances.go.html to see an example of how to use ListManagedInstancesRequest.
type ListManagedInstancesRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the managed instance for which to list resources.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// A filter to return only instances whose managed instance status matches the given status.
	Status []ManagedInstanceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only instances whose architecture type matches the given architecture.
	ArchType []ArchTypeEnum `contributesTo:"query" name:"archType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only instances whose OS family type matches the given OS family.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only managed instances acting as management stations.
	IsManagementStation *bool `mandatory:"false" contributesTo:"query" name:"isManagementStation"`

	// A filter to return only managed instances that are attached to the specified group.
	Group *string `mandatory:"false" contributesTo:"query" name:"group"`

	// A filter to return only managed instances that are NOT attached to the specified group.
	GroupNotEqualTo *string `mandatory:"false" contributesTo:"query" name:"groupNotEqualTo"`

	// A filter to return only managed instances that are associated with the specified lifecycle environment.
	LifecycleStage *string `mandatory:"false" contributesTo:"query" name:"lifecycleStage"`

	// A filter to return only managed instances that are NOT associated with the specified lifecycle environment.
	LifecycleStageNotEqualTo *string `mandatory:"false" contributesTo:"query" name:"lifecycleStageNotEqualTo"`

	// A filter to return only managed instances that are attached to the specified group or lifecycle environment.
	IsAttachedToGroupOrLifecycleStage *bool `mandatory:"false" contributesTo:"query" name:"isAttachedToGroupOrLifecycleStage"`

	// The OCID for the software source.
	SoftwareSourceId *string `mandatory:"false" contributesTo:"query" name:"softwareSourceId"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	AdvisoryName []string `contributesTo:"query" name:"advisoryName" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListManagedInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Status {
		if _, ok := GetMappingManagedInstanceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ArchType {
		if _, ok := GetMappingArchTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", val, strings.Join(GetArchTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListManagedInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstancesResponse wrapper for the ListManagedInstances operation
type ListManagedInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceCollection instances
	ManagedInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstancesSortOrderEnum Enum with underlying type: string
type ListManagedInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstancesSortOrderEnum
const (
	ListManagedInstancesSortOrderAsc  ListManagedInstancesSortOrderEnum = "ASC"
	ListManagedInstancesSortOrderDesc ListManagedInstancesSortOrderEnum = "DESC"
)

var mappingListManagedInstancesSortOrderEnum = map[string]ListManagedInstancesSortOrderEnum{
	"ASC":  ListManagedInstancesSortOrderAsc,
	"DESC": ListManagedInstancesSortOrderDesc,
}

var mappingListManagedInstancesSortOrderEnumLowerCase = map[string]ListManagedInstancesSortOrderEnum{
	"asc":  ListManagedInstancesSortOrderAsc,
	"desc": ListManagedInstancesSortOrderDesc,
}

// GetListManagedInstancesSortOrderEnumValues Enumerates the set of values for ListManagedInstancesSortOrderEnum
func GetListManagedInstancesSortOrderEnumValues() []ListManagedInstancesSortOrderEnum {
	values := make([]ListManagedInstancesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstancesSortOrderEnum
func GetListManagedInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesSortOrderEnum(val string) (ListManagedInstancesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstancesSortByEnum Enum with underlying type: string
type ListManagedInstancesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstancesSortByEnum
const (
	ListManagedInstancesSortByTimecreated ListManagedInstancesSortByEnum = "timeCreated"
	ListManagedInstancesSortByDisplayname ListManagedInstancesSortByEnum = "displayName"
)

var mappingListManagedInstancesSortByEnum = map[string]ListManagedInstancesSortByEnum{
	"timeCreated": ListManagedInstancesSortByTimecreated,
	"displayName": ListManagedInstancesSortByDisplayname,
}

var mappingListManagedInstancesSortByEnumLowerCase = map[string]ListManagedInstancesSortByEnum{
	"timecreated": ListManagedInstancesSortByTimecreated,
	"displayname": ListManagedInstancesSortByDisplayname,
}

// GetListManagedInstancesSortByEnumValues Enumerates the set of values for ListManagedInstancesSortByEnum
func GetListManagedInstancesSortByEnumValues() []ListManagedInstancesSortByEnum {
	values := make([]ListManagedInstancesSortByEnum, 0)
	for _, v := range mappingListManagedInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstancesSortByEnum
func GetListManagedInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesSortByEnum(val string) (ListManagedInstancesSortByEnum, bool) {
	enum, ok := mappingListManagedInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
