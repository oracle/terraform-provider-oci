// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListContainersRequest wrapper for the ListContainers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListContainers.go.html to see an example of how to use ListContainersRequest.
type ListContainersRequest struct {

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// if true, this returns latest version of container.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the container name.
	ContainerName *string `mandatory:"false" contributesTo:"query" name:"containerName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	//   state for the resource type.
	LifecycleState ListContainersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the target workload.
	TargetWorkload ListContainersTargetWorkloadEnum `mandatory:"false" contributesTo:"query" name:"targetWorkload" omitEmpty:"true"`

	// <b>Filter</b> results by the usage.
	UsageQueryParam ListContainersUsageQueryParamEnum `mandatory:"false" contributesTo:"query" name:"usageQueryParam" omitEmpty:"true"`

	// <b>Filter</b> results by the container version tag.
	TagQueryParam *string `mandatory:"false" contributesTo:"query" name:"tagQueryParam"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContainersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListContainersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListContainersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainersTargetWorkloadEnum(string(request.TargetWorkload)); !ok && request.TargetWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetWorkload: %s. Supported values are: %s.", request.TargetWorkload, strings.Join(GetListContainersTargetWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainersUsageQueryParamEnum(string(request.UsageQueryParam)); !ok && request.UsageQueryParam != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageQueryParam: %s. Supported values are: %s.", request.UsageQueryParam, strings.Join(GetListContainersUsageQueryParamEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListContainersResponse wrapper for the ListContainers operation
type ListContainersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ContainerSummary instances
	Items []ContainerSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListContainersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainersLifecycleStateEnum Enum with underlying type: string
type ListContainersLifecycleStateEnum string

// Set of constants representing the allowable values for ListContainersLifecycleStateEnum
const (
	ListContainersLifecycleStateActive   ListContainersLifecycleStateEnum = "ACTIVE"
	ListContainersLifecycleStateInactive ListContainersLifecycleStateEnum = "INACTIVE"
)

var mappingListContainersLifecycleStateEnum = map[string]ListContainersLifecycleStateEnum{
	"ACTIVE":   ListContainersLifecycleStateActive,
	"INACTIVE": ListContainersLifecycleStateInactive,
}

var mappingListContainersLifecycleStateEnumLowerCase = map[string]ListContainersLifecycleStateEnum{
	"active":   ListContainersLifecycleStateActive,
	"inactive": ListContainersLifecycleStateInactive,
}

// GetListContainersLifecycleStateEnumValues Enumerates the set of values for ListContainersLifecycleStateEnum
func GetListContainersLifecycleStateEnumValues() []ListContainersLifecycleStateEnum {
	values := make([]ListContainersLifecycleStateEnum, 0)
	for _, v := range mappingListContainersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersLifecycleStateEnumStringValues Enumerates the set of values in String for ListContainersLifecycleStateEnum
func GetListContainersLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListContainersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersLifecycleStateEnum(val string) (ListContainersLifecycleStateEnum, bool) {
	enum, ok := mappingListContainersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainersTargetWorkloadEnum Enum with underlying type: string
type ListContainersTargetWorkloadEnum string

// Set of constants representing the allowable values for ListContainersTargetWorkloadEnum
const (
	ListContainersTargetWorkloadModelDeployment ListContainersTargetWorkloadEnum = "MODEL_DEPLOYMENT"
	ListContainersTargetWorkloadJobRun          ListContainersTargetWorkloadEnum = "JOB_RUN"
)

var mappingListContainersTargetWorkloadEnum = map[string]ListContainersTargetWorkloadEnum{
	"MODEL_DEPLOYMENT": ListContainersTargetWorkloadModelDeployment,
	"JOB_RUN":          ListContainersTargetWorkloadJobRun,
}

var mappingListContainersTargetWorkloadEnumLowerCase = map[string]ListContainersTargetWorkloadEnum{
	"model_deployment": ListContainersTargetWorkloadModelDeployment,
	"job_run":          ListContainersTargetWorkloadJobRun,
}

// GetListContainersTargetWorkloadEnumValues Enumerates the set of values for ListContainersTargetWorkloadEnum
func GetListContainersTargetWorkloadEnumValues() []ListContainersTargetWorkloadEnum {
	values := make([]ListContainersTargetWorkloadEnum, 0)
	for _, v := range mappingListContainersTargetWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersTargetWorkloadEnumStringValues Enumerates the set of values in String for ListContainersTargetWorkloadEnum
func GetListContainersTargetWorkloadEnumStringValues() []string {
	return []string{
		"MODEL_DEPLOYMENT",
		"JOB_RUN",
	}
}

// GetMappingListContainersTargetWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersTargetWorkloadEnum(val string) (ListContainersTargetWorkloadEnum, bool) {
	enum, ok := mappingListContainersTargetWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainersUsageQueryParamEnum Enum with underlying type: string
type ListContainersUsageQueryParamEnum string

// Set of constants representing the allowable values for ListContainersUsageQueryParamEnum
const (
	ListContainersUsageQueryParamInference      ListContainersUsageQueryParamEnum = "INFERENCE"
	ListContainersUsageQueryParamFineTune       ListContainersUsageQueryParamEnum = "FINE_TUNE"
	ListContainersUsageQueryParamEvaluation     ListContainersUsageQueryParamEnum = "EVALUATION"
	ListContainersUsageQueryParamBatchInference ListContainersUsageQueryParamEnum = "BATCH_INFERENCE"
	ListContainersUsageQueryParamOther          ListContainersUsageQueryParamEnum = "OTHER"
)

var mappingListContainersUsageQueryParamEnum = map[string]ListContainersUsageQueryParamEnum{
	"INFERENCE":       ListContainersUsageQueryParamInference,
	"FINE_TUNE":       ListContainersUsageQueryParamFineTune,
	"EVALUATION":      ListContainersUsageQueryParamEvaluation,
	"BATCH_INFERENCE": ListContainersUsageQueryParamBatchInference,
	"OTHER":           ListContainersUsageQueryParamOther,
}

var mappingListContainersUsageQueryParamEnumLowerCase = map[string]ListContainersUsageQueryParamEnum{
	"inference":       ListContainersUsageQueryParamInference,
	"fine_tune":       ListContainersUsageQueryParamFineTune,
	"evaluation":      ListContainersUsageQueryParamEvaluation,
	"batch_inference": ListContainersUsageQueryParamBatchInference,
	"other":           ListContainersUsageQueryParamOther,
}

// GetListContainersUsageQueryParamEnumValues Enumerates the set of values for ListContainersUsageQueryParamEnum
func GetListContainersUsageQueryParamEnumValues() []ListContainersUsageQueryParamEnum {
	values := make([]ListContainersUsageQueryParamEnum, 0)
	for _, v := range mappingListContainersUsageQueryParamEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainersUsageQueryParamEnumStringValues Enumerates the set of values in String for ListContainersUsageQueryParamEnum
func GetListContainersUsageQueryParamEnumStringValues() []string {
	return []string{
		"INFERENCE",
		"FINE_TUNE",
		"EVALUATION",
		"BATCH_INFERENCE",
		"OTHER",
	}
}

// GetMappingListContainersUsageQueryParamEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainersUsageQueryParamEnum(val string) (ListContainersUsageQueryParamEnum, bool) {
	enum, ok := mappingListContainersUsageQueryParamEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
