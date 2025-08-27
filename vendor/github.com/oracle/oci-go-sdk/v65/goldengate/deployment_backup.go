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

// DeploymentBackup A backup of the current state of the GoldenGate deployment. Can be used to restore a deployment, or create a new deployment with that state as the starting deployment state.
type DeploymentBackup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"true" json:"deploymentId"`

	// The type of deployment, which can be any one of the Allowed values.
	// NOTE: Use of the value 'OGG' is maintained for backward compatibility purposes.
	//     Its use is discouraged in favor of 'DATABASE_ORACLE'.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Version of OGG
	OggVersion *string `mandatory:"true" json:"oggVersion"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// True if this object is automatically created
	IsAutomatic *bool `mandatory:"false" json:"isAutomatic"`

	// Possible deployment backup source types.
	BackupSourceType DeploymentBackupBackupSourceTypeEnum `mandatory:"false" json:"backupSourceType,omitempty"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time of the resource backup. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeOfBackup *common.SDKTime `mandatory:"false" json:"timeOfBackup"`

	// The time of the resource backup finish. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeBackupFinished *common.SDKTime `mandatory:"false" json:"timeBackupFinished"`

	// The size of the backup stored in object storage (in bytes)
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Possible Deployment backup types.
	BackupType DeploymentBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Name of the object to be uploaded to object storage
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Parameter to allow users to create backup without trails
	IsMetadataOnly *bool `mandatory:"false" json:"isMetadataOnly"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m DeploymentBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeploymentBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeploymentBackupBackupSourceTypeEnum(string(m.BackupSourceType)); !ok && m.BackupSourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupSourceType: %s. Supported values are: %s.", m.BackupSourceType, strings.Join(GetDeploymentBackupBackupSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetDeploymentBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeploymentBackupBackupSourceTypeEnum Enum with underlying type: string
type DeploymentBackupBackupSourceTypeEnum string

// Set of constants representing the allowable values for DeploymentBackupBackupSourceTypeEnum
const (
	DeploymentBackupBackupSourceTypeManual    DeploymentBackupBackupSourceTypeEnum = "MANUAL"
	DeploymentBackupBackupSourceTypeAutomatic DeploymentBackupBackupSourceTypeEnum = "AUTOMATIC"
	DeploymentBackupBackupSourceTypeScheduled DeploymentBackupBackupSourceTypeEnum = "SCHEDULED"
)

var mappingDeploymentBackupBackupSourceTypeEnum = map[string]DeploymentBackupBackupSourceTypeEnum{
	"MANUAL":    DeploymentBackupBackupSourceTypeManual,
	"AUTOMATIC": DeploymentBackupBackupSourceTypeAutomatic,
	"SCHEDULED": DeploymentBackupBackupSourceTypeScheduled,
}

var mappingDeploymentBackupBackupSourceTypeEnumLowerCase = map[string]DeploymentBackupBackupSourceTypeEnum{
	"manual":    DeploymentBackupBackupSourceTypeManual,
	"automatic": DeploymentBackupBackupSourceTypeAutomatic,
	"scheduled": DeploymentBackupBackupSourceTypeScheduled,
}

// GetDeploymentBackupBackupSourceTypeEnumValues Enumerates the set of values for DeploymentBackupBackupSourceTypeEnum
func GetDeploymentBackupBackupSourceTypeEnumValues() []DeploymentBackupBackupSourceTypeEnum {
	values := make([]DeploymentBackupBackupSourceTypeEnum, 0)
	for _, v := range mappingDeploymentBackupBackupSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentBackupBackupSourceTypeEnumStringValues Enumerates the set of values in String for DeploymentBackupBackupSourceTypeEnum
func GetDeploymentBackupBackupSourceTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATIC",
		"SCHEDULED",
	}
}

// GetMappingDeploymentBackupBackupSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentBackupBackupSourceTypeEnum(val string) (DeploymentBackupBackupSourceTypeEnum, bool) {
	enum, ok := mappingDeploymentBackupBackupSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
