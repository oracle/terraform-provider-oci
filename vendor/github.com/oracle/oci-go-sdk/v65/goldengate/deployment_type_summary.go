// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DeploymentTypeSummary The meta-data specific on particular deployment type represented by deploymentType field.
// Deprecated properties: `sourceTechnologies` and `targetTechnologies` are not populated. They will be removed after September 15 2025.
// The list of supported source and target technologies can be accessed using the url provided in `supportedTechnologiesUrl` property.
type DeploymentTypeSummary struct {

	// The deployment category defines the broad separation of the deployment type into three categories.
	// Currently the separation is 'DATA_REPLICATION', 'STREAM_ANALYTICS' and 'DATA_TRANSFORMS'.
	Category DeploymentCategoryEnum `mandatory:"true" json:"category"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of deployment, which can be any one of the Allowed values.
	// NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.
	//     Its use is discouraged in favor of 'DATABASE_ORACLE'.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// An array of connectionTypes.
	ConnectionTypes []ConnectionTypeEnum `mandatory:"false" json:"connectionTypes,omitempty"`

	// List of the supported technologies generally.  The value is a freeform text string generally consisting
	// of a description of the technology and optionally the speific version(s) support.  For example,
	// [ "Oracle Database 19c", "Oracle Exadata", "OCI Streaming" ]
	SourceTechnologies []string `mandatory:"false" json:"sourceTechnologies"`

	// List of the supported technologies generally.  The value is a freeform text string generally consisting
	// of a description of the technology and optionally the speific version(s) support.  For example,
	// [ "Oracle Database 19c", "Oracle Exadata", "OCI Streaming" ]
	TargetTechnologies []string `mandatory:"false" json:"targetTechnologies"`

	// Version of OGG
	OggVersion *string `mandatory:"false" json:"oggVersion"`

	// The URL to the webpage listing the supported technologies.
	SupportedTechnologiesUrl *string `mandatory:"false" json:"supportedTechnologiesUrl"`

	// The default admin username used by deployment.
	DefaultUsername *string `mandatory:"false" json:"defaultUsername"`

	// Specifies supported capabilities or features by a deployment type .
	SupportedCapabilities []SupportedCapabilitiesEnum `mandatory:"false" json:"supportedCapabilities,omitempty"`
}

func (m DeploymentTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetDeploymentCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}

	for _, val := range m.ConnectionTypes {
		if _, ok := GetMappingConnectionTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionTypes: %s. Supported values are: %s.", val, strings.Join(GetConnectionTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range m.SupportedCapabilities {
		if _, ok := GetMappingSupportedCapabilitiesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedCapabilities: %s. Supported values are: %s.", val, strings.Join(GetSupportedCapabilitiesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
