// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of an asynchronous work request.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet.
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The operation type of the work request.
	OperationType ListWorkRequestsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// The status of the work request.
	Status []OperationStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListWorkRequestsOperationTypeEnumStringValues(), ",")))
	}
	for _, val := range request.Status {
		if _, ok := GetMappingOperationStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetOperationStatusEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestCollection instances
	WorkRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsOperationTypeEnum Enum with underlying type: string
type ListWorkRequestsOperationTypeEnum string

// Set of constants representing the allowable values for ListWorkRequestsOperationTypeEnum
const (
	ListWorkRequestsOperationTypeCreateFleet                      ListWorkRequestsOperationTypeEnum = "CREATE_FLEET"
	ListWorkRequestsOperationTypeDeleteFleet                      ListWorkRequestsOperationTypeEnum = "DELETE_FLEET"
	ListWorkRequestsOperationTypeMoveFleet                        ListWorkRequestsOperationTypeEnum = "MOVE_FLEET"
	ListWorkRequestsOperationTypeUpdateFleet                      ListWorkRequestsOperationTypeEnum = "UPDATE_FLEET"
	ListWorkRequestsOperationTypeUpdateFleetAgentConfiguration    ListWorkRequestsOperationTypeEnum = "UPDATE_FLEET_AGENT_CONFIGURATION"
	ListWorkRequestsOperationTypeDeleteJavaInstallation           ListWorkRequestsOperationTypeEnum = "DELETE_JAVA_INSTALLATION"
	ListWorkRequestsOperationTypeCreateJavaInstallation           ListWorkRequestsOperationTypeEnum = "CREATE_JAVA_INSTALLATION"
	ListWorkRequestsOperationTypeCollectJfr                       ListWorkRequestsOperationTypeEnum = "COLLECT_JFR"
	ListWorkRequestsOperationTypeRequestCryptoEventAnalysis       ListWorkRequestsOperationTypeEnum = "REQUEST_CRYPTO_EVENT_ANALYSIS"
	ListWorkRequestsOperationTypeRequestPerformanceTuningAnalysis ListWorkRequestsOperationTypeEnum = "REQUEST_PERFORMANCE_TUNING_ANALYSIS"
	ListWorkRequestsOperationTypeRequestJavaMigrationAnalysis     ListWorkRequestsOperationTypeEnum = "REQUEST_JAVA_MIGRATION_ANALYSIS"
	ListWorkRequestsOperationTypeDeleteJmsReport                  ListWorkRequestsOperationTypeEnum = "DELETE_JMS_REPORT"
	ListWorkRequestsOperationTypeScanJavaServerUsage              ListWorkRequestsOperationTypeEnum = "SCAN_JAVA_SERVER_USAGE"
	ListWorkRequestsOperationTypeScanLibraryUsage                 ListWorkRequestsOperationTypeEnum = "SCAN_LIBRARY_USAGE"
	ListWorkRequestsOperationTypeExportDataCsv                    ListWorkRequestsOperationTypeEnum = "EXPORT_DATA_CSV"
	ListWorkRequestsOperationTypeCreateDrsFile                    ListWorkRequestsOperationTypeEnum = "CREATE_DRS_FILE"
	ListWorkRequestsOperationTypeUpdateDrsFile                    ListWorkRequestsOperationTypeEnum = "UPDATE_DRS_FILE"
	ListWorkRequestsOperationTypeDeleteDrsFile                    ListWorkRequestsOperationTypeEnum = "DELETE_DRS_FILE"
	ListWorkRequestsOperationTypeEnableDrs                        ListWorkRequestsOperationTypeEnum = "ENABLE_DRS"
	ListWorkRequestsOperationTypeDisableDrs                       ListWorkRequestsOperationTypeEnum = "DISABLE_DRS"
)

var mappingListWorkRequestsOperationTypeEnum = map[string]ListWorkRequestsOperationTypeEnum{
	"CREATE_FLEET":                        ListWorkRequestsOperationTypeCreateFleet,
	"DELETE_FLEET":                        ListWorkRequestsOperationTypeDeleteFleet,
	"MOVE_FLEET":                          ListWorkRequestsOperationTypeMoveFleet,
	"UPDATE_FLEET":                        ListWorkRequestsOperationTypeUpdateFleet,
	"UPDATE_FLEET_AGENT_CONFIGURATION":    ListWorkRequestsOperationTypeUpdateFleetAgentConfiguration,
	"DELETE_JAVA_INSTALLATION":            ListWorkRequestsOperationTypeDeleteJavaInstallation,
	"CREATE_JAVA_INSTALLATION":            ListWorkRequestsOperationTypeCreateJavaInstallation,
	"COLLECT_JFR":                         ListWorkRequestsOperationTypeCollectJfr,
	"REQUEST_CRYPTO_EVENT_ANALYSIS":       ListWorkRequestsOperationTypeRequestCryptoEventAnalysis,
	"REQUEST_PERFORMANCE_TUNING_ANALYSIS": ListWorkRequestsOperationTypeRequestPerformanceTuningAnalysis,
	"REQUEST_JAVA_MIGRATION_ANALYSIS":     ListWorkRequestsOperationTypeRequestJavaMigrationAnalysis,
	"DELETE_JMS_REPORT":                   ListWorkRequestsOperationTypeDeleteJmsReport,
	"SCAN_JAVA_SERVER_USAGE":              ListWorkRequestsOperationTypeScanJavaServerUsage,
	"SCAN_LIBRARY_USAGE":                  ListWorkRequestsOperationTypeScanLibraryUsage,
	"EXPORT_DATA_CSV":                     ListWorkRequestsOperationTypeExportDataCsv,
	"CREATE_DRS_FILE":                     ListWorkRequestsOperationTypeCreateDrsFile,
	"UPDATE_DRS_FILE":                     ListWorkRequestsOperationTypeUpdateDrsFile,
	"DELETE_DRS_FILE":                     ListWorkRequestsOperationTypeDeleteDrsFile,
	"ENABLE_DRS":                          ListWorkRequestsOperationTypeEnableDrs,
	"DISABLE_DRS":                         ListWorkRequestsOperationTypeDisableDrs,
}

var mappingListWorkRequestsOperationTypeEnumLowerCase = map[string]ListWorkRequestsOperationTypeEnum{
	"create_fleet":                        ListWorkRequestsOperationTypeCreateFleet,
	"delete_fleet":                        ListWorkRequestsOperationTypeDeleteFleet,
	"move_fleet":                          ListWorkRequestsOperationTypeMoveFleet,
	"update_fleet":                        ListWorkRequestsOperationTypeUpdateFleet,
	"update_fleet_agent_configuration":    ListWorkRequestsOperationTypeUpdateFleetAgentConfiguration,
	"delete_java_installation":            ListWorkRequestsOperationTypeDeleteJavaInstallation,
	"create_java_installation":            ListWorkRequestsOperationTypeCreateJavaInstallation,
	"collect_jfr":                         ListWorkRequestsOperationTypeCollectJfr,
	"request_crypto_event_analysis":       ListWorkRequestsOperationTypeRequestCryptoEventAnalysis,
	"request_performance_tuning_analysis": ListWorkRequestsOperationTypeRequestPerformanceTuningAnalysis,
	"request_java_migration_analysis":     ListWorkRequestsOperationTypeRequestJavaMigrationAnalysis,
	"delete_jms_report":                   ListWorkRequestsOperationTypeDeleteJmsReport,
	"scan_java_server_usage":              ListWorkRequestsOperationTypeScanJavaServerUsage,
	"scan_library_usage":                  ListWorkRequestsOperationTypeScanLibraryUsage,
	"export_data_csv":                     ListWorkRequestsOperationTypeExportDataCsv,
	"create_drs_file":                     ListWorkRequestsOperationTypeCreateDrsFile,
	"update_drs_file":                     ListWorkRequestsOperationTypeUpdateDrsFile,
	"delete_drs_file":                     ListWorkRequestsOperationTypeDeleteDrsFile,
	"enable_drs":                          ListWorkRequestsOperationTypeEnableDrs,
	"disable_drs":                         ListWorkRequestsOperationTypeDisableDrs,
}

// GetListWorkRequestsOperationTypeEnumValues Enumerates the set of values for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumValues() []ListWorkRequestsOperationTypeEnum {
	values := make([]ListWorkRequestsOperationTypeEnum, 0)
	for _, v := range mappingListWorkRequestsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsOperationTypeEnumStringValues Enumerates the set of values in String for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumStringValues() []string {
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

// GetMappingListWorkRequestsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOperationTypeEnum(val string) (ListWorkRequestsOperationTypeEnum, bool) {
	enum, ok := mappingListWorkRequestsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
