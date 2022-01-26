// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseRegistrationCollection A list of DatabaseRegistrations.
type DatabaseRegistrationCollection struct {

	// An array of DatabaseRegistration summaries.
	Items []DatabaseRegistrationSummary `mandatory:"true" json:"items"`
}

func (m DatabaseRegistrationCollection) String() string {
	return common.PointerString(m)
}
