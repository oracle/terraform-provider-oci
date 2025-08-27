// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystemSummary A summary of a DB System.
type DbSystemSummary struct {

	// The OCID of the DB System.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the DB System.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the MySQL Version in use for the DB System.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The date and time the DB System was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DB System was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The database mode indicating the types of statements that are allowed to run in the DB system.
	// This mode applies only to statements run by user connections. Replicated write statements continue
	// to be allowed regardless of the DatabaseMode.
	//   - READ_WRITE: allow running read and write statements on the DB system;
	//   - READ_ONLY: only allow running read statements on the DB system.
	DatabaseMode DbSystemDatabaseModeEnum `mandatory:"true" json:"databaseMode"`

	// The access mode indicating if the database access is unrestricted (to all MySQL user accounts),
	// or restricted (to only certain users with specific privileges):
	//  - UNRESTRICTED: the access to the database is not restricted;
	//  - RESTRICTED: the access is allowed only to users with specific privileges;
	//    RESTRICTED will correspond to setting the MySQL system variable
	//    offline_mode (https://dev.mysql.com/doc/en/server-system-variables.html#sysvar_offline_mode) to ON.
	AccessMode DbSystemAccessModeEnum `mandatory:"true" json:"accessMode"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the compartment the DB System belongs in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Rest *RestDetails `mandatory:"false" json:"rest"`

	// Specifies if the DB System is highly available.
	IsHighlyAvailable *bool `mandatory:"false" json:"isHighlyAvailable"`

	CurrentPlacement *DbSystemPlacement `mandatory:"false" json:"currentPlacement"`

	// If the DB System has a HeatWave Cluster attached.
	IsHeatWaveClusterAttached *bool `mandatory:"false" json:"isHeatWaveClusterAttached"`

	HeatWaveCluster *HeatWaveClusterSummary `mandatory:"false" json:"heatWaveCluster"`

	// The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.
	// In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains
	// and the MySQL instance in that domain is promoted to the primary instance.
	// This redirection does not affect the IP address of the DB System in any way.
	// For a standalone DB System, this defines the availability domain in which the DB System is placed.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The fault domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.
	// In a failover scenario, the Read/Write endpoint is redirected to one of the other fault domains
	// and the MySQL instance in that domain is promoted to the primary instance.
	// This redirection does not affect the IP address of the DB System in any way.
	// For a standalone DB System, this defines the fault domain in which the DB System is placed.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The network endpoints available for this DB System.
	Endpoints []DbSystemEndpoint `mandatory:"false" json:"endpoints"`

	DeletionPolicy *DeletionPolicyDetails `mandatory:"false" json:"deletionPolicy"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	BackupPolicy *BackupPolicy `mandatory:"false" json:"backupPolicy"`

	// The shape of the primary instances of the DB System. The shape
	// determines resources allocated to a DB System - CPU cores
	// and memory for VM shapes; CPU cores, memory and storage for non-VM
	// (or bare metal) shapes. To get a list of shapes, use (the
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled,
	// and whether to enable or disable syncing of the Binary Logs.
	CrashRecovery CrashRecoveryStatusEnum `mandatory:"false" json:"crashRecovery,omitempty"`

	// Whether to enable monitoring via the Database Management service.
	DatabaseManagement DatabaseManagementStatusEnum `mandatory:"false" json:"databaseManagement,omitempty"`

	ReadEndpoint *ReadEndpointDetails `mandatory:"false" json:"readEndpoint"`
}

func (m DbSystemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemDatabaseModeEnum(string(m.DatabaseMode)); !ok && m.DatabaseMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseMode: %s. Supported values are: %s.", m.DatabaseMode, strings.Join(GetDbSystemDatabaseModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemAccessModeEnum(string(m.AccessMode)); !ok && m.AccessMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessMode: %s. Supported values are: %s.", m.AccessMode, strings.Join(GetDbSystemAccessModeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCrashRecoveryStatusEnum(string(m.CrashRecovery)); !ok && m.CrashRecovery != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CrashRecovery: %s. Supported values are: %s.", m.CrashRecovery, strings.Join(GetCrashRecoveryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseManagementStatusEnum(string(m.DatabaseManagement)); !ok && m.DatabaseManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseManagement: %s. Supported values are: %s.", m.DatabaseManagement, strings.Join(GetDatabaseManagementStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
