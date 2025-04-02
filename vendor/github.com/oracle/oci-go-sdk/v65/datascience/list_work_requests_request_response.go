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

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by the type of the operation associated with the work request.
	OperationType ListWorkRequestsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// <b>Filter</b> results by work request status.
	Status ListWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, the results are shown in descending order. All other fields default to ascending order.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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
	if _, ok := GetMappingListWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
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

	// A list of []WorkRequestSummary instances
	Items []WorkRequestSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
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
	ListWorkRequestsOperationTypeNotebookSessionCreate             ListWorkRequestsOperationTypeEnum = "NOTEBOOK_SESSION_CREATE"
	ListWorkRequestsOperationTypeNotebookSessionDelete             ListWorkRequestsOperationTypeEnum = "NOTEBOOK_SESSION_DELETE"
	ListWorkRequestsOperationTypeNotebookSessionActivate           ListWorkRequestsOperationTypeEnum = "NOTEBOOK_SESSION_ACTIVATE"
	ListWorkRequestsOperationTypeNotebookSessionDeactivate         ListWorkRequestsOperationTypeEnum = "NOTEBOOK_SESSION_DEACTIVATE"
	ListWorkRequestsOperationTypeModelversionsetDelete             ListWorkRequestsOperationTypeEnum = "MODELVERSIONSET_DELETE"
	ListWorkRequestsOperationTypeExportModelArtifact               ListWorkRequestsOperationTypeEnum = "EXPORT_MODEL_ARTIFACT"
	ListWorkRequestsOperationTypeImportModelArtifact               ListWorkRequestsOperationTypeEnum = "IMPORT_MODEL_ARTIFACT"
	ListWorkRequestsOperationTypeModelDeploymentCreate             ListWorkRequestsOperationTypeEnum = "MODEL_DEPLOYMENT_CREATE"
	ListWorkRequestsOperationTypeModelDeploymentDelete             ListWorkRequestsOperationTypeEnum = "MODEL_DEPLOYMENT_DELETE"
	ListWorkRequestsOperationTypeModelDeploymentActivate           ListWorkRequestsOperationTypeEnum = "MODEL_DEPLOYMENT_ACTIVATE"
	ListWorkRequestsOperationTypeModelDeploymentDeactivate         ListWorkRequestsOperationTypeEnum = "MODEL_DEPLOYMENT_DEACTIVATE"
	ListWorkRequestsOperationTypeModelDeploymentUpdate             ListWorkRequestsOperationTypeEnum = "MODEL_DEPLOYMENT_UPDATE"
	ListWorkRequestsOperationTypeProjectDelete                     ListWorkRequestsOperationTypeEnum = "PROJECT_DELETE"
	ListWorkRequestsOperationTypeWorkrequestCancel                 ListWorkRequestsOperationTypeEnum = "WORKREQUEST_CANCEL"
	ListWorkRequestsOperationTypeJobDelete                         ListWorkRequestsOperationTypeEnum = "JOB_DELETE"
	ListWorkRequestsOperationTypePipelineCreate                    ListWorkRequestsOperationTypeEnum = "PIPELINE_CREATE"
	ListWorkRequestsOperationTypePipelineDelete                    ListWorkRequestsOperationTypeEnum = "PIPELINE_DELETE"
	ListWorkRequestsOperationTypePipelineRunCreate                 ListWorkRequestsOperationTypeEnum = "PIPELINE_RUN_CREATE"
	ListWorkRequestsOperationTypePipelineRunCancel                 ListWorkRequestsOperationTypeEnum = "PIPELINE_RUN_CANCEL"
	ListWorkRequestsOperationTypePipelineRunDelete                 ListWorkRequestsOperationTypeEnum = "PIPELINE_RUN_DELETE"
	ListWorkRequestsOperationTypeMlApplicationPackageUpload        ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_PACKAGE_UPLOAD"
	ListWorkRequestsOperationTypeMlApplicationTriggerStart         ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_TRIGGER_START"
	ListWorkRequestsOperationTypeMlApplicationImplementationDelete ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_IMPLEMENTATION_DELETE"
	ListWorkRequestsOperationTypeMlApplicationImplementationUpdate ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_IMPLEMENTATION_UPDATE"
	ListWorkRequestsOperationTypeMlApplicationImplementationMove   ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_IMPLEMENTATION_MOVE"
	ListWorkRequestsOperationTypeMlApplicationInstanceCreate       ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_CREATE"
	ListWorkRequestsOperationTypeMlApplicationInstanceUpdate       ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_UPDATE"
	ListWorkRequestsOperationTypeMlApplicationInstanceDelete       ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_DELETE"
	ListWorkRequestsOperationTypeMlApplicationInstanceMove         ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_MOVE"
	ListWorkRequestsOperationTypeMlApplicationInstanceViewCreate   ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_VIEW_CREATE"
	ListWorkRequestsOperationTypeMlApplicationInstanceViewUpdate   ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_VIEW_UPDATE"
	ListWorkRequestsOperationTypeMlApplicationInstanceViewDelete   ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_VIEW_DELETE"
	ListWorkRequestsOperationTypeMlApplicationInstanceViewUpgrade  ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_VIEW_UPGRADE"
	ListWorkRequestsOperationTypeMlApplicationInstanceViewMove     ListWorkRequestsOperationTypeEnum = "ML_APPLICATION_INSTANCE_VIEW_MOVE"
	ListWorkRequestsOperationTypePrivateEndpointCreate             ListWorkRequestsOperationTypeEnum = "PRIVATE_ENDPOINT_CREATE"
	ListWorkRequestsOperationTypePrivateEndpointDelete             ListWorkRequestsOperationTypeEnum = "PRIVATE_ENDPOINT_DELETE"
	ListWorkRequestsOperationTypePrivateEndpointMove               ListWorkRequestsOperationTypeEnum = "PRIVATE_ENDPOINT_MOVE"
	ListWorkRequestsOperationTypePrivateEndpointUpdate             ListWorkRequestsOperationTypeEnum = "PRIVATE_ENDPOINT_UPDATE"
	ListWorkRequestsOperationTypeScheduleCreate                    ListWorkRequestsOperationTypeEnum = "SCHEDULE_CREATE"
	ListWorkRequestsOperationTypeScheduleUpdate                    ListWorkRequestsOperationTypeEnum = "SCHEDULE_UPDATE"
	ListWorkRequestsOperationTypeScheduleDelete                    ListWorkRequestsOperationTypeEnum = "SCHEDULE_DELETE"
	ListWorkRequestsOperationTypeScheduleMove                      ListWorkRequestsOperationTypeEnum = "SCHEDULE_MOVE"
	ListWorkRequestsOperationTypeScheduleActivate                  ListWorkRequestsOperationTypeEnum = "SCHEDULE_ACTIVATE"
	ListWorkRequestsOperationTypeScheduleDeactivate                ListWorkRequestsOperationTypeEnum = "SCHEDULE_DEACTIVATE"
	ListWorkRequestsOperationTypeRegisterModelArtifact             ListWorkRequestsOperationTypeEnum = "REGISTER_MODEL_ARTIFACT"
	ListWorkRequestsOperationTypeRestoreArchivedModel              ListWorkRequestsOperationTypeEnum = "RESTORE_ARCHIVED_MODEL"
)

