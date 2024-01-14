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

// ListSecurityFeatureAnalyticsRequest wrapper for the ListSecurityFeatureAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityFeatureAnalytics.go.html to see an example of how to use ListSecurityFeatureAnalyticsRequest.
type ListSecurityFeatureAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityFeatureAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityFeatureAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityFeatureAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityFeatureAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityFeatureAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityFeatureAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityFeatureAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityFeatureAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityFeatureAnalyticsResponse wrapper for the ListSecurityFeatureAnalytics operation
type ListSecurityFeatureAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SecurityFeatureAnalyticsCollection instance
	SecurityFeatureAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSecurityFeatureAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityFeatureAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityFeatureAnalyticsAccessLevelEnum Enum with underlying type: string
type ListSecurityFeatureAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityFeatureAnalyticsAccessLevelEnum
const (
	ListSecurityFeatureAnalyticsAccessLevelRestricted ListSecurityFeatureAnalyticsAccessLevelEnum = "RESTRICTED"
	ListSecurityFeatureAnalyticsAccessLevelAccessible ListSecurityFeatureAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityFeatureAnalyticsAccessLevelEnum = map[string]ListSecurityFeatureAnalyticsAccessLevelEnum{
	"RESTRICTED": ListSecurityFeatureAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityFeatureAnalyticsAccessLevelAccessible,
}

var mappingListSecurityFeatureAnalyticsAccessLevelEnumLowerCase = map[string]ListSecurityFeatureAnalyticsAccessLevelEnum{
	"restricted": ListSecurityFeatureAnalyticsAccessLevelRestricted,
	"accessible": ListSecurityFeatureAnalyticsAccessLevelAccessible,
}

// GetListSecurityFeatureAnalyticsAccessLevelEnumValues Enumerates the set of values for ListSecurityFeatureAnalyticsAccessLevelEnum
func GetListSecurityFeatureAnalyticsAccessLevelEnumValues() []ListSecurityFeatureAnalyticsAccessLevelEnum {
	values := make([]ListSecurityFeatureAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListSecurityFeatureAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeatureAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityFeatureAnalyticsAccessLevelEnum
func GetListSecurityFeatureAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityFeatureAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeatureAnalyticsAccessLevelEnum(val string) (ListSecurityFeatureAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityFeatureAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
