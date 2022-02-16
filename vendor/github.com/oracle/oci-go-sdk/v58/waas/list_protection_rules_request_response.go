// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListProtectionRulesRequest wrapper for the ListProtectionRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListProtectionRules.go.html to see an example of how to use ListProtectionRulesRequest.
type ListProtectionRulesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WAAS policy.
	WaasPolicyId *string `mandatory:"true" contributesTo:"path" name:"waasPolicyId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter rules using a list of ModSecurity rule IDs.
	ModSecurityRuleId []string `contributesTo:"query" name:"modSecurityRuleId" collectionFormat:"multi"`

	// Filter rules using a list of actions.
	Action []ListProtectionRulesActionEnum `contributesTo:"query" name:"action" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProtectionRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Action {
		if _, ok := GetMappingListProtectionRulesActionEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", val, strings.Join(GetListProtectionRulesActionEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProtectionRulesResponse wrapper for the ListProtectionRules operation
type ListProtectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ProtectionRule instances
	Items []ProtectionRule `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListProtectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectionRulesActionEnum Enum with underlying type: string
type ListProtectionRulesActionEnum string

// Set of constants representing the allowable values for ListProtectionRulesActionEnum
const (
	ListProtectionRulesActionOff    ListProtectionRulesActionEnum = "OFF"
	ListProtectionRulesActionDetect ListProtectionRulesActionEnum = "DETECT"
	ListProtectionRulesActionBlock  ListProtectionRulesActionEnum = "BLOCK"
)

var mappingListProtectionRulesActionEnum = map[string]ListProtectionRulesActionEnum{
	"OFF":    ListProtectionRulesActionOff,
	"DETECT": ListProtectionRulesActionDetect,
	"BLOCK":  ListProtectionRulesActionBlock,
}

// GetListProtectionRulesActionEnumValues Enumerates the set of values for ListProtectionRulesActionEnum
func GetListProtectionRulesActionEnumValues() []ListProtectionRulesActionEnum {
	values := make([]ListProtectionRulesActionEnum, 0)
	for _, v := range mappingListProtectionRulesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionRulesActionEnumStringValues Enumerates the set of values in String for ListProtectionRulesActionEnum
func GetListProtectionRulesActionEnumStringValues() []string {
	return []string{
		"OFF",
		"DETECT",
		"BLOCK",
	}
}

// GetMappingListProtectionRulesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionRulesActionEnum(val string) (ListProtectionRulesActionEnum, bool) {
	mappingListProtectionRulesActionEnumIgnoreCase := make(map[string]ListProtectionRulesActionEnum)
	for k, v := range mappingListProtectionRulesActionEnum {
		mappingListProtectionRulesActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListProtectionRulesActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