var mappingListWorkRequestsOperationTypeEnum = map[string]ListWorkRequestsOperationTypeEnum{
	"NOTEBOOK_SESSION_CREATE":              ListWorkRequestsOperationTypeNotebookSessionCreate,
	"NOTEBOOK_SESSION_DELETE":              ListWorkRequestsOperationTypeNotebookSessionDelete,
	"NOTEBOOK_SESSION_ACTIVATE":            ListWorkRequestsOperationTypeNotebookSessionActivate,
	"NOTEBOOK_SESSION_DEACTIVATE":          ListWorkRequestsOperationTypeNotebookSessionDeactivate,
	"MODELVERSIONSET_DELETE":               ListWorkRequestsOperationTypeModelversionsetDelete,
	"EXPORT_MODEL_ARTIFACT":                ListWorkRequestsOperationTypeExportModelArtifact,
	"IMPORT_MODEL_ARTIFACT":                ListWorkRequestsOperationTypeImportModelArtifact,
	"MODEL_DEPLOYMENT_CREATE":              ListWorkRequestsOperationTypeModelDeploymentCreate,
	"MODEL_DEPLOYMENT_DELETE":              ListWorkRequestsOperationTypeModelDeploymentDelete,
	"MODEL_DEPLOYMENT_ACTIVATE":            ListWorkRequestsOperationTypeModelDeploymentActivate,
	"MODEL_DEPLOYMENT_DEACTIVATE":          ListWorkRequestsOperationTypeModelDeploymentDeactivate,
	"MODEL_DEPLOYMENT_UPDATE":              ListWorkRequestsOperationTypeModelDeploymentUpdate,
	"PROJECT_DELETE":                       ListWorkRequestsOperationTypeProjectDelete,
	"WORKREQUEST_CANCEL":                   ListWorkRequestsOperationTypeWorkrequestCancel,
	"JOB_DELETE":                           ListWorkRequestsOperationTypeJobDelete,
	"PIPELINE_CREATE":                      ListWorkRequestsOperationTypePipelineCreate,
	"PIPELINE_DELETE":                      ListWorkRequestsOperationTypePipelineDelete,
	"PIPELINE_RUN_CREATE":                  ListWorkRequestsOperationTypePipelineRunCreate,
	"PIPELINE_RUN_CANCEL":                  ListWorkRequestsOperationTypePipelineRunCancel,
	"PIPELINE_RUN_DELETE":                  ListWorkRequestsOperationTypePipelineRunDelete,
	"ML_APPLICATION_PACKAGE_UPLOAD":        ListWorkRequestsOperationTypeMlApplicationPackageUpload,
	"ML_APPLICATION_TRIGGER_START":         ListWorkRequestsOperationTypeMlApplicationTriggerStart,
	"ML_APPLICATION_IMPLEMENTATION_DELETE": ListWorkRequestsOperationTypeMlApplicationImplementationDelete,
	"ML_APPLICATION_IMPLEMENTATION_UPDATE": ListWorkRequestsOperationTypeMlApplicationImplementationUpdate,
	"ML_APPLICATION_IMPLEMENTATION_MOVE":   ListWorkRequestsOperationTypeMlApplicationImplementationMove,
	"ML_APPLICATION_INSTANCE_CREATE":       ListWorkRequestsOperationTypeMlApplicationInstanceCreate,
	"ML_APPLICATION_INSTANCE_UPDATE":       ListWorkRequestsOperationTypeMlApplicationInstanceUpdate,
	"ML_APPLICATION_INSTANCE_DELETE":       ListWorkRequestsOperationTypeMlApplicationInstanceDelete,
	"ML_APPLICATION_INSTANCE_MOVE":         ListWorkRequestsOperationTypeMlApplicationInstanceMove,
	"ML_APPLICATION_INSTANCE_VIEW_CREATE":  ListWorkRequestsOperationTypeMlApplicationInstanceViewCreate,
	"ML_APPLICATION_INSTANCE_VIEW_UPDATE":  ListWorkRequestsOperationTypeMlApplicationInstanceViewUpdate,
	"ML_APPLICATION_INSTANCE_VIEW_DELETE":  ListWorkRequestsOperationTypeMlApplicationInstanceViewDelete,
	"ML_APPLICATION_INSTANCE_VIEW_UPGRADE": ListWorkRequestsOperationTypeMlApplicationInstanceViewUpgrade,
	"ML_APPLICATION_INSTANCE_VIEW_MOVE":    ListWorkRequestsOperationTypeMlApplicationInstanceViewMove,
	"PRIVATE_ENDPOINT_CREATE":              ListWorkRequestsOperationTypePrivateEndpointCreate,
	"PRIVATE_ENDPOINT_DELETE":              ListWorkRequestsOperationTypePrivateEndpointDelete,
	"PRIVATE_ENDPOINT_MOVE":                ListWorkRequestsOperationTypePrivateEndpointMove,
	"PRIVATE_ENDPOINT_UPDATE":              ListWorkRequestsOperationTypePrivateEndpointUpdate,
	"SCHEDULE_CREATE":                      ListWorkRequestsOperationTypeScheduleCreate,
	"SCHEDULE_UPDATE":                      ListWorkRequestsOperationTypeScheduleUpdate,
	"SCHEDULE_DELETE":                      ListWorkRequestsOperationTypeScheduleDelete,
	"SCHEDULE_MOVE":                        ListWorkRequestsOperationTypeScheduleMove,
	"SCHEDULE_ACTIVATE":                    ListWorkRequestsOperationTypeScheduleActivate,
	"SCHEDULE_DEACTIVATE":                  ListWorkRequestsOperationTypeScheduleDeactivate,
	"REGISTER_MODEL_ARTIFACT":              ListWorkRequestsOperationTypeRegisterModelArtifact,
	"RESTORE_ARCHIVED_MODEL":               ListWorkRequestsOperationTypeRestoreArchivedModel,
}

