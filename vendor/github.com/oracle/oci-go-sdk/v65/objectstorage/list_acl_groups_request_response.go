// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAclGroupsRequest wrapper for the ListAclGroups operation
type ListAclGroupsRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list buckets.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the ACL Group.
	AclGroupId *string `mandatory:"false" contributesTo:"query" name:"aclGroupId"`

	// The name of the ACL Group.
	AclGroupName *string `mandatory:"false" contributesTo:"query" name:"aclGroupName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order can be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListAclGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAclGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAclGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAclGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAclGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListAclGroupsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["namespaceName"] != nil {
		templateParam := mandatoryParamMap["namespaceName"]
		for _, template := range templateParam {
			replacementParam := *request.NamespaceName
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
	if mandatoryParamMap["compartmentId"] != nil {
		templateParam := mandatoryParamMap["compartmentId"]
		for _, template := range templateParam {
			replacementParam := *request.CompartmentId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAclGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAclGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAclGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAclGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAclGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAclGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAclGroupsResponse wrapper for the ListAclGroups operation
type ListAclGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AclGroupCollection instances
	AclGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Paginating a list of ACL Groups.
	// In the GET request, set the limit to the number of items that you want returned in the response.
	// If the `opc-next-page` header appears in the response, it indicates that this is a partial list
	// of ACL Groups and there are additional ACL Groups to get. Include the value of this header as the `page`
	// parameter in a subsequent GET request to get the next set of ACL Groups. Repeat this process to retrieve
	// the entire list of ACL Groups.
	// For more details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAclGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAclGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAclGroupsSortByEnum Enum with underlying type: string
type ListAclGroupsSortByEnum string

// Set of constants representing the allowable values for ListAclGroupsSortByEnum
const (
	ListAclGroupsSortByTimecreated ListAclGroupsSortByEnum = "TIMECREATED"
	ListAclGroupsSortByName        ListAclGroupsSortByEnum = "NAME"
)

var mappingListAclGroupsSortByEnum = map[string]ListAclGroupsSortByEnum{
	"TIMECREATED": ListAclGroupsSortByTimecreated,
	"NAME":        ListAclGroupsSortByName,
}

var mappingListAclGroupsSortByEnumLowerCase = map[string]ListAclGroupsSortByEnum{
	"timecreated": ListAclGroupsSortByTimecreated,
	"name":        ListAclGroupsSortByName,
}

// GetListAclGroupsSortByEnumValues Enumerates the set of values for ListAclGroupsSortByEnum
func GetListAclGroupsSortByEnumValues() []ListAclGroupsSortByEnum {
	values := make([]ListAclGroupsSortByEnum, 0)
	for _, v := range mappingListAclGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAclGroupsSortByEnumStringValues Enumerates the set of values in String for ListAclGroupsSortByEnum
func GetListAclGroupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListAclGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAclGroupsSortByEnum(val string) (ListAclGroupsSortByEnum, bool) {
	enum, ok := mappingListAclGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAclGroupsSortOrderEnum Enum with underlying type: string
type ListAclGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListAclGroupsSortOrderEnum
const (
	ListAclGroupsSortOrderAsc  ListAclGroupsSortOrderEnum = "ASC"
	ListAclGroupsSortOrderDesc ListAclGroupsSortOrderEnum = "DESC"
)

var mappingListAclGroupsSortOrderEnum = map[string]ListAclGroupsSortOrderEnum{
	"ASC":  ListAclGroupsSortOrderAsc,
	"DESC": ListAclGroupsSortOrderDesc,
}

var mappingListAclGroupsSortOrderEnumLowerCase = map[string]ListAclGroupsSortOrderEnum{
	"asc":  ListAclGroupsSortOrderAsc,
	"desc": ListAclGroupsSortOrderDesc,
}

// GetListAclGroupsSortOrderEnumValues Enumerates the set of values for ListAclGroupsSortOrderEnum
func GetListAclGroupsSortOrderEnumValues() []ListAclGroupsSortOrderEnum {
	values := make([]ListAclGroupsSortOrderEnum, 0)
	for _, v := range mappingListAclGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAclGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListAclGroupsSortOrderEnum
func GetListAclGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAclGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAclGroupsSortOrderEnum(val string) (ListAclGroupsSortOrderEnum, bool) {
	enum, ok := mappingListAclGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
