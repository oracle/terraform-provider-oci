// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRecommendationsRequest wrapper for the ListRecommendations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListRecommendations.go.html to see an example of how to use ListRecommendationsRequest.
type ListRecommendationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WAAS policy.
	WaasPolicyId *string `mandatory:"true" contributesTo:"path" name:"waasPolicyId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that matches recommended protection rules based on the selected action. If unspecified, rules with any action type are returned.
	RecommendedAction ListRecommendationsRecommendedActionEnum `mandatory:"false" contributesTo:"query" name:"recommendedAction" omitEmpty:"true"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecommendationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecommendationsRecommendedActionEnum(string(request.RecommendedAction)); !ok && request.RecommendedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecommendedAction: %s. Supported values are: %s.", request.RecommendedAction, strings.Join(GetListRecommendationsRecommendedActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecommendationsResponse wrapper for the ListRecommendations operation
type ListRecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []Recommendation instances
	Items []Recommendation `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendationsRecommendedActionEnum Enum with underlying type: string
type ListRecommendationsRecommendedActionEnum string

// Set of constants representing the allowable values for ListRecommendationsRecommendedActionEnum
const (
	ListRecommendationsRecommendedActionDetect ListRecommendationsRecommendedActionEnum = "DETECT"
	ListRecommendationsRecommendedActionBlock  ListRecommendationsRecommendedActionEnum = "BLOCK"
)

var mappingListRecommendationsRecommendedActionEnum = map[string]ListRecommendationsRecommendedActionEnum{
	"DETECT": ListRecommendationsRecommendedActionDetect,
	"BLOCK":  ListRecommendationsRecommendedActionBlock,
}

var mappingListRecommendationsRecommendedActionEnumLowerCase = map[string]ListRecommendationsRecommendedActionEnum{
	"detect": ListRecommendationsRecommendedActionDetect,
	"block":  ListRecommendationsRecommendedActionBlock,
}

// GetListRecommendationsRecommendedActionEnumValues Enumerates the set of values for ListRecommendationsRecommendedActionEnum
func GetListRecommendationsRecommendedActionEnumValues() []ListRecommendationsRecommendedActionEnum {
	values := make([]ListRecommendationsRecommendedActionEnum, 0)
	for _, v := range mappingListRecommendationsRecommendedActionEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendationsRecommendedActionEnumStringValues Enumerates the set of values in String for ListRecommendationsRecommendedActionEnum
func GetListRecommendationsRecommendedActionEnumStringValues() []string {
	return []string{
		"DETECT",
		"BLOCK",
	}
}

// GetMappingListRecommendationsRecommendedActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendationsRecommendedActionEnum(val string) (ListRecommendationsRecommendedActionEnum, bool) {
	enum, ok := mappingListRecommendationsRecommendedActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
