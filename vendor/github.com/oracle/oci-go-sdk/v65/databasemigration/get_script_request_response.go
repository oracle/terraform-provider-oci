// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetScriptRequest wrapper for the GetScript operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetScript.go.html to see an example of how to use GetScriptRequest.
type GetScriptRequest struct {

	// The ID of the script to download.
	ScriptId GetScriptScriptIdEnum `mandatory:"true" contributesTo:"path" name:"scriptId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetScriptRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetScriptRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetScriptRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetScriptRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetScriptRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetScriptScriptIdEnum(string(request.ScriptId)); !ok && request.ScriptId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScriptId: %s. Supported values are: %s.", request.ScriptId, strings.Join(GetGetScriptScriptIdEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetScriptResponse wrapper for the GetScript operation
type GetScriptResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetScriptResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetScriptResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetScriptScriptIdEnum Enum with underlying type: string
type GetScriptScriptIdEnum string

// Set of constants representing the allowable values for GetScriptScriptIdEnum
const (
	GetScriptScriptIdUserCreationSqlScript GetScriptScriptIdEnum = "USER_CREATION_SQL_SCRIPT"
)

var mappingGetScriptScriptIdEnum = map[string]GetScriptScriptIdEnum{
	"USER_CREATION_SQL_SCRIPT": GetScriptScriptIdUserCreationSqlScript,
}

var mappingGetScriptScriptIdEnumLowerCase = map[string]GetScriptScriptIdEnum{
	"user_creation_sql_script": GetScriptScriptIdUserCreationSqlScript,
}

// GetGetScriptScriptIdEnumValues Enumerates the set of values for GetScriptScriptIdEnum
func GetGetScriptScriptIdEnumValues() []GetScriptScriptIdEnum {
	values := make([]GetScriptScriptIdEnum, 0)
	for _, v := range mappingGetScriptScriptIdEnum {
		values = append(values, v)
	}
	return values
}

// GetGetScriptScriptIdEnumStringValues Enumerates the set of values in String for GetScriptScriptIdEnum
func GetGetScriptScriptIdEnumStringValues() []string {
	return []string{
		"USER_CREATION_SQL_SCRIPT",
	}
}

// GetMappingGetScriptScriptIdEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetScriptScriptIdEnum(val string) (GetScriptScriptIdEnum, bool) {
	enum, ok := mappingGetScriptScriptIdEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
