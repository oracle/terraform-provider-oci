// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetWorkflowOrReplicationTargetRequest wrapper for the GetWorkflowOrReplicationTarget operation
type GetWorkflowOrReplicationTargetRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	ReplicationId *string `mandatory:"true" contributesTo:"query" name:"replicationId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target file system.
	FileSystemId *string `mandatory:"true" contributesTo:"query" name:"fileSystemId"`

	// The type of replication target workflow
	WorkflowType GetWorkflowOrReplicationTargetWorkflowTypeEnum `mandatory:"true" contributesTo:"query" name:"workflowType" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetWorkflowOrReplicationTargetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetWorkflowOrReplicationTargetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetWorkflowOrReplicationTargetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request GetWorkflowOrReplicationTargetRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["replicationId"] != nil {
		templateParam := mandatoryParamMap["replicationId"]
		for _, template := range templateParam {
			replacementParam := *request.ReplicationId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
	if mandatoryParamMap["fileSystemId"] != nil {
		templateParam := mandatoryParamMap["fileSystemId"]
		for _, template := range templateParam {
			replacementParam := *request.FileSystemId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetWorkflowOrReplicationTargetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetWorkflowOrReplicationTargetRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetWorkflowOrReplicationTargetWorkflowTypeEnum(string(request.WorkflowType)); !ok && request.WorkflowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkflowType: %s. Supported values are: %s.", request.WorkflowType, strings.Join(GetGetWorkflowOrReplicationTargetWorkflowTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetWorkflowOrReplicationTargetResponse wrapper for the GetWorkflowOrReplicationTarget operation
type GetWorkflowOrReplicationTargetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GetWorkflowOrReplicationTargetDetails instance
	GetWorkflowOrReplicationTargetDetails `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetWorkflowOrReplicationTargetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetWorkflowOrReplicationTargetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetWorkflowOrReplicationTargetWorkflowTypeEnum Enum with underlying type: string
type GetWorkflowOrReplicationTargetWorkflowTypeEnum string

// Set of constants representing the allowable values for GetWorkflowOrReplicationTargetWorkflowTypeEnum
const (
	GetWorkflowOrReplicationTargetWorkflowTypeCreate GetWorkflowOrReplicationTargetWorkflowTypeEnum = "CREATE"
	GetWorkflowOrReplicationTargetWorkflowTypeUpdate GetWorkflowOrReplicationTargetWorkflowTypeEnum = "UPDATE"
)

var mappingGetWorkflowOrReplicationTargetWorkflowTypeEnum = map[string]GetWorkflowOrReplicationTargetWorkflowTypeEnum{
	"CREATE": GetWorkflowOrReplicationTargetWorkflowTypeCreate,
	"UPDATE": GetWorkflowOrReplicationTargetWorkflowTypeUpdate,
}

var mappingGetWorkflowOrReplicationTargetWorkflowTypeEnumLowerCase = map[string]GetWorkflowOrReplicationTargetWorkflowTypeEnum{
	"create": GetWorkflowOrReplicationTargetWorkflowTypeCreate,
	"update": GetWorkflowOrReplicationTargetWorkflowTypeUpdate,
}

// GetGetWorkflowOrReplicationTargetWorkflowTypeEnumValues Enumerates the set of values for GetWorkflowOrReplicationTargetWorkflowTypeEnum
func GetGetWorkflowOrReplicationTargetWorkflowTypeEnumValues() []GetWorkflowOrReplicationTargetWorkflowTypeEnum {
	values := make([]GetWorkflowOrReplicationTargetWorkflowTypeEnum, 0)
	for _, v := range mappingGetWorkflowOrReplicationTargetWorkflowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetWorkflowOrReplicationTargetWorkflowTypeEnumStringValues Enumerates the set of values in String for GetWorkflowOrReplicationTargetWorkflowTypeEnum
func GetGetWorkflowOrReplicationTargetWorkflowTypeEnumStringValues() []string {
	return []string{
		"CREATE",
		"UPDATE",
	}
}

// GetMappingGetWorkflowOrReplicationTargetWorkflowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetWorkflowOrReplicationTargetWorkflowTypeEnum(val string) (GetWorkflowOrReplicationTargetWorkflowTypeEnum, bool) {
	enum, ok := mappingGetWorkflowOrReplicationTargetWorkflowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
