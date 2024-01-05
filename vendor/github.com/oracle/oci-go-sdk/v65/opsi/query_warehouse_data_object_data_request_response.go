// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// QueryWarehouseDataObjectDataRequest wrapper for the QueryWarehouseDataObjectData operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/QueryWarehouseDataObjectData.go.html to see an example of how to use QueryWarehouseDataObjectDataRequest.
type QueryWarehouseDataObjectDataRequest struct {

	// Type of the Warehouse.
	WarehouseType QueryWarehouseDataObjectDataWarehouseTypeEnum `mandatory:"true" contributesTo:"path" name:"warehouseType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Warehouse.
	WarehouseId *string `mandatory:"true" contributesTo:"path" name:"warehouseId"`

	// The information to be used for querying a Warehouse.
	QueryWarehouseDataObjectDataDetails `contributesTo:"body"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request QueryWarehouseDataObjectDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request QueryWarehouseDataObjectDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request QueryWarehouseDataObjectDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request QueryWarehouseDataObjectDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request QueryWarehouseDataObjectDataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryWarehouseDataObjectDataWarehouseTypeEnum(string(request.WarehouseType)); !ok && request.WarehouseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WarehouseType: %s. Supported values are: %s.", request.WarehouseType, strings.Join(GetQueryWarehouseDataObjectDataWarehouseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryWarehouseDataObjectDataResponse wrapper for the QueryWarehouseDataObjectData operation
type QueryWarehouseDataObjectDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QueryDataObjectResultSetRowsCollection instances
	QueryDataObjectResultSetRowsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response QueryWarehouseDataObjectDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response QueryWarehouseDataObjectDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// QueryWarehouseDataObjectDataWarehouseTypeEnum Enum with underlying type: string
type QueryWarehouseDataObjectDataWarehouseTypeEnum string

// Set of constants representing the allowable values for QueryWarehouseDataObjectDataWarehouseTypeEnum
const (
	QueryWarehouseDataObjectDataWarehouseTypeAwrhubs QueryWarehouseDataObjectDataWarehouseTypeEnum = "awrHubs"
)

var mappingQueryWarehouseDataObjectDataWarehouseTypeEnum = map[string]QueryWarehouseDataObjectDataWarehouseTypeEnum{
	"awrHubs": QueryWarehouseDataObjectDataWarehouseTypeAwrhubs,
}

var mappingQueryWarehouseDataObjectDataWarehouseTypeEnumLowerCase = map[string]QueryWarehouseDataObjectDataWarehouseTypeEnum{
	"awrhubs": QueryWarehouseDataObjectDataWarehouseTypeAwrhubs,
}

// GetQueryWarehouseDataObjectDataWarehouseTypeEnumValues Enumerates the set of values for QueryWarehouseDataObjectDataWarehouseTypeEnum
func GetQueryWarehouseDataObjectDataWarehouseTypeEnumValues() []QueryWarehouseDataObjectDataWarehouseTypeEnum {
	values := make([]QueryWarehouseDataObjectDataWarehouseTypeEnum, 0)
	for _, v := range mappingQueryWarehouseDataObjectDataWarehouseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryWarehouseDataObjectDataWarehouseTypeEnumStringValues Enumerates the set of values in String for QueryWarehouseDataObjectDataWarehouseTypeEnum
func GetQueryWarehouseDataObjectDataWarehouseTypeEnumStringValues() []string {
	return []string{
		"awrHubs",
	}
}

// GetMappingQueryWarehouseDataObjectDataWarehouseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryWarehouseDataObjectDataWarehouseTypeEnum(val string) (QueryWarehouseDataObjectDataWarehouseTypeEnum, bool) {
	enum, ok := mappingQueryWarehouseDataObjectDataWarehouseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
