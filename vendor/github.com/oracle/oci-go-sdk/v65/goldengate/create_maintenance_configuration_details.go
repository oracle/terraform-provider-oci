// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMaintenanceConfigurationDetails Defines the maintenance configuration for create operation.
type CreateMaintenanceConfigurationDetails struct {

	// By default auto upgrade for interim releases are not enabled. If auto-upgrade is enabled for interim release,
	// you have to specify interimReleaseUpgradePeriodInDays too.
	IsInterimReleaseAutoUpgradeEnabled *bool `mandatory:"false" json:"isInterimReleaseAutoUpgradeEnabled"`

	// Defines auto upgrade period for interim releases. This period must be shorter or equal to bundle release upgrade period.
	InterimReleaseUpgradePeriodInDays *int `mandatory:"false" json:"interimReleaseUpgradePeriodInDays"`

	// Defines auto upgrade period for bundle releases. Manually configured period cannot be longer than service defined period for bundle releases.
	// This period must be shorter or equal to major release upgrade period. Not passing this field during create will equate to using the service default.
	BundleReleaseUpgradePeriodInDays *int `mandatory:"false" json:"bundleReleaseUpgradePeriodInDays"`

	// Defines auto upgrade period for major releases. Manually configured period cannot be longer than service defined period for major releases.
	// Not passing this field during create will equate to using the service default.
	MajorReleaseUpgradePeriodInDays *int `mandatory:"false" json:"majorReleaseUpgradePeriodInDays"`

	// Defines auto upgrade period for releases with security fix. Manually configured period cannot be longer than service defined period for security releases.
	// Not passing this field during create will equate to using the service default.
	SecurityPatchUpgradePeriodInDays *int `mandatory:"false" json:"securityPatchUpgradePeriodInDays"`
}

func (m CreateMaintenanceConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMaintenanceConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
