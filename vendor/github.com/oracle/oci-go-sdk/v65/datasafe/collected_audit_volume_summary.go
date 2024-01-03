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

// CollectedAuditVolumeSummary The volume data point for audit data collected by datasafe.
type CollectedAuditVolumeSummary struct {

	// The OCID of the audit profile resource.
	AuditProfileId *string `mandatory:"true" json:"auditProfileId"`

	// Represents the month under consideration in which the aggregated audit data volume collected by Data Safe is displayed.
	// This field will be the UTC start of the day of the first day of the month for which the aggregate count corresponds to, in the format defined by RFC3339..
	// For instance, the value of 01-01-2021T00:00:00Z represents Jan 2021.
	MonthInConsideration *common.SDKTime `mandatory:"true" json:"monthInConsideration"`

	// The audit data volume collected by Data Safe and is available online in repository.
	OnlineVolume *int64 `mandatory:"true" json:"onlineVolume"`

	// The audit data volume collected by Data Safe and is available in archive storage.
	ArchivedVolume *int64 `mandatory:"true" json:"archivedVolume"`
}

func (m CollectedAuditVolumeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CollectedAuditVolumeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
