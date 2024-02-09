// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetAllC3DrgAttachmentsRequest wrapper for the GetAllC3DrgAttachments operation
type GetAllC3DrgAttachmentsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId *string `mandatory:"true" contributesTo:"path" name:"drgId"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The type for the network resource attached to the DRG.
	AttachmentType GetAllC3DrgAttachmentsAttachmentTypeEnum `mandatory:"false" contributesTo:"query" name:"attachmentType" omitEmpty:"true"`

	// Whether the DRG attachment lives in a different tenancy than the DRG.
	IsCrossTenancy *bool `mandatory:"false" contributesTo:"query" name:"isCrossTenancy"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAllC3DrgAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAllC3DrgAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAllC3DrgAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAllC3DrgAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAllC3DrgAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAllC3DrgAttachmentsAttachmentTypeEnum(string(request.AttachmentType)); !ok && request.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", request.AttachmentType, strings.Join(GetGetAllC3DrgAttachmentsAttachmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAllC3DrgAttachmentsResponse wrapper for the GetAllC3DrgAttachments operation
type GetAllC3DrgAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DrgAttachmentInfo instances
	Items []DrgAttachmentInfo `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAllC3DrgAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAllC3DrgAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAllC3DrgAttachmentsAttachmentTypeEnum Enum with underlying type: string
type GetAllC3DrgAttachmentsAttachmentTypeEnum string

// Set of constants representing the allowable values for GetAllC3DrgAttachmentsAttachmentTypeEnum
const (
	GetAllC3DrgAttachmentsAttachmentTypeVcn                     GetAllC3DrgAttachmentsAttachmentTypeEnum = "VCN"
	GetAllC3DrgAttachmentsAttachmentTypeVirtualCircuit          GetAllC3DrgAttachmentsAttachmentTypeEnum = "VIRTUAL_CIRCUIT"
	GetAllC3DrgAttachmentsAttachmentTypeRemotePeeringConnection GetAllC3DrgAttachmentsAttachmentTypeEnum = "REMOTE_PEERING_CONNECTION"
	GetAllC3DrgAttachmentsAttachmentTypeIpsecTunnel             GetAllC3DrgAttachmentsAttachmentTypeEnum = "IPSEC_TUNNEL"
	GetAllC3DrgAttachmentsAttachmentTypeAll                     GetAllC3DrgAttachmentsAttachmentTypeEnum = "ALL"
)

var mappingGetAllC3DrgAttachmentsAttachmentTypeEnum = map[string]GetAllC3DrgAttachmentsAttachmentTypeEnum{
	"VCN":                       GetAllC3DrgAttachmentsAttachmentTypeVcn,
	"VIRTUAL_CIRCUIT":           GetAllC3DrgAttachmentsAttachmentTypeVirtualCircuit,
	"REMOTE_PEERING_CONNECTION": GetAllC3DrgAttachmentsAttachmentTypeRemotePeeringConnection,
	"IPSEC_TUNNEL":              GetAllC3DrgAttachmentsAttachmentTypeIpsecTunnel,
	"ALL":                       GetAllC3DrgAttachmentsAttachmentTypeAll,
}

var mappingGetAllC3DrgAttachmentsAttachmentTypeEnumLowerCase = map[string]GetAllC3DrgAttachmentsAttachmentTypeEnum{
	"vcn":                       GetAllC3DrgAttachmentsAttachmentTypeVcn,
	"virtual_circuit":           GetAllC3DrgAttachmentsAttachmentTypeVirtualCircuit,
	"remote_peering_connection": GetAllC3DrgAttachmentsAttachmentTypeRemotePeeringConnection,
	"ipsec_tunnel":              GetAllC3DrgAttachmentsAttachmentTypeIpsecTunnel,
	"all":                       GetAllC3DrgAttachmentsAttachmentTypeAll,
}

// GetGetAllC3DrgAttachmentsAttachmentTypeEnumValues Enumerates the set of values for GetAllC3DrgAttachmentsAttachmentTypeEnum
func GetGetAllC3DrgAttachmentsAttachmentTypeEnumValues() []GetAllC3DrgAttachmentsAttachmentTypeEnum {
	values := make([]GetAllC3DrgAttachmentsAttachmentTypeEnum, 0)
	for _, v := range mappingGetAllC3DrgAttachmentsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAllC3DrgAttachmentsAttachmentTypeEnumStringValues Enumerates the set of values in String for GetAllC3DrgAttachmentsAttachmentTypeEnum
func GetGetAllC3DrgAttachmentsAttachmentTypeEnumStringValues() []string {
	return []string{
		"VCN",
		"VIRTUAL_CIRCUIT",
		"REMOTE_PEERING_CONNECTION",
		"IPSEC_TUNNEL",
		"ALL",
	}
}

// GetMappingGetAllC3DrgAttachmentsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAllC3DrgAttachmentsAttachmentTypeEnum(val string) (GetAllC3DrgAttachmentsAttachmentTypeEnum, bool) {
	enum, ok := mappingGetAllC3DrgAttachmentsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
