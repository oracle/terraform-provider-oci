// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// AnalyzeAssetsRequest wrapper for the AnalyzeAssets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/AnalyzeAssets.go.html to see an example of how to use AnalyzeAssetsRequest.
type AnalyzeAssetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// An array of properties on which to aggregate.
	AggregationProperties []string `contributesTo:"query" name:"aggregationProperties" collectionFormat:"multi"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only assets whose lifecycleState matches the given lifecycleState.
	LifecycleState AssetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Source key from where the assets originate.
	SourceKey *string `mandatory:"false" contributesTo:"query" name:"sourceKey"`

	// External asset key.
	ExternalAssetKey *string `mandatory:"false" contributesTo:"query" name:"externalAssetKey"`

	// The type of asset.
	AssetType AnalyzeAssetsAssetTypeEnum `mandatory:"false" contributesTo:"query" name:"assetType" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder AnalyzeAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The dimensions in which to group the aggregations.
	GroupBy []string `contributesTo:"query" name:"groupBy" collectionFormat:"multi"`

	// Unique Inventory identifier.
	InventoryId *string `mandatory:"false" contributesTo:"query" name:"inventoryId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request AnalyzeAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AnalyzeAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request AnalyzeAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AnalyzeAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request AnalyzeAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAssetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAnalyzeAssetsAssetTypeEnum(string(request.AssetType)); !ok && request.AssetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetType: %s. Supported values are: %s.", request.AssetType, strings.Join(GetAnalyzeAssetsAssetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAnalyzeAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetAnalyzeAssetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AnalyzeAssetsResponse wrapper for the AnalyzeAssets operation
type AnalyzeAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssetAggregationCollection instances
	AssetAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response AnalyzeAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AnalyzeAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// AnalyzeAssetsAssetTypeEnum Enum with underlying type: string
type AnalyzeAssetsAssetTypeEnum string

// Set of constants representing the allowable values for AnalyzeAssetsAssetTypeEnum
const (
	AnalyzeAssetsAssetTypeVmwareVm AnalyzeAssetsAssetTypeEnum = "VMWARE_VM"
	AnalyzeAssetsAssetTypeVm       AnalyzeAssetsAssetTypeEnum = "VM"
)

var mappingAnalyzeAssetsAssetTypeEnum = map[string]AnalyzeAssetsAssetTypeEnum{
	"VMWARE_VM": AnalyzeAssetsAssetTypeVmwareVm,
	"VM":        AnalyzeAssetsAssetTypeVm,
}

var mappingAnalyzeAssetsAssetTypeEnumLowerCase = map[string]AnalyzeAssetsAssetTypeEnum{
	"vmware_vm": AnalyzeAssetsAssetTypeVmwareVm,
	"vm":        AnalyzeAssetsAssetTypeVm,
}

// GetAnalyzeAssetsAssetTypeEnumValues Enumerates the set of values for AnalyzeAssetsAssetTypeEnum
func GetAnalyzeAssetsAssetTypeEnumValues() []AnalyzeAssetsAssetTypeEnum {
	values := make([]AnalyzeAssetsAssetTypeEnum, 0)
	for _, v := range mappingAnalyzeAssetsAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAnalyzeAssetsAssetTypeEnumStringValues Enumerates the set of values in String for AnalyzeAssetsAssetTypeEnum
func GetAnalyzeAssetsAssetTypeEnumStringValues() []string {
	return []string{
		"VMWARE_VM",
		"VM",
	}
}

// GetMappingAnalyzeAssetsAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnalyzeAssetsAssetTypeEnum(val string) (AnalyzeAssetsAssetTypeEnum, bool) {
	enum, ok := mappingAnalyzeAssetsAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AnalyzeAssetsSortOrderEnum Enum with underlying type: string
type AnalyzeAssetsSortOrderEnum string

// Set of constants representing the allowable values for AnalyzeAssetsSortOrderEnum
const (
	AnalyzeAssetsSortOrderAsc  AnalyzeAssetsSortOrderEnum = "ASC"
	AnalyzeAssetsSortOrderDesc AnalyzeAssetsSortOrderEnum = "DESC"
)

var mappingAnalyzeAssetsSortOrderEnum = map[string]AnalyzeAssetsSortOrderEnum{
	"ASC":  AnalyzeAssetsSortOrderAsc,
	"DESC": AnalyzeAssetsSortOrderDesc,
}

var mappingAnalyzeAssetsSortOrderEnumLowerCase = map[string]AnalyzeAssetsSortOrderEnum{
	"asc":  AnalyzeAssetsSortOrderAsc,
	"desc": AnalyzeAssetsSortOrderDesc,
}

// GetAnalyzeAssetsSortOrderEnumValues Enumerates the set of values for AnalyzeAssetsSortOrderEnum
func GetAnalyzeAssetsSortOrderEnumValues() []AnalyzeAssetsSortOrderEnum {
	values := make([]AnalyzeAssetsSortOrderEnum, 0)
	for _, v := range mappingAnalyzeAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetAnalyzeAssetsSortOrderEnumStringValues Enumerates the set of values in String for AnalyzeAssetsSortOrderEnum
func GetAnalyzeAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingAnalyzeAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnalyzeAssetsSortOrderEnum(val string) (AnalyzeAssetsSortOrderEnum, bool) {
	enum, ok := mappingAnalyzeAssetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
