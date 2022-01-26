// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListContainerImageSignaturesRequest wrapper for the ListContainerImageSignatures operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListContainerImageSignatures.go.html to see an example of how to use ListContainerImageSignaturesRequest.
type ListContainerImageSignaturesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// inspected depending on the the setting of `accessLevel`.
	// Default is false. Can only be set to true when calling the API
	// on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return a container image summary only for the specified container image OCID.
	ImageId *string `mandatory:"false" contributesTo:"query" name:"imageId"`

	// A filter to return container images only for the specified container repository OCID.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter to return container images or container image signatures that match the repository name.
	// Example: `foo` or `foo*`
	RepositoryName *string `mandatory:"false" contributesTo:"query" name:"repositoryName"`

	// The digest of the container image.
	// Example: `sha256:e7d38b3517548a1c71e41bffe9c8ae6d6d29546ce46bf62159837aad072c90aa`
	ImageDigest *string `mandatory:"false" contributesTo:"query" name:"imageDigest"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.
	// Example: `ocid1.keyversion.oc1..exampleuniqueID`
	KmsKeyId *string `mandatory:"false" contributesTo:"query" name:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.
	// Example: `ocid1.keyversion.oc1..exampleuniqueID`
	KmsKeyVersionId *string `mandatory:"false" contributesTo:"query" name:"kmsKeyVersionId"`

	// The algorithm to be used for signing. These are the only supported signing algorithms for container images.
	SigningAlgorithm ListContainerImageSignaturesSigningAlgorithmEnum `mandatory:"false" contributesTo:"query" name:"signingAlgorithm" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListContainerImageSignaturesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListContainerImageSignaturesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContainerImageSignaturesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContainerImageSignaturesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContainerImageSignaturesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContainerImageSignaturesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListContainerImageSignaturesResponse wrapper for the ListContainerImageSignatures operation
type ListContainerImageSignaturesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContainerImageSignatureCollection instances
	ContainerImageSignatureCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListContainerImageSignaturesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContainerImageSignaturesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContainerImageSignaturesSigningAlgorithmEnum Enum with underlying type: string
type ListContainerImageSignaturesSigningAlgorithmEnum string

// Set of constants representing the allowable values for ListContainerImageSignaturesSigningAlgorithmEnum
const (
	ListContainerImageSignaturesSigningAlgorithm224RsaPkcsPss ListContainerImageSignaturesSigningAlgorithmEnum = "SHA_224_RSA_PKCS_PSS"
	ListContainerImageSignaturesSigningAlgorithm256RsaPkcsPss ListContainerImageSignaturesSigningAlgorithmEnum = "SHA_256_RSA_PKCS_PSS"
	ListContainerImageSignaturesSigningAlgorithm384RsaPkcsPss ListContainerImageSignaturesSigningAlgorithmEnum = "SHA_384_RSA_PKCS_PSS"
	ListContainerImageSignaturesSigningAlgorithm512RsaPkcsPss ListContainerImageSignaturesSigningAlgorithmEnum = "SHA_512_RSA_PKCS_PSS"
)

var mappingListContainerImageSignaturesSigningAlgorithm = map[string]ListContainerImageSignaturesSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm512RsaPkcsPss,
}

// GetListContainerImageSignaturesSigningAlgorithmEnumValues Enumerates the set of values for ListContainerImageSignaturesSigningAlgorithmEnum
func GetListContainerImageSignaturesSigningAlgorithmEnumValues() []ListContainerImageSignaturesSigningAlgorithmEnum {
	values := make([]ListContainerImageSignaturesSigningAlgorithmEnum, 0)
	for _, v := range mappingListContainerImageSignaturesSigningAlgorithm {
		values = append(values, v)
	}
	return values
}

// ListContainerImageSignaturesSortByEnum Enum with underlying type: string
type ListContainerImageSignaturesSortByEnum string

// Set of constants representing the allowable values for ListContainerImageSignaturesSortByEnum
const (
	ListContainerImageSignaturesSortByTimecreated ListContainerImageSignaturesSortByEnum = "TIMECREATED"
	ListContainerImageSignaturesSortByDisplayname ListContainerImageSignaturesSortByEnum = "DISPLAYNAME"
)

var mappingListContainerImageSignaturesSortBy = map[string]ListContainerImageSignaturesSortByEnum{
	"TIMECREATED": ListContainerImageSignaturesSortByTimecreated,
	"DISPLAYNAME": ListContainerImageSignaturesSortByDisplayname,
}

// GetListContainerImageSignaturesSortByEnumValues Enumerates the set of values for ListContainerImageSignaturesSortByEnum
func GetListContainerImageSignaturesSortByEnumValues() []ListContainerImageSignaturesSortByEnum {
	values := make([]ListContainerImageSignaturesSortByEnum, 0)
	for _, v := range mappingListContainerImageSignaturesSortBy {
		values = append(values, v)
	}
	return values
}

// ListContainerImageSignaturesSortOrderEnum Enum with underlying type: string
type ListContainerImageSignaturesSortOrderEnum string

// Set of constants representing the allowable values for ListContainerImageSignaturesSortOrderEnum
const (
	ListContainerImageSignaturesSortOrderAsc  ListContainerImageSignaturesSortOrderEnum = "ASC"
	ListContainerImageSignaturesSortOrderDesc ListContainerImageSignaturesSortOrderEnum = "DESC"
)

var mappingListContainerImageSignaturesSortOrder = map[string]ListContainerImageSignaturesSortOrderEnum{
	"ASC":  ListContainerImageSignaturesSortOrderAsc,
	"DESC": ListContainerImageSignaturesSortOrderDesc,
}

// GetListContainerImageSignaturesSortOrderEnumValues Enumerates the set of values for ListContainerImageSignaturesSortOrderEnum
func GetListContainerImageSignaturesSortOrderEnumValues() []ListContainerImageSignaturesSortOrderEnum {
	values := make([]ListContainerImageSignaturesSortOrderEnum, 0)
	for _, v := range mappingListContainerImageSignaturesSortOrder {
		values = append(values, v)
	}
	return values
}
