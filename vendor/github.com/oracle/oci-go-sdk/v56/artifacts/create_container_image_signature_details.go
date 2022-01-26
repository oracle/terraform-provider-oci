// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
}

func (m CreateContainerImageSignatureDetails) String() string {
	return common.PointerString(m)
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

var mappingCreateContainerImageSignatureDetailsSigningAlgorithm = map[string]CreateContainerImageSignatureDetailsSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": CreateContainerImageSignatureDetailsSigningAlgorithm512RsaPkcsPss,
}

// GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumValues Enumerates the set of values for CreateContainerImageSignatureDetailsSigningAlgorithmEnum
func GetCreateContainerImageSignatureDetailsSigningAlgorithmEnumValues() []CreateContainerImageSignatureDetailsSigningAlgorithmEnum {
	values := make([]CreateContainerImageSignatureDetailsSigningAlgorithmEnum, 0)
	for _, v := range mappingCreateContainerImageSignatureDetailsSigningAlgorithm {
		values = append(values, v)
	}
	return values
}
