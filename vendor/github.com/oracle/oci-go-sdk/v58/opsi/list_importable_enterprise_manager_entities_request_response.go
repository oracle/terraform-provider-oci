// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListImportableEnterpriseManagerEntitiesRequest wrapper for the ListImportableEnterpriseManagerEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableEnterpriseManagerEntities.go.html to see an example of how to use ListImportableEnterpriseManagerEntitiesRequest.
type ListImportableEnterpriseManagerEntitiesRequest struct {

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"true" contributesTo:"path" name:"enterpriseManagerBridgeId"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter by one or more Enterprise Manager entity types. Currently, the supported types are "oracle_pdb", "oracle_database", "host", "oracle_dbmachine", "oracle_exa_cloud_service", and "oracle_oci_exadata_cloud_service". If this parameter is not specified, targets of all supported entity types are returned by default.
	EnterpriseManagerEntityType []string `contributesTo:"query" name:"enterpriseManagerEntityType" collectionFormat:"multi"`

	// Used in combination with enterpriseManagerParentEntityIdentifier to return the members of a particular Enterprise Manager parent entity. Both enterpriseManagerIdentifier and enterpriseManagerParentEntityIdentifier must be specified to identify a particular Enterprise Manager parent entity.
	EnterpriseManagerIdentifier *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerIdentifier"`

	// Used in combination with enterpriseManagerIdentifier to return the members of a particular Enterprise Manager parent entity. Both enterpriseManagerIdentifier and enterpriseManagerParentEntityIdentifier must be specified to identify a particular  Enterprise Manager parent entity.
	EnterpriseManagerParentEntityIdentifier *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerParentEntityIdentifier"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImportableEnterpriseManagerEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImportableEnterpriseManagerEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImportableEnterpriseManagerEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImportableEnterpriseManagerEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListImportableEnterpriseManagerEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListImportableEnterpriseManagerEntitiesResponse wrapper for the ListImportableEnterpriseManagerEntities operation
type ListImportableEnterpriseManagerEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ImportableEnterpriseManagerEntityCollection instances
	ImportableEnterpriseManagerEntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImportableEnterpriseManagerEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImportableEnterpriseManagerEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
