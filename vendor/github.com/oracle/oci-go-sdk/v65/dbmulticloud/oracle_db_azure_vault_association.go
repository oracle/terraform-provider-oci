// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbAzureVaultAssociation Oracle DB Azure Vault Association Resource Object.
type OracleDbAzureVaultAssociation struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault Association Resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains Oracle DB Azure Vault Association Resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of Oracle DB Azure Vault Association.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Connector.
	OracleDbAzureConnectorId *string `mandatory:"true" json:"oracleDbAzureConnectorId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault.
	OracleDbAzureVaultId *string `mandatory:"true" json:"oracleDbAzureVaultId"`

	// The Associated Resources are accessible or not.
	IsResourceAccessible *bool `mandatory:"false" json:"isResourceAccessible"`

	// The current lifecycle state of the Oracle DB Azure Vault Association Resource.
	LifecycleState OracleDbAzureVaultAssociationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Oracle DB Azure Vault Association was created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Oracle DB Azure Vault Association was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Description of the latest modification of the Oracle DB Azure Vault Association Resource.
	LastModification *string `mandatory:"false" json:"lastModification"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OracleDbAzureVaultAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbAzureVaultAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbAzureVaultAssociationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOracleDbAzureVaultAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbAzureVaultAssociationLifecycleStateEnum Enum with underlying type: string
type OracleDbAzureVaultAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for OracleDbAzureVaultAssociationLifecycleStateEnum
const (
	OracleDbAzureVaultAssociationLifecycleStateCreating OracleDbAzureVaultAssociationLifecycleStateEnum = "CREATING"
	OracleDbAzureVaultAssociationLifecycleStateActive   OracleDbAzureVaultAssociationLifecycleStateEnum = "ACTIVE"
	OracleDbAzureVaultAssociationLifecycleStateUpdating OracleDbAzureVaultAssociationLifecycleStateEnum = "UPDATING"
	OracleDbAzureVaultAssociationLifecycleStateDeleting OracleDbAzureVaultAssociationLifecycleStateEnum = "DELETING"
	OracleDbAzureVaultAssociationLifecycleStateDeleted  OracleDbAzureVaultAssociationLifecycleStateEnum = "DELETED"
	OracleDbAzureVaultAssociationLifecycleStateFailed   OracleDbAzureVaultAssociationLifecycleStateEnum = "FAILED"
)

var mappingOracleDbAzureVaultAssociationLifecycleStateEnum = map[string]OracleDbAzureVaultAssociationLifecycleStateEnum{
	"CREATING": OracleDbAzureVaultAssociationLifecycleStateCreating,
	"ACTIVE":   OracleDbAzureVaultAssociationLifecycleStateActive,
	"UPDATING": OracleDbAzureVaultAssociationLifecycleStateUpdating,
	"DELETING": OracleDbAzureVaultAssociationLifecycleStateDeleting,
	"DELETED":  OracleDbAzureVaultAssociationLifecycleStateDeleted,
	"FAILED":   OracleDbAzureVaultAssociationLifecycleStateFailed,
}

var mappingOracleDbAzureVaultAssociationLifecycleStateEnumLowerCase = map[string]OracleDbAzureVaultAssociationLifecycleStateEnum{
	"creating": OracleDbAzureVaultAssociationLifecycleStateCreating,
	"active":   OracleDbAzureVaultAssociationLifecycleStateActive,
	"updating": OracleDbAzureVaultAssociationLifecycleStateUpdating,
	"deleting": OracleDbAzureVaultAssociationLifecycleStateDeleting,
	"deleted":  OracleDbAzureVaultAssociationLifecycleStateDeleted,
	"failed":   OracleDbAzureVaultAssociationLifecycleStateFailed,
}

// GetOracleDbAzureVaultAssociationLifecycleStateEnumValues Enumerates the set of values for OracleDbAzureVaultAssociationLifecycleStateEnum
func GetOracleDbAzureVaultAssociationLifecycleStateEnumValues() []OracleDbAzureVaultAssociationLifecycleStateEnum {
	values := make([]OracleDbAzureVaultAssociationLifecycleStateEnum, 0)
	for _, v := range mappingOracleDbAzureVaultAssociationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbAzureVaultAssociationLifecycleStateEnumStringValues Enumerates the set of values in String for OracleDbAzureVaultAssociationLifecycleStateEnum
func GetOracleDbAzureVaultAssociationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOracleDbAzureVaultAssociationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbAzureVaultAssociationLifecycleStateEnum(val string) (OracleDbAzureVaultAssociationLifecycleStateEnum, bool) {
	enum, ok := mappingOracleDbAzureVaultAssociationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
