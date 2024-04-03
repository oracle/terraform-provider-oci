// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// References References to the sections of STIG, CIS, GDPR and/or OBP relevant to the current finding.
type References struct {

	// Relevant section from STIG.
	Stig *string `mandatory:"false" json:"stig"`

	// Relevant section from CIS.
	Cis *string `mandatory:"false" json:"cis"`

	// Relevant section from GDPR.
	Gdpr *string `mandatory:"false" json:"gdpr"`

	// Relevant section from OBP.
	Obp *string `mandatory:"false" json:"obp"`
}

func (m References) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m References) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
