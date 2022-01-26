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

// EnableDataSafeConfigurationDetails The details used to enable Data Safe in the tenancy and region.
type EnableDataSafeConfigurationDetails struct {

	// Indicates if Data Safe is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`
}

func (m EnableDataSafeConfigurationDetails) String() string {
	return common.PointerString(m)
}
