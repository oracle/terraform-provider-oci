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

// ListDatabaseTableAccessEntriesRequest wrapper for the ListDatabaseTableAccessEntries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDatabaseTableAccessEntries.go.html to see an example of how to use ListDatabaseTableAccessEntriesRequest.
type ListDatabaseTableAccessEntriesRequest struct {

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

	// The field to sort by. Only one sort parameter should be provided.
	SortBy ListDatabaseTableAccessEntriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDatabaseTableAccessEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseTableAccessEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseTableAccessEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseTableAccessEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseTableAccessEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseTableAccessEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseTableAccessEntriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseTableAccessEntriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseTableAccessEntriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseTableAccessEntriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseTableAccessEntriesResponse wrapper for the ListDatabaseTableAccessEntries operation
type ListDatabaseTableAccessEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseTableAccessEntryCollection instances
	DatabaseTableAccessEntryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDatabaseTableAccessEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseTableAccessEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseTableAccessEntriesSortByEnum Enum with underlying type: string
type ListDatabaseTableAccessEntriesSortByEnum string

// Set of constants representing the allowable values for ListDatabaseTableAccessEntriesSortByEnum
const (
	ListDatabaseTableAccessEntriesSortByKey                                          ListDatabaseTableAccessEntriesSortByEnum = "key"
	ListDatabaseTableAccessEntriesSortByGrantee                                      ListDatabaseTableAccessEntriesSortByEnum = "grantee"
	ListDatabaseTableAccessEntriesSortByAccesstype                                   ListDatabaseTableAccessEntriesSortByEnum = "accessType"
	ListDatabaseTableAccessEntriesSortByTableschema                                  ListDatabaseTableAccessEntriesSortByEnum = "tableSchema"
	ListDatabaseTableAccessEntriesSortByTablename                                    ListDatabaseTableAccessEntriesSortByEnum = "tableName"
	ListDatabaseTableAccessEntriesSortByPrivilegetype                                ListDatabaseTableAccessEntriesSortByEnum = "privilegeType"
	ListDatabaseTableAccessEntriesSortByPrivilege                                    ListDatabaseTableAccessEntriesSortByEnum = "privilege"
	ListDatabaseTableAccessEntriesSortByPrivilegegrantable                           ListDatabaseTableAccessEntriesSortByEnum = "privilegeGrantable"
	ListDatabaseTableAccessEntriesSortByGrantfromrole                                ListDatabaseTableAccessEntriesSortByEnum = "grantFromRole"
	ListDatabaseTableAccessEntriesSortByAccessthroughobject                          ListDatabaseTableAccessEntriesSortByEnum = "accessThroughObject"
	ListDatabaseTableAccessEntriesSortByColumnname                                   ListDatabaseTableAccessEntriesSortByEnum = "columnName"
	ListDatabaseTableAccessEntriesSortByGrantor                                      ListDatabaseTableAccessEntriesSortByEnum = "grantor"
	ListDatabaseTableAccessEntriesSortByArealltablesaccessible                       ListDatabaseTableAccessEntriesSortByEnum = "areAllTablesAccessible"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyview                    ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedByView"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbylabelsecurity           ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedByLabelSecurity"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbydatabasevault           ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedByDatabaseVault"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyvirtualprivatedatabase  ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedByVirtualPrivateDatabase"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyredaction               ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedByRedaction"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyrealapplicationsecurity ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedByRealApplicationSecurity"
	ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbysqlfirewall             ListDatabaseTableAccessEntriesSortByEnum = "isAccessConstrainedBySqlFirewall"
	ListDatabaseTableAccessEntriesSortByIssensitive                                  ListDatabaseTableAccessEntriesSortByEnum = "isSensitive"
)

var mappingListDatabaseTableAccessEntriesSortByEnum = map[string]ListDatabaseTableAccessEntriesSortByEnum{
	"key":                                ListDatabaseTableAccessEntriesSortByKey,
	"grantee":                            ListDatabaseTableAccessEntriesSortByGrantee,
	"accessType":                         ListDatabaseTableAccessEntriesSortByAccesstype,
	"tableSchema":                        ListDatabaseTableAccessEntriesSortByTableschema,
	"tableName":                          ListDatabaseTableAccessEntriesSortByTablename,
	"privilegeType":                      ListDatabaseTableAccessEntriesSortByPrivilegetype,
	"privilege":                          ListDatabaseTableAccessEntriesSortByPrivilege,
	"privilegeGrantable":                 ListDatabaseTableAccessEntriesSortByPrivilegegrantable,
	"grantFromRole":                      ListDatabaseTableAccessEntriesSortByGrantfromrole,
	"accessThroughObject":                ListDatabaseTableAccessEntriesSortByAccessthroughobject,
	"columnName":                         ListDatabaseTableAccessEntriesSortByColumnname,
	"grantor":                            ListDatabaseTableAccessEntriesSortByGrantor,
	"areAllTablesAccessible":             ListDatabaseTableAccessEntriesSortByArealltablesaccessible,
	"isAccessConstrainedByView":          ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyview,
	"isAccessConstrainedByLabelSecurity": ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbylabelsecurity,
	"isAccessConstrainedByDatabaseVault": ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbydatabasevault,
	"isAccessConstrainedByVirtualPrivateDatabase":  ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyvirtualprivatedatabase,
	"isAccessConstrainedByRedaction":               ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyredaction,
	"isAccessConstrainedByRealApplicationSecurity": ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyrealapplicationsecurity,
	"isAccessConstrainedBySqlFirewall":             ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbysqlfirewall,
	"isSensitive":                                  ListDatabaseTableAccessEntriesSortByIssensitive,
}

var mappingListDatabaseTableAccessEntriesSortByEnumLowerCase = map[string]ListDatabaseTableAccessEntriesSortByEnum{
	"key":                                ListDatabaseTableAccessEntriesSortByKey,
	"grantee":                            ListDatabaseTableAccessEntriesSortByGrantee,
	"accesstype":                         ListDatabaseTableAccessEntriesSortByAccesstype,
	"tableschema":                        ListDatabaseTableAccessEntriesSortByTableschema,
	"tablename":                          ListDatabaseTableAccessEntriesSortByTablename,
	"privilegetype":                      ListDatabaseTableAccessEntriesSortByPrivilegetype,
	"privilege":                          ListDatabaseTableAccessEntriesSortByPrivilege,
	"privilegegrantable":                 ListDatabaseTableAccessEntriesSortByPrivilegegrantable,
	"grantfromrole":                      ListDatabaseTableAccessEntriesSortByGrantfromrole,
	"accessthroughobject":                ListDatabaseTableAccessEntriesSortByAccessthroughobject,
	"columnname":                         ListDatabaseTableAccessEntriesSortByColumnname,
	"grantor":                            ListDatabaseTableAccessEntriesSortByGrantor,
	"arealltablesaccessible":             ListDatabaseTableAccessEntriesSortByArealltablesaccessible,
	"isaccessconstrainedbyview":          ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyview,
	"isaccessconstrainedbylabelsecurity": ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbylabelsecurity,
	"isaccessconstrainedbydatabasevault": ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbydatabasevault,
	"isaccessconstrainedbyvirtualprivatedatabase":  ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyvirtualprivatedatabase,
	"isaccessconstrainedbyredaction":               ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyredaction,
	"isaccessconstrainedbyrealapplicationsecurity": ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbyrealapplicationsecurity,
	"isaccessconstrainedbysqlfirewall":             ListDatabaseTableAccessEntriesSortByIsaccessconstrainedbysqlfirewall,
	"issensitive":                                  ListDatabaseTableAccessEntriesSortByIssensitive,
}

// GetListDatabaseTableAccessEntriesSortByEnumValues Enumerates the set of values for ListDatabaseTableAccessEntriesSortByEnum
func GetListDatabaseTableAccessEntriesSortByEnumValues() []ListDatabaseTableAccessEntriesSortByEnum {
	values := make([]ListDatabaseTableAccessEntriesSortByEnum, 0)
	for _, v := range mappingListDatabaseTableAccessEntriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseTableAccessEntriesSortByEnumStringValues Enumerates the set of values in String for ListDatabaseTableAccessEntriesSortByEnum
func GetListDatabaseTableAccessEntriesSortByEnumStringValues() []string {
	return []string{
		"key",
		"grantee",
		"accessType",
		"tableSchema",
		"tableName",
		"privilegeType",
		"privilege",
		"privilegeGrantable",
		"grantFromRole",
		"accessThroughObject",
		"columnName",
		"grantor",
		"areAllTablesAccessible",
		"isAccessConstrainedByView",
		"isAccessConstrainedByLabelSecurity",
		"isAccessConstrainedByDatabaseVault",
		"isAccessConstrainedByVirtualPrivateDatabase",
		"isAccessConstrainedByRedaction",
		"isAccessConstrainedByRealApplicationSecurity",
		"isAccessConstrainedBySqlFirewall",
		"isSensitive",
	}
}

// GetMappingListDatabaseTableAccessEntriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseTableAccessEntriesSortByEnum(val string) (ListDatabaseTableAccessEntriesSortByEnum, bool) {
	enum, ok := mappingListDatabaseTableAccessEntriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseTableAccessEntriesSortOrderEnum Enum with underlying type: string
type ListDatabaseTableAccessEntriesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseTableAccessEntriesSortOrderEnum
const (
	ListDatabaseTableAccessEntriesSortOrderAsc  ListDatabaseTableAccessEntriesSortOrderEnum = "ASC"
	ListDatabaseTableAccessEntriesSortOrderDesc ListDatabaseTableAccessEntriesSortOrderEnum = "DESC"
)

var mappingListDatabaseTableAccessEntriesSortOrderEnum = map[string]ListDatabaseTableAccessEntriesSortOrderEnum{
	"ASC":  ListDatabaseTableAccessEntriesSortOrderAsc,
	"DESC": ListDatabaseTableAccessEntriesSortOrderDesc,
}

var mappingListDatabaseTableAccessEntriesSortOrderEnumLowerCase = map[string]ListDatabaseTableAccessEntriesSortOrderEnum{
	"asc":  ListDatabaseTableAccessEntriesSortOrderAsc,
	"desc": ListDatabaseTableAccessEntriesSortOrderDesc,
}

// GetListDatabaseTableAccessEntriesSortOrderEnumValues Enumerates the set of values for ListDatabaseTableAccessEntriesSortOrderEnum
func GetListDatabaseTableAccessEntriesSortOrderEnumValues() []ListDatabaseTableAccessEntriesSortOrderEnum {
	values := make([]ListDatabaseTableAccessEntriesSortOrderEnum, 0)
	for _, v := range mappingListDatabaseTableAccessEntriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseTableAccessEntriesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseTableAccessEntriesSortOrderEnum
func GetListDatabaseTableAccessEntriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseTableAccessEntriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseTableAccessEntriesSortOrderEnum(val string) (ListDatabaseTableAccessEntriesSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseTableAccessEntriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
