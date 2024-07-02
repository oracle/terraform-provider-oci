// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBackupsRequest wrapper for the ListBackups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListBackups.go.html to see an example of how to use ListBackupsRequest.
type ListBackupsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" contributesTo:"query" name:"databaseId"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// If provided, filters the results to the set of database versions which are supported for the given shape family.
	ShapeFamily ListBackupsShapeFamilyEnum `mandatory:"false" contributesTo:"query" name:"shapeFamily" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBackupsShapeFamilyEnum(string(request.ShapeFamily)); !ok && request.ShapeFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeFamily: %s. Supported values are: %s.", request.ShapeFamily, strings.Join(GetListBackupsShapeFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBackupsResponse wrapper for the ListBackups operation
type ListBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BackupSummary instances
	Items []BackupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBackupsShapeFamilyEnum Enum with underlying type: string
type ListBackupsShapeFamilyEnum string

// Set of constants representing the allowable values for ListBackupsShapeFamilyEnum
const (
	ListBackupsShapeFamilySinglenode     ListBackupsShapeFamilyEnum = "SINGLENODE"
	ListBackupsShapeFamilyYoda           ListBackupsShapeFamilyEnum = "YODA"
	ListBackupsShapeFamilyVirtualmachine ListBackupsShapeFamilyEnum = "VIRTUALMACHINE"
	ListBackupsShapeFamilyExadata        ListBackupsShapeFamilyEnum = "EXADATA"
	ListBackupsShapeFamilyExacc          ListBackupsShapeFamilyEnum = "EXACC"
	ListBackupsShapeFamilyExadbXs        ListBackupsShapeFamilyEnum = "EXADB_XS"
)

var mappingListBackupsShapeFamilyEnum = map[string]ListBackupsShapeFamilyEnum{
	"SINGLENODE":     ListBackupsShapeFamilySinglenode,
	"YODA":           ListBackupsShapeFamilyYoda,
	"VIRTUALMACHINE": ListBackupsShapeFamilyVirtualmachine,
	"EXADATA":        ListBackupsShapeFamilyExadata,
	"EXACC":          ListBackupsShapeFamilyExacc,
	"EXADB_XS":       ListBackupsShapeFamilyExadbXs,
}

var mappingListBackupsShapeFamilyEnumLowerCase = map[string]ListBackupsShapeFamilyEnum{
	"singlenode":     ListBackupsShapeFamilySinglenode,
	"yoda":           ListBackupsShapeFamilyYoda,
	"virtualmachine": ListBackupsShapeFamilyVirtualmachine,
	"exadata":        ListBackupsShapeFamilyExadata,
	"exacc":          ListBackupsShapeFamilyExacc,
	"exadb_xs":       ListBackupsShapeFamilyExadbXs,
}

// GetListBackupsShapeFamilyEnumValues Enumerates the set of values for ListBackupsShapeFamilyEnum
func GetListBackupsShapeFamilyEnumValues() []ListBackupsShapeFamilyEnum {
	values := make([]ListBackupsShapeFamilyEnum, 0)
	for _, v := range mappingListBackupsShapeFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListBackupsShapeFamilyEnumStringValues Enumerates the set of values in String for ListBackupsShapeFamilyEnum
func GetListBackupsShapeFamilyEnumStringValues() []string {
	return []string{
		"SINGLENODE",
		"YODA",
		"VIRTUALMACHINE",
		"EXADATA",
		"EXACC",
		"EXADB_XS",
	}
}

// GetMappingListBackupsShapeFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBackupsShapeFamilyEnum(val string) (ListBackupsShapeFamilyEnum, bool) {
	enum, ok := mappingListBackupsShapeFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
