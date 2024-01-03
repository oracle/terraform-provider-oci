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

// DeploymentVersionSummary The summary data of a specific deployment version.
type DeploymentVersionSummary struct {

	// Version of OGG
	OggVersion *string `mandatory:"true" json:"oggVersion"`

	// The type of deployment, which can be any one of the Allowed values.
	// NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.
	//     Its use is discouraged in favor of 'DATABASE_ORACLE'.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// The time the resource was released. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	// The type of release.
	ReleaseType ReleaseTypeEnum `mandatory:"false" json:"releaseType,omitempty"`

	// Indicates if OGG release contains security fix.
	IsSecurityFix *bool `mandatory:"false" json:"isSecurityFix"`

	// The time until OGG version is supported. After this date has passed OGG version will not be available anymore. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeSupportedUntil *common.SDKTime `mandatory:"false" json:"timeSupportedUntil"`
}

func (m DeploymentVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReleaseTypeEnum(string(m.ReleaseType)); !ok && m.ReleaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReleaseType: %s. Supported values are: %s.", m.ReleaseType, strings.Join(GetReleaseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
