// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VanityDomain Vanity Domain resource
type VanityDomain struct {

	// The unique identifier (OCID) of the VanityDomain. Can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Vanity domain
	VanityDomain *string `mandatory:"true" json:"vanityDomain"`

	// The OCID of the Fusion environment that the VanityDomain is created on
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// The current lifecycleState of the VanityDomain
	LifecycleState VanityDomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current lifecycleDetails of the VanityDomain
	LifecycleDetails VanityDomainLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// The time the VanityDomain was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the VanityDomain was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time the VanityDomain is scheduled to enable. An RFC3339 formatted datetime string
	TimeEnabled *common.SDKTime `mandatory:"true" json:"timeEnabled"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The prefix value of the DnsPrefix. Can't be changed after creation
	Prefix *string `mandatory:"false" json:"prefix"`

	// The origin request type for which the certificate is generated
	OriginCertRequestType VanityDomainOriginCertRequestTypeEnum `mandatory:"false" json:"originCertRequestType,omitempty"`

	// The cdn request type for which the certificate is generated
	CdnCertRequestType VanityDomainCdnCertRequestTypeEnum `mandatory:"false" json:"cdnCertRequestType,omitempty"`

	// The dns is managed by the customer or Oracle
	DnsManagedBy VanityDomainDnsManagedByEnum `mandatory:"false" json:"dnsManagedBy,omitempty"`

	CertificateInfo *CertificateInfo `mandatory:"false" json:"certificateInfo"`

	// The cm link that was used to create the DNS prefix
	ChangeManagementLink *string `mandatory:"false" json:"changeManagementLink"`

	// The ID of the VanityDomainActivity is scheduled
	ScheduledActivityId *string `mandatory:"false" json:"scheduledActivityId"`

	// List of dns records, comma separated
	CustomerDnsRecords []VanityDnsRecord `mandatory:"false" json:"customerDnsRecords"`

	// Identify if this dns is inactive or active
	IsDnsStatusReady *bool `mandatory:"false" json:"isDnsStatusReady"`

	// Identify if this origin cert is inactive or active
	IsOriginCertStatusReady *bool `mandatory:"false" json:"isOriginCertStatusReady"`

	// The origin cert status
	OriginCertStatus VanityDomainOriginCertStatusEnum `mandatory:"false" json:"originCertStatus,omitempty"`

	// The origin cert expiry date
	TimeOriginCertExpired *common.SDKTime `mandatory:"false" json:"timeOriginCertExpired"`

	// Identify if this cdn cert is inactive or active
	IsCdnCertStatusReady *bool `mandatory:"false" json:"isCdnCertStatusReady"`

	// The cdn cert status
	CdnCertStatus VanityDomainCdnCertStatusEnum `mandatory:"false" json:"cdnCertStatus,omitempty"`

	// The cdn cert expiry date
	TimeCdnCertExpired *common.SDKTime `mandatory:"false" json:"timeCdnCertExpired"`
}

func (m VanityDomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VanityDomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVanityDomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVanityDomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetVanityDomainLifecycleDetailsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVanityDomainOriginCertRequestTypeEnum(string(m.OriginCertRequestType)); !ok && m.OriginCertRequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OriginCertRequestType: %s. Supported values are: %s.", m.OriginCertRequestType, strings.Join(GetVanityDomainOriginCertRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainCdnCertRequestTypeEnum(string(m.CdnCertRequestType)); !ok && m.CdnCertRequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CdnCertRequestType: %s. Supported values are: %s.", m.CdnCertRequestType, strings.Join(GetVanityDomainCdnCertRequestTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainDnsManagedByEnum(string(m.DnsManagedBy)); !ok && m.DnsManagedBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsManagedBy: %s. Supported values are: %s.", m.DnsManagedBy, strings.Join(GetVanityDomainDnsManagedByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainOriginCertStatusEnum(string(m.OriginCertStatus)); !ok && m.OriginCertStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OriginCertStatus: %s. Supported values are: %s.", m.OriginCertStatus, strings.Join(GetVanityDomainOriginCertStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVanityDomainCdnCertStatusEnum(string(m.CdnCertStatus)); !ok && m.CdnCertStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CdnCertStatus: %s. Supported values are: %s.", m.CdnCertStatus, strings.Join(GetVanityDomainCdnCertStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VanityDomainLifecycleStateEnum Enum with underlying type: string
type VanityDomainLifecycleStateEnum string

// Set of constants representing the allowable values for VanityDomainLifecycleStateEnum
const (
	VanityDomainLifecycleStateInactive       VanityDomainLifecycleStateEnum = "INACTIVE"
	VanityDomainLifecycleStateActive         VanityDomainLifecycleStateEnum = "ACTIVE"
	VanityDomainLifecycleStateUpdating       VanityDomainLifecycleStateEnum = "UPDATING"
	VanityDomainLifecycleStateNeedsAttention VanityDomainLifecycleStateEnum = "NEEDS_ATTENTION"
	VanityDomainLifecycleStateDeleting       VanityDomainLifecycleStateEnum = "DELETING"
	VanityDomainLifecycleStateDeleted        VanityDomainLifecycleStateEnum = "DELETED"
	VanityDomainLifecycleStateFailed         VanityDomainLifecycleStateEnum = "FAILED"
)

var mappingVanityDomainLifecycleStateEnum = map[string]VanityDomainLifecycleStateEnum{
	"INACTIVE":        VanityDomainLifecycleStateInactive,
	"ACTIVE":          VanityDomainLifecycleStateActive,
	"UPDATING":        VanityDomainLifecycleStateUpdating,
	"NEEDS_ATTENTION": VanityDomainLifecycleStateNeedsAttention,
	"DELETING":        VanityDomainLifecycleStateDeleting,
	"DELETED":         VanityDomainLifecycleStateDeleted,
	"FAILED":          VanityDomainLifecycleStateFailed,
}

var mappingVanityDomainLifecycleStateEnumLowerCase = map[string]VanityDomainLifecycleStateEnum{
	"inactive":        VanityDomainLifecycleStateInactive,
	"active":          VanityDomainLifecycleStateActive,
	"updating":        VanityDomainLifecycleStateUpdating,
	"needs_attention": VanityDomainLifecycleStateNeedsAttention,
	"deleting":        VanityDomainLifecycleStateDeleting,
	"deleted":         VanityDomainLifecycleStateDeleted,
	"failed":          VanityDomainLifecycleStateFailed,
}

// GetVanityDomainLifecycleStateEnumValues Enumerates the set of values for VanityDomainLifecycleStateEnum
func GetVanityDomainLifecycleStateEnumValues() []VanityDomainLifecycleStateEnum {
	values := make([]VanityDomainLifecycleStateEnum, 0)
	for _, v := range mappingVanityDomainLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainLifecycleStateEnumStringValues Enumerates the set of values in String for VanityDomainLifecycleStateEnum
func GetVanityDomainLifecycleStateEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVanityDomainLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainLifecycleStateEnum(val string) (VanityDomainLifecycleStateEnum, bool) {
	enum, ok := mappingVanityDomainLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainLifecycleDetailsEnum Enum with underlying type: string
type VanityDomainLifecycleDetailsEnum string

// Set of constants representing the allowable values for VanityDomainLifecycleDetailsEnum
const (
	VanityDomainLifecycleDetailsNone                       VanityDomainLifecycleDetailsEnum = "NONE"
	VanityDomainLifecycleDetailsDomainValidationInProgress VanityDomainLifecycleDetailsEnum = "DOMAIN_VALIDATION_IN_PROGRESS"
	VanityDomainLifecycleDetailsDomainValidationCompleted  VanityDomainLifecycleDetailsEnum = "DOMAIN_VALIDATION_COMPLETED"
	VanityDomainLifecycleDetailsDomainValidationFailed     VanityDomainLifecycleDetailsEnum = "DOMAIN_VALIDATION_FAILED"
	VanityDomainLifecycleDetailsCertConfigInProgress       VanityDomainLifecycleDetailsEnum = "CERT_CONFIG_IN_PROGRESS"
	VanityDomainLifecycleDetailsCertConfigCompleted        VanityDomainLifecycleDetailsEnum = "CERT_CONFIG_COMPLETED"
	VanityDomainLifecycleDetailsCertConfigFailed           VanityDomainLifecycleDetailsEnum = "CERT_CONFIG_FAILED"
	VanityDomainLifecycleDetailsEnableScheduled            VanityDomainLifecycleDetailsEnum = "ENABLE_SCHEDULED"
	VanityDomainLifecycleDetailsEnabling                   VanityDomainLifecycleDetailsEnum = "ENABLING"
	VanityDomainLifecycleDetailsEnabled                    VanityDomainLifecycleDetailsEnum = "ENABLED"
	VanityDomainLifecycleDetailsEnableFailed               VanityDomainLifecycleDetailsEnum = "ENABLE_FAILED"
	VanityDomainLifecycleDetailsEnableCancelled            VanityDomainLifecycleDetailsEnum = "ENABLE_CANCELLED"
	VanityDomainLifecycleDetailsDeleteFailed               VanityDomainLifecycleDetailsEnum = "DELETE_FAILED"
	VanityDomainLifecycleDetailsUnknown                    VanityDomainLifecycleDetailsEnum = "UNKNOWN"
)

var mappingVanityDomainLifecycleDetailsEnum = map[string]VanityDomainLifecycleDetailsEnum{
	"NONE":                          VanityDomainLifecycleDetailsNone,
	"DOMAIN_VALIDATION_IN_PROGRESS": VanityDomainLifecycleDetailsDomainValidationInProgress,
	"DOMAIN_VALIDATION_COMPLETED":   VanityDomainLifecycleDetailsDomainValidationCompleted,
	"DOMAIN_VALIDATION_FAILED":      VanityDomainLifecycleDetailsDomainValidationFailed,
	"CERT_CONFIG_IN_PROGRESS":       VanityDomainLifecycleDetailsCertConfigInProgress,
	"CERT_CONFIG_COMPLETED":         VanityDomainLifecycleDetailsCertConfigCompleted,
	"CERT_CONFIG_FAILED":            VanityDomainLifecycleDetailsCertConfigFailed,
	"ENABLE_SCHEDULED":              VanityDomainLifecycleDetailsEnableScheduled,
	"ENABLING":                      VanityDomainLifecycleDetailsEnabling,
	"ENABLED":                       VanityDomainLifecycleDetailsEnabled,
	"ENABLE_FAILED":                 VanityDomainLifecycleDetailsEnableFailed,
	"ENABLE_CANCELLED":              VanityDomainLifecycleDetailsEnableCancelled,
	"DELETE_FAILED":                 VanityDomainLifecycleDetailsDeleteFailed,
	"UNKNOWN":                       VanityDomainLifecycleDetailsUnknown,
}

var mappingVanityDomainLifecycleDetailsEnumLowerCase = map[string]VanityDomainLifecycleDetailsEnum{
	"none":                          VanityDomainLifecycleDetailsNone,
	"domain_validation_in_progress": VanityDomainLifecycleDetailsDomainValidationInProgress,
	"domain_validation_completed":   VanityDomainLifecycleDetailsDomainValidationCompleted,
	"domain_validation_failed":      VanityDomainLifecycleDetailsDomainValidationFailed,
	"cert_config_in_progress":       VanityDomainLifecycleDetailsCertConfigInProgress,
	"cert_config_completed":         VanityDomainLifecycleDetailsCertConfigCompleted,
	"cert_config_failed":            VanityDomainLifecycleDetailsCertConfigFailed,
	"enable_scheduled":              VanityDomainLifecycleDetailsEnableScheduled,
	"enabling":                      VanityDomainLifecycleDetailsEnabling,
	"enabled":                       VanityDomainLifecycleDetailsEnabled,
	"enable_failed":                 VanityDomainLifecycleDetailsEnableFailed,
	"enable_cancelled":              VanityDomainLifecycleDetailsEnableCancelled,
	"delete_failed":                 VanityDomainLifecycleDetailsDeleteFailed,
	"unknown":                       VanityDomainLifecycleDetailsUnknown,
}

// GetVanityDomainLifecycleDetailsEnumValues Enumerates the set of values for VanityDomainLifecycleDetailsEnum
func GetVanityDomainLifecycleDetailsEnumValues() []VanityDomainLifecycleDetailsEnum {
	values := make([]VanityDomainLifecycleDetailsEnum, 0)
	for _, v := range mappingVanityDomainLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainLifecycleDetailsEnumStringValues Enumerates the set of values in String for VanityDomainLifecycleDetailsEnum
func GetVanityDomainLifecycleDetailsEnumStringValues() []string {
	return []string{
		"NONE",
		"DOMAIN_VALIDATION_IN_PROGRESS",
		"DOMAIN_VALIDATION_COMPLETED",
		"DOMAIN_VALIDATION_FAILED",
		"CERT_CONFIG_IN_PROGRESS",
		"CERT_CONFIG_COMPLETED",
		"CERT_CONFIG_FAILED",
		"ENABLE_SCHEDULED",
		"ENABLING",
		"ENABLED",
		"ENABLE_FAILED",
		"ENABLE_CANCELLED",
		"DELETE_FAILED",
		"UNKNOWN",
	}
}

// GetMappingVanityDomainLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainLifecycleDetailsEnum(val string) (VanityDomainLifecycleDetailsEnum, bool) {
	enum, ok := mappingVanityDomainLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainOriginCertRequestTypeEnum Enum with underlying type: string
type VanityDomainOriginCertRequestTypeEnum string

// Set of constants representing the allowable values for VanityDomainOriginCertRequestTypeEnum
const (
	VanityDomainOriginCertRequestTypeCsr VanityDomainOriginCertRequestTypeEnum = "REQUEST_CSR"
	VanityDomainOriginCertRequestTypeDv  VanityDomainOriginCertRequestTypeEnum = "REQUEST_DV"
)

var mappingVanityDomainOriginCertRequestTypeEnum = map[string]VanityDomainOriginCertRequestTypeEnum{
	"REQUEST_CSR": VanityDomainOriginCertRequestTypeCsr,
	"REQUEST_DV":  VanityDomainOriginCertRequestTypeDv,
}

var mappingVanityDomainOriginCertRequestTypeEnumLowerCase = map[string]VanityDomainOriginCertRequestTypeEnum{
	"request_csr": VanityDomainOriginCertRequestTypeCsr,
	"request_dv":  VanityDomainOriginCertRequestTypeDv,
}

// GetVanityDomainOriginCertRequestTypeEnumValues Enumerates the set of values for VanityDomainOriginCertRequestTypeEnum
func GetVanityDomainOriginCertRequestTypeEnumValues() []VanityDomainOriginCertRequestTypeEnum {
	values := make([]VanityDomainOriginCertRequestTypeEnum, 0)
	for _, v := range mappingVanityDomainOriginCertRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainOriginCertRequestTypeEnumStringValues Enumerates the set of values in String for VanityDomainOriginCertRequestTypeEnum
func GetVanityDomainOriginCertRequestTypeEnumStringValues() []string {
	return []string{
		"REQUEST_CSR",
		"REQUEST_DV",
	}
}

// GetMappingVanityDomainOriginCertRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainOriginCertRequestTypeEnum(val string) (VanityDomainOriginCertRequestTypeEnum, bool) {
	enum, ok := mappingVanityDomainOriginCertRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainCdnCertRequestTypeEnum Enum with underlying type: string
type VanityDomainCdnCertRequestTypeEnum string

// Set of constants representing the allowable values for VanityDomainCdnCertRequestTypeEnum
const (
	VanityDomainCdnCertRequestTypeCsr VanityDomainCdnCertRequestTypeEnum = "REQUEST_CSR"
	VanityDomainCdnCertRequestTypeDv  VanityDomainCdnCertRequestTypeEnum = "REQUEST_DV"
)

var mappingVanityDomainCdnCertRequestTypeEnum = map[string]VanityDomainCdnCertRequestTypeEnum{
	"REQUEST_CSR": VanityDomainCdnCertRequestTypeCsr,
	"REQUEST_DV":  VanityDomainCdnCertRequestTypeDv,
}

var mappingVanityDomainCdnCertRequestTypeEnumLowerCase = map[string]VanityDomainCdnCertRequestTypeEnum{
	"request_csr": VanityDomainCdnCertRequestTypeCsr,
	"request_dv":  VanityDomainCdnCertRequestTypeDv,
}

// GetVanityDomainCdnCertRequestTypeEnumValues Enumerates the set of values for VanityDomainCdnCertRequestTypeEnum
func GetVanityDomainCdnCertRequestTypeEnumValues() []VanityDomainCdnCertRequestTypeEnum {
	values := make([]VanityDomainCdnCertRequestTypeEnum, 0)
	for _, v := range mappingVanityDomainCdnCertRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainCdnCertRequestTypeEnumStringValues Enumerates the set of values in String for VanityDomainCdnCertRequestTypeEnum
func GetVanityDomainCdnCertRequestTypeEnumStringValues() []string {
	return []string{
		"REQUEST_CSR",
		"REQUEST_DV",
	}
}

// GetMappingVanityDomainCdnCertRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainCdnCertRequestTypeEnum(val string) (VanityDomainCdnCertRequestTypeEnum, bool) {
	enum, ok := mappingVanityDomainCdnCertRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainDnsManagedByEnum Enum with underlying type: string
type VanityDomainDnsManagedByEnum string

// Set of constants representing the allowable values for VanityDomainDnsManagedByEnum
const (
	VanityDomainDnsManagedByOracleManaged   VanityDomainDnsManagedByEnum = "ORACLE_MANAGED"
	VanityDomainDnsManagedByCustomerManaged VanityDomainDnsManagedByEnum = "CUSTOMER_MANAGED"
)

var mappingVanityDomainDnsManagedByEnum = map[string]VanityDomainDnsManagedByEnum{
	"ORACLE_MANAGED":   VanityDomainDnsManagedByOracleManaged,
	"CUSTOMER_MANAGED": VanityDomainDnsManagedByCustomerManaged,
}

var mappingVanityDomainDnsManagedByEnumLowerCase = map[string]VanityDomainDnsManagedByEnum{
	"oracle_managed":   VanityDomainDnsManagedByOracleManaged,
	"customer_managed": VanityDomainDnsManagedByCustomerManaged,
}

// GetVanityDomainDnsManagedByEnumValues Enumerates the set of values for VanityDomainDnsManagedByEnum
func GetVanityDomainDnsManagedByEnumValues() []VanityDomainDnsManagedByEnum {
	values := make([]VanityDomainDnsManagedByEnum, 0)
	for _, v := range mappingVanityDomainDnsManagedByEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainDnsManagedByEnumStringValues Enumerates the set of values in String for VanityDomainDnsManagedByEnum
func GetVanityDomainDnsManagedByEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED",
	}
}

// GetMappingVanityDomainDnsManagedByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainDnsManagedByEnum(val string) (VanityDomainDnsManagedByEnum, bool) {
	enum, ok := mappingVanityDomainDnsManagedByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainOriginCertStatusEnum Enum with underlying type: string
type VanityDomainOriginCertStatusEnum string

// Set of constants representing the allowable values for VanityDomainOriginCertStatusEnum
const (
	VanityDomainOriginCertStatusActive   VanityDomainOriginCertStatusEnum = "ACTIVE"
	VanityDomainOriginCertStatusInactive VanityDomainOriginCertStatusEnum = "INACTIVE"
	VanityDomainOriginCertStatusExpiring VanityDomainOriginCertStatusEnum = "EXPIRING"
	VanityDomainOriginCertStatusExpired  VanityDomainOriginCertStatusEnum = "EXPIRED"
)

var mappingVanityDomainOriginCertStatusEnum = map[string]VanityDomainOriginCertStatusEnum{
	"ACTIVE":   VanityDomainOriginCertStatusActive,
	"INACTIVE": VanityDomainOriginCertStatusInactive,
	"EXPIRING": VanityDomainOriginCertStatusExpiring,
	"EXPIRED":  VanityDomainOriginCertStatusExpired,
}

var mappingVanityDomainOriginCertStatusEnumLowerCase = map[string]VanityDomainOriginCertStatusEnum{
	"active":   VanityDomainOriginCertStatusActive,
	"inactive": VanityDomainOriginCertStatusInactive,
	"expiring": VanityDomainOriginCertStatusExpiring,
	"expired":  VanityDomainOriginCertStatusExpired,
}

// GetVanityDomainOriginCertStatusEnumValues Enumerates the set of values for VanityDomainOriginCertStatusEnum
func GetVanityDomainOriginCertStatusEnumValues() []VanityDomainOriginCertStatusEnum {
	values := make([]VanityDomainOriginCertStatusEnum, 0)
	for _, v := range mappingVanityDomainOriginCertStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainOriginCertStatusEnumStringValues Enumerates the set of values in String for VanityDomainOriginCertStatusEnum
func GetVanityDomainOriginCertStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"EXPIRING",
		"EXPIRED",
	}
}

// GetMappingVanityDomainOriginCertStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainOriginCertStatusEnum(val string) (VanityDomainOriginCertStatusEnum, bool) {
	enum, ok := mappingVanityDomainOriginCertStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VanityDomainCdnCertStatusEnum Enum with underlying type: string
type VanityDomainCdnCertStatusEnum string

// Set of constants representing the allowable values for VanityDomainCdnCertStatusEnum
const (
	VanityDomainCdnCertStatusActive   VanityDomainCdnCertStatusEnum = "ACTIVE"
	VanityDomainCdnCertStatusInactive VanityDomainCdnCertStatusEnum = "INACTIVE"
	VanityDomainCdnCertStatusExpiring VanityDomainCdnCertStatusEnum = "EXPIRING"
	VanityDomainCdnCertStatusExpired  VanityDomainCdnCertStatusEnum = "EXPIRED"
)

var mappingVanityDomainCdnCertStatusEnum = map[string]VanityDomainCdnCertStatusEnum{
	"ACTIVE":   VanityDomainCdnCertStatusActive,
	"INACTIVE": VanityDomainCdnCertStatusInactive,
	"EXPIRING": VanityDomainCdnCertStatusExpiring,
	"EXPIRED":  VanityDomainCdnCertStatusExpired,
}

var mappingVanityDomainCdnCertStatusEnumLowerCase = map[string]VanityDomainCdnCertStatusEnum{
	"active":   VanityDomainCdnCertStatusActive,
	"inactive": VanityDomainCdnCertStatusInactive,
	"expiring": VanityDomainCdnCertStatusExpiring,
	"expired":  VanityDomainCdnCertStatusExpired,
}

// GetVanityDomainCdnCertStatusEnumValues Enumerates the set of values for VanityDomainCdnCertStatusEnum
func GetVanityDomainCdnCertStatusEnumValues() []VanityDomainCdnCertStatusEnum {
	values := make([]VanityDomainCdnCertStatusEnum, 0)
	for _, v := range mappingVanityDomainCdnCertStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetVanityDomainCdnCertStatusEnumStringValues Enumerates the set of values in String for VanityDomainCdnCertStatusEnum
func GetVanityDomainCdnCertStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"EXPIRING",
		"EXPIRED",
	}
}

// GetMappingVanityDomainCdnCertStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVanityDomainCdnCertStatusEnum(val string) (VanityDomainCdnCertStatusEnum, bool) {
	enum, ok := mappingVanityDomainCdnCertStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
