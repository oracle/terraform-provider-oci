// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data Plane Integration
//
// <b>Microsoft Azure:</b> <br>
// <b>Oracle Azure Connector Resource:</b>:&nbsp;&nbsp;The Oracle Azure Connector Resource is used to install the Azure Arc Server on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
//  The supported method to install the Azure Arc Server (Azure Identity) on the Exadata VM cluster:
// <ul>
//  <li>Using a Bearer Access Token</li>
// </ul>
// <b>Oracle Azure Blob Container Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Container Resource is used to capture the details of an Azure Blob Container.
// This resource can then be reused across multiple Exadata VM clusters in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D) to mount the Azure container.
// <b>Oracle Azure Blob Mount Resource:</b>&nbsp;&nbsp;The Oracle Azure Blob Mount Resource is used to mount an Azure Blob Container on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// It relies on both the Oracle Azure Connector and the Oracle Azure Blob Container Resource to perform the mount operation.
// <b>Discover Azure Vaults and Keys Resource:</b>&nbsp;&nbsp;The Discover Oracle Azure Vaults and Azure Keys Resource is used to discover Azure Vaults and the associated encryption keys available in your Azure project.
// <b>Oracle Azure Vault:</b>&nbsp;&nbsp;The Oracle Azure Vault Resource is used to manage Azure Vaults within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Oracle Azure Key:</b>&nbsp;&nbsp;Oracle Azure Key Resource is used to register and manage a Oracle Azure Key Key within Oracle Cloud Infrastructure (OCI) under an associated Azure Vault.
// <br>
// <b>Google Cloud:</b><br>
// <b>Oracle Google Cloud Connector Resource:</b>&nbsp;&nbsp;The Oracle Google Cloud Connector Resource is used to install the Google Cloud Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Discover Google Key Rings and Keys Resource:</b>&nbsp;&nbsp;The Discover Google Key Rings and Keys Resource is used to discover Google Cloud Key Rings and the associated encryption keys available in your Google Cloud project.
// <b>Google Key Rings Resource:</b>&nbsp;&nbsp;The Google Key Rings Resource is used to register and manage Google Cloud Key Rings within Oracle Cloud Infrastructure (OCI) for use with services such as Oracle Exadata Database Service on Dedicated Infrastructure.
// <b>Google Key Resource:</b>&nbsp;&nbsp;The Google Key Resource is used to register and manage a Google Cloud Key within Oracle Cloud Infrastructure (OCI) under an associated Google Key Ring.
// <br>
// <b>AWS</b>:<br>
// <b>Oracle AWS Connector Resource:</b>&nbsp;&nbsp;The Oracle AWS Connector Resource is used to install the AWS Identity Connector on an Exadata VM cluster in Oracle Exadata Database Service on Dedicated Infrastructure (ExaDB-D).
// <b>Oracle AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbAwsKeyPool Oracle DB AWS Key Pool is a resource that represents a collection of AWS KMS keys.
// These keys are used for Oracle Database Transparent Data Encryption (TDE) operations.
type OracleDbAwsKeyPool struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB AWS Key Pool resource.
	Id *string `mandatory:"true" json:"id"`

	// The Compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that contains this Oracle DB AWS Pool resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Oracle DB AWS key pool resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Indicates whether AWS KMS keys in this key pool can be reused.
	IsReuseAllowed *bool `mandatory:"false" json:"isReuseAllowed"`

	// Indicates whether this key pool accepts only multi-region AWS KMS keys. When false, the key pool accepts only single-region AWS KMS keys.
	IsMultiRegionKeyPool *bool `mandatory:"false" json:"isMultiRegionKeyPool"`

	// The lifecycle state of the Oracle DB AWS Key Pool resource.
	LifecycleState OracleDbAwsKeyPoolLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	ReplicationMetadata *ReplicationMetadata `mandatory:"false" json:"replicationMetadata"`

	// Time when the DB AWS Key Pool resource was created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the DB AWS Key Pool resource was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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

func (m OracleDbAwsKeyPool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbAwsKeyPool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbAwsKeyPoolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOracleDbAwsKeyPoolLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbAwsKeyPoolLifecycleStateEnum Enum with underlying type: string
type OracleDbAwsKeyPoolLifecycleStateEnum string

// Set of constants representing the allowable values for OracleDbAwsKeyPoolLifecycleStateEnum
const (
	OracleDbAwsKeyPoolLifecycleStateCreating OracleDbAwsKeyPoolLifecycleStateEnum = "CREATING"
	OracleDbAwsKeyPoolLifecycleStateActive   OracleDbAwsKeyPoolLifecycleStateEnum = "ACTIVE"
	OracleDbAwsKeyPoolLifecycleStateUpdating OracleDbAwsKeyPoolLifecycleStateEnum = "UPDATING"
	OracleDbAwsKeyPoolLifecycleStateDeleting OracleDbAwsKeyPoolLifecycleStateEnum = "DELETING"
	OracleDbAwsKeyPoolLifecycleStateDeleted  OracleDbAwsKeyPoolLifecycleStateEnum = "DELETED"
	OracleDbAwsKeyPoolLifecycleStateFailed   OracleDbAwsKeyPoolLifecycleStateEnum = "FAILED"
)

var mappingOracleDbAwsKeyPoolLifecycleStateEnum = map[string]OracleDbAwsKeyPoolLifecycleStateEnum{
	"CREATING": OracleDbAwsKeyPoolLifecycleStateCreating,
	"ACTIVE":   OracleDbAwsKeyPoolLifecycleStateActive,
	"UPDATING": OracleDbAwsKeyPoolLifecycleStateUpdating,
	"DELETING": OracleDbAwsKeyPoolLifecycleStateDeleting,
	"DELETED":  OracleDbAwsKeyPoolLifecycleStateDeleted,
	"FAILED":   OracleDbAwsKeyPoolLifecycleStateFailed,
}

var mappingOracleDbAwsKeyPoolLifecycleStateEnumLowerCase = map[string]OracleDbAwsKeyPoolLifecycleStateEnum{
	"creating": OracleDbAwsKeyPoolLifecycleStateCreating,
	"active":   OracleDbAwsKeyPoolLifecycleStateActive,
	"updating": OracleDbAwsKeyPoolLifecycleStateUpdating,
	"deleting": OracleDbAwsKeyPoolLifecycleStateDeleting,
	"deleted":  OracleDbAwsKeyPoolLifecycleStateDeleted,
	"failed":   OracleDbAwsKeyPoolLifecycleStateFailed,
}

// GetOracleDbAwsKeyPoolLifecycleStateEnumValues Enumerates the set of values for OracleDbAwsKeyPoolLifecycleStateEnum
func GetOracleDbAwsKeyPoolLifecycleStateEnumValues() []OracleDbAwsKeyPoolLifecycleStateEnum {
	values := make([]OracleDbAwsKeyPoolLifecycleStateEnum, 0)
	for _, v := range mappingOracleDbAwsKeyPoolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbAwsKeyPoolLifecycleStateEnumStringValues Enumerates the set of values in String for OracleDbAwsKeyPoolLifecycleStateEnum
func GetOracleDbAwsKeyPoolLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOracleDbAwsKeyPoolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbAwsKeyPoolLifecycleStateEnum(val string) (OracleDbAwsKeyPoolLifecycleStateEnum, bool) {
	enum, ok := mappingOracleDbAwsKeyPoolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
