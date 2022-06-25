// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDatabaseCharacterSetsRequest wrapper for the ListAutonomousDatabaseCharacterSets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDatabaseCharacterSets.go.html to see an example of how to use ListAutonomousDatabaseCharacterSetsRequest.
type ListAutonomousDatabaseCharacterSetsRequest struct {

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies whether this request is for Autonomous Database on Shared infrastructure. By default, this request will be for Autonomous Database on Dedicated Exadata Infrastructure.
	IsShared *bool `mandatory:"false" contributesTo:"query" name:"isShared"`

	// Specifies whether this request pertains to database character sets or national character sets.
	CharacterSetType ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum `mandatory:"false" contributesTo:"query" name:"characterSetType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDatabaseCharacterSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDatabaseCharacterSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDatabaseCharacterSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDatabaseCharacterSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDatabaseCharacterSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum(string(request.CharacterSetType)); !ok && request.CharacterSetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CharacterSetType: %s. Supported values are: %s.", request.CharacterSetType, strings.Join(GetListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDatabaseCharacterSetsResponse wrapper for the ListAutonomousDatabaseCharacterSets operation
type ListAutonomousDatabaseCharacterSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []AutonomousDatabaseCharacterSets instance
	Items []AutonomousDatabaseCharacterSets `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDatabaseCharacterSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDatabaseCharacterSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum Enum with underlying type: string
type ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum string

// Set of constants representing the allowable values for ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum
const (
	ListAutonomousDatabaseCharacterSetsCharacterSetTypeDatabase ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum = "DATABASE"
	ListAutonomousDatabaseCharacterSetsCharacterSetTypeNational ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum = "NATIONAL"
)

var mappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum = map[string]ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum{
	"DATABASE": ListAutonomousDatabaseCharacterSetsCharacterSetTypeDatabase,
	"NATIONAL": ListAutonomousDatabaseCharacterSetsCharacterSetTypeNational,
}

var mappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumLowerCase = map[string]ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum{
	"database": ListAutonomousDatabaseCharacterSetsCharacterSetTypeDatabase,
	"national": ListAutonomousDatabaseCharacterSetsCharacterSetTypeNational,
}

// GetListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumValues Enumerates the set of values for ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum
func GetListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumValues() []ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum {
	values := make([]ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum, 0)
	for _, v := range mappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumStringValues Enumerates the set of values in String for ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum
func GetListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"NATIONAL",
	}
}

// GetMappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum(val string) (ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum, bool) {
	enum, ok := mappingListAutonomousDatabaseCharacterSetsCharacterSetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
