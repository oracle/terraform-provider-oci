// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAclsRequest wrapper for the ListAcls operation
type ListAclsRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list buckets.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the ACL.
	AclId *string `mandatory:"false" contributesTo:"query" name:"aclId"`

	// The name of the ACL.
	AclName *string `mandatory:"false" contributesTo:"query" name:"aclName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order can be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListAclsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAclsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAclsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAclsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAclsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListAclsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
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
func (request ListAclsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAclsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAclsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAclsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAclsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAclsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAclsResponse wrapper for the ListAcls operation
type ListAclsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AclCollection instances
	AclCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Paginating a list of ACLs.
	// In the GET request, set the limit to the number of items that you want returned in the response.
	// If the `opc-next-page` header appears in the response, it indicates that this is a partial list
	// of ACLs and there are additional ACLs to get. Include the value of this header as the `page` parameter
	// in a subsequent GET request to get the next set of ACLs. Repeat this process to retrieve the entire list
	// of ACLs.
	// For more details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAclsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAclsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAclsSortByEnum Enum with underlying type: string
type ListAclsSortByEnum string

// Set of constants representing the allowable values for ListAclsSortByEnum
const (
	ListAclsSortByTimecreated ListAclsSortByEnum = "TIMECREATED"
	ListAclsSortByName        ListAclsSortByEnum = "NAME"
)

var mappingListAclsSortByEnum = map[string]ListAclsSortByEnum{
	"TIMECREATED": ListAclsSortByTimecreated,
	"NAME":        ListAclsSortByName,
}

var mappingListAclsSortByEnumLowerCase = map[string]ListAclsSortByEnum{
	"timecreated": ListAclsSortByTimecreated,
	"name":        ListAclsSortByName,
}

// GetListAclsSortByEnumValues Enumerates the set of values for ListAclsSortByEnum
func GetListAclsSortByEnumValues() []ListAclsSortByEnum {
	values := make([]ListAclsSortByEnum, 0)
	for _, v := range mappingListAclsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAclsSortByEnumStringValues Enumerates the set of values in String for ListAclsSortByEnum
func GetListAclsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListAclsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAclsSortByEnum(val string) (ListAclsSortByEnum, bool) {
	enum, ok := mappingListAclsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAclsSortOrderEnum Enum with underlying type: string
type ListAclsSortOrderEnum string

// Set of constants representing the allowable values for ListAclsSortOrderEnum
const (
	ListAclsSortOrderAsc  ListAclsSortOrderEnum = "ASC"
	ListAclsSortOrderDesc ListAclsSortOrderEnum = "DESC"
)

var mappingListAclsSortOrderEnum = map[string]ListAclsSortOrderEnum{
	"ASC":  ListAclsSortOrderAsc,
	"DESC": ListAclsSortOrderDesc,
}

var mappingListAclsSortOrderEnumLowerCase = map[string]ListAclsSortOrderEnum{
	"asc":  ListAclsSortOrderAsc,
	"desc": ListAclsSortOrderDesc,
}

// GetListAclsSortOrderEnumValues Enumerates the set of values for ListAclsSortOrderEnum
func GetListAclsSortOrderEnumValues() []ListAclsSortOrderEnum {
	values := make([]ListAclsSortOrderEnum, 0)
	for _, v := range mappingListAclsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAclsSortOrderEnumStringValues Enumerates the set of values in String for ListAclsSortOrderEnum
func GetListAclsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAclsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAclsSortOrderEnum(val string) (ListAclsSortOrderEnum, bool) {
	enum, ok := mappingListAclsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
