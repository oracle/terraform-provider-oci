// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListC3DrgRouteRulesRequest wrapper for the ListC3DrgRouteRules operation
type ListC3DrgRouteRulesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
	DrgRouteTableId *string `mandatory:"true" contributesTo:"path" name:"drgRouteTableId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Static routes are specified through the DRG route table API.
	// Dynamic routes are learned by the DRG from the DRG attachments through various routing protocols.
	RouteType ListC3DrgRouteRulesRouteTypeEnum `mandatory:"false" contributesTo:"query" name:"routeType" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListC3DrgRouteRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListC3DrgRouteRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListC3DrgRouteRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListC3DrgRouteRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListC3DrgRouteRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListC3DrgRouteRulesRouteTypeEnum(string(request.RouteType)); !ok && request.RouteType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RouteType: %s. Supported values are: %s.", request.RouteType, strings.Join(GetListC3DrgRouteRulesRouteTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListC3DrgRouteRulesResponse wrapper for the ListC3DrgRouteRules operation
type ListC3DrgRouteRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DrgRouteRule instances
	Items []DrgRouteRule `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListC3DrgRouteRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListC3DrgRouteRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListC3DrgRouteRulesRouteTypeEnum Enum with underlying type: string
type ListC3DrgRouteRulesRouteTypeEnum string

// Set of constants representing the allowable values for ListC3DrgRouteRulesRouteTypeEnum
const (
	ListC3DrgRouteRulesRouteTypeStatic  ListC3DrgRouteRulesRouteTypeEnum = "STATIC"
	ListC3DrgRouteRulesRouteTypeDynamic ListC3DrgRouteRulesRouteTypeEnum = "DYNAMIC"
)

var mappingListC3DrgRouteRulesRouteTypeEnum = map[string]ListC3DrgRouteRulesRouteTypeEnum{
	"STATIC":  ListC3DrgRouteRulesRouteTypeStatic,
	"DYNAMIC": ListC3DrgRouteRulesRouteTypeDynamic,
}

var mappingListC3DrgRouteRulesRouteTypeEnumLowerCase = map[string]ListC3DrgRouteRulesRouteTypeEnum{
	"static":  ListC3DrgRouteRulesRouteTypeStatic,
	"dynamic": ListC3DrgRouteRulesRouteTypeDynamic,
}

// GetListC3DrgRouteRulesRouteTypeEnumValues Enumerates the set of values for ListC3DrgRouteRulesRouteTypeEnum
func GetListC3DrgRouteRulesRouteTypeEnumValues() []ListC3DrgRouteRulesRouteTypeEnum {
	values := make([]ListC3DrgRouteRulesRouteTypeEnum, 0)
	for _, v := range mappingListC3DrgRouteRulesRouteTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListC3DrgRouteRulesRouteTypeEnumStringValues Enumerates the set of values in String for ListC3DrgRouteRulesRouteTypeEnum
func GetListC3DrgRouteRulesRouteTypeEnumStringValues() []string {
	return []string{
		"STATIC",
		"DYNAMIC",
	}
}

// GetMappingListC3DrgRouteRulesRouteTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListC3DrgRouteRulesRouteTypeEnum(val string) (ListC3DrgRouteRulesRouteTypeEnum, bool) {
	enum, ok := mappingListC3DrgRouteRulesRouteTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
