// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEgressDisintermediatedRoutesRequest wrapper for the ListEgressDisintermediatedRoutes operation
type ListEgressDisintermediatedRoutesRequest struct {

	// Route table label
	Label *int `mandatory:"true" contributesTo:"query" name:"label"`

	// ad name
	AdName ListEgressDisintermediatedRoutesAdNameEnum `mandatory:"true" contributesTo:"query" name:"adName" omitEmpty:"true"`

	// shard id
	ShardId *int `mandatory:"true" contributesTo:"query" name:"shardId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table assigned to the DRG attachment.
	DrgRouteTableId *string `mandatory:"false" contributesTo:"query" name:"drgRouteTableId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEgressDisintermediatedRoutesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEgressDisintermediatedRoutesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEgressDisintermediatedRoutesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListEgressDisintermediatedRoutesRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEgressDisintermediatedRoutesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEgressDisintermediatedRoutesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEgressDisintermediatedRoutesAdNameEnum(string(request.AdName)); !ok && request.AdName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdName: %s. Supported values are: %s.", request.AdName, strings.Join(GetListEgressDisintermediatedRoutesAdNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEgressDisintermediatedRoutesResponse wrapper for the ListEgressDisintermediatedRoutes operation
type ListEgressDisintermediatedRoutesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []EgressDisintermediatedRoute instances
	Items []EgressDisintermediatedRoute `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEgressDisintermediatedRoutesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEgressDisintermediatedRoutesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEgressDisintermediatedRoutesAdNameEnum Enum with underlying type: string
type ListEgressDisintermediatedRoutesAdNameEnum string

// Set of constants representing the allowable values for ListEgressDisintermediatedRoutesAdNameEnum
const (
	ListEgressDisintermediatedRoutesAdNameAd1  ListEgressDisintermediatedRoutesAdNameEnum = "AD1"
	ListEgressDisintermediatedRoutesAdNameAd2  ListEgressDisintermediatedRoutesAdNameEnum = "AD2"
	ListEgressDisintermediatedRoutesAdNameAd3  ListEgressDisintermediatedRoutesAdNameEnum = "AD3"
	ListEgressDisintermediatedRoutesAdNamePop1 ListEgressDisintermediatedRoutesAdNameEnum = "POP1"
	ListEgressDisintermediatedRoutesAdNamePop2 ListEgressDisintermediatedRoutesAdNameEnum = "POP2"
)

var mappingListEgressDisintermediatedRoutesAdNameEnum = map[string]ListEgressDisintermediatedRoutesAdNameEnum{
	"AD1":  ListEgressDisintermediatedRoutesAdNameAd1,
	"AD2":  ListEgressDisintermediatedRoutesAdNameAd2,
	"AD3":  ListEgressDisintermediatedRoutesAdNameAd3,
	"POP1": ListEgressDisintermediatedRoutesAdNamePop1,
	"POP2": ListEgressDisintermediatedRoutesAdNamePop2,
}

var mappingListEgressDisintermediatedRoutesAdNameEnumLowerCase = map[string]ListEgressDisintermediatedRoutesAdNameEnum{
	"ad1":  ListEgressDisintermediatedRoutesAdNameAd1,
	"ad2":  ListEgressDisintermediatedRoutesAdNameAd2,
	"ad3":  ListEgressDisintermediatedRoutesAdNameAd3,
	"pop1": ListEgressDisintermediatedRoutesAdNamePop1,
	"pop2": ListEgressDisintermediatedRoutesAdNamePop2,
}

// GetListEgressDisintermediatedRoutesAdNameEnumValues Enumerates the set of values for ListEgressDisintermediatedRoutesAdNameEnum
func GetListEgressDisintermediatedRoutesAdNameEnumValues() []ListEgressDisintermediatedRoutesAdNameEnum {
	values := make([]ListEgressDisintermediatedRoutesAdNameEnum, 0)
	for _, v := range mappingListEgressDisintermediatedRoutesAdNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListEgressDisintermediatedRoutesAdNameEnumStringValues Enumerates the set of values in String for ListEgressDisintermediatedRoutesAdNameEnum
func GetListEgressDisintermediatedRoutesAdNameEnumStringValues() []string {
	return []string{
		"AD1",
		"AD2",
		"AD3",
		"POP1",
		"POP2",
	}
}

// GetMappingListEgressDisintermediatedRoutesAdNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEgressDisintermediatedRoutesAdNameEnum(val string) (ListEgressDisintermediatedRoutesAdNameEnum, bool) {
	enum, ok := mappingListEgressDisintermediatedRoutesAdNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
