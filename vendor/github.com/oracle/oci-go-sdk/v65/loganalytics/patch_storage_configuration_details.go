// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchStorageConfigurationDetails Details for partially updating storage configuration settings. Only specified properties are updated.
type PatchStorageConfigurationDetails struct {

	// Array of archiving configurations by data type to update when specified.
	ArchivingConfigurations []ArchivingConfigurationByDataType `mandatory:"false" json:"archivingConfigurations"`

	// Array of archiving configurations for specific log indexes to update when specified.
	ArchivingConfigurationsLogIndexOverrides []ArchivingConfigurationsLogIndexOverride `mandatory:"false" json:"archivingConfigurationsLogIndexOverrides"`

	// Array of log index configurations by data type to update when specified.
	LogIndexConfigurations []LogIndexConfigurationSummary `mandatory:"false" json:"logIndexConfigurations"`
}

func (m PatchStorageConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchStorageConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
