// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInventoryRulesRequest wrapper for the ListInventoryRules operation
type ListInventoryRulesRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInventoryRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInventoryRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInventoryRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListInventoryRulesRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["namespaceName"] != nil {
		templateParam := mandatoryParamMap["namespaceName"]
		for _, template := range templateParam {
			replacementParam := *request.NamespaceName
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInventoryRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInventoryRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInventoryRulesResponse wrapper for the ListInventoryRules operation
type ListInventoryRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InventoryRuleCollection instances
	InventoryRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Paginating a list of Inventory rules.
	// If the `opc-next-page` header appears in the response, it indicates that this is a partial list
	// of inventory rules and there are additional rules to get. Include the value of this header as
	// the `page` parameter in a subsequent GET request to get the next set of inventory rules.
	// Repeat this process to retrieve the entire list of inventory rules.
	// For more details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInventoryRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInventoryRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
