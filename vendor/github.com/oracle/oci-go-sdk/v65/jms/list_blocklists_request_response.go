// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBlocklistsRequest wrapper for the ListBlocklists operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListBlocklists.go.html to see an example of how to use ListBlocklistsRequest.
type ListBlocklistsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The operation type.
	Operation ListBlocklistsOperationEnum `mandatory:"false" contributesTo:"query" name:"operation" omitEmpty:"true"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListBlocklistsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used to sort blocklist records. Only one sort order may be provided.
	// Default order for _operation_ is **ascending**.
	// If no value is specified, _operation_ is default.
	SortBy ListBlocklistsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBlocklistsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBlocklistsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBlocklistsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBlocklistsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBlocklistsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBlocklistsOperationEnum(string(request.Operation)); !ok && request.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", request.Operation, strings.Join(GetListBlocklistsOperationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBlocklistsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBlocklistsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBlocklistsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBlocklistsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBlocklistsResponse wrapper for the ListBlocklists operation
type ListBlocklistsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BlocklistCollection instances
	BlocklistCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBlocklistsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBlocklistsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBlocklistsOperationEnum Enum with underlying type: string
type ListBlocklistsOperationEnum string

// Set of constants representing the allowable values for ListBlocklistsOperationEnum
const (
	ListBlocklistsOperationCreateFleet                      ListBlocklistsOperationEnum = "CREATE_FLEET"
	ListBlocklistsOperationDeleteFleet                      ListBlocklistsOperationEnum = "DELETE_FLEET"
	ListBlocklistsOperationMoveFleet                        ListBlocklistsOperationEnum = "MOVE_FLEET"
	ListBlocklistsOperationUpdateFleet                      ListBlocklistsOperationEnum = "UPDATE_FLEET"
	ListBlocklistsOperationUpdateFleetAgentConfiguration    ListBlocklistsOperationEnum = "UPDATE_FLEET_AGENT_CONFIGURATION"
	ListBlocklistsOperationDeleteJavaInstallation           ListBlocklistsOperationEnum = "DELETE_JAVA_INSTALLATION"
	ListBlocklistsOperationCreateJavaInstallation           ListBlocklistsOperationEnum = "CREATE_JAVA_INSTALLATION"
	ListBlocklistsOperationCollectJfr                       ListBlocklistsOperationEnum = "COLLECT_JFR"
	ListBlocklistsOperationRequestCryptoEventAnalysis       ListBlocklistsOperationEnum = "REQUEST_CRYPTO_EVENT_ANALYSIS"
	ListBlocklistsOperationRequestPerformanceTuningAnalysis ListBlocklistsOperationEnum = "REQUEST_PERFORMANCE_TUNING_ANALYSIS"
	ListBlocklistsOperationRequestJavaMigrationAnalysis     ListBlocklistsOperationEnum = "REQUEST_JAVA_MIGRATION_ANALYSIS"
	ListBlocklistsOperationDeleteJmsReport                  ListBlocklistsOperationEnum = "DELETE_JMS_REPORT"
	ListBlocklistsOperationScanJavaServerUsage              ListBlocklistsOperationEnum = "SCAN_JAVA_SERVER_USAGE"
	ListBlocklistsOperationScanLibraryUsage                 ListBlocklistsOperationEnum = "SCAN_LIBRARY_USAGE"
	ListBlocklistsOperationExportDataCsv                    ListBlocklistsOperationEnum = "EXPORT_DATA_CSV"
	ListBlocklistsOperationCreateDrsFile                    ListBlocklistsOperationEnum = "CREATE_DRS_FILE"
	ListBlocklistsOperationUpdateDrsFile                    ListBlocklistsOperationEnum = "UPDATE_DRS_FILE"
	ListBlocklistsOperationDeleteDrsFile                    ListBlocklistsOperationEnum = "DELETE_DRS_FILE"
	ListBlocklistsOperationEnableDrs                        ListBlocklistsOperationEnum = "ENABLE_DRS"
	ListBlocklistsOperationDisableDrs                       ListBlocklistsOperationEnum = "DISABLE_DRS"
)

var mappingListBlocklistsOperationEnum = map[string]ListBlocklistsOperationEnum{
	"CREATE_FLEET":                        ListBlocklistsOperationCreateFleet,
	"DELETE_FLEET":                        ListBlocklistsOperationDeleteFleet,
	"MOVE_FLEET":                          ListBlocklistsOperationMoveFleet,
	"UPDATE_FLEET":                        ListBlocklistsOperationUpdateFleet,
	"UPDATE_FLEET_AGENT_CONFIGURATION":    ListBlocklistsOperationUpdateFleetAgentConfiguration,
	"DELETE_JAVA_INSTALLATION":            ListBlocklistsOperationDeleteJavaInstallation,
	"CREATE_JAVA_INSTALLATION":            ListBlocklistsOperationCreateJavaInstallation,
	"COLLECT_JFR":                         ListBlocklistsOperationCollectJfr,
	"REQUEST_CRYPTO_EVENT_ANALYSIS":       ListBlocklistsOperationRequestCryptoEventAnalysis,
	"REQUEST_PERFORMANCE_TUNING_ANALYSIS": ListBlocklistsOperationRequestPerformanceTuningAnalysis,
	"REQUEST_JAVA_MIGRATION_ANALYSIS":     ListBlocklistsOperationRequestJavaMigrationAnalysis,
	"DELETE_JMS_REPORT":                   ListBlocklistsOperationDeleteJmsReport,
	"SCAN_JAVA_SERVER_USAGE":              ListBlocklistsOperationScanJavaServerUsage,
	"SCAN_LIBRARY_USAGE":                  ListBlocklistsOperationScanLibraryUsage,
	"EXPORT_DATA_CSV":                     ListBlocklistsOperationExportDataCsv,
	"CREATE_DRS_FILE":                     ListBlocklistsOperationCreateDrsFile,
	"UPDATE_DRS_FILE":                     ListBlocklistsOperationUpdateDrsFile,
	"DELETE_DRS_FILE":                     ListBlocklistsOperationDeleteDrsFile,
	"ENABLE_DRS":                          ListBlocklistsOperationEnableDrs,
	"DISABLE_DRS":                         ListBlocklistsOperationDisableDrs,
}

var mappingListBlocklistsOperationEnumLowerCase = map[string]ListBlocklistsOperationEnum{
	"create_fleet":                        ListBlocklistsOperationCreateFleet,
	"delete_fleet":                        ListBlocklistsOperationDeleteFleet,
	"move_fleet":                          ListBlocklistsOperationMoveFleet,
	"update_fleet":                        ListBlocklistsOperationUpdateFleet,
	"update_fleet_agent_configuration":    ListBlocklistsOperationUpdateFleetAgentConfiguration,
	"delete_java_installation":            ListBlocklistsOperationDeleteJavaInstallation,
	"create_java_installation":            ListBlocklistsOperationCreateJavaInstallation,
	"collect_jfr":                         ListBlocklistsOperationCollectJfr,
	"request_crypto_event_analysis":       ListBlocklistsOperationRequestCryptoEventAnalysis,
	"request_performance_tuning_analysis": ListBlocklistsOperationRequestPerformanceTuningAnalysis,
	"request_java_migration_analysis":     ListBlocklistsOperationRequestJavaMigrationAnalysis,
	"delete_jms_report":                   ListBlocklistsOperationDeleteJmsReport,
	"scan_java_server_usage":              ListBlocklistsOperationScanJavaServerUsage,
	"scan_library_usage":                  ListBlocklistsOperationScanLibraryUsage,
	"export_data_csv":                     ListBlocklistsOperationExportDataCsv,
	"create_drs_file":                     ListBlocklistsOperationCreateDrsFile,
	"update_drs_file":                     ListBlocklistsOperationUpdateDrsFile,
	"delete_drs_file":                     ListBlocklistsOperationDeleteDrsFile,
	"enable_drs":                          ListBlocklistsOperationEnableDrs,
	"disable_drs":                         ListBlocklistsOperationDisableDrs,
}

// GetListBlocklistsOperationEnumValues Enumerates the set of values for ListBlocklistsOperationEnum
func GetListBlocklistsOperationEnumValues() []ListBlocklistsOperationEnum {
	values := make([]ListBlocklistsOperationEnum, 0)
	for _, v := range mappingListBlocklistsOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetListBlocklistsOperationEnumStringValues Enumerates the set of values in String for ListBlocklistsOperationEnum
func GetListBlocklistsOperationEnumStringValues() []string {
	return []string{
		"CREATE_FLEET",
		"DELETE_FLEET",
		"MOVE_FLEET",
		"UPDATE_FLEET",
		"UPDATE_FLEET_AGENT_CONFIGURATION",
		"DELETE_JAVA_INSTALLATION",
		"CREATE_JAVA_INSTALLATION",
		"COLLECT_JFR",
		"REQUEST_CRYPTO_EVENT_ANALYSIS",
		"REQUEST_PERFORMANCE_TUNING_ANALYSIS",
		"REQUEST_JAVA_MIGRATION_ANALYSIS",
		"DELETE_JMS_REPORT",
		"SCAN_JAVA_SERVER_USAGE",
		"SCAN_LIBRARY_USAGE",
		"EXPORT_DATA_CSV",
		"CREATE_DRS_FILE",
		"UPDATE_DRS_FILE",
		"DELETE_DRS_FILE",
		"ENABLE_DRS",
		"DISABLE_DRS",
	}
}

// GetMappingListBlocklistsOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBlocklistsOperationEnum(val string) (ListBlocklistsOperationEnum, bool) {
	enum, ok := mappingListBlocklistsOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBlocklistsSortOrderEnum Enum with underlying type: string
type ListBlocklistsSortOrderEnum string

// Set of constants representing the allowable values for ListBlocklistsSortOrderEnum
const (
	ListBlocklistsSortOrderAsc  ListBlocklistsSortOrderEnum = "ASC"
	ListBlocklistsSortOrderDesc ListBlocklistsSortOrderEnum = "DESC"
)

var mappingListBlocklistsSortOrderEnum = map[string]ListBlocklistsSortOrderEnum{
	"ASC":  ListBlocklistsSortOrderAsc,
	"DESC": ListBlocklistsSortOrderDesc,
}

var mappingListBlocklistsSortOrderEnumLowerCase = map[string]ListBlocklistsSortOrderEnum{
	"asc":  ListBlocklistsSortOrderAsc,
	"desc": ListBlocklistsSortOrderDesc,
}

// GetListBlocklistsSortOrderEnumValues Enumerates the set of values for ListBlocklistsSortOrderEnum
func GetListBlocklistsSortOrderEnumValues() []ListBlocklistsSortOrderEnum {
	values := make([]ListBlocklistsSortOrderEnum, 0)
	for _, v := range mappingListBlocklistsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBlocklistsSortOrderEnumStringValues Enumerates the set of values in String for ListBlocklistsSortOrderEnum
func GetListBlocklistsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBlocklistsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBlocklistsSortOrderEnum(val string) (ListBlocklistsSortOrderEnum, bool) {
	enum, ok := mappingListBlocklistsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBlocklistsSortByEnum Enum with underlying type: string
type ListBlocklistsSortByEnum string

// Set of constants representing the allowable values for ListBlocklistsSortByEnum
const (
	ListBlocklistsSortByOperation ListBlocklistsSortByEnum = "operation"
)

var mappingListBlocklistsSortByEnum = map[string]ListBlocklistsSortByEnum{
	"operation": ListBlocklistsSortByOperation,
}

var mappingListBlocklistsSortByEnumLowerCase = map[string]ListBlocklistsSortByEnum{
	"operation": ListBlocklistsSortByOperation,
}

// GetListBlocklistsSortByEnumValues Enumerates the set of values for ListBlocklistsSortByEnum
func GetListBlocklistsSortByEnumValues() []ListBlocklistsSortByEnum {
	values := make([]ListBlocklistsSortByEnum, 0)
	for _, v := range mappingListBlocklistsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBlocklistsSortByEnumStringValues Enumerates the set of values in String for ListBlocklistsSortByEnum
func GetListBlocklistsSortByEnumStringValues() []string {
	return []string{
		"operation",
	}
}

// GetMappingListBlocklistsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBlocklistsSortByEnum(val string) (ListBlocklistsSortByEnum, bool) {
	enum, ok := mappingListBlocklistsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
