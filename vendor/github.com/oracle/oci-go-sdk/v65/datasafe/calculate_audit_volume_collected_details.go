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

// CalculateAuditVolumeCollectedDetails The details for calculating audit data volume collected by data safe.
type CalculateAuditVolumeCollectedDetails struct {

	// The date from which the audit volume collected by data safe has to be calculated, in the format defined by RFC3339.
	TimeFromMonth *common.SDKTime `mandatory:"true" json:"timeFromMonth"`

	// The date from which the audit volume collected by data safe has to be calculated, in the format defined by RFC3339. If not specified, this will default to the current date.
	TimeToMonth *common.SDKTime `mandatory:"false" json:"timeToMonth"`
}

func (m CalculateAuditVolumeCollectedDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CalculateAuditVolumeCollectedDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
