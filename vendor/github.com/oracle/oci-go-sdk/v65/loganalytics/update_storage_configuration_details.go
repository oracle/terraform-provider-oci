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

// UpdateStorageConfigurationDetails Details for replacing storage configuration settings.
type UpdateStorageConfigurationDetails struct {

	// Array of archiving configurations by data type.
	ArchivingConfigurations []ArchivingConfigurationByDataType `mandatory:"true" json:"archivingConfigurations"`

	// Array of archiving configurations for specific log indexes.
	ArchivingConfigurationsLogIndexOverrides []ArchivingConfigurationsLogIndexOverride `mandatory:"true" json:"archivingConfigurationsLogIndexOverrides"`

	// Array of log index configurations by data type.
	LogIndexConfigurations []LogIndexConfigurationSummary `mandatory:"true" json:"logIndexConfigurations"`
}

func (m UpdateStorageConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateStorageConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
