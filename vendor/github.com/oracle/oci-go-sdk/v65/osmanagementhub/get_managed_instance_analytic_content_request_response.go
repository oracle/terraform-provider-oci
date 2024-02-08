// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetManagedInstanceAnalyticContentRequest wrapper for the GetManagedInstanceAnalyticContent operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagedInstanceAnalyticContent.go.html to see an example of how to use GetManagedInstanceAnalyticContentRequest.
type GetManagedInstanceAnalyticContentRequest struct {

	// This compartmentId is used to list managed instances within a compartment.
	// Or serve as an additional filter to restrict only managed instances with in certain compartment if other filter presents.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID of the managed instance group for which to list resources.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The OCID of the lifecycle environment.
	LifecycleEnvironmentId *string `mandatory:"false" contributesTo:"query" name:"lifecycleEnvironmentId"`

	// The OCID of the lifecycle stage for which to list resources.
	LifecycleStageId *string `mandatory:"false" contributesTo:"query" name:"lifecycleStageId"`

	// A filter to return only instances whose managed instance status matches the given status.
	Status []ManagedInstanceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Filter instances by Location. Used when report target type is compartment or group.
	InstanceLocation GetManagedInstanceAnalyticContentInstanceLocationEnum `mandatory:"false" contributesTo:"query" name:"instanceLocation" omitEmpty:"true"`

	// A filter to return instances with number of available security updates equals to the number specified.
	SecurityUpdatesAvailableEqualsTo *int `mandatory:"false" contributesTo:"query" name:"securityUpdatesAvailableEqualsTo"`

	// A filter to return instances with number of available bug updates equals to the number specified.
	BugUpdatesAvailableEqualsTo *int `mandatory:"false" contributesTo:"query" name:"bugUpdatesAvailableEqualsTo"`

	// A filter to return instances with number of available security updates greater than the number specified.
	SecurityUpdatesAvailableGreaterThan *int `mandatory:"false" contributesTo:"query" name:"securityUpdatesAvailableGreaterThan"`

	// A filter to return instances with number of available bug updates greater than the number specified.
	BugUpdatesAvailableGreaterThan *int `mandatory:"false" contributesTo:"query" name:"bugUpdatesAvailableGreaterThan"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetManagedInstanceAnalyticContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetManagedInstanceAnalyticContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetManagedInstanceAnalyticContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetManagedInstanceAnalyticContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetManagedInstanceAnalyticContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Status {
		if _, ok := GetMappingManagedInstanceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingGetManagedInstanceAnalyticContentInstanceLocationEnum(string(request.InstanceLocation)); !ok && request.InstanceLocation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceLocation: %s. Supported values are: %s.", request.InstanceLocation, strings.Join(GetGetManagedInstanceAnalyticContentInstanceLocationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetManagedInstanceAnalyticContentResponse wrapper for the GetManagedInstanceAnalyticContent operation
type GetManagedInstanceAnalyticContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetManagedInstanceAnalyticContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetManagedInstanceAnalyticContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetManagedInstanceAnalyticContentInstanceLocationEnum Enum with underlying type: string
type GetManagedInstanceAnalyticContentInstanceLocationEnum string

// Set of constants representing the allowable values for GetManagedInstanceAnalyticContentInstanceLocationEnum
const (
	GetManagedInstanceAnalyticContentInstanceLocationOnPremise  GetManagedInstanceAnalyticContentInstanceLocationEnum = "ON_PREMISE"
	GetManagedInstanceAnalyticContentInstanceLocationOciCompute GetManagedInstanceAnalyticContentInstanceLocationEnum = "OCI_COMPUTE"
	GetManagedInstanceAnalyticContentInstanceLocationAzure      GetManagedInstanceAnalyticContentInstanceLocationEnum = "AZURE"
	GetManagedInstanceAnalyticContentInstanceLocationEc2        GetManagedInstanceAnalyticContentInstanceLocationEnum = "EC2"
)

var mappingGetManagedInstanceAnalyticContentInstanceLocationEnum = map[string]GetManagedInstanceAnalyticContentInstanceLocationEnum{
	"ON_PREMISE":  GetManagedInstanceAnalyticContentInstanceLocationOnPremise,
	"OCI_COMPUTE": GetManagedInstanceAnalyticContentInstanceLocationOciCompute,
	"AZURE":       GetManagedInstanceAnalyticContentInstanceLocationAzure,
	"EC2":         GetManagedInstanceAnalyticContentInstanceLocationEc2,
}

var mappingGetManagedInstanceAnalyticContentInstanceLocationEnumLowerCase = map[string]GetManagedInstanceAnalyticContentInstanceLocationEnum{
	"on_premise":  GetManagedInstanceAnalyticContentInstanceLocationOnPremise,
	"oci_compute": GetManagedInstanceAnalyticContentInstanceLocationOciCompute,
	"azure":       GetManagedInstanceAnalyticContentInstanceLocationAzure,
	"ec2":         GetManagedInstanceAnalyticContentInstanceLocationEc2,
}

// GetGetManagedInstanceAnalyticContentInstanceLocationEnumValues Enumerates the set of values for GetManagedInstanceAnalyticContentInstanceLocationEnum
func GetGetManagedInstanceAnalyticContentInstanceLocationEnumValues() []GetManagedInstanceAnalyticContentInstanceLocationEnum {
	values := make([]GetManagedInstanceAnalyticContentInstanceLocationEnum, 0)
	for _, v := range mappingGetManagedInstanceAnalyticContentInstanceLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetGetManagedInstanceAnalyticContentInstanceLocationEnumStringValues Enumerates the set of values in String for GetManagedInstanceAnalyticContentInstanceLocationEnum
func GetGetManagedInstanceAnalyticContentInstanceLocationEnumStringValues() []string {
	return []string{
		"ON_PREMISE",
		"OCI_COMPUTE",
		"AZURE",
		"EC2",
	}
}

// GetMappingGetManagedInstanceAnalyticContentInstanceLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetManagedInstanceAnalyticContentInstanceLocationEnum(val string) (GetManagedInstanceAnalyticContentInstanceLocationEnum, bool) {
	enum, ok := mappingGetManagedInstanceAnalyticContentInstanceLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
