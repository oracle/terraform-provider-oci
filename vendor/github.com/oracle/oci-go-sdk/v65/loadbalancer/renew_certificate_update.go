// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RenewCertificateUpdate The representation of RenewCertificateUpdate
type RenewCertificateUpdate struct {

	// The OCID of the certificate in OCI Certificates service.
	// Example: ocid1.certificate.oc1.phx.<unique_ID>
	CertificateId *string `mandatory:"true" json:"certificateId"`

	// The certificate version observed by GAXCP and sent to LBCP for optimization/skipping.
	// The dataplane compares this value against the stored currentVersion for this load balancer and certificate.
	// If currentVersion >= expectedVersion, the dataplane skips re-download; otherwise it pulls and applies the latest from OCI Certificates.
	ExpectedVersion *int `mandatory:"true" json:"expectedVersion"`
}

func (m RenewCertificateUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RenewCertificateUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
