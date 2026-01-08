// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPsaInternalDataRequest wrapper for the ListPsaInternalData operation
type ListPsaInternalDataRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"true" contributesTo:"query" name:"vcnId"`

	// A list of serviceId filter to return only resources that match the given serviceIds separated by comma (https://swagger.io/docs/specification/v2_0/describing-parameters/#array-and-multi-value-parameters). The state value is case-insensitive.
	ServiceIds []string `contributesTo:"query" name:"serviceIds" collectionFormat:"csv"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ListPsaInternalDataLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPsaInternalDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPsaInternalDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPsaInternalDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPsaInternalDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPsaInternalDataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPsaInternalDataLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListPsaInternalDataLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPsaInternalDataResponse wrapper for the ListPsaInternalData operation
type ListPsaInternalDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PsaInternalDataCollection instance
	PsaInternalDataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPsaInternalDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPsaInternalDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPsaInternalDataLifecycleStateEnum Enum with underlying type: string
type ListPsaInternalDataLifecycleStateEnum string

// Set of constants representing the allowable values for ListPsaInternalDataLifecycleStateEnum
const (
	ListPsaInternalDataLifecycleStateCreating ListPsaInternalDataLifecycleStateEnum = "CREATING"
	ListPsaInternalDataLifecycleStateUpdating ListPsaInternalDataLifecycleStateEnum = "UPDATING"
	ListPsaInternalDataLifecycleStateActive   ListPsaInternalDataLifecycleStateEnum = "ACTIVE"
	ListPsaInternalDataLifecycleStateDeleting ListPsaInternalDataLifecycleStateEnum = "DELETING"
	ListPsaInternalDataLifecycleStateDeleted  ListPsaInternalDataLifecycleStateEnum = "DELETED"
	ListPsaInternalDataLifecycleStateFailed   ListPsaInternalDataLifecycleStateEnum = "FAILED"
)

var mappingListPsaInternalDataLifecycleStateEnum = map[string]ListPsaInternalDataLifecycleStateEnum{
	"CREATING": ListPsaInternalDataLifecycleStateCreating,
	"UPDATING": ListPsaInternalDataLifecycleStateUpdating,
	"ACTIVE":   ListPsaInternalDataLifecycleStateActive,
	"DELETING": ListPsaInternalDataLifecycleStateDeleting,
	"DELETED":  ListPsaInternalDataLifecycleStateDeleted,
	"FAILED":   ListPsaInternalDataLifecycleStateFailed,
}

var mappingListPsaInternalDataLifecycleStateEnumLowerCase = map[string]ListPsaInternalDataLifecycleStateEnum{
	"creating": ListPsaInternalDataLifecycleStateCreating,
	"updating": ListPsaInternalDataLifecycleStateUpdating,
	"active":   ListPsaInternalDataLifecycleStateActive,
	"deleting": ListPsaInternalDataLifecycleStateDeleting,
	"deleted":  ListPsaInternalDataLifecycleStateDeleted,
	"failed":   ListPsaInternalDataLifecycleStateFailed,
}

// GetListPsaInternalDataLifecycleStateEnumValues Enumerates the set of values for ListPsaInternalDataLifecycleStateEnum
func GetListPsaInternalDataLifecycleStateEnumValues() []ListPsaInternalDataLifecycleStateEnum {
	values := make([]ListPsaInternalDataLifecycleStateEnum, 0)
	for _, v := range mappingListPsaInternalDataLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaInternalDataLifecycleStateEnumStringValues Enumerates the set of values in String for ListPsaInternalDataLifecycleStateEnum
func GetListPsaInternalDataLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListPsaInternalDataLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaInternalDataLifecycleStateEnum(val string) (ListPsaInternalDataLifecycleStateEnum, bool) {
	enum, ok := mappingListPsaInternalDataLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
