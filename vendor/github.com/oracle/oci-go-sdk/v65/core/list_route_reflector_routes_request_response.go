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

// ListRouteReflectorRoutesRequest wrapper for the ListRouteReflectorRoutes operation
type ListRouteReflectorRoutesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
	DrgRouteTableId *string `mandatory:"true" contributesTo:"path" name:"drgRouteTableId"`

	// The infobase
	Infobase ListRouteReflectorRoutesInfobaseEnum `mandatory:"true" contributesTo:"query" name:"infobase" omitEmpty:"true"`

	// The vrf label
	Vrf *int `mandatory:"true" contributesTo:"query" name:"vrf"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRouteReflectorRoutesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRouteReflectorRoutesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRouteReflectorRoutesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListRouteReflectorRoutesRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["drgRouteTableId"] != nil {
		templateParam := mandatoryParamMap["drgRouteTableId"]
		for _, template := range templateParam {
			replacementParam := *request.DrgRouteTableId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRouteReflectorRoutesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRouteReflectorRoutesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRouteReflectorRoutesInfobaseEnum(string(request.Infobase)); !ok && request.Infobase != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Infobase: %s. Supported values are: %s.", request.Infobase, strings.Join(GetListRouteReflectorRoutesInfobaseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRouteReflectorRoutesResponse wrapper for the ListRouteReflectorRoutes operation
type ListRouteReflectorRoutesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []RouteReflectorRoute instances
	Items []RouteReflectorRoute `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRouteReflectorRoutesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRouteReflectorRoutesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRouteReflectorRoutesInfobaseEnum Enum with underlying type: string
type ListRouteReflectorRoutesInfobaseEnum string

// Set of constants representing the allowable values for ListRouteReflectorRoutesInfobaseEnum
const (
	ListRouteReflectorRoutesInfobaseFib ListRouteReflectorRoutesInfobaseEnum = "FIB"
	ListRouteReflectorRoutesInfobaseRib ListRouteReflectorRoutesInfobaseEnum = "RIB"
)

var mappingListRouteReflectorRoutesInfobaseEnum = map[string]ListRouteReflectorRoutesInfobaseEnum{
	"FIB": ListRouteReflectorRoutesInfobaseFib,
	"RIB": ListRouteReflectorRoutesInfobaseRib,
}

var mappingListRouteReflectorRoutesInfobaseEnumLowerCase = map[string]ListRouteReflectorRoutesInfobaseEnum{
	"fib": ListRouteReflectorRoutesInfobaseFib,
	"rib": ListRouteReflectorRoutesInfobaseRib,
}

// GetListRouteReflectorRoutesInfobaseEnumValues Enumerates the set of values for ListRouteReflectorRoutesInfobaseEnum
func GetListRouteReflectorRoutesInfobaseEnumValues() []ListRouteReflectorRoutesInfobaseEnum {
	values := make([]ListRouteReflectorRoutesInfobaseEnum, 0)
	for _, v := range mappingListRouteReflectorRoutesInfobaseEnum {
		values = append(values, v)
	}
	return values
}

// GetListRouteReflectorRoutesInfobaseEnumStringValues Enumerates the set of values in String for ListRouteReflectorRoutesInfobaseEnum
func GetListRouteReflectorRoutesInfobaseEnumStringValues() []string {
	return []string{
		"FIB",
		"RIB",
	}
}

// GetMappingListRouteReflectorRoutesInfobaseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRouteReflectorRoutesInfobaseEnum(val string) (ListRouteReflectorRoutesInfobaseEnum, bool) {
	enum, ok := mappingListRouteReflectorRoutesInfobaseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
