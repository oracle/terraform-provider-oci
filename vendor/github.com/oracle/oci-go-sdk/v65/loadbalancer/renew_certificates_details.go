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

// RenewCertificatesDetails The representation of RenewCertificatesDetails
type RenewCertificatesDetails struct {

	// The list of certificate updates to apply to the load balancer.
	// Each entry specifies a certificate OCID and the expected version observed by GAXCP. The service
	// creates a work request that triggers the dataplane to download and apply only those certificates
	// whose current version is less than the expected version.
	// Limits: Up to 200 certificates per load balancer in a single request.
	// If isForce=false, must contain at least one entry.
	// If isForce=true, may be omitted or empty to renew all certificates attached to the LB.
	Certificates []RenewCertificateUpdate `mandatory:"false" json:"certificates"`
}

func (m RenewCertificatesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RenewCertificatesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
