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

// DeploymentTypeSummary The meta-data specific on particular deployment type represented by deploymentType field.
type DeploymentTypeSummary struct {

	// The deployment category defines the broad separation of the deployment type into three categories.
	// Currently the separation is 'DATA_REPLICATION', 'STREAM_ANALYTICS' and 'DATA_TRANSFORMS'.
	Category DeploymentTypeSummaryCategoryEnum `mandatory:"true" json:"category"`

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
}

func (m DeploymentTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentTypeSummaryCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetDeploymentTypeSummaryCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}

	for _, val := range m.ConnectionTypes {
		if _, ok := GetMappingConnectionTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionTypes: %s. Supported values are: %s.", val, strings.Join(GetConnectionTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeploymentTypeSummaryCategoryEnum Enum with underlying type: string
type DeploymentTypeSummaryCategoryEnum string

// Set of constants representing the allowable values for DeploymentTypeSummaryCategoryEnum
const (
	DeploymentTypeSummaryCategoryDataReplication DeploymentTypeSummaryCategoryEnum = "DATA_REPLICATION"
	DeploymentTypeSummaryCategoryStreamAnalytics DeploymentTypeSummaryCategoryEnum = "STREAM_ANALYTICS"
	DeploymentTypeSummaryCategoryDataTransforms  DeploymentTypeSummaryCategoryEnum = "DATA_TRANSFORMS"
)

var mappingDeploymentTypeSummaryCategoryEnum = map[string]DeploymentTypeSummaryCategoryEnum{
	"DATA_REPLICATION": DeploymentTypeSummaryCategoryDataReplication,
	"STREAM_ANALYTICS": DeploymentTypeSummaryCategoryStreamAnalytics,
	"DATA_TRANSFORMS":  DeploymentTypeSummaryCategoryDataTransforms,
}

var mappingDeploymentTypeSummaryCategoryEnumLowerCase = map[string]DeploymentTypeSummaryCategoryEnum{
	"data_replication": DeploymentTypeSummaryCategoryDataReplication,
	"stream_analytics": DeploymentTypeSummaryCategoryStreamAnalytics,
	"data_transforms":  DeploymentTypeSummaryCategoryDataTransforms,
}

// GetDeploymentTypeSummaryCategoryEnumValues Enumerates the set of values for DeploymentTypeSummaryCategoryEnum
func GetDeploymentTypeSummaryCategoryEnumValues() []DeploymentTypeSummaryCategoryEnum {
	values := make([]DeploymentTypeSummaryCategoryEnum, 0)
	for _, v := range mappingDeploymentTypeSummaryCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentTypeSummaryCategoryEnumStringValues Enumerates the set of values in String for DeploymentTypeSummaryCategoryEnum
func GetDeploymentTypeSummaryCategoryEnumStringValues() []string {
	return []string{
		"DATA_REPLICATION",
		"STREAM_ANALYTICS",
		"DATA_TRANSFORMS",
	}
}

// GetMappingDeploymentTypeSummaryCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentTypeSummaryCategoryEnum(val string) (DeploymentTypeSummaryCategoryEnum, bool) {
	enum, ok := mappingDeploymentTypeSummaryCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
