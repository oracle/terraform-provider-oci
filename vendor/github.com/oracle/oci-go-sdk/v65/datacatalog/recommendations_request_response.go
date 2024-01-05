// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// RecommendationsRequest wrapper for the Recommendations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/Recommendations.go.html to see an example of how to use RecommendationsRequest.
type RecommendationsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter used to return only recommendations of the specified type.
	RecommendationType []RecommendationTypeEnum `contributesTo:"query" name:"recommendationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter used to provide the unique identifier of the source object, for which a list of recommendations will be returned for review.
	SourceObjectKey *string `mandatory:"true" contributesTo:"query" name:"sourceObjectKey"`

	// A filter used to provide the type of the source object, for which a list of recommendations will be returned for review.
	SourceObjectType RecommendationsSourceObjectTypeEnum `mandatory:"true" contributesTo:"query" name:"sourceObjectType" omitEmpty:"true"`

	// A filter used to return only recommendations having the requested status.
	RecommendationStatus RecommendationsRecommendationStatusEnum `mandatory:"false" contributesTo:"query" name:"recommendationStatus" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RecommendationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.RecommendationType {
		if _, ok := GetMappingRecommendationTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecommendationType: %s. Supported values are: %s.", val, strings.Join(GetRecommendationTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingRecommendationsSourceObjectTypeEnum(string(request.SourceObjectType)); !ok && request.SourceObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceObjectType: %s. Supported values are: %s.", request.SourceObjectType, strings.Join(GetRecommendationsSourceObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRecommendationsRecommendationStatusEnum(string(request.RecommendationStatus)); !ok && request.RecommendationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecommendationStatus: %s. Supported values are: %s.", request.RecommendationStatus, strings.Join(GetRecommendationsRecommendationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecommendationsResponse wrapper for the Recommendations operation
type RecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RecommendationCollection instance
	RecommendationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RecommendationsSourceObjectTypeEnum Enum with underlying type: string
type RecommendationsSourceObjectTypeEnum string

// Set of constants representing the allowable values for RecommendationsSourceObjectTypeEnum
const (
	RecommendationsSourceObjectTypeDataEntity RecommendationsSourceObjectTypeEnum = "DATA_ENTITY"
	RecommendationsSourceObjectTypeAttribute  RecommendationsSourceObjectTypeEnum = "ATTRIBUTE"
	RecommendationsSourceObjectTypeTerm       RecommendationsSourceObjectTypeEnum = "TERM"
	RecommendationsSourceObjectTypeCategory   RecommendationsSourceObjectTypeEnum = "CATEGORY"
)

var mappingRecommendationsSourceObjectTypeEnum = map[string]RecommendationsSourceObjectTypeEnum{
	"DATA_ENTITY": RecommendationsSourceObjectTypeDataEntity,
	"ATTRIBUTE":   RecommendationsSourceObjectTypeAttribute,
	"TERM":        RecommendationsSourceObjectTypeTerm,
	"CATEGORY":    RecommendationsSourceObjectTypeCategory,
}

var mappingRecommendationsSourceObjectTypeEnumLowerCase = map[string]RecommendationsSourceObjectTypeEnum{
	"data_entity": RecommendationsSourceObjectTypeDataEntity,
	"attribute":   RecommendationsSourceObjectTypeAttribute,
	"term":        RecommendationsSourceObjectTypeTerm,
	"category":    RecommendationsSourceObjectTypeCategory,
}

// GetRecommendationsSourceObjectTypeEnumValues Enumerates the set of values for RecommendationsSourceObjectTypeEnum
func GetRecommendationsSourceObjectTypeEnumValues() []RecommendationsSourceObjectTypeEnum {
	values := make([]RecommendationsSourceObjectTypeEnum, 0)
	for _, v := range mappingRecommendationsSourceObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationsSourceObjectTypeEnumStringValues Enumerates the set of values in String for RecommendationsSourceObjectTypeEnum
func GetRecommendationsSourceObjectTypeEnumStringValues() []string {
	return []string{
		"DATA_ENTITY",
		"ATTRIBUTE",
		"TERM",
		"CATEGORY",
	}
}

// GetMappingRecommendationsSourceObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationsSourceObjectTypeEnum(val string) (RecommendationsSourceObjectTypeEnum, bool) {
	enum, ok := mappingRecommendationsSourceObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RecommendationsRecommendationStatusEnum Enum with underlying type: string
type RecommendationsRecommendationStatusEnum string

// Set of constants representing the allowable values for RecommendationsRecommendationStatusEnum
const (
	RecommendationsRecommendationStatusAccepted RecommendationsRecommendationStatusEnum = "ACCEPTED"
	RecommendationsRecommendationStatusRejected RecommendationsRecommendationStatusEnum = "REJECTED"
	RecommendationsRecommendationStatusInferred RecommendationsRecommendationStatusEnum = "INFERRED"
)

var mappingRecommendationsRecommendationStatusEnum = map[string]RecommendationsRecommendationStatusEnum{
	"ACCEPTED": RecommendationsRecommendationStatusAccepted,
	"REJECTED": RecommendationsRecommendationStatusRejected,
	"INFERRED": RecommendationsRecommendationStatusInferred,
}

var mappingRecommendationsRecommendationStatusEnumLowerCase = map[string]RecommendationsRecommendationStatusEnum{
	"accepted": RecommendationsRecommendationStatusAccepted,
	"rejected": RecommendationsRecommendationStatusRejected,
	"inferred": RecommendationsRecommendationStatusInferred,
}

// GetRecommendationsRecommendationStatusEnumValues Enumerates the set of values for RecommendationsRecommendationStatusEnum
func GetRecommendationsRecommendationStatusEnumValues() []RecommendationsRecommendationStatusEnum {
	values := make([]RecommendationsRecommendationStatusEnum, 0)
	for _, v := range mappingRecommendationsRecommendationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationsRecommendationStatusEnumStringValues Enumerates the set of values in String for RecommendationsRecommendationStatusEnum
func GetRecommendationsRecommendationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"REJECTED",
		"INFERRED",
	}
}

// GetMappingRecommendationsRecommendationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationsRecommendationStatusEnum(val string) (RecommendationsRecommendationStatusEnum, bool) {
	enum, ok := mappingRecommendationsRecommendationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
