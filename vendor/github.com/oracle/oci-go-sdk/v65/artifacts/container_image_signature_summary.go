// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Artifacts and Container Images API
//
// API covering the Artifacts and Registry (https://docs.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as generic artifacts and container images.
//

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerImageSignatureSummary Container image signature summary.
type ContainerImageSignatureSummary struct {

	// The OCID of the compartment in which the container repository exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The last 10 characters of the kmsKeyId, the last 10 characters of the kmsKeyVersionId, the signingAlgorithm, and the last 10 characters of the signatureId.
	// Example: `wrmz22sixa::qdwyc2ptun::SHA_256_RSA_PKCS_PSS::2vwmobasva`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image signature.
	// Example: `ocid1.containerimagesignature.oc1..exampleuniqueID`
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image.
	// Example: `ocid1.containerimage.oc1..exampleuniqueID`
	ImageId *string `mandatory:"true" json:"imageId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyId used to sign the container image.
	// Example: `ocid1.key.oc1..exampleuniqueID`
	KmsKeyId *string `mandatory:"true" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the kmsKeyVersionId used to sign the container image.
	// Example: `ocid1.keyversion.oc1..exampleuniqueID`
	KmsKeyVersionId *string `mandatory:"true" json:"kmsKeyVersionId"`

	// The base64 encoded signature payload that was signed.
	Message *string `mandatory:"true" json:"message"`

	// The signature of the message field using the kmsKeyId, the kmsKeyVersionId, and the signingAlgorithm.
	Signature *string `mandatory:"true" json:"signature"`

	// The algorithm to be used for signing. These are the only supported signing algorithms for container images.
	SigningAlgorithm ContainerImageSignatureSummarySigningAlgorithmEnum `mandatory:"true" json:"signingAlgorithm"`

	// An RFC 3339 timestamp indicating when the image was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the container image signature.
	LifecycleState ContainerImageSignatureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The system tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
}

func (m ContainerImageSignatureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerImageSignatureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContainerImageSignatureSummarySigningAlgorithmEnum(string(m.SigningAlgorithm)); !ok && m.SigningAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SigningAlgorithm: %s. Supported values are: %s.", m.SigningAlgorithm, strings.Join(GetContainerImageSignatureSummarySigningAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContainerImageSignatureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetContainerImageSignatureLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContainerImageSignatureSummarySigningAlgorithmEnum Enum with underlying type: string
type ContainerImageSignatureSummarySigningAlgorithmEnum string

// Set of constants representing the allowable values for ContainerImageSignatureSummarySigningAlgorithmEnum
const (
	ContainerImageSignatureSummarySigningAlgorithm224RsaPkcsPss ContainerImageSignatureSummarySigningAlgorithmEnum = "SHA_224_RSA_PKCS_PSS"
	ContainerImageSignatureSummarySigningAlgorithm256RsaPkcsPss ContainerImageSignatureSummarySigningAlgorithmEnum = "SHA_256_RSA_PKCS_PSS"
	ContainerImageSignatureSummarySigningAlgorithm384RsaPkcsPss ContainerImageSignatureSummarySigningAlgorithmEnum = "SHA_384_RSA_PKCS_PSS"
	ContainerImageSignatureSummarySigningAlgorithm512RsaPkcsPss ContainerImageSignatureSummarySigningAlgorithmEnum = "SHA_512_RSA_PKCS_PSS"
)

var mappingContainerImageSignatureSummarySigningAlgorithmEnum = map[string]ContainerImageSignatureSummarySigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": ContainerImageSignatureSummarySigningAlgorithm224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": ContainerImageSignatureSummarySigningAlgorithm256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": ContainerImageSignatureSummarySigningAlgorithm384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": ContainerImageSignatureSummarySigningAlgorithm512RsaPkcsPss,
}

var mappingContainerImageSignatureSummarySigningAlgorithmEnumLowerCase = map[string]ContainerImageSignatureSummarySigningAlgorithmEnum{
	"sha_224_rsa_pkcs_pss": ContainerImageSignatureSummarySigningAlgorithm224RsaPkcsPss,
	"sha_256_rsa_pkcs_pss": ContainerImageSignatureSummarySigningAlgorithm256RsaPkcsPss,
	"sha_384_rsa_pkcs_pss": ContainerImageSignatureSummarySigningAlgorithm384RsaPkcsPss,
	"sha_512_rsa_pkcs_pss": ContainerImageSignatureSummarySigningAlgorithm512RsaPkcsPss,
}

// GetContainerImageSignatureSummarySigningAlgorithmEnumValues Enumerates the set of values for ContainerImageSignatureSummarySigningAlgorithmEnum
func GetContainerImageSignatureSummarySigningAlgorithmEnumValues() []ContainerImageSignatureSummarySigningAlgorithmEnum {
	values := make([]ContainerImageSignatureSummarySigningAlgorithmEnum, 0)
	for _, v := range mappingContainerImageSignatureSummarySigningAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerImageSignatureSummarySigningAlgorithmEnumStringValues Enumerates the set of values in String for ContainerImageSignatureSummarySigningAlgorithmEnum
func GetContainerImageSignatureSummarySigningAlgorithmEnumStringValues() []string {
	return []string{
		"SHA_224_RSA_PKCS_PSS",
		"SHA_256_RSA_PKCS_PSS",
		"SHA_384_RSA_PKCS_PSS",
		"SHA_512_RSA_PKCS_PSS",
	}
}

// GetMappingContainerImageSignatureSummarySigningAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerImageSignatureSummarySigningAlgorithmEnum(val string) (ContainerImageSignatureSummarySigningAlgorithmEnum, bool) {
	enum, ok := mappingContainerImageSignatureSummarySigningAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
