// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DeleteProtectedDatabaseRequest wrapper for the DeleteProtectedDatabase operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/DeleteProtectedDatabase.go.html to see an example of how to use DeleteProtectedDatabaseRequest.
type DeleteProtectedDatabaseRequest struct {

	// The protected database OCID.
	ProtectedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"protectedDatabaseId"`

	// Defines a preferred schedule to delete a protected database after you terminate the source database.
	// * The default schedule is DELETE_AFTER_72_HOURS, so that the delete operation can occur 72 hours (3 days) after the source database is terminated .
	// * The alternate schedule is DELETE_AFTER_RETENTION_PERIOD. Specify this option if you want to delete a protected database only after the policy-defined backup retention period expires.
	DeletionSchedule DeleteProtectedDatabaseDeletionScheduleEnum `mandatory:"false" contributesTo:"query" name:"deletionSchedule" omitEmpty:"true"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteProtectedDatabaseRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteProtectedDatabaseRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteProtectedDatabaseRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteProtectedDatabaseRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteProtectedDatabaseRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeleteProtectedDatabaseDeletionScheduleEnum(string(request.DeletionSchedule)); !ok && request.DeletionSchedule != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeletionSchedule: %s. Supported values are: %s.", request.DeletionSchedule, strings.Join(GetDeleteProtectedDatabaseDeletionScheduleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteProtectedDatabaseResponse wrapper for the DeleteProtectedDatabase operation
type DeleteProtectedDatabaseResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier of the work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteProtectedDatabaseResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteProtectedDatabaseResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteProtectedDatabaseDeletionScheduleEnum Enum with underlying type: string
type DeleteProtectedDatabaseDeletionScheduleEnum string

// Set of constants representing the allowable values for DeleteProtectedDatabaseDeletionScheduleEnum
const (
	DeleteProtectedDatabaseDeletionScheduleRetentionPeriod DeleteProtectedDatabaseDeletionScheduleEnum = "DELETE_AFTER_RETENTION_PERIOD"
	DeleteProtectedDatabaseDeletionSchedule72Hours         DeleteProtectedDatabaseDeletionScheduleEnum = "DELETE_AFTER_72_HOURS"
)

var mappingDeleteProtectedDatabaseDeletionScheduleEnum = map[string]DeleteProtectedDatabaseDeletionScheduleEnum{
	"DELETE_AFTER_RETENTION_PERIOD": DeleteProtectedDatabaseDeletionScheduleRetentionPeriod,
	"DELETE_AFTER_72_HOURS":         DeleteProtectedDatabaseDeletionSchedule72Hours,
}

var mappingDeleteProtectedDatabaseDeletionScheduleEnumLowerCase = map[string]DeleteProtectedDatabaseDeletionScheduleEnum{
	"delete_after_retention_period": DeleteProtectedDatabaseDeletionScheduleRetentionPeriod,
	"delete_after_72_hours":         DeleteProtectedDatabaseDeletionSchedule72Hours,
}

// GetDeleteProtectedDatabaseDeletionScheduleEnumValues Enumerates the set of values for DeleteProtectedDatabaseDeletionScheduleEnum
func GetDeleteProtectedDatabaseDeletionScheduleEnumValues() []DeleteProtectedDatabaseDeletionScheduleEnum {
	values := make([]DeleteProtectedDatabaseDeletionScheduleEnum, 0)
	for _, v := range mappingDeleteProtectedDatabaseDeletionScheduleEnum {
		values = append(values, v)
	}
	return values
}

// GetDeleteProtectedDatabaseDeletionScheduleEnumStringValues Enumerates the set of values in String for DeleteProtectedDatabaseDeletionScheduleEnum
func GetDeleteProtectedDatabaseDeletionScheduleEnumStringValues() []string {
	return []string{
		"DELETE_AFTER_RETENTION_PERIOD",
		"DELETE_AFTER_72_HOURS",
	}
}

// GetMappingDeleteProtectedDatabaseDeletionScheduleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeleteProtectedDatabaseDeletionScheduleEnum(val string) (DeleteProtectedDatabaseDeletionScheduleEnum, bool) {
	enum, ok := mappingDeleteProtectedDatabaseDeletionScheduleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
