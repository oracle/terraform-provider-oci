// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring (APM) Control Plane API
//
// Provide a set of APIs for tenant to perform operations like create, update, delete and list APM domains, and also
// work request APIs to monitor progress of these operations.
//

package apmcontrolplane

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// ChangeApmDomainCompartmentDetails The information needed for the change compartment operation.
type ChangeApmDomainCompartmentDetails struct {

	// The OCID of the destination compartment for the APM Domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeApmDomainCompartmentDetails) String() string {
	return common.PointerString(m)
}
