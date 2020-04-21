// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CertificateExtensions The representation of CertificateExtensions
type CertificateExtensions struct {

	// The certificate extension name.
	Name *string `mandatory:"false" json:"name"`

	// The critical flag of the extension. Critical extensions must be processed, non-critical extensions can be ignored.
	IsCritical *bool `mandatory:"false" json:"isCritical"`

	// The certificate extension value.
	Value *string `mandatory:"false" json:"value"`
}

func (m CertificateExtensions) String() string {
	return common.PointerString(m)
}
