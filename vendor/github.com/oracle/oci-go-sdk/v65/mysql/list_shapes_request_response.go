// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListShapesRequest wrapper for the ListShapes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListShapes.go.html to see an example of how to use ListShapesRequest.
type ListShapesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Return shapes that are supported by the service feature.
	IsSupportedFor []ListShapesIsSupportedForEnum `contributesTo:"query" name:"isSupportedFor" omitEmpty:"true" collectionFormat:"multi"`

	// The name of the Availability Domain.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// Name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListShapesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListShapesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListShapesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListShapesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListShapesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.IsSupportedFor {
		if _, ok := GetMappingListShapesIsSupportedForEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsSupportedFor: %s. Supported values are: %s.", val, strings.Join(GetListShapesIsSupportedForEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListShapesResponse wrapper for the ListShapes operation
type ListShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []ShapeSummary instance
	Items []ShapeSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListShapesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListShapesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListShapesIsSupportedForEnum Enum with underlying type: string
type ListShapesIsSupportedForEnum string

// Set of constants representing the allowable values for ListShapesIsSupportedForEnum
const (
	ListShapesIsSupportedForDbsystem         ListShapesIsSupportedForEnum = "DBSYSTEM"
	ListShapesIsSupportedForAnalyticscluster ListShapesIsSupportedForEnum = "ANALYTICSCLUSTER"
	ListShapesIsSupportedForHeatwavecluster  ListShapesIsSupportedForEnum = "HEATWAVECLUSTER"
)

var mappingListShapesIsSupportedForEnum = map[string]ListShapesIsSupportedForEnum{
	"DBSYSTEM":         ListShapesIsSupportedForDbsystem,
	"ANALYTICSCLUSTER": ListShapesIsSupportedForAnalyticscluster,
	"HEATWAVECLUSTER":  ListShapesIsSupportedForHeatwavecluster,
}

var mappingListShapesIsSupportedForEnumLowerCase = map[string]ListShapesIsSupportedForEnum{
	"dbsystem":         ListShapesIsSupportedForDbsystem,
	"analyticscluster": ListShapesIsSupportedForAnalyticscluster,
	"heatwavecluster":  ListShapesIsSupportedForHeatwavecluster,
}

// GetListShapesIsSupportedForEnumValues Enumerates the set of values for ListShapesIsSupportedForEnum
func GetListShapesIsSupportedForEnumValues() []ListShapesIsSupportedForEnum {
	values := make([]ListShapesIsSupportedForEnum, 0)
	for _, v := range mappingListShapesIsSupportedForEnum {
		values = append(values, v)
	}
	return values
}

// GetListShapesIsSupportedForEnumStringValues Enumerates the set of values in String for ListShapesIsSupportedForEnum
func GetListShapesIsSupportedForEnumStringValues() []string {
	return []string{
		"DBSYSTEM",
		"ANALYTICSCLUSTER",
		"HEATWAVECLUSTER",
	}
}

// GetMappingListShapesIsSupportedForEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShapesIsSupportedForEnum(val string) (ListShapesIsSupportedForEnum, bool) {
	enum, ok := mappingListShapesIsSupportedForEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
