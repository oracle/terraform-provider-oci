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

// DeleteReplicationRequest wrapper for the DeleteReplication operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/DeleteReplication.go.html to see an example of how to use DeleteReplicationRequest.
type DeleteReplicationRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	ReplicationId *string `mandatory:"true" contributesTo:"path" name:"replicationId"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// You can choose a mode for deleting the replication resource.
	// - `FINISH_CYCLE_IF_CAPTURING_OR_APPLYING` Before deleting, complete the current delta cycle. If cycle is idle, delete immediately. Safest option.
	// - `ONE_MORE_CYCLE` Before deleting, complete the current delta cycle, and initiate one more cycle. If cycle is idle, initiate one more cycle. Use for lossless failover.
	// - `FINISH_CYCLE_IF_APPLYING` Before deleting, finish applying. If cycle is idle or capturing, delete immediately. Fastest option.
	DeleteMode DeleteReplicationDeleteModeEnum `mandatory:"false" contributesTo:"query" name:"deleteMode" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteReplicationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteReplicationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteReplicationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteReplicationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteReplicationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeleteReplicationDeleteModeEnum(string(request.DeleteMode)); !ok && request.DeleteMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeleteMode: %s. Supported values are: %s.", request.DeleteMode, strings.Join(GetDeleteReplicationDeleteModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteReplicationResponse wrapper for the DeleteReplication operation
type DeleteReplicationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteReplicationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteReplicationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteReplicationDeleteModeEnum Enum with underlying type: string
type DeleteReplicationDeleteModeEnum string

// Set of constants representing the allowable values for DeleteReplicationDeleteModeEnum
const (
	DeleteReplicationDeleteModeFinishCycleIfCapturingOrApplying DeleteReplicationDeleteModeEnum = "FINISH_CYCLE_IF_CAPTURING_OR_APPLYING"
	DeleteReplicationDeleteModeOneMoreCycle                     DeleteReplicationDeleteModeEnum = "ONE_MORE_CYCLE"
	DeleteReplicationDeleteModeFinishCycleIfApplying            DeleteReplicationDeleteModeEnum = "FINISH_CYCLE_IF_APPLYING"
)

var mappingDeleteReplicationDeleteModeEnum = map[string]DeleteReplicationDeleteModeEnum{
	"FINISH_CYCLE_IF_CAPTURING_OR_APPLYING": DeleteReplicationDeleteModeFinishCycleIfCapturingOrApplying,
	"ONE_MORE_CYCLE":                        DeleteReplicationDeleteModeOneMoreCycle,
	"FINISH_CYCLE_IF_APPLYING":              DeleteReplicationDeleteModeFinishCycleIfApplying,
}

var mappingDeleteReplicationDeleteModeEnumLowerCase = map[string]DeleteReplicationDeleteModeEnum{
	"finish_cycle_if_capturing_or_applying": DeleteReplicationDeleteModeFinishCycleIfCapturingOrApplying,
	"one_more_cycle":                        DeleteReplicationDeleteModeOneMoreCycle,
	"finish_cycle_if_applying":              DeleteReplicationDeleteModeFinishCycleIfApplying,
}

// GetDeleteReplicationDeleteModeEnumValues Enumerates the set of values for DeleteReplicationDeleteModeEnum
func GetDeleteReplicationDeleteModeEnumValues() []DeleteReplicationDeleteModeEnum {
	values := make([]DeleteReplicationDeleteModeEnum, 0)
	for _, v := range mappingDeleteReplicationDeleteModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeleteReplicationDeleteModeEnumStringValues Enumerates the set of values in String for DeleteReplicationDeleteModeEnum
func GetDeleteReplicationDeleteModeEnumStringValues() []string {
	return []string{
		"FINISH_CYCLE_IF_CAPTURING_OR_APPLYING",
		"ONE_MORE_CYCLE",
		"FINISH_CYCLE_IF_APPLYING",
	}
}

// GetMappingDeleteReplicationDeleteModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeleteReplicationDeleteModeEnum(val string) (DeleteReplicationDeleteModeEnum, bool) {
	enum, ok := mappingDeleteReplicationDeleteModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
