// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummaryMetricItem A single metric card (e.g., 'CPU Usage').
type SummaryMetricItem struct {

	// Stable ID for the metric (e.g., 'cpu_util').
	Key *string `mandatory:"true" json:"key"`

	// Raw numeric value.
	Value *float64 `mandatory:"true" json:"value"`

	// User-facing name (e.g., 'CPU Utilization').
	Label *string `mandatory:"false" json:"label"`

	// Unit suffix (e.g., 'GB', 'ms').
	Unit *string `mandatory:"false" json:"unit"`

	// Visual hint: 'OK', 'WARNING', 'CRITICAL', 'UNKNOWN'.
	Status *string `mandatory:"false" json:"status"`
}

func (m SummaryMetricItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummaryMetricItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
