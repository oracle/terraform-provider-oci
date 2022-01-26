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

// CertificateRevocationListDetails The details of the certificate revocation list (CRL).
type CertificateRevocationListDetails struct {
	ObjectStorageConfig *ObjectStorageBucketConfigDetails `mandatory:"true" json:"objectStorageConfig"`

	// Optional CRL access points, expressed using a format where the version number of the issuing CA is inserted wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2.
	CustomFormattedUrls []string `mandatory:"false" json:"customFormattedUrls"`
}

func (m CertificateRevocationListDetails) String() string {
	return common.PointerString(m)
}
