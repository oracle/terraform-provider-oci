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

// ListRecommendedScheduledActionsRequest wrapper for the ListRecommendedScheduledActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListRecommendedScheduledActions.go.html to see an example of how to use ListRecommendedScheduledActionsRequest.
type ListRecommendedScheduledActionsRequest struct {

	// The Scheduling Policy OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SchedulingPolicyId *string `mandatory:"true" contributesTo:"path" name:"schedulingPolicyId"`

	// The target resource OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) the scheduled actions will be for.
	SchedulingPolicyTargetResourceId *string `mandatory:"true" contributesTo:"query" name:"schedulingPolicyTargetResourceId"`

	// The scheduling plan intent the scheduled actions will be for.
	PlanIntent ListRecommendedScheduledActionsPlanIntentEnum `mandatory:"true" contributesTo:"query" name:"planIntent" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendedScheduledActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendedScheduledActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendedScheduledActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendedScheduledActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecommendedScheduledActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecommendedScheduledActionsPlanIntentEnum(string(request.PlanIntent)); !ok && request.PlanIntent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanIntent: %s. Supported values are: %s.", request.PlanIntent, strings.Join(GetListRecommendedScheduledActionsPlanIntentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecommendedScheduledActionsResponse wrapper for the ListRecommendedScheduledActions operation
type ListRecommendedScheduledActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecommendedScheduledActionsCollection instances
	RecommendedScheduledActionsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRecommendedScheduledActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendedScheduledActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendedScheduledActionsPlanIntentEnum Enum with underlying type: string
type ListRecommendedScheduledActionsPlanIntentEnum string

// Set of constants representing the allowable values for ListRecommendedScheduledActionsPlanIntentEnum
const (
	ListRecommendedScheduledActionsPlanIntentExadataInfrastructureFullSoftwareUpdate ListRecommendedScheduledActionsPlanIntentEnum = "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"
)

var mappingListRecommendedScheduledActionsPlanIntentEnum = map[string]ListRecommendedScheduledActionsPlanIntentEnum{
	"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE": ListRecommendedScheduledActionsPlanIntentExadataInfrastructureFullSoftwareUpdate,
}

var mappingListRecommendedScheduledActionsPlanIntentEnumLowerCase = map[string]ListRecommendedScheduledActionsPlanIntentEnum{
	"exadata_infrastructure_full_software_update": ListRecommendedScheduledActionsPlanIntentExadataInfrastructureFullSoftwareUpdate,
}

// GetListRecommendedScheduledActionsPlanIntentEnumValues Enumerates the set of values for ListRecommendedScheduledActionsPlanIntentEnum
func GetListRecommendedScheduledActionsPlanIntentEnumValues() []ListRecommendedScheduledActionsPlanIntentEnum {
	values := make([]ListRecommendedScheduledActionsPlanIntentEnum, 0)
	for _, v := range mappingListRecommendedScheduledActionsPlanIntentEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecommendedScheduledActionsPlanIntentEnumStringValues Enumerates the set of values in String for ListRecommendedScheduledActionsPlanIntentEnum
func GetListRecommendedScheduledActionsPlanIntentEnumStringValues() []string {
	return []string{
		"EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingListRecommendedScheduledActionsPlanIntentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecommendedScheduledActionsPlanIntentEnum(val string) (ListRecommendedScheduledActionsPlanIntentEnum, bool) {
	enum, ok := mappingListRecommendedScheduledActionsPlanIntentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
