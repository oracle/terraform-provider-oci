// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContainerImageSignaturesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListContainerImageSignaturesSigningAlgorithmEnum(string(request.SigningAlgorithm)); !ok && request.SigningAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SigningAlgorithm: %s. Supported values are: %s.", request.SigningAlgorithm, strings.Join(GetListContainerImageSignaturesSigningAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainerImageSignaturesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListContainerImageSignaturesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContainerImageSignaturesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListContainerImageSignaturesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListContainerImageSignaturesSigningAlgorithmEnum = map[string]ListContainerImageSignaturesSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": ListContainerImageSignaturesSigningAlgorithm512RsaPkcsPss,
}

// GetListContainerImageSignaturesSigningAlgorithmEnumValues Enumerates the set of values for ListContainerImageSignaturesSigningAlgorithmEnum
func GetListContainerImageSignaturesSigningAlgorithmEnumValues() []ListContainerImageSignaturesSigningAlgorithmEnum {
	values := make([]ListContainerImageSignaturesSigningAlgorithmEnum, 0)
	for _, v := range mappingListContainerImageSignaturesSigningAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerImageSignaturesSigningAlgorithmEnumStringValues Enumerates the set of values in String for ListContainerImageSignaturesSigningAlgorithmEnum
func GetListContainerImageSignaturesSigningAlgorithmEnumStringValues() []string {
	return []string{
		"SHA_224_RSA_PKCS_PSS",
		"SHA_256_RSA_PKCS_PSS",
		"SHA_384_RSA_PKCS_PSS",
		"SHA_512_RSA_PKCS_PSS",
	}
}

// GetMappingListContainerImageSignaturesSigningAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerImageSignaturesSigningAlgorithmEnum(val string) (ListContainerImageSignaturesSigningAlgorithmEnum, bool) {
	mappingListContainerImageSignaturesSigningAlgorithmEnumIgnoreCase := make(map[string]ListContainerImageSignaturesSigningAlgorithmEnum)
	for k, v := range mappingListContainerImageSignaturesSigningAlgorithmEnum {
		mappingListContainerImageSignaturesSigningAlgorithmEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListContainerImageSignaturesSigningAlgorithmEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainerImageSignaturesSortByEnum Enum with underlying type: string
type ListContainerImageSignaturesSortByEnum string

// Set of constants representing the allowable values for ListContainerImageSignaturesSortByEnum
const (
	ListContainerImageSignaturesSortByTimecreated ListContainerImageSignaturesSortByEnum = "TIMECREATED"
	ListContainerImageSignaturesSortByDisplayname ListContainerImageSignaturesSortByEnum = "DISPLAYNAME"
)

var mappingListContainerImageSignaturesSortByEnum = map[string]ListContainerImageSignaturesSortByEnum{
	"TIMECREATED": ListContainerImageSignaturesSortByTimecreated,
	"DISPLAYNAME": ListContainerImageSignaturesSortByDisplayname,
}

// GetListContainerImageSignaturesSortByEnumValues Enumerates the set of values for ListContainerImageSignaturesSortByEnum
func GetListContainerImageSignaturesSortByEnumValues() []ListContainerImageSignaturesSortByEnum {
	values := make([]ListContainerImageSignaturesSortByEnum, 0)
	for _, v := range mappingListContainerImageSignaturesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerImageSignaturesSortByEnumStringValues Enumerates the set of values in String for ListContainerImageSignaturesSortByEnum
func GetListContainerImageSignaturesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListContainerImageSignaturesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerImageSignaturesSortByEnum(val string) (ListContainerImageSignaturesSortByEnum, bool) {
	mappingListContainerImageSignaturesSortByEnumIgnoreCase := make(map[string]ListContainerImageSignaturesSortByEnum)
	for k, v := range mappingListContainerImageSignaturesSortByEnum {
		mappingListContainerImageSignaturesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListContainerImageSignaturesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListContainerImageSignaturesSortOrderEnum Enum with underlying type: string
type ListContainerImageSignaturesSortOrderEnum string

// Set of constants representing the allowable values for ListContainerImageSignaturesSortOrderEnum
const (
	ListContainerImageSignaturesSortOrderAsc  ListContainerImageSignaturesSortOrderEnum = "ASC"
	ListContainerImageSignaturesSortOrderDesc ListContainerImageSignaturesSortOrderEnum = "DESC"
)

var mappingListContainerImageSignaturesSortOrderEnum = map[string]ListContainerImageSignaturesSortOrderEnum{
	"ASC":  ListContainerImageSignaturesSortOrderAsc,
	"DESC": ListContainerImageSignaturesSortOrderDesc,
}

// GetListContainerImageSignaturesSortOrderEnumValues Enumerates the set of values for ListContainerImageSignaturesSortOrderEnum
func GetListContainerImageSignaturesSortOrderEnumValues() []ListContainerImageSignaturesSortOrderEnum {
	values := make([]ListContainerImageSignaturesSortOrderEnum, 0)
	for _, v := range mappingListContainerImageSignaturesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListContainerImageSignaturesSortOrderEnumStringValues Enumerates the set of values in String for ListContainerImageSignaturesSortOrderEnum
func GetListContainerImageSignaturesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListContainerImageSignaturesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContainerImageSignaturesSortOrderEnum(val string) (ListContainerImageSignaturesSortOrderEnum, bool) {
	mappingListContainerImageSignaturesSortOrderEnumIgnoreCase := make(map[string]ListContainerImageSignaturesSortOrderEnum)
	for k, v := range mappingListContainerImageSignaturesSortOrderEnum {
		mappingListContainerImageSignaturesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListContainerImageSignaturesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
