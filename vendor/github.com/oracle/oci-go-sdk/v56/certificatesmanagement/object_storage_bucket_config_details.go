// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ObjectStorageBucketConfigDetails The details of the Object Storage bucket configured to store the certificate revocation list (CRL).
type ObjectStorageBucketConfigDetails struct {

	// The name of the bucket where the CRL is stored.
	ObjectStorageBucketName *string `mandatory:"true" json:"objectStorageBucketName"`

	// The object name in the bucket where the CRL is stored, expressed using a format where the version number of the issuing CA is inserted as part of the Object Storage object name wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2.
	ObjectStorageObjectNameFormat *string `mandatory:"true" json:"objectStorageObjectNameFormat"`

	// The tenancy of the bucket where the CRL is stored.
	ObjectStorageNamespace *string `mandatory:"false" json:"objectStorageNamespace"`
}

func (m ObjectStorageBucketConfigDetails) String() string {
	return common.PointerString(m)
}
