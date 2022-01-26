// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListBucketsRequest wrapper for the ListBuckets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/ListBuckets.go.html to see an example of how to use ListBucketsRequest.
type ListBucketsRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list buckets.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Bucket summary in list of buckets includes the 'namespace', 'name', 'compartmentId', 'createdBy', 'timeCreated',
	// and 'etag' fields. This parameter can also include 'tags' (freeformTags and definedTags). The only supported value of this parameter is 'tags' for now. Example 'tags'.
	Fields []ListBucketsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"csv"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBucketsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBucketsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBucketsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBucketsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBucketsResponse wrapper for the ListBuckets operation
type ListBucketsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BucketSummary instances
	Items []BucketSummary `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For paginating a list of buckets.
	// In the GET request, set the limit to the number of buckets items that you want returned in the response.
	// If the `opc-next-page` header appears in the response, then this is a partial list and there are additional
	// buckets to get. Include the header's value as the `page` parameter in the subsequent GET request to get the
	// next batch of buckets. Repeat this process to retrieve the entire list of buckets.
	// By default, the page limit is set to 25 buckets per page, but you can specify a value from 1 to 1000.
	// For more details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBucketsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBucketsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBucketsFieldsEnum Enum with underlying type: string
type ListBucketsFieldsEnum string

// Set of constants representing the allowable values for ListBucketsFieldsEnum
const (
	ListBucketsFieldsTags ListBucketsFieldsEnum = "tags"
)

var mappingListBucketsFields = map[string]ListBucketsFieldsEnum{
	"tags": ListBucketsFieldsTags,
}

// GetListBucketsFieldsEnumValues Enumerates the set of values for ListBucketsFieldsEnum
func GetListBucketsFieldsEnumValues() []ListBucketsFieldsEnum {
	values := make([]ListBucketsFieldsEnum, 0)
	for _, v := range mappingListBucketsFields {
		values = append(values, v)
	}
	return values
}
