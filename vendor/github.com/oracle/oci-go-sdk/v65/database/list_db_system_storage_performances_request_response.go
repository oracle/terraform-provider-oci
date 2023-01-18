// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemStoragePerformances.go.html to see an example of how to use ListDbSystemStoragePerformancesRequest.
type ListDbSystemStoragePerformancesRequest struct {

	// The DB system storage management option. Used to list database versions available for that storage manager. Valid values are `ASM` and `LVM`.
	// * ASM specifies Oracle Automatic Storage Management
	// * LVM specifies logical volume manager, sometimes called logical disk manager.
	StorageManagement DbSystemOptionsStorageManagementEnum `mandatory:"true" contributesTo:"query" name:"storageManagement" omitEmpty:"true"`

	// Optional. Filters the performance results by shape type.
	ShapeType *string `mandatory:"false" contributesTo:"query" name:"shapeType"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemStoragePerformancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemStoragePerformancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
