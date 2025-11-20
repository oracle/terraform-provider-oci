// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
// <b>Google AWS Key Resource:</b>&nbsp;&nbsp;The Oracle AWS Key Resource is used to register and manage a AWS Key within Oracle Cloud Infrastructure (OCI).
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbGcpKeyRing Oracle DB GCP Key Ring resource object.
type OracleDbGcpKeyRing struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB GCP Key Ring resource.
	Id *string `mandatory:"true" json:"id"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the Oracle DB GCP Key Ring resource resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Oracle DB GCP Identity Connector resource resides.
	OracleDbConnectorId *string `mandatory:"false" json:"oracleDbConnectorId"`

	// Display name of DB GCP Key Ring resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// GCP Key Ring ID.
	GcpKeyRingId *string `mandatory:"false" json:"gcpKeyRingId"`

	// Oracle DB GCP Key Ring resource Type.
	Type *string `mandatory:"false" json:"type"`

	// Location of the GCP Key Ring resource.
	Location *string `mandatory:"false" json:"location"`

	// Oracle DB GCP Key Ring resource's properties.
	Properties map[string]string `mandatory:"false" json:"properties"`

	// The lifecycle state of the Oracle DB GCP Key Ring resource.
	LifecycleState OracleDbGcpKeyRingLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the DB GCP Key Ring resource was created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the DB GCP Key Ring resource was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'
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

func (m OracleDbGcpKeyRing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbGcpKeyRing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbGcpKeyRingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOracleDbGcpKeyRingLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbGcpKeyRingLifecycleStateEnum Enum with underlying type: string
type OracleDbGcpKeyRingLifecycleStateEnum string

// Set of constants representing the allowable values for OracleDbGcpKeyRingLifecycleStateEnum
const (
	OracleDbGcpKeyRingLifecycleStateCreating OracleDbGcpKeyRingLifecycleStateEnum = "CREATING"
	OracleDbGcpKeyRingLifecycleStateActive   OracleDbGcpKeyRingLifecycleStateEnum = "ACTIVE"
	OracleDbGcpKeyRingLifecycleStateUpdating OracleDbGcpKeyRingLifecycleStateEnum = "UPDATING"
	OracleDbGcpKeyRingLifecycleStateDeleting OracleDbGcpKeyRingLifecycleStateEnum = "DELETING"
	OracleDbGcpKeyRingLifecycleStateDeleted  OracleDbGcpKeyRingLifecycleStateEnum = "DELETED"
	OracleDbGcpKeyRingLifecycleStateFailed   OracleDbGcpKeyRingLifecycleStateEnum = "FAILED"
)

var mappingOracleDbGcpKeyRingLifecycleStateEnum = map[string]OracleDbGcpKeyRingLifecycleStateEnum{
	"CREATING": OracleDbGcpKeyRingLifecycleStateCreating,
	"ACTIVE":   OracleDbGcpKeyRingLifecycleStateActive,
	"UPDATING": OracleDbGcpKeyRingLifecycleStateUpdating,
	"DELETING": OracleDbGcpKeyRingLifecycleStateDeleting,
	"DELETED":  OracleDbGcpKeyRingLifecycleStateDeleted,
	"FAILED":   OracleDbGcpKeyRingLifecycleStateFailed,
}

var mappingOracleDbGcpKeyRingLifecycleStateEnumLowerCase = map[string]OracleDbGcpKeyRingLifecycleStateEnum{
	"creating": OracleDbGcpKeyRingLifecycleStateCreating,
	"active":   OracleDbGcpKeyRingLifecycleStateActive,
	"updating": OracleDbGcpKeyRingLifecycleStateUpdating,
	"deleting": OracleDbGcpKeyRingLifecycleStateDeleting,
	"deleted":  OracleDbGcpKeyRingLifecycleStateDeleted,
	"failed":   OracleDbGcpKeyRingLifecycleStateFailed,
}

// GetOracleDbGcpKeyRingLifecycleStateEnumValues Enumerates the set of values for OracleDbGcpKeyRingLifecycleStateEnum
func GetOracleDbGcpKeyRingLifecycleStateEnumValues() []OracleDbGcpKeyRingLifecycleStateEnum {
	values := make([]OracleDbGcpKeyRingLifecycleStateEnum, 0)
	for _, v := range mappingOracleDbGcpKeyRingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbGcpKeyRingLifecycleStateEnumStringValues Enumerates the set of values in String for OracleDbGcpKeyRingLifecycleStateEnum
func GetOracleDbGcpKeyRingLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOracleDbGcpKeyRingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbGcpKeyRingLifecycleStateEnum(val string) (OracleDbGcpKeyRingLifecycleStateEnum, bool) {
	enum, ok := mappingOracleDbGcpKeyRingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
