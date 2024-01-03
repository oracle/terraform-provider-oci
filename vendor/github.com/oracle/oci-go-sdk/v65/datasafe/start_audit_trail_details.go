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

// StartAuditTrailDetails The details used to  start an audit trail.
type StartAuditTrailDetails struct {

	// The date from which the audit trail must start collecting data, in the format defined by RFC3339.
	AuditCollectionStartTime *common.SDKTime `mandatory:"true" json:"auditCollectionStartTime"`

	// Indicates if auto purge is enabled on the target database, which helps delete audit data in the
	// target database every seven days so that the database's audit trail does not become too large.
	IsAutoPurgeEnabled *bool `mandatory:"false" json:"isAutoPurgeEnabled"`
}

func (m StartAuditTrailDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StartAuditTrailDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
