// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ChangeRetentionDetails Details for the audit retention months to be modified.
type ChangeRetentionDetails struct {

	// Indicates the number of months the audit records will be stored online in Oracle Data Safe audit repository for
	// immediate reporting and analysis. Minimum: 1; Maximum:12 months
	OnlineMonths *int `mandatory:"false" json:"onlineMonths"`

	// Indicates the number of months the audit records will be stored offline in the Data Safe audit archive.
	// Minimum: 0; Maximum: 72 months.
	// If you have a requirement to store the audit data even longer in archive, please contact the Oracle Support.
	OfflineMonths *int `mandatory:"false" json:"offlineMonths"`

	// Indicates whether audit retention settings like online and offline months is set at the
	// target level overriding the global audit retention settings.
	IsOverrideGlobalRetentionSetting *bool `mandatory:"false" json:"isOverrideGlobalRetentionSetting"`
}

func (m ChangeRetentionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeRetentionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
