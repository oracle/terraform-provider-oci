// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCloudDbSystemDiscoveryDetails The details required to create a cloud DB system discovery resource.
type CreateCloudDbSystemDiscoveryDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the cloud DB system discovery.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the cloud DB system resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent cloud DB Infrastructure. For VM Dbsystems ,
	// it will be the DBSystem Id. For ExaCS and ExaCC,  it will be the cloudVmClusterId and vmClusterId respectively.
	DbaasParentInfrastructureId *string `mandatory:"true" json:"dbaasParentInfrastructureId"`

	// The deployment type of cloud dbsystem.
	DeploymentType CloudDbSystemDeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCloudDbSystemDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCloudDbSystemDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudDbSystemDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetCloudDbSystemDeploymentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