var mappingListWorkRequestsOperationTypeEnumLowerCase = map[string]ListWorkRequestsOperationTypeEnum{
	"notebook_session_create":              ListWorkRequestsOperationTypeNotebookSessionCreate,
	"notebook_session_delete":              ListWorkRequestsOperationTypeNotebookSessionDelete,
	"notebook_session_activate":            ListWorkRequestsOperationTypeNotebookSessionActivate,
	"notebook_session_deactivate":          ListWorkRequestsOperationTypeNotebookSessionDeactivate,
	"modelversionset_delete":               ListWorkRequestsOperationTypeModelversionsetDelete,
	"export_model_artifact":                ListWorkRequestsOperationTypeExportModelArtifact,
	"import_model_artifact":                ListWorkRequestsOperationTypeImportModelArtifact,
	"model_deployment_create":              ListWorkRequestsOperationTypeModelDeploymentCreate,
	"model_deployment_delete":              ListWorkRequestsOperationTypeModelDeploymentDelete,
	"model_deployment_activate":            ListWorkRequestsOperationTypeModelDeploymentActivate,
	"model_deployment_deactivate":          ListWorkRequestsOperationTypeModelDeploymentDeactivate,
	"model_deployment_update":              ListWorkRequestsOperationTypeModelDeploymentUpdate,
	"project_delete":                       ListWorkRequestsOperationTypeProjectDelete,
	"workrequest_cancel":                   ListWorkRequestsOperationTypeWorkrequestCancel,
	"job_delete":                           ListWorkRequestsOperationTypeJobDelete,
	"pipeline_create":                      ListWorkRequestsOperationTypePipelineCreate,
	"pipeline_delete":                      ListWorkRequestsOperationTypePipelineDelete,
	"pipeline_run_create":                  ListWorkRequestsOperationTypePipelineRunCreate,
	"pipeline_run_cancel":                  ListWorkRequestsOperationTypePipelineRunCancel,
	"pipeline_run_delete":                  ListWorkRequestsOperationTypePipelineRunDelete,
	"ml_application_package_upload":        ListWorkRequestsOperationTypeMlApplicationPackageUpload,
	"ml_application_trigger_start":         ListWorkRequestsOperationTypeMlApplicationTriggerStart,
	"ml_application_implementation_delete": ListWorkRequestsOperationTypeMlApplicationImplementationDelete,
	"ml_application_implementation_update": ListWorkRequestsOperationTypeMlApplicationImplementationUpdate,
	"ml_application_implementation_move":   ListWorkRequestsOperationTypeMlApplicationImplementationMove,
	"ml_application_instance_create":       ListWorkRequestsOperationTypeMlApplicationInstanceCreate,
	"ml_application_instance_update":       ListWorkRequestsOperationTypeMlApplicationInstanceUpdate,
	"ml_application_instance_delete":       ListWorkRequestsOperationTypeMlApplicationInstanceDelete,
	"ml_application_instance_move":         ListWorkRequestsOperationTypeMlApplicationInstanceMove,
	"ml_application_instance_view_create":  ListWorkRequestsOperationTypeMlApplicationInstanceViewCreate,
	"ml_application_instance_view_update":  ListWorkRequestsOperationTypeMlApplicationInstanceViewUpdate,
	"ml_application_instance_view_delete":  ListWorkRequestsOperationTypeMlApplicationInstanceViewDelete,
	"ml_application_instance_view_upgrade": ListWorkRequestsOperationTypeMlApplicationInstanceViewUpgrade,
	"ml_application_instance_view_move":    ListWorkRequestsOperationTypeMlApplicationInstanceViewMove,
	"private_endpoint_create":              ListWorkRequestsOperationTypePrivateEndpointCreate,
	"private_endpoint_delete":              ListWorkRequestsOperationTypePrivateEndpointDelete,
	"private_endpoint_move":                ListWorkRequestsOperationTypePrivateEndpointMove,
	"private_endpoint_update":              ListWorkRequestsOperationTypePrivateEndpointUpdate,
	"schedule_create":                      ListWorkRequestsOperationTypeScheduleCreate,
	"schedule_update":                      ListWorkRequestsOperationTypeScheduleUpdate,
	"schedule_delete":                      ListWorkRequestsOperationTypeScheduleDelete,
	"schedule_move":                        ListWorkRequestsOperationTypeScheduleMove,
	"schedule_activate":                    ListWorkRequestsOperationTypeScheduleActivate,
	"schedule_deactivate":                  ListWorkRequestsOperationTypeScheduleDeactivate,
	"register_model_artifact":              ListWorkRequestsOperationTypeRegisterModelArtifact,
	"restore_archived_model":               ListWorkRequestsOperationTypeRestoreArchivedModel,
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
		"NOTEBOOK_SESSION_CREATE",
		"NOTEBOOK_SESSION_DELETE",
		"NOTEBOOK_SESSION_ACTIVATE",
		"NOTEBOOK_SESSION_DEACTIVATE",
		"MODELVERSIONSET_DELETE",
		"EXPORT_MODEL_ARTIFACT",
		"IMPORT_MODEL_ARTIFACT",
		"MODEL_DEPLOYMENT_CREATE",
		"MODEL_DEPLOYMENT_DELETE",
		"MODEL_DEPLOYMENT_ACTIVATE",
		"MODEL_DEPLOYMENT_DEACTIVATE",
		"MODEL_DEPLOYMENT_UPDATE",
		"PROJECT_DELETE",
		"WORKREQUEST_CANCEL",
		"JOB_DELETE",
		"PIPELINE_CREATE",
		"PIPELINE_DELETE",
		"PIPELINE_RUN_CREATE",
		"PIPELINE_RUN_CANCEL",
		"PIPELINE_RUN_DELETE",
		"ML_APPLICATION_PACKAGE_UPLOAD",
		"ML_APPLICATION_TRIGGER_START",
		"ML_APPLICATION_IMPLEMENTATION_DELETE",
		"ML_APPLICATION_IMPLEMENTATION_UPDATE",
		"ML_APPLICATION_IMPLEMENTATION_MOVE",
		"ML_APPLICATION_INSTANCE_CREATE",
		"ML_APPLICATION_INSTANCE_UPDATE",
		"ML_APPLICATION_INSTANCE_DELETE",
		"ML_APPLICATION_INSTANCE_MOVE",
		"ML_APPLICATION_INSTANCE_VIEW_CREATE",
		"ML_APPLICATION_INSTANCE_VIEW_UPDATE",
		"ML_APPLICATION_INSTANCE_VIEW_DELETE",
		"ML_APPLICATION_INSTANCE_VIEW_UPGRADE",
		"ML_APPLICATION_INSTANCE_VIEW_MOVE",
		"PRIVATE_ENDPOINT_CREATE",
		"PRIVATE_ENDPOINT_DELETE",
		"PRIVATE_ENDPOINT_MOVE",
		"PRIVATE_ENDPOINT_UPDATE",
		"SCHEDULE_CREATE",
		"SCHEDULE_UPDATE",
		"SCHEDULE_DELETE",
		"SCHEDULE_MOVE",
		"SCHEDULE_ACTIVATE",
		"SCHEDULE_DEACTIVATE",
		"REGISTER_MODEL_ARTIFACT",
		"RESTORE_ARCHIVED_MODEL",
	}
}

// GetMappingListWorkRequestsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOperationTypeEnum(val string) (ListWorkRequestsOperationTypeEnum, bool) {
	enum, ok := mappingListWorkRequestsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsStatusEnum Enum with underlying type: string
type ListWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsStatusEnum
const (
	ListWorkRequestsStatusAccepted   ListWorkRequestsStatusEnum = "ACCEPTED"
	ListWorkRequestsStatusInProgress ListWorkRequestsStatusEnum = "IN_PROGRESS"
	ListWorkRequestsStatusFailed     ListWorkRequestsStatusEnum = "FAILED"
	ListWorkRequestsStatusSucceeded  ListWorkRequestsStatusEnum = "SUCCEEDED"
	ListWorkRequestsStatusCanceling  ListWorkRequestsStatusEnum = "CANCELING"
	ListWorkRequestsStatusCanceled   ListWorkRequestsStatusEnum = "CANCELED"
)

var mappingListWorkRequestsStatusEnum = map[string]ListWorkRequestsStatusEnum{
	"ACCEPTED":    ListWorkRequestsStatusAccepted,
	"IN_PROGRESS": ListWorkRequestsStatusInProgress,
	"FAILED":      ListWorkRequestsStatusFailed,
	"SUCCEEDED":   ListWorkRequestsStatusSucceeded,
	"CANCELING":   ListWorkRequestsStatusCanceling,
	"CANCELED":    ListWorkRequestsStatusCanceled,
}

