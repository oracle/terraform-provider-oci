// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbSystemStoragePerformancesRequest wrapper for the ListDbSystemStoragePerformances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemStoragePerformances.go.html to see an example of how to use ListDbSystemStoragePerformancesRequest.
type ListDbSystemStoragePerformancesRequest struct {

	// The DB system storage management option. Used to list database versions available for that storage manager. Valid values are `ASM` and `LVM`.
	// * ASM specifies Oracle Automatic Storage Management
	// * LVM specifies logical volume manager, sometimes called logical disk manager.
	StorageManagement DbSystemOptionsStorageManagementEnum `mandatory:"true" contributesTo:"query" name:"storageManagement" omitEmpty:"true"`

	// Optional. Filters the performance results by shape type.
	ShapeType *string `mandatory:"false" contributesTo:"query" name:"shapeType"`

	// The database edition of quota (STANDARD_EDITION/ENTERPRISE_EDITION/ENTERPRISE_EDITION_HIGH_PERFORMANCE/ENTERPRISE_EDITION_EXTREME_PERFORMANCE/ENTERPRISE_EDITION_DEVELOPER)
	DatabaseEdition ListDbSystemStoragePerformancesDatabaseEditionEnum `mandatory:"false" contributesTo:"query" name:"databaseEdition" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbSystemStoragePerformancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbSystemStoragePerformancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbSystemStoragePerformancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbSystemStoragePerformancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbSystemStoragePerformancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemOptionsStorageManagementEnum(string(request.StorageManagement)); !ok && request.StorageManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageManagement: %s. Supported values are: %s.", request.StorageManagement, strings.Join(GetDbSystemOptionsStorageManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbSystemStoragePerformancesDatabaseEditionEnum(string(request.DatabaseEdition)); !ok && request.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", request.DatabaseEdition, strings.Join(GetListDbSystemStoragePerformancesDatabaseEditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbSystemStoragePerformancesResponse wrapper for the ListDbSystemStoragePerformances operation
type ListDbSystemStoragePerformancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []DbSystemStoragePerformanceSummary instance
	Items []DbSystemStoragePerformanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemStoragePerformancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemStoragePerformancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbSystemStoragePerformancesDatabaseEditionEnum Enum with underlying type: string
type ListDbSystemStoragePerformancesDatabaseEditionEnum string

// Set of constants representing the allowable values for ListDbSystemStoragePerformancesDatabaseEditionEnum
const (
	ListDbSystemStoragePerformancesDatabaseEditionStandardEdition                     ListDbSystemStoragePerformancesDatabaseEditionEnum = "STANDARD_EDITION"
	ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEdition                   ListDbSystemStoragePerformancesDatabaseEditionEnum = "ENTERPRISE_EDITION"
	ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionHighPerformance    ListDbSystemStoragePerformancesDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionExtremePerformance ListDbSystemStoragePerformancesDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionDeveloper          ListDbSystemStoragePerformancesDatabaseEditionEnum = "ENTERPRISE_EDITION_DEVELOPER"
)

var mappingListDbSystemStoragePerformancesDatabaseEditionEnum = map[string]ListDbSystemStoragePerformancesDatabaseEditionEnum{
	"STANDARD_EDITION":                       ListDbSystemStoragePerformancesDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_DEVELOPER":           ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionDeveloper,
}

var mappingListDbSystemStoragePerformancesDatabaseEditionEnumLowerCase = map[string]ListDbSystemStoragePerformancesDatabaseEditionEnum{
	"standard_edition":                       ListDbSystemStoragePerformancesDatabaseEditionStandardEdition,
	"enterprise_edition":                     ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionExtremePerformance,
	"enterprise_edition_developer":           ListDbSystemStoragePerformancesDatabaseEditionEnterpriseEditionDeveloper,
}

// GetListDbSystemStoragePerformancesDatabaseEditionEnumValues Enumerates the set of values for ListDbSystemStoragePerformancesDatabaseEditionEnum
func GetListDbSystemStoragePerformancesDatabaseEditionEnumValues() []ListDbSystemStoragePerformancesDatabaseEditionEnum {
	values := make([]ListDbSystemStoragePerformancesDatabaseEditionEnum, 0)
	for _, v := range mappingListDbSystemStoragePerformancesDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemStoragePerformancesDatabaseEditionEnumStringValues Enumerates the set of values in String for ListDbSystemStoragePerformancesDatabaseEditionEnum
func GetListDbSystemStoragePerformancesDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
		"ENTERPRISE_EDITION_DEVELOPER",
	}
}

// GetMappingListDbSystemStoragePerformancesDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemStoragePerformancesDatabaseEditionEnum(val string) (ListDbSystemStoragePerformancesDatabaseEditionEnum, bool) {
	enum, ok := mappingListDbSystemStoragePerformancesDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
