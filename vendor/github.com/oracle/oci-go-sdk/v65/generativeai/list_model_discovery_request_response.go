// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModelDiscoveryRequest wrapper for the ListModelDiscovery operation
type ListModelDiscoveryRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose realm matches the given realm.
	Realm []string `contributesTo:"query" name:"realm" collectionFormat:"multi"`

	// A filter to return only resources whose region matches the given region.
	Region []string `contributesTo:"query" name:"region" collectionFormat:"multi"`

	// A filter to return only resources whose model identifier matches the given modelId.
	ModelId *string `mandatory:"false" contributesTo:"query" name:"modelId"`

	// A filter to return only resources whose serving modes match the given servingModes.
	ServingMode []ListModelDiscoveryServingModeEnum `contributesTo:"query" name:"servingMode" omitEmpty:"true" collectionFormat:"multi"`

	// Filter models that support any of the specified API capabilities.
	ApiCapability []string `contributesTo:"query" name:"apiCapability" collectionFormat:"multi"`

	// A filter to return only resources their capability matches the given capability.
	Capability []ModelCapabilityEnum `contributesTo:"query" name:"capability" omitEmpty:"true" collectionFormat:"multi"`

	// Filter models by access type.
	ModelAccess []ListModelDiscoveryModelAccessEnum `contributesTo:"query" name:"modelAccess" omitEmpty:"true" collectionFormat:"multi"`

	// If true, return only deprecated models; if false, exclude deprecated models.
	IsDeprecated *bool `mandatory:"false" contributesTo:"query" name:"isDeprecated"`

	// Filter models based on on-demand retirement status.
	IsOnDemandRetired *bool `mandatory:"false" contributesTo:"query" name:"isOnDemandRetired"`

	// Filter models based on dedicated retirement status.
	IsDedicatedRetired *bool `mandatory:"false" contributesTo:"query" name:"isDedicatedRetired"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelDiscoveryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelDiscoveryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelDiscoveryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelDiscoveryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelDiscoveryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ServingMode {
		if _, ok := GetMappingListModelDiscoveryServingModeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServingMode: %s. Supported values are: %s.", val, strings.Join(GetListModelDiscoveryServingModeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Capability {
		if _, ok := GetMappingModelCapabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Capability: %s. Supported values are: %s.", val, strings.Join(GetModelCapabilityEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ModelAccess {
		if _, ok := GetMappingListModelDiscoveryModelAccessEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelAccess: %s. Supported values are: %s.", val, strings.Join(GetListModelDiscoveryModelAccessEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelDiscoveryResponse wrapper for the ListModelDiscovery operation
type ListModelDiscoveryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ModelDiscoveryCollection instances
	ModelDiscoveryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListModelDiscoveryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelDiscoveryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelDiscoveryServingModeEnum Enum with underlying type: string
type ListModelDiscoveryServingModeEnum string

// Set of constants representing the allowable values for ListModelDiscoveryServingModeEnum
const (
	ListModelDiscoveryServingModeOnDemand  ListModelDiscoveryServingModeEnum = "ON_DEMAND"
	ListModelDiscoveryServingModeDedicated ListModelDiscoveryServingModeEnum = "DEDICATED"
)

var mappingListModelDiscoveryServingModeEnum = map[string]ListModelDiscoveryServingModeEnum{
	"ON_DEMAND": ListModelDiscoveryServingModeOnDemand,
	"DEDICATED": ListModelDiscoveryServingModeDedicated,
}

var mappingListModelDiscoveryServingModeEnumLowerCase = map[string]ListModelDiscoveryServingModeEnum{
	"on_demand": ListModelDiscoveryServingModeOnDemand,
	"dedicated": ListModelDiscoveryServingModeDedicated,
}

// GetListModelDiscoveryServingModeEnumValues Enumerates the set of values for ListModelDiscoveryServingModeEnum
func GetListModelDiscoveryServingModeEnumValues() []ListModelDiscoveryServingModeEnum {
	values := make([]ListModelDiscoveryServingModeEnum, 0)
	for _, v := range mappingListModelDiscoveryServingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelDiscoveryServingModeEnumStringValues Enumerates the set of values in String for ListModelDiscoveryServingModeEnum
func GetListModelDiscoveryServingModeEnumStringValues() []string {
	return []string{
		"ON_DEMAND",
		"DEDICATED",
	}
}

// GetMappingListModelDiscoveryServingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelDiscoveryServingModeEnum(val string) (ListModelDiscoveryServingModeEnum, bool) {
	enum, ok := mappingListModelDiscoveryServingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelDiscoveryModelAccessEnum Enum with underlying type: string
type ListModelDiscoveryModelAccessEnum string

// Set of constants representing the allowable values for ListModelDiscoveryModelAccessEnum
const (
	ListModelDiscoveryModelAccessHosted ListModelDiscoveryModelAccessEnum = "HOSTED"
	ListModelDiscoveryModelAccessProxy  ListModelDiscoveryModelAccessEnum = "PROXY"
)

var mappingListModelDiscoveryModelAccessEnum = map[string]ListModelDiscoveryModelAccessEnum{
	"HOSTED": ListModelDiscoveryModelAccessHosted,
	"PROXY":  ListModelDiscoveryModelAccessProxy,
}

var mappingListModelDiscoveryModelAccessEnumLowerCase = map[string]ListModelDiscoveryModelAccessEnum{
	"hosted": ListModelDiscoveryModelAccessHosted,
	"proxy":  ListModelDiscoveryModelAccessProxy,
}

// GetListModelDiscoveryModelAccessEnumValues Enumerates the set of values for ListModelDiscoveryModelAccessEnum
func GetListModelDiscoveryModelAccessEnumValues() []ListModelDiscoveryModelAccessEnum {
	values := make([]ListModelDiscoveryModelAccessEnum, 0)
	for _, v := range mappingListModelDiscoveryModelAccessEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelDiscoveryModelAccessEnumStringValues Enumerates the set of values in String for ListModelDiscoveryModelAccessEnum
func GetListModelDiscoveryModelAccessEnumStringValues() []string {
	return []string{
		"HOSTED",
		"PROXY",
	}
}

// GetMappingListModelDiscoveryModelAccessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelDiscoveryModelAccessEnum(val string) (ListModelDiscoveryModelAccessEnum, bool) {
	enum, ok := mappingListModelDiscoveryModelAccessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
