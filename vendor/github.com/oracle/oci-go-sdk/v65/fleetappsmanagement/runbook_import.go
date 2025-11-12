// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunbookImport Runbook Import Status check details.
type RunbookImport struct {

	// Tracking Id.
	TrackingId *string `mandatory:"true" json:"trackingId"`

	// The OCID of the resource.
	RunbookId *string `mandatory:"true" json:"runbookId"`

	// Runbook name.
	RunbookName *string `mandatory:"true" json:"runbookName"`

	// Runbook version.
	RunbookVersion *string `mandatory:"true" json:"runbookVersion"`

	// Status.
	Status *string `mandatory:"true" json:"status"`

	// Map of details
	Details map[string]string `mandatory:"true" json:"details"`
}

func (m RunbookImport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunbookImport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
