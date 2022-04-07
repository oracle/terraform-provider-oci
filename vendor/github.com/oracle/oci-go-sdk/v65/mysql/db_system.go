// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystem A DB System is the core logical unit of MySQL Database Service.
type DbSystem struct {

	// The OCID of the DB System.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment the DB System belongs in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet the DB System is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Name of the MySQL Version in use for the DB System.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGBs *int `mandatory:"true" json:"dataStorageSizeInGBs"`

	// The current state of the DB System.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Maintenance *MaintenanceDetails `mandatory:"true" json:"maintenance"`

	DeletionPolicy *DeletionPolicyDetails `mandatory:"true" json:"deletionPolicy"`

	// The date and time the DB System was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DB System was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// If the policy is to enable high availability of the instance, by
	// maintaining secondary/failover capacity as necessary.
	IsHighlyAvailable *bool `mandatory:"false" json:"isHighlyAvailable"`

	CurrentPlacement *DbSystemPlacement `mandatory:"false" json:"currentPlacement"`

	// DEPRECATED -- please use `isHeatWaveClusterAttached` instead.
	// If the DB System has an Analytics Cluster attached.
	IsAnalyticsClusterAttached *bool `mandatory:"false" json:"isAnalyticsClusterAttached"`

	AnalyticsCluster *AnalyticsClusterSummary `mandatory:"false" json:"analyticsCluster"`

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

	// The shape of the primary instances of the DB System. The shape
	// determines resources allocated to a DB System - CPU cores
	// and memory for VM shapes; CPU cores, memory and storage for non-VM
	// (or bare metal) shapes. To get a list of shapes, use (the
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	BackupPolicy *BackupPolicy `mandatory:"false" json:"backupPolicy"`

	Source DbSystemSource `mandatory:"false" json:"source"`

	// The OCID of the Configuration to be used for Instances in this DB System.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// The hostname for the primary endpoint of the DB System. Used for DNS.
	// The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, "dbsystem-1" in FQDN "dbsystem-1.subnet123.vcn1.oraclevcn.com").
	// Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The IP address the DB System is configured to listen on. A private
	// IP address of the primary endpoint of the DB System. Must be an
	// available IP address within the subnet's CIDR. This will be a
	// "dotted-quad" style IPv4 address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The port for primary endpoint of the DB System to listen on.
	Port *int `mandatory:"false" json:"port"`

	// The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.
	PortX *int `mandatory:"false" json:"portX"`

	// The network endpoints available for this DB System.
	Endpoints []DbSystemEndpoint `mandatory:"false" json:"endpoints"`

	// A list with a summary of all the Channels attached to the DB System.
	Channels []ChannelSummary `mandatory:"false" json:"channels"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Whether to run the DB System with InnoDB Redo Logs and the Double Write Buffer enabled or disabled,
	// and whether to enable or disable syncing of the Binary Logs.
	CrashRecovery CrashRecoveryStatusEnum `mandatory:"false" json:"crashRecovery,omitempty"`
}

func (m DbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCrashRecoveryStatusEnum(string(m.CrashRecovery)); !ok && m.CrashRecovery != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CrashRecovery: %s. Supported values are: %s.", m.CrashRecovery, strings.Join(GetCrashRecoveryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DbSystem) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                           `json:"description"`
		IsHighlyAvailable          *bool                             `json:"isHighlyAvailable"`
		CurrentPlacement           *DbSystemPlacement                `json:"currentPlacement"`
		IsAnalyticsClusterAttached *bool                             `json:"isAnalyticsClusterAttached"`
		AnalyticsCluster           *AnalyticsClusterSummary          `json:"analyticsCluster"`
		IsHeatWaveClusterAttached  *bool                             `json:"isHeatWaveClusterAttached"`
		HeatWaveCluster            *HeatWaveClusterSummary           `json:"heatWaveCluster"`
		AvailabilityDomain         *string                           `json:"availabilityDomain"`
		FaultDomain                *string                           `json:"faultDomain"`
		ShapeName                  *string                           `json:"shapeName"`
		BackupPolicy               *BackupPolicy                     `json:"backupPolicy"`
		Source                     dbsystemsource                    `json:"source"`
		ConfigurationId            *string                           `json:"configurationId"`
		HostnameLabel              *string                           `json:"hostnameLabel"`
		IpAddress                  *string                           `json:"ipAddress"`
		Port                       *int                              `json:"port"`
		PortX                      *int                              `json:"portX"`
		Endpoints                  []DbSystemEndpoint                `json:"endpoints"`
		Channels                   []ChannelSummary                  `json:"channels"`
		LifecycleDetails           *string                           `json:"lifecycleDetails"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		CrashRecovery              CrashRecoveryStatusEnum           `json:"crashRecovery"`
		Id                         *string                           `json:"id"`
		DisplayName                *string                           `json:"displayName"`
		CompartmentId              *string                           `json:"compartmentId"`
		SubnetId                   *string                           `json:"subnetId"`
		MysqlVersion               *string                           `json:"mysqlVersion"`
		DataStorageSizeInGBs       *int                              `json:"dataStorageSizeInGBs"`
		LifecycleState             DbSystemLifecycleStateEnum        `json:"lifecycleState"`
		Maintenance                *MaintenanceDetails               `json:"maintenance"`
		DeletionPolicy             *DeletionPolicyDetails            `json:"deletionPolicy"`
		TimeCreated                *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.IsHighlyAvailable = model.IsHighlyAvailable

	m.CurrentPlacement = model.CurrentPlacement

	m.IsAnalyticsClusterAttached = model.IsAnalyticsClusterAttached

	m.AnalyticsCluster = model.AnalyticsCluster

	m.IsHeatWaveClusterAttached = model.IsHeatWaveClusterAttached

	m.HeatWaveCluster = model.HeatWaveCluster

	m.AvailabilityDomain = model.AvailabilityDomain

	m.FaultDomain = model.FaultDomain

	m.ShapeName = model.ShapeName

	m.BackupPolicy = model.BackupPolicy

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(DbSystemSource)
	} else {
		m.Source = nil
	}

	m.ConfigurationId = model.ConfigurationId

	m.HostnameLabel = model.HostnameLabel

	m.IpAddress = model.IpAddress

	m.Port = model.Port

	m.PortX = model.PortX

	m.Endpoints = make([]DbSystemEndpoint, len(model.Endpoints))
	for i, n := range model.Endpoints {
		m.Endpoints[i] = n
	}

	m.Channels = make([]ChannelSummary, len(model.Channels))
	for i, n := range model.Channels {
		m.Channels[i] = n
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CrashRecovery = model.CrashRecovery

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.SubnetId = model.SubnetId

	m.MysqlVersion = model.MysqlVersion

	m.DataStorageSizeInGBs = model.DataStorageSizeInGBs

	m.LifecycleState = model.LifecycleState

	m.Maintenance = model.Maintenance

	m.DeletionPolicy = model.DeletionPolicy

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// DbSystemLifecycleStateEnum Enum with underlying type: string
type DbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemLifecycleStateEnum
const (
	DbSystemLifecycleStateCreating DbSystemLifecycleStateEnum = "CREATING"
	DbSystemLifecycleStateActive   DbSystemLifecycleStateEnum = "ACTIVE"
	DbSystemLifecycleStateInactive DbSystemLifecycleStateEnum = "INACTIVE"
	DbSystemLifecycleStateUpdating DbSystemLifecycleStateEnum = "UPDATING"
	DbSystemLifecycleStateDeleting DbSystemLifecycleStateEnum = "DELETING"
	DbSystemLifecycleStateDeleted  DbSystemLifecycleStateEnum = "DELETED"
	DbSystemLifecycleStateFailed   DbSystemLifecycleStateEnum = "FAILED"
)

var mappingDbSystemLifecycleStateEnum = map[string]DbSystemLifecycleStateEnum{
	"CREATING": DbSystemLifecycleStateCreating,
	"ACTIVE":   DbSystemLifecycleStateActive,
	"INACTIVE": DbSystemLifecycleStateInactive,
	"UPDATING": DbSystemLifecycleStateUpdating,
	"DELETING": DbSystemLifecycleStateDeleting,
	"DELETED":  DbSystemLifecycleStateDeleted,
	"FAILED":   DbSystemLifecycleStateFailed,
}

var mappingDbSystemLifecycleStateEnumLowerCase = map[string]DbSystemLifecycleStateEnum{
	"creating": DbSystemLifecycleStateCreating,
	"active":   DbSystemLifecycleStateActive,
	"inactive": DbSystemLifecycleStateInactive,
	"updating": DbSystemLifecycleStateUpdating,
	"deleting": DbSystemLifecycleStateDeleting,
	"deleted":  DbSystemLifecycleStateDeleted,
	"failed":   DbSystemLifecycleStateFailed,
}

// GetDbSystemLifecycleStateEnumValues Enumerates the set of values for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumValues() []DbSystemLifecycleStateEnum {
	values := make([]DbSystemLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDbSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemLifecycleStateEnum(val string) (DbSystemLifecycleStateEnum, bool) {
	enum, ok := mappingDbSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
