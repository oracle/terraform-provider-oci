// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmClusterProperties OLVM Cluster properties.
type OlvmClusterProperties struct {

	// A human-readable name in plain text.
	ClusterName *string `mandatory:"true" json:"clusterName"`

	// A human-readable description in plain text.
	ClusterDescription *string `mandatory:"false" json:"clusterDescription"`

	// Free text containing comments about this object.
	Comment *string `mandatory:"false" json:"comment"`

	// Whether ballooning is enabled.
	IsBallooningEnabled *bool `mandatory:"false" json:"isBallooningEnabled"`

	BiosType *OlvmBiosType `mandatory:"false" json:"biosType"`

	Cpu *OlvmCpu `mandatory:"false" json:"cpu"`

	// Custom scheduling policy properties of the cluster.
	CustomSchedulingPolicyProperties []OlvmProperty `mandatory:"false" json:"customSchedulingPolicyProperties"`

	Display *OlvmDisplay `mandatory:"false" json:"display"`

	ErrorHandling *OlvmErrorHandling `mandatory:"false" json:"errorHandling"`

	FencingPolicy *OlvmFencingPolicy `mandatory:"false" json:"fencingPolicy"`

	// FIPS mode of the cluster.
	FipsMode OlvmClusterPropertiesFipsModeEnum `mandatory:"false" json:"fipsMode,omitempty"`

	// The type of firewall to be used on hosts in this cluster.
	FirewallType OlvmClusterPropertiesFirewallTypeEnum `mandatory:"false" json:"firewallType,omitempty"`

	// Indicates if Gluster service is used.
	IsGlusterService *bool `mandatory:"false" json:"isGlusterService"`

	// The name of the tuned profile.
	GlusterTunedProfile *string `mandatory:"false" json:"glusterTunedProfile"`

	// Indicates whether HA reservation is enabled.
	IsHaReservation *bool `mandatory:"false" json:"isHaReservation"`

	Ksm *Ksm `mandatory:"false" json:"ksm"`

	// The memory consumption threshold for logging audit log events
	LogMaxMemoryUsedThreshold *int `mandatory:"false" json:"logMaxMemoryUsedThreshold"`

	// The memory consumption threshold type for logging audit log events.
	LogMaxMemoryUsedThresholdType OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum `mandatory:"false" json:"logMaxMemoryUsedThresholdType,omitempty"`

	MemoryPolicy *OlvmMemoryPolicy `mandatory:"false" json:"memoryPolicy"`

	MigrationPolicy *OlvmMigrationOptions `mandatory:"false" json:"migrationPolicy"`

	// Set of random number generator (RNG) sources required from each host in the cluster.
	RequiredRngSources []RngSourceEnum `mandatory:"false" json:"requiredRngSources"`

	SerialNumber *OlvmSerialNumber `mandatory:"false" json:"serialNumber"`

	// List of supported versions.
	SupportedVersions []OlvmVersion `mandatory:"false" json:"supportedVersions"`

	// Describes all switch types supported by the Manager
	SwitchType OlvmClusterPropertiesSwitchTypeEnum `mandatory:"false" json:"switchType,omitempty"`

	// Indicates if threads should be used as cores
	IsUseThreadsAsCores *bool `mandatory:"false" json:"isUseThreadsAsCores"`

	// Indicates if trusted service is enabled
	IsTrustedService *bool `mandatory:"false" json:"isTrustedService"`

	// Indicates if tunnel migration is enabled
	IsTunnelMigration *bool `mandatory:"false" json:"isTunnelMigration"`

	// The upgrade correlation identifier.
	UpgradeCorrelationId *string `mandatory:"false" json:"upgradeCorrelationId"`

	// Indicates if an upgrade has been started for the cluster.
	IsUpdateInProgress *bool `mandatory:"false" json:"isUpdateInProgress"`

	// If an upgrade is in progress, the upgrade???s reported percent complete.
	UpgradePercentComplete *int `mandatory:"false" json:"upgradePercentComplete"`

	Version *OlvmVersion `mandatory:"false" json:"version"`

	// Indicates if virt service is enabled.
	IsVirtService *bool `mandatory:"false" json:"isVirtService"`

	// Indicates if VNC encryption is enabled.
	IsVncEncryption *bool `mandatory:"false" json:"isVncEncryption"`

	// List of data centers where storage domain belongs
	DataCenters []OlvmDataCenter `mandatory:"false" json:"dataCenters"`
}

