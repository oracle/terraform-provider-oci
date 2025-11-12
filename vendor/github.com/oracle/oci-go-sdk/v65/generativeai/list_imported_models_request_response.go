// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListImportedModelsRequest wrapper for the ListImportedModels operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListImportedModels.go.html to see an example of how to use ListImportedModelsRequest.
type ListImportedModelsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire vendor given.
	Vendor *string `mandatory:"false" contributesTo:"query" name:"vendor"`

	// A filter to return only resources their capability matches the given capability.
	Capability []ImportedModelCapabilityEnum `contributesTo:"query" name:"capability" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ImportedModelLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the importedModel.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListImportedModelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated` is descending.
	SortBy ListImportedModelsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImportedModelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImportedModelsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImportedModelsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImportedModelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListImportedModelsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Capability {
		if _, ok := GetMappingImportedModelCapabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Capability: %s. Supported values are: %s.", val, strings.Join(GetImportedModelCapabilityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingImportedModelLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetImportedModelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportedModelsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListImportedModelsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportedModelsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListImportedModelsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListImportedModelsResponse wrapper for the ListImportedModels operation
type ListImportedModelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ImportedModelCollection instances
	ImportedModelCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImportedModelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImportedModelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListImportedModelsSortOrderEnum Enum with underlying type: string
type ListImportedModelsSortOrderEnum string

// Set of constants representing the allowable values for ListImportedModelsSortOrderEnum
const (
	ListImportedModelsSortOrderAsc  ListImportedModelsSortOrderEnum = "ASC"
	ListImportedModelsSortOrderDesc ListImportedModelsSortOrderEnum = "DESC"
)

var mappingListImportedModelsSortOrderEnum = map[string]ListImportedModelsSortOrderEnum{
	"ASC":  ListImportedModelsSortOrderAsc,
	"DESC": ListImportedModelsSortOrderDesc,
}

var mappingListImportedModelsSortOrderEnumLowerCase = map[string]ListImportedModelsSortOrderEnum{
	"asc":  ListImportedModelsSortOrderAsc,
	"desc": ListImportedModelsSortOrderDesc,
}

// GetListImportedModelsSortOrderEnumValues Enumerates the set of values for ListImportedModelsSortOrderEnum
func GetListImportedModelsSortOrderEnumValues() []ListImportedModelsSortOrderEnum {
	values := make([]ListImportedModelsSortOrderEnum, 0)
	for _, v := range mappingListImportedModelsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportedModelsSortOrderEnumStringValues Enumerates the set of values in String for ListImportedModelsSortOrderEnum
func GetListImportedModelsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListImportedModelsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportedModelsSortOrderEnum(val string) (ListImportedModelsSortOrderEnum, bool) {
	enum, ok := mappingListImportedModelsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportedModelsSortByEnum Enum with underlying type: string
type ListImportedModelsSortByEnum string

// Set of constants representing the allowable values for ListImportedModelsSortByEnum
const (
	ListImportedModelsSortByDisplayname    ListImportedModelsSortByEnum = "displayName"
	ListImportedModelsSortByLifecyclestate ListImportedModelsSortByEnum = "lifecycleState"
	ListImportedModelsSortByTimecreated    ListImportedModelsSortByEnum = "timeCreated"
)

var mappingListImportedModelsSortByEnum = map[string]ListImportedModelsSortByEnum{
	"displayName":    ListImportedModelsSortByDisplayname,
	"lifecycleState": ListImportedModelsSortByLifecyclestate,
	"timeCreated":    ListImportedModelsSortByTimecreated,
}

var mappingListImportedModelsSortByEnumLowerCase = map[string]ListImportedModelsSortByEnum{
	"displayname":    ListImportedModelsSortByDisplayname,
	"lifecyclestate": ListImportedModelsSortByLifecyclestate,
	"timecreated":    ListImportedModelsSortByTimecreated,
}

// GetListImportedModelsSortByEnumValues Enumerates the set of values for ListImportedModelsSortByEnum
func GetListImportedModelsSortByEnumValues() []ListImportedModelsSortByEnum {
	values := make([]ListImportedModelsSortByEnum, 0)
	for _, v := range mappingListImportedModelsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportedModelsSortByEnumStringValues Enumerates the set of values in String for ListImportedModelsSortByEnum
func GetListImportedModelsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"lifecycleState",
		"timeCreated",
	}
}

// GetMappingListImportedModelsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportedModelsSortByEnum(val string) (ListImportedModelsSortByEnum, bool) {
	enum, ok := mappingListImportedModelsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
