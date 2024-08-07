// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportSetting Export settings of a fleet.
type ExportSetting struct {

	// ExportSetting flag to store enabled or disabled status.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The internal identifier of the export setting.
	ExportSettingKey *string `mandatory:"false" json:"exportSettingKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the fleet.
	FleetId *string `mandatory:"false" json:"fleetId"`

	// The duration of data to be exported for fleets.
	ExportDuration ExportDurationEnum `mandatory:"false" json:"exportDuration,omitempty"`

	// Resource to export data associated from the fleets.
	ExportResources ExportResourcesEnum `mandatory:"false" json:"exportResources,omitempty"`

	// Acknowledgement for cross region target bucket configuration.
	IsCrossRegionAcknowledged *bool `mandatory:"false" json:"isCrossRegionAcknowledged"`

	// The name of the bucket where data will be exported.
	TargetBucketName *string `mandatory:"false" json:"targetBucketName"`

	// The namespace of the bucket where data will be exported.
	TargetBucketNamespace *string `mandatory:"false" json:"targetBucketNamespace"`

	// The id of the region of the target bucket.
	TargetBucketRegion *string `mandatory:"false" json:"targetBucketRegion"`

	// Schedule at which data will be exported.
	ExportFrequency ExportFrequencyEnum `mandatory:"false" json:"exportFrequency,omitempty"`

	// The creation date and time of the export setting (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The update date and time of the export setting (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`
}

func (m ExportSetting) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportSetting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExportDurationEnum(string(m.ExportDuration)); !ok && m.ExportDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExportDuration: %s. Supported values are: %s.", m.ExportDuration, strings.Join(GetExportDurationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExportResourcesEnum(string(m.ExportResources)); !ok && m.ExportResources != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExportResources: %s. Supported values are: %s.", m.ExportResources, strings.Join(GetExportResourcesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExportFrequencyEnum(string(m.ExportFrequency)); !ok && m.ExportFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExportFrequency: %s. Supported values are: %s.", m.ExportFrequency, strings.Join(GetExportFrequencyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
