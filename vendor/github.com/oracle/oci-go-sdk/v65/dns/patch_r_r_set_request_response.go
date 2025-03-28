// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// PatchRRSetRequest wrapper for the PatchRRSet operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/PatchRRSet.go.html to see an example of how to use PatchRRSetRequest.
type PatchRRSetRequest struct {

	// The name or OCID of the target zone.
	ZoneNameOrId *string `mandatory:"true" contributesTo:"path" name:"zoneNameOrId"`

	// The target fully-qualified domain name (FQDN) within the target zone.
	Domain *string `mandatory:"true" contributesTo:"path" name:"domain"`

	// The type of the target RRSet within the target zone.
	Rtype *string `mandatory:"true" contributesTo:"path" name:"rtype"`

	// Operations describing how to modify the collection of records.
	PatchRrSetDetails `contributesTo:"body"`

	// The `If-Match` header field makes the request method conditional on the
	// existence of at least one current representation of the target resource,
	// when the field-value is `*`, or having a current representation of the
	// target resource that has an entity-tag matching a member of the list of
	// entity-tags provided in the field-value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

	// The `If-Unmodified-Since` header field makes the request method
	// conditional on the selected representation's last modification date being
	// earlier than or equal to the date provided in the field-value.  This
	// field accomplishes the same purpose as If-Match for cases where the user
	// agent does not have an entity-tag for the representation.
	IfUnmodifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Unmodified-Since"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope PatchRRSetScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the zone is associated with. Required when accessing a private zone by name.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// The OCID of the compartment the zone belongs to.
	// This parameter is deprecated and should be omitted.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PatchRRSetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PatchRRSetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request PatchRRSetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PatchRRSetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PatchRRSetRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchRRSetScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetPatchRRSetScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchRRSetResponse wrapper for the PatchRRSet operation
type PatchRRSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RecordCollection instance
	RecordCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response PatchRRSetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PatchRRSetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// PatchRRSetScopeEnum Enum with underlying type: string
type PatchRRSetScopeEnum string

// Set of constants representing the allowable values for PatchRRSetScopeEnum
const (
	PatchRRSetScopeGlobal  PatchRRSetScopeEnum = "GLOBAL"
	PatchRRSetScopePrivate PatchRRSetScopeEnum = "PRIVATE"
)

var mappingPatchRRSetScopeEnum = map[string]PatchRRSetScopeEnum{
	"GLOBAL":  PatchRRSetScopeGlobal,
	"PRIVATE": PatchRRSetScopePrivate,
}

var mappingPatchRRSetScopeEnumLowerCase = map[string]PatchRRSetScopeEnum{
	"global":  PatchRRSetScopeGlobal,
	"private": PatchRRSetScopePrivate,
}

// GetPatchRRSetScopeEnumValues Enumerates the set of values for PatchRRSetScopeEnum
func GetPatchRRSetScopeEnumValues() []PatchRRSetScopeEnum {
	values := make([]PatchRRSetScopeEnum, 0)
	for _, v := range mappingPatchRRSetScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchRRSetScopeEnumStringValues Enumerates the set of values in String for PatchRRSetScopeEnum
func GetPatchRRSetScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingPatchRRSetScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchRRSetScopeEnum(val string) (PatchRRSetScopeEnum, bool) {
	enum, ok := mappingPatchRRSetScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
