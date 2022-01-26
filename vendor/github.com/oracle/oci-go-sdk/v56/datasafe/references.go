// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// References References to the sections of STIG, CIS, and/or GDPR relevant to the current finding.
type References struct {

	// Relevant section from STIG.
	Stig *string `mandatory:"false" json:"stig"`

	// Relevant section from CIS.
	Cis *string `mandatory:"false" json:"cis"`

	// Relevant section from GDPR.
	Gdpr *string `mandatory:"false" json:"gdpr"`
}

func (m References) String() string {
	return common.PointerString(m)
}
