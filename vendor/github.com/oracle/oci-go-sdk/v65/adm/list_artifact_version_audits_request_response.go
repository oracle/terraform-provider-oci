// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListArtifactVersionAuditsRequest wrapper for the ListArtifactVersionAudits operation
type ListArtifactVersionAuditsRequest struct {

	// A filter to return only resources that match the specified identifier.
	// Required only if the compartmentId query parameter is not specified.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that belong to the specified compartment identifier.
	// Required only if the id query param is not specified.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Audits that were created against the specified knowledge base.
	KnowledgeBaseId *string `mandatory:"false" contributesTo:"query" name:"knowledgeBaseId"`

	// A filter to return only Artifact Version Audits that match the specified lifecycleState.
	LifecycleState ArtifactVersionAuditLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListArtifactVersionAuditsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field used to sort Artifact Version Audits. Only one sort order is allowed.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _displayName_ is **ascending**.
	SortBy ListArtifactVersionAuditsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only Vulnerability Audits with timeCreated greater or equal to the specified value.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter to return only Vulnerability Audits with timeCreated less or equal to the specified value.
	TimeCreatedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThanOrEqualTo"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListArtifactVersionAuditsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListArtifactVersionAuditsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListArtifactVersionAuditsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListArtifactVersionAuditsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListArtifactVersionAuditsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactVersionAuditLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetArtifactVersionAuditLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListArtifactVersionAuditsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListArtifactVersionAuditsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListArtifactVersionAuditsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListArtifactVersionAuditsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListArtifactVersionAuditsResponse wrapper for the ListArtifactVersionAudits operation
type ListArtifactVersionAuditsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ArtifactVersionAuditCollection instances
	ArtifactVersionAuditCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListArtifactVersionAuditsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListArtifactVersionAuditsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListArtifactVersionAuditsSortOrderEnum Enum with underlying type: string
type ListArtifactVersionAuditsSortOrderEnum string

// Set of constants representing the allowable values for ListArtifactVersionAuditsSortOrderEnum
const (
	ListArtifactVersionAuditsSortOrderAsc  ListArtifactVersionAuditsSortOrderEnum = "ASC"
	ListArtifactVersionAuditsSortOrderDesc ListArtifactVersionAuditsSortOrderEnum = "DESC"
)

var mappingListArtifactVersionAuditsSortOrderEnum = map[string]ListArtifactVersionAuditsSortOrderEnum{
	"ASC":  ListArtifactVersionAuditsSortOrderAsc,
	"DESC": ListArtifactVersionAuditsSortOrderDesc,
}

var mappingListArtifactVersionAuditsSortOrderEnumLowerCase = map[string]ListArtifactVersionAuditsSortOrderEnum{
	"asc":  ListArtifactVersionAuditsSortOrderAsc,
	"desc": ListArtifactVersionAuditsSortOrderDesc,
}

// GetListArtifactVersionAuditsSortOrderEnumValues Enumerates the set of values for ListArtifactVersionAuditsSortOrderEnum
func GetListArtifactVersionAuditsSortOrderEnumValues() []ListArtifactVersionAuditsSortOrderEnum {
	values := make([]ListArtifactVersionAuditsSortOrderEnum, 0)
	for _, v := range mappingListArtifactVersionAuditsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListArtifactVersionAuditsSortOrderEnumStringValues Enumerates the set of values in String for ListArtifactVersionAuditsSortOrderEnum
func GetListArtifactVersionAuditsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListArtifactVersionAuditsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListArtifactVersionAuditsSortOrderEnum(val string) (ListArtifactVersionAuditsSortOrderEnum, bool) {
	enum, ok := mappingListArtifactVersionAuditsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListArtifactVersionAuditsSortByEnum Enum with underlying type: string
type ListArtifactVersionAuditsSortByEnum string

// Set of constants representing the allowable values for ListArtifactVersionAuditsSortByEnum
const (
	ListArtifactVersionAuditsSortByTimecreated ListArtifactVersionAuditsSortByEnum = "timeCreated"
	ListArtifactVersionAuditsSortByDisplayname ListArtifactVersionAuditsSortByEnum = "displayName"
)

var mappingListArtifactVersionAuditsSortByEnum = map[string]ListArtifactVersionAuditsSortByEnum{
	"timeCreated": ListArtifactVersionAuditsSortByTimecreated,
	"displayName": ListArtifactVersionAuditsSortByDisplayname,
}

var mappingListArtifactVersionAuditsSortByEnumLowerCase = map[string]ListArtifactVersionAuditsSortByEnum{
	"timecreated": ListArtifactVersionAuditsSortByTimecreated,
	"displayname": ListArtifactVersionAuditsSortByDisplayname,
}

// GetListArtifactVersionAuditsSortByEnumValues Enumerates the set of values for ListArtifactVersionAuditsSortByEnum
func GetListArtifactVersionAuditsSortByEnumValues() []ListArtifactVersionAuditsSortByEnum {
	values := make([]ListArtifactVersionAuditsSortByEnum, 0)
	for _, v := range mappingListArtifactVersionAuditsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListArtifactVersionAuditsSortByEnumStringValues Enumerates the set of values in String for ListArtifactVersionAuditsSortByEnum
func GetListArtifactVersionAuditsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListArtifactVersionAuditsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListArtifactVersionAuditsSortByEnum(val string) (ListArtifactVersionAuditsSortByEnum, bool) {
	enum, ok := mappingListArtifactVersionAuditsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