func (m OlvmClusterProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmClusterProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmClusterPropertiesFipsModeEnum(string(m.FipsMode)); !ok && m.FipsMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FipsMode: %s. Supported values are: %s.", m.FipsMode, strings.Join(GetOlvmClusterPropertiesFipsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmClusterPropertiesFirewallTypeEnum(string(m.FirewallType)); !ok && m.FirewallType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FirewallType: %s. Supported values are: %s.", m.FirewallType, strings.Join(GetOlvmClusterPropertiesFirewallTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum(string(m.LogMaxMemoryUsedThresholdType)); !ok && m.LogMaxMemoryUsedThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogMaxMemoryUsedThresholdType: %s. Supported values are: %s.", m.LogMaxMemoryUsedThresholdType, strings.Join(GetOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmClusterPropertiesSwitchTypeEnum(string(m.SwitchType)); !ok && m.SwitchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SwitchType: %s. Supported values are: %s.", m.SwitchType, strings.Join(GetOlvmClusterPropertiesSwitchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmClusterPropertiesFipsModeEnum Enum with underlying type: string
type OlvmClusterPropertiesFipsModeEnum string

// Set of constants representing the allowable values for OlvmClusterPropertiesFipsModeEnum
const (
	OlvmClusterPropertiesFipsModeDisabled  OlvmClusterPropertiesFipsModeEnum = "DISABLED"
	OlvmClusterPropertiesFipsModeEnabled   OlvmClusterPropertiesFipsModeEnum = "ENABLED"
	OlvmClusterPropertiesFipsModeUndefined OlvmClusterPropertiesFipsModeEnum = "UNDEFINED"
)

var mappingOlvmClusterPropertiesFipsModeEnum = map[string]OlvmClusterPropertiesFipsModeEnum{
	"DISABLED":  OlvmClusterPropertiesFipsModeDisabled,
	"ENABLED":   OlvmClusterPropertiesFipsModeEnabled,
	"UNDEFINED": OlvmClusterPropertiesFipsModeUndefined,
}

var mappingOlvmClusterPropertiesFipsModeEnumLowerCase = map[string]OlvmClusterPropertiesFipsModeEnum{
	"disabled":  OlvmClusterPropertiesFipsModeDisabled,
	"enabled":   OlvmClusterPropertiesFipsModeEnabled,
	"undefined": OlvmClusterPropertiesFipsModeUndefined,
}

// GetOlvmClusterPropertiesFipsModeEnumValues Enumerates the set of values for OlvmClusterPropertiesFipsModeEnum
func GetOlvmClusterPropertiesFipsModeEnumValues() []OlvmClusterPropertiesFipsModeEnum {
	values := make([]OlvmClusterPropertiesFipsModeEnum, 0)
	for _, v := range mappingOlvmClusterPropertiesFipsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmClusterPropertiesFipsModeEnumStringValues Enumerates the set of values in String for OlvmClusterPropertiesFipsModeEnum
func GetOlvmClusterPropertiesFipsModeEnumStringValues() []string {
	return []string{
		"DISABLED",
		"ENABLED",
		"UNDEFINED",
	}
}

// GetMappingOlvmClusterPropertiesFipsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmClusterPropertiesFipsModeEnum(val string) (OlvmClusterPropertiesFipsModeEnum, bool) {
	enum, ok := mappingOlvmClusterPropertiesFipsModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmClusterPropertiesFirewallTypeEnum Enum with underlying type: string
type OlvmClusterPropertiesFirewallTypeEnum string

// Set of constants representing the allowable values for OlvmClusterPropertiesFirewallTypeEnum
const (
	OlvmClusterPropertiesFirewallTypeFirewalld OlvmClusterPropertiesFirewallTypeEnum = "FIREWALLD"
	OlvmClusterPropertiesFirewallTypeIptables  OlvmClusterPropertiesFirewallTypeEnum = "IPTABLES"
)

var mappingOlvmClusterPropertiesFirewallTypeEnum = map[string]OlvmClusterPropertiesFirewallTypeEnum{
	"FIREWALLD": OlvmClusterPropertiesFirewallTypeFirewalld,
	"IPTABLES":  OlvmClusterPropertiesFirewallTypeIptables,
}

var mappingOlvmClusterPropertiesFirewallTypeEnumLowerCase = map[string]OlvmClusterPropertiesFirewallTypeEnum{
	"firewalld": OlvmClusterPropertiesFirewallTypeFirewalld,
	"iptables":  OlvmClusterPropertiesFirewallTypeIptables,
}

// GetOlvmClusterPropertiesFirewallTypeEnumValues Enumerates the set of values for OlvmClusterPropertiesFirewallTypeEnum
func GetOlvmClusterPropertiesFirewallTypeEnumValues() []OlvmClusterPropertiesFirewallTypeEnum {
	values := make([]OlvmClusterPropertiesFirewallTypeEnum, 0)
	for _, v := range mappingOlvmClusterPropertiesFirewallTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmClusterPropertiesFirewallTypeEnumStringValues Enumerates the set of values in String for OlvmClusterPropertiesFirewallTypeEnum
func GetOlvmClusterPropertiesFirewallTypeEnumStringValues() []string {
	return []string{
		"FIREWALLD",
		"IPTABLES",
	}
}

// GetMappingOlvmClusterPropertiesFirewallTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmClusterPropertiesFirewallTypeEnum(val string) (OlvmClusterPropertiesFirewallTypeEnum, bool) {
	enum, ok := mappingOlvmClusterPropertiesFirewallTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum Enum with underlying type: string
type OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum string

// Set of constants representing the allowable values for OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum
const (
	OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeAbsoluteValueInMb OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum = "ABSOLUTE_VALUE_IN_MB"
	OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypePercentage        OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum = "PERCENTAGE"
)

var mappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum = map[string]OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum{
	"ABSOLUTE_VALUE_IN_MB": OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeAbsoluteValueInMb,
	"PERCENTAGE":           OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypePercentage,
}

var mappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumLowerCase = map[string]OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum{
	"absolute_value_in_mb": OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeAbsoluteValueInMb,
	"percentage":           OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypePercentage,
}

// GetOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumValues Enumerates the set of values for OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum
func GetOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumValues() []OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum {
	values := make([]OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum, 0)
	for _, v := range mappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumStringValues Enumerates the set of values in String for OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum
func GetOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumStringValues() []string {
	return []string{
		"ABSOLUTE_VALUE_IN_MB",
		"PERCENTAGE",
	}
}

// GetMappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum(val string) (OlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnum, bool) {
	enum, ok := mappingOlvmClusterPropertiesLogMaxMemoryUsedThresholdTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmClusterPropertiesSwitchTypeEnum Enum with underlying type: string
type OlvmClusterPropertiesSwitchTypeEnum string

// Set of constants representing the allowable values for OlvmClusterPropertiesSwitchTypeEnum
const (
	OlvmClusterPropertiesSwitchTypeLegacy OlvmClusterPropertiesSwitchTypeEnum = "LEGACY"
	OlvmClusterPropertiesSwitchTypeOvs    OlvmClusterPropertiesSwitchTypeEnum = "OVS"
)

var mappingOlvmClusterPropertiesSwitchTypeEnum = map[string]OlvmClusterPropertiesSwitchTypeEnum{
	"LEGACY": OlvmClusterPropertiesSwitchTypeLegacy,
	"OVS":    OlvmClusterPropertiesSwitchTypeOvs,
}

var mappingOlvmClusterPropertiesSwitchTypeEnumLowerCase = map[string]OlvmClusterPropertiesSwitchTypeEnum{
	"legacy": OlvmClusterPropertiesSwitchTypeLegacy,
	"ovs":    OlvmClusterPropertiesSwitchTypeOvs,
}

// GetOlvmClusterPropertiesSwitchTypeEnumValues Enumerates the set of values for OlvmClusterPropertiesSwitchTypeEnum
func GetOlvmClusterPropertiesSwitchTypeEnumValues() []OlvmClusterPropertiesSwitchTypeEnum {
	values := make([]OlvmClusterPropertiesSwitchTypeEnum, 0)
	for _, v := range mappingOlvmClusterPropertiesSwitchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmClusterPropertiesSwitchTypeEnumStringValues Enumerates the set of values in String for OlvmClusterPropertiesSwitchTypeEnum
func GetOlvmClusterPropertiesSwitchTypeEnumStringValues() []string {
	return []string{
		"LEGACY",
		"OVS",
	}
}

// GetMappingOlvmClusterPropertiesSwitchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmClusterPropertiesSwitchTypeEnum(val string) (OlvmClusterPropertiesSwitchTypeEnum, bool) {
	enum, ok := mappingOlvmClusterPropertiesSwitchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
