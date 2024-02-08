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

// ListDatabaseViewAccessEntriesRequest wrapper for the ListDatabaseViewAccessEntries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDatabaseViewAccessEntries.go.html to see an example of how to use ListDatabaseViewAccessEntriesRequest.
type ListDatabaseViewAccessEntriesRequest struct {

	// The OCID of the security policy report resource.
	SecurityPolicyReportId *string `mandatory:"true" contributesTo:"path" name:"securityPolicyReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// **Example:** query=(accessType eq 'SELECT') and (grantee eq 'ADMIN')
	ScimQuery *string `mandatory:"false" contributesTo:"query" name:"scimQuery"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The field to sort by. Only one sort parameter should be provided.
	SortBy ListDatabaseViewAccessEntriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDatabaseViewAccessEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseViewAccessEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseViewAccessEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseViewAccessEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseViewAccessEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseViewAccessEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseViewAccessEntriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseViewAccessEntriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseViewAccessEntriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseViewAccessEntriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseViewAccessEntriesResponse wrapper for the ListDatabaseViewAccessEntries operation
type ListDatabaseViewAccessEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseViewAccessEntryCollection instances
	DatabaseViewAccessEntryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDatabaseViewAccessEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseViewAccessEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseViewAccessEntriesSortByEnum Enum with underlying type: string
type ListDatabaseViewAccessEntriesSortByEnum string

// Set of constants representing the allowable values for ListDatabaseViewAccessEntriesSortByEnum
const (
	ListDatabaseViewAccessEntriesSortByKey                                          ListDatabaseViewAccessEntriesSortByEnum = "key"
	ListDatabaseViewAccessEntriesSortByGrantee                                      ListDatabaseViewAccessEntriesSortByEnum = "grantee"
	ListDatabaseViewAccessEntriesSortByAccesstype                                   ListDatabaseViewAccessEntriesSortByEnum = "accessType"
	ListDatabaseViewAccessEntriesSortByTableschema                                  ListDatabaseViewAccessEntriesSortByEnum = "tableSchema"
	ListDatabaseViewAccessEntriesSortByTablename                                    ListDatabaseViewAccessEntriesSortByEnum = "tableName"
	ListDatabaseViewAccessEntriesSortByViewschema                                   ListDatabaseViewAccessEntriesSortByEnum = "viewSchema"
	ListDatabaseViewAccessEntriesSortByViewname                                     ListDatabaseViewAccessEntriesSortByEnum = "viewName"
	ListDatabaseViewAccessEntriesSortByPrivilegetype                                ListDatabaseViewAccessEntriesSortByEnum = "privilegeType"
	ListDatabaseViewAccessEntriesSortByPrivilege                                    ListDatabaseViewAccessEntriesSortByEnum = "privilege"
	ListDatabaseViewAccessEntriesSortByPrivilegegrantable                           ListDatabaseViewAccessEntriesSortByEnum = "privilegeGrantable"
	ListDatabaseViewAccessEntriesSortByGrantfromrole                                ListDatabaseViewAccessEntriesSortByEnum = "grantFromRole"
	ListDatabaseViewAccessEntriesSortByAccessthroughobject                          ListDatabaseViewAccessEntriesSortByEnum = "accessThroughObject"
	ListDatabaseViewAccessEntriesSortByColumnname                                   ListDatabaseViewAccessEntriesSortByEnum = "columnName"
	ListDatabaseViewAccessEntriesSortByGrantor                                      ListDatabaseViewAccessEntriesSortByEnum = "grantor"
	ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbydatabasevault           ListDatabaseViewAccessEntriesSortByEnum = "isAccessConstrainedByDatabaseVault"
	ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyvirtualprivatedatabase  ListDatabaseViewAccessEntriesSortByEnum = "isAccessConstrainedByVirtualPrivateDatabase"
	ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyredaction               ListDatabaseViewAccessEntriesSortByEnum = "isAccessConstrainedByRedaction"
	ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyrealapplicationsecurity ListDatabaseViewAccessEntriesSortByEnum = "isAccessConstrainedByRealApplicationSecurity"
	ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbysqlfirewall             ListDatabaseViewAccessEntriesSortByEnum = "isAccessConstrainedBySqlFirewall"
)

var mappingListDatabaseViewAccessEntriesSortByEnum = map[string]ListDatabaseViewAccessEntriesSortByEnum{
	"key":                                ListDatabaseViewAccessEntriesSortByKey,
	"grantee":                            ListDatabaseViewAccessEntriesSortByGrantee,
	"accessType":                         ListDatabaseViewAccessEntriesSortByAccesstype,
	"tableSchema":                        ListDatabaseViewAccessEntriesSortByTableschema,
	"tableName":                          ListDatabaseViewAccessEntriesSortByTablename,
	"viewSchema":                         ListDatabaseViewAccessEntriesSortByViewschema,
	"viewName":                           ListDatabaseViewAccessEntriesSortByViewname,
	"privilegeType":                      ListDatabaseViewAccessEntriesSortByPrivilegetype,
	"privilege":                          ListDatabaseViewAccessEntriesSortByPrivilege,
	"privilegeGrantable":                 ListDatabaseViewAccessEntriesSortByPrivilegegrantable,
	"grantFromRole":                      ListDatabaseViewAccessEntriesSortByGrantfromrole,
	"accessThroughObject":                ListDatabaseViewAccessEntriesSortByAccessthroughobject,
	"columnName":                         ListDatabaseViewAccessEntriesSortByColumnname,
	"grantor":                            ListDatabaseViewAccessEntriesSortByGrantor,
	"isAccessConstrainedByDatabaseVault": ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbydatabasevault,
	"isAccessConstrainedByVirtualPrivateDatabase":  ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyvirtualprivatedatabase,
	"isAccessConstrainedByRedaction":               ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyredaction,
	"isAccessConstrainedByRealApplicationSecurity": ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyrealapplicationsecurity,
	"isAccessConstrainedBySqlFirewall":             ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbysqlfirewall,
}

var mappingListDatabaseViewAccessEntriesSortByEnumLowerCase = map[string]ListDatabaseViewAccessEntriesSortByEnum{
	"key":                                ListDatabaseViewAccessEntriesSortByKey,
	"grantee":                            ListDatabaseViewAccessEntriesSortByGrantee,
	"accesstype":                         ListDatabaseViewAccessEntriesSortByAccesstype,
	"tableschema":                        ListDatabaseViewAccessEntriesSortByTableschema,
	"tablename":                          ListDatabaseViewAccessEntriesSortByTablename,
	"viewschema":                         ListDatabaseViewAccessEntriesSortByViewschema,
	"viewname":                           ListDatabaseViewAccessEntriesSortByViewname,
	"privilegetype":                      ListDatabaseViewAccessEntriesSortByPrivilegetype,
	"privilege":                          ListDatabaseViewAccessEntriesSortByPrivilege,
	"privilegegrantable":                 ListDatabaseViewAccessEntriesSortByPrivilegegrantable,
	"grantfromrole":                      ListDatabaseViewAccessEntriesSortByGrantfromrole,
	"accessthroughobject":                ListDatabaseViewAccessEntriesSortByAccessthroughobject,
	"columnname":                         ListDatabaseViewAccessEntriesSortByColumnname,
	"grantor":                            ListDatabaseViewAccessEntriesSortByGrantor,
	"isaccessconstrainedbydatabasevault": ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbydatabasevault,
	"isaccessconstrainedbyvirtualprivatedatabase":  ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyvirtualprivatedatabase,
	"isaccessconstrainedbyredaction":               ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyredaction,
	"isaccessconstrainedbyrealapplicationsecurity": ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbyrealapplicationsecurity,
	"isaccessconstrainedbysqlfirewall":             ListDatabaseViewAccessEntriesSortByIsaccessconstrainedbysqlfirewall,
}

// GetListDatabaseViewAccessEntriesSortByEnumValues Enumerates the set of values for ListDatabaseViewAccessEntriesSortByEnum
func GetListDatabaseViewAccessEntriesSortByEnumValues() []ListDatabaseViewAccessEntriesSortByEnum {
	values := make([]ListDatabaseViewAccessEntriesSortByEnum, 0)
	for _, v := range mappingListDatabaseViewAccessEntriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseViewAccessEntriesSortByEnumStringValues Enumerates the set of values in String for ListDatabaseViewAccessEntriesSortByEnum
func GetListDatabaseViewAccessEntriesSortByEnumStringValues() []string {
	return []string{
		"key",
		"grantee",
		"accessType",
		"tableSchema",
		"tableName",
		"viewSchema",
		"viewName",
		"privilegeType",
		"privilege",
		"privilegeGrantable",
		"grantFromRole",
		"accessThroughObject",
		"columnName",
		"grantor",
		"isAccessConstrainedByDatabaseVault",
		"isAccessConstrainedByVirtualPrivateDatabase",
		"isAccessConstrainedByRedaction",
		"isAccessConstrainedByRealApplicationSecurity",
		"isAccessConstrainedBySqlFirewall",
	}
}

// GetMappingListDatabaseViewAccessEntriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseViewAccessEntriesSortByEnum(val string) (ListDatabaseViewAccessEntriesSortByEnum, bool) {
	enum, ok := mappingListDatabaseViewAccessEntriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseViewAccessEntriesSortOrderEnum Enum with underlying type: string
type ListDatabaseViewAccessEntriesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseViewAccessEntriesSortOrderEnum
const (
	ListDatabaseViewAccessEntriesSortOrderAsc  ListDatabaseViewAccessEntriesSortOrderEnum = "ASC"
	ListDatabaseViewAccessEntriesSortOrderDesc ListDatabaseViewAccessEntriesSortOrderEnum = "DESC"
)

var mappingListDatabaseViewAccessEntriesSortOrderEnum = map[string]ListDatabaseViewAccessEntriesSortOrderEnum{
	"ASC":  ListDatabaseViewAccessEntriesSortOrderAsc,
	"DESC": ListDatabaseViewAccessEntriesSortOrderDesc,
}

var mappingListDatabaseViewAccessEntriesSortOrderEnumLowerCase = map[string]ListDatabaseViewAccessEntriesSortOrderEnum{
	"asc":  ListDatabaseViewAccessEntriesSortOrderAsc,
	"desc": ListDatabaseViewAccessEntriesSortOrderDesc,
}

// GetListDatabaseViewAccessEntriesSortOrderEnumValues Enumerates the set of values for ListDatabaseViewAccessEntriesSortOrderEnum
func GetListDatabaseViewAccessEntriesSortOrderEnumValues() []ListDatabaseViewAccessEntriesSortOrderEnum {
	values := make([]ListDatabaseViewAccessEntriesSortOrderEnum, 0)
	for _, v := range mappingListDatabaseViewAccessEntriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseViewAccessEntriesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseViewAccessEntriesSortOrderEnum
func GetListDatabaseViewAccessEntriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseViewAccessEntriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseViewAccessEntriesSortOrderEnum(val string) (ListDatabaseViewAccessEntriesSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseViewAccessEntriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
