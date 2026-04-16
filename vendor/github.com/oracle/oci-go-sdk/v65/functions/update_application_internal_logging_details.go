// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateApplicationInternalLoggingDetails Request payload to update (replace) the internal-only Lumberjack logging configuration for an application.
// Availability Domains must be specified using *physical* AD identifiers.
// Physical AD name formats vary by region. Example values: `iad-ad-1`, `phx-ad-2`, `ap-hyderabad-1-ad-1`.
// Example 1: `{"compartmentId": "ocid1.tenancy.oc1..aaaa...xyz", "namespace": "functions-app-logs", "logGroup": "default"}`
// Example 2: `{"compartmentId": "ocid1.tenancy.oc1..aaaa...xyz", "namespace": "functions-app-logs", "logGroup": "default", "enabledAds": ["iad-ad-1", "iad-ad-3"], "adOverrides": {"iad-ad-3": {"namespace": "functions-app-logs2", "logGroup": "custLogGroup"}}}`
// Example 3: `{"compartmentId": "ocid1.tenancy.oc1..aaaa...xyz", "namespace": "functions-app-logs", "logGroup": "default", "enabledAds": ["iad-ad-1", "iad-ad-2", "iad-ad-3"], "adOverrides": {"iad-ad-2": {"namespace": "functions-app-logs2", "logGroup": "custLogGroup2"}, "iad-ad-3": {"namespace": "functions-app-logs3", "logGroup": "custLogGroup3"}}}`
type UpdateApplicationInternalLoggingDetails struct {

	// OCID of the compartment that owns the Lumberjack namespace and log group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Lumberjack namespace name. The namespace must exist in each enabled AD unless overridden in adOverrides.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The Lumberjack log group within the namespace.
	LogGroup *string `mandatory:"true" json:"logGroup"`

	// List of Physical Availability Domains (ADs) for which the customer namespace is registered and logs should be written.
	// If omitted, all ADs in the region are considered enabled for logging (subject to validation).
	EnabledAds []string `mandatory:"false" json:"enabledAds"`

	// Map of Availability Domain (AD) identifiers to per-AD override settings.
	// Use this when the Lumberjack `namespace` and/or `logGroup` differs for a specific AD within the region.
	// The map key is the physical AD identifier. Example values: `iad-ad-1`, `phx-ad-2`, `ap-hyderabad-1-ad-1`.
	// For any AD not present in this map, the top-level `namespace` and `logGroup` values apply.
	AdOverrides map[string]ApplicationInternalLoggingAdOverride `mandatory:"false" json:"adOverrides"`
}

func (m UpdateApplicationInternalLoggingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateApplicationInternalLoggingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