var mappingListWorkRequestsStatusEnumLowerCase = map[string]ListWorkRequestsStatusEnum{
	"accepted":    ListWorkRequestsStatusAccepted,
	"in_progress": ListWorkRequestsStatusInProgress,
	"failed":      ListWorkRequestsStatusFailed,
	"succeeded":   ListWorkRequestsStatusSucceeded,
	"canceling":   ListWorkRequestsStatusCanceling,
	"canceled":    ListWorkRequestsStatusCanceled,
}

// GetListWorkRequestsStatusEnumValues Enumerates the set of values for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumValues() []ListWorkRequestsStatusEnum {
	values := make([]ListWorkRequestsStatusEnum, 0)
	for _, v := range mappingListWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsStatusEnum(val string) (ListWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrderEnum = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

var mappingListWorkRequestsSortOrderEnumLowerCase = map[string]ListWorkRequestsSortOrderEnum{
	"asc":  ListWorkRequestsSortOrderAsc,
	"desc": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortOrderEnum(val string) (ListWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByOperationtype ListWorkRequestsSortByEnum = "operationType"
	ListWorkRequestsSortByStatus        ListWorkRequestsSortByEnum = "status"
	ListWorkRequestsSortByTimeaccepted  ListWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListWorkRequestsSortByEnum = map[string]ListWorkRequestsSortByEnum{
	"operationType": ListWorkRequestsSortByOperationtype,
	"status":        ListWorkRequestsSortByStatus,
	"timeAccepted":  ListWorkRequestsSortByTimeaccepted,
}

var mappingListWorkRequestsSortByEnumLowerCase = map[string]ListWorkRequestsSortByEnum{
	"operationtype": ListWorkRequestsSortByOperationtype,
	"status":        ListWorkRequestsSortByStatus,
	"timeaccepted":  ListWorkRequestsSortByTimeaccepted,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"operationType",
		"status",
		"timeAccepted",
	}
}

// GetMappingListWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortByEnum(val string) (ListWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
