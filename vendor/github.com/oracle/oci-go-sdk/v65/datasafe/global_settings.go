// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GlobalSettings Details of the tenancy level global settings in Data Safe.
type GlobalSettings struct {

	// The paid usage option chosen by the customer admin.
	IsPaidUsage *bool `mandatory:"false" json:"isPaidUsage"`

	// The online retention period in months.
	OnlineRetentionPeriod *int `mandatory:"false" json:"onlineRetentionPeriod"`

	// The offline retention period in months.
	OfflineRetentionPeriod *int `mandatory:"false" json:"offlineRetentionPeriod"`
}

func (m GlobalSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GlobalSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
