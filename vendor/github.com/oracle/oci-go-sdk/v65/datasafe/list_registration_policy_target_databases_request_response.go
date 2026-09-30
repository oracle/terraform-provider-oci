// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRegistrationPolicyTargetDatabasesRequest wrapper for the ListRegistrationPolicyTargetDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListRegistrationPolicyTargetDatabases.go.html to see an example of how to use ListRegistrationPolicyTargetDatabasesRequest.
type ListRegistrationPolicyTargetDatabasesRequest struct {

	// The OCID of the registration policy to be used for identification
	RegistrationPolicyId *string `mandatory:"true" contributesTo:"path" name:"registrationPolicyId"`

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return the target database only if it is registered via the registration policy.
	TargetDatabaseId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseId"`

	// Filters registered targets associated with this registration policy by membership status.
	// - OPTIN: returns targets that are included (opted in) by the policy.
	// - OPTOUT: returns targets that are explicitly excluded (opted out) by the policy.
	MembershipStatus ListRegistrationPolicyTargetDatabasesMembershipStatusEnum `mandatory:"false" contributesTo:"query" name:"membershipStatus" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRegistrationPolicyTargetDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRegistrationPolicyTargetDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRegistrationPolicyTargetDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRegistrationPolicyTargetDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRegistrationPolicyTargetDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRegistrationPolicyTargetDatabasesMembershipStatusEnum(string(request.MembershipStatus)); !ok && request.MembershipStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MembershipStatus: %s. Supported values are: %s.", request.MembershipStatus, strings.Join(GetListRegistrationPolicyTargetDatabasesMembershipStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRegistrationPolicyTargetDatabasesResponse wrapper for the ListRegistrationPolicyTargetDatabases operation
type ListRegistrationPolicyTargetDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RegistrationPolicyTargetDatabaseSummaryCollection instances
	RegistrationPolicyTargetDatabaseSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRegistrationPolicyTargetDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRegistrationPolicyTargetDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRegistrationPolicyTargetDatabasesMembershipStatusEnum Enum with underlying type: string
type ListRegistrationPolicyTargetDatabasesMembershipStatusEnum string

// Set of constants representing the allowable values for ListRegistrationPolicyTargetDatabasesMembershipStatusEnum
const (
	ListRegistrationPolicyTargetDatabasesMembershipStatusOptin  ListRegistrationPolicyTargetDatabasesMembershipStatusEnum = "OPTIN"
	ListRegistrationPolicyTargetDatabasesMembershipStatusOptout ListRegistrationPolicyTargetDatabasesMembershipStatusEnum = "OPTOUT"
)

var mappingListRegistrationPolicyTargetDatabasesMembershipStatusEnum = map[string]ListRegistrationPolicyTargetDatabasesMembershipStatusEnum{
	"OPTIN":  ListRegistrationPolicyTargetDatabasesMembershipStatusOptin,
	"OPTOUT": ListRegistrationPolicyTargetDatabasesMembershipStatusOptout,
}

var mappingListRegistrationPolicyTargetDatabasesMembershipStatusEnumLowerCase = map[string]ListRegistrationPolicyTargetDatabasesMembershipStatusEnum{
	"optin":  ListRegistrationPolicyTargetDatabasesMembershipStatusOptin,
	"optout": ListRegistrationPolicyTargetDatabasesMembershipStatusOptout,
}

// GetListRegistrationPolicyTargetDatabasesMembershipStatusEnumValues Enumerates the set of values for ListRegistrationPolicyTargetDatabasesMembershipStatusEnum
func GetListRegistrationPolicyTargetDatabasesMembershipStatusEnumValues() []ListRegistrationPolicyTargetDatabasesMembershipStatusEnum {
	values := make([]ListRegistrationPolicyTargetDatabasesMembershipStatusEnum, 0)
	for _, v := range mappingListRegistrationPolicyTargetDatabasesMembershipStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListRegistrationPolicyTargetDatabasesMembershipStatusEnumStringValues Enumerates the set of values in String for ListRegistrationPolicyTargetDatabasesMembershipStatusEnum
func GetListRegistrationPolicyTargetDatabasesMembershipStatusEnumStringValues() []string {
	return []string{
		"OPTIN",
		"OPTOUT",
	}
}

// GetMappingListRegistrationPolicyTargetDatabasesMembershipStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRegistrationPolicyTargetDatabasesMembershipStatusEnum(val string) (ListRegistrationPolicyTargetDatabasesMembershipStatusEnum, bool) {
	enum, ok := mappingListRegistrationPolicyTargetDatabasesMembershipStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
