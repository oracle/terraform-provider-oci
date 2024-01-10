// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Artifacts and Container Images API
//
// API covering the Artifacts and Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as generic artifacts and container images.
//

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateContainerImageSignatureDetails Upload container image signature request details.
type CreateContainerImageSignatureDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the container repository exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the container image.
	// Example: `ocid1.containerimage.oc1..exampleuniqueID`
	ImageId *string `mandatory:"true" json:"imageId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the kmsKeyId used to sign the container image.
	// Example: `ocid1.key.oc1..exampleuniqueID`
	KmsKeyId *string `mandatory:"true" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.
	// Example: `ocid1.keyversion.oc1..exampleuniqueID`
	KmsKeyVersionId *string `mandatory:"true" json:"kmsKeyVersionId"`

	// The base64 encoded signature payload that was signed.
	Message *string `mandatory:"true" json:"message"`

	// The signature of the message field using the kmsKeyId, the kmsKeyVersionId, and the signingAlgorithm.
	Signature *string `mandatory:"true" json:"signature"`

	// The algorithm to be used for signing. These are the only supported signing algorithms for container images.
	SigningAlgorithm CreateContainerImageSignatureDetailsSigningAlgorithmEnum `mandatory:"true" json:"signingAlgorithm"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateContainerImageSignatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerImageSignatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateContainerImageSignatureDetailsSigningAlgorithmEnum(string(m.SigningAlgorithm)); !ok && m.SigningAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SigningAlgorithm: %s. Supported values are: %s.", m.SigningAlgorithm, strings.Join(GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateContainerImageSignatureDetailsSigningAlgorithmEnum Enum with underlying type: string
type CreateContainerImageSignatureDetailsSigningAlgorithmEnum string

// Set of constants representing the allowable values for CreateContainerImageSignatureDetailsSigningAlgorithmEnum
const (
	CreateContainerImageSignatureDetailsSigningAlgorithm224RsaPkcsPss CreateContainerImageSignatureDetailsSigningAlgorithmEnum = "SHA_224_RSA_PKCS_PSS"
	CreateContainerImageSignatureDetailsSigningAlgorithm256RsaPkcsPss CreateContainerImageSignatureDetailsSigningAlgorithmEnum = "SHA_256_RSA_PKCS_PSS"
	CreateContainerImageSignatureDetailsSigningAlgorithm384RsaPkcsPss CreateContainerImageSignatureDetailsSigningAlgorithmEnum = "SHA_384_RSA_PKCS_PSS"
	CreateContainerImageSignatureDetailsSigningAlgorithm512RsaPkcsPss CreateContainerImageSignatureDetailsSigningAlgorithmEnum = "SHA_512_RSA_PKCS_PSS"
)

var mappingCreateContainerImageSignatureDetailsSigningAlgorithmEnum = map[string]CreateContainerImageSignatureDetailsSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm512RsaPkcsPss,
}

var mappingCreateContainerImageSignatureDetailsSigningAlgorithmEnumLowerCase = map[string]CreateContainerImageSignatureDetailsSigningAlgorithmEnum{
	"sha_224_rsa_pkcs_pss": CreateContainerImageSignatureDetailsSigningAlgorithm224RsaPkcsPss,
	"sha_256_rsa_pkcs_pss": CreateContainerImageSignatureDetailsSigningAlgorithm256RsaPkcsPss,
	"sha_384_rsa_pkcs_pss": CreateContainerImageSignatureDetailsSigningAlgorithm384RsaPkcsPss,
	"sha_512_rsa_pkcs_pss": CreateContainerImageSignatureDetailsSigningAlgorithm512RsaPkcsPss,
}

// GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumValues Enumerates the set of values for CreateContainerImageSignatureDetailsSigningAlgorithmEnum
func GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumValues() []CreateContainerImageSignatureDetailsSigningAlgorithmEnum {
	values := make([]CreateContainerImageSignatureDetailsSigningAlgorithmEnum, 0)
	for _, v := range mappingCreateContainerImageSignatureDetailsSigningAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumStringValues Enumerates the set of values in String for CreateContainerImageSignatureDetailsSigningAlgorithmEnum
func GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumStringValues() []string {
	return []string{
		"SHA_224_RSA_PKCS_PSS",
		"SHA_256_RSA_PKCS_PSS",
		"SHA_384_RSA_PKCS_PSS",
		"SHA_512_RSA_PKCS_PSS",
	}
}

// GetMappingCreateContainerImageSignatureDetailsSigningAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateContainerImageSignatureDetailsSigningAlgorithmEnum(val string) (CreateContainerImageSignatureDetailsSigningAlgorithmEnum, bool) {
	enum, ok := mappingCreateContainerImageSignatureDetailsSigningAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
