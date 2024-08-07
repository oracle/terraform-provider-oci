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

// DeploymentEnvironmentSummary The metadata specific to a production or development/testing environment.
type DeploymentEnvironmentSummary struct {

	// Specifies whether the deployment is used in a production or development/testing environment.
	EnvironmentType EnvironmentTypeEnum `mandatory:"true" json:"environmentType"`

	// The deployment category defines the broad separation of the deployment type into three categories.
	// Currently the separation is 'DATA_REPLICATION', 'STREAM_ANALYTICS' and 'DATA_TRANSFORMS'.
	Category DeploymentCategoryEnum `mandatory:"true" json:"category"`

	// The minimum CPU core count.
	MinCpuCoreCount *int `mandatory:"true" json:"minCpuCoreCount"`

	// The default CPU core count.
	DefaultCpuCoreCount *int `mandatory:"true" json:"defaultCpuCoreCount"`

	// The maximum CPU core count.
	MaxCpuCoreCount *int `mandatory:"true" json:"maxCpuCoreCount"`

	// Specifies whether the "Auto scaling" option should be enabled by default or not.
	IsAutoScalingEnabledByDefault *bool `mandatory:"true" json:"isAutoScalingEnabledByDefault"`

	// The multiplier value between CPU core count and network bandwidth.
	NetworkBandwidthPerOcpuInGbps *int `mandatory:"true" json:"networkBandwidthPerOcpuInGbps"`

	// The multiplier value between CPU core count and memory size.
	MemoryPerOcpuInGBs *int `mandatory:"true" json:"memoryPerOcpuInGBs"`

	// The multiplier value between CPU core count and storage usage limit size.
	StorageUsageLimitPerOcpuInGBs *int `mandatory:"true" json:"storageUsageLimitPerOcpuInGBs"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m DeploymentEnvironmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentEnvironmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnvironmentTypeEnum(string(m.EnvironmentType)); !ok && m.EnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentType: %s. Supported values are: %s.", m.EnvironmentType, strings.Join(GetEnvironmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetDeploymentCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
