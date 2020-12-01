// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v30/common"
	"net/http"
)

// GetConditionMetadataTypeRequest wrapper for the GetConditionMetadataType operation
type GetConditionMetadataTypeRequest struct {

	// The type of the condition meta data.
	ConditionMetadataTypeId GetConditionMetadataTypeConditionMetadataTypeIdEnum `mandatory:"true" contributesTo:"path" name:"conditionMetadataTypeId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// ServiceType filter for the condition meta data.
	ServiceType *string `mandatory:"false" contributesTo:"query" name:"serviceType"`

	// Resource filter for the condition meta data.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConditionMetadataTypeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetConditionMetadataTypeRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetConditionMetadataTypeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetConditionMetadataTypeResponse wrapper for the GetConditionMetadataType operation
type GetConditionMetadataTypeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConditionMetadataType instance
	ConditionMetadataType `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetConditionMetadataTypeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetConditionMetadataTypeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetConditionMetadataTypeConditionMetadataTypeIdEnum Enum with underlying type: string
type GetConditionMetadataTypeConditionMetadataTypeIdEnum string

// Set of constants representing the allowable values for GetConditionMetadataTypeConditionMetadataTypeIdEnum
const (
	GetConditionMetadataTypeConditionMetadataTypeIdActivitycondition   GetConditionMetadataTypeConditionMetadataTypeIdEnum = "ActivityCondition"
	GetConditionMetadataTypeConditionMetadataTypeIdSecuritycondition   GetConditionMetadataTypeConditionMetadataTypeIdEnum = "SecurityCondition"
	GetConditionMetadataTypeConditionMetadataTypeIdCloudguardcondition GetConditionMetadataTypeConditionMetadataTypeIdEnum = "CloudGuardCondition"
)

var mappingGetConditionMetadataTypeConditionMetadataTypeId = map[string]GetConditionMetadataTypeConditionMetadataTypeIdEnum{
	"ActivityCondition":   GetConditionMetadataTypeConditionMetadataTypeIdActivitycondition,
	"SecurityCondition":   GetConditionMetadataTypeConditionMetadataTypeIdSecuritycondition,
	"CloudGuardCondition": GetConditionMetadataTypeConditionMetadataTypeIdCloudguardcondition,
}

// GetGetConditionMetadataTypeConditionMetadataTypeIdEnumValues Enumerates the set of values for GetConditionMetadataTypeConditionMetadataTypeIdEnum
func GetGetConditionMetadataTypeConditionMetadataTypeIdEnumValues() []GetConditionMetadataTypeConditionMetadataTypeIdEnum {
	values := make([]GetConditionMetadataTypeConditionMetadataTypeIdEnum, 0)
	for _, v := range mappingGetConditionMetadataTypeConditionMetadataTypeId {
		values = append(values, v)
	}
	return values
}
