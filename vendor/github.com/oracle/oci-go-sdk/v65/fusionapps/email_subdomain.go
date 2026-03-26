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

// EmailSubdomain email subdomain details for a marketing brand
type EmailSubdomain struct {

	// The unique identifier (OCID) of emailsubdomain. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// The name for email subdomain for a brand
	Name *string `mandatory:"true" json:"name"`

	// Marketing Brand Identifier
	MarketingBrandId *string `mandatory:"true" json:"marketingBrandId"`

	// Fusion Environment Identifier
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// email subdomain lifecyclestate
	LifecycleState EmailSubdomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// dns management type for email subdomain
	DnsManagement EmailSubdomainDnsManagementEnum `mandatory:"true" json:"dnsManagement"`

	// dns status for email subdomain
	DnsStatus EmailSubdomainDnsStatusEnum `mandatory:"true" json:"dnsStatus"`

	// dns management type for email subdomain
	CertificateManagement EmailSubdomainCertificateManagementEnum `mandatory:"true" json:"certificateManagement"`

	// certificate status for email subdomain
	CertificateStatus EmailSubdomainCertificateStatusEnum `mandatory:"true" json:"certificateStatus"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Email subdomain intermediate states
	LifecycleDetails EmailSubdomainLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// certification expiration date
	TimeCertificateExpiration *common.SDKTime `mandatory:"false" json:"timeCertificateExpiration"`

	// The time the Email Subdomain was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m EmailSubdomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailSubdomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailSubdomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailSubdomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailSubdomainDnsManagementEnum(string(m.DnsManagement)); !ok && m.DnsManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsManagement: %s. Supported values are: %s.", m.DnsManagement, strings.Join(GetEmailSubdomainDnsManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailSubdomainDnsStatusEnum(string(m.DnsStatus)); !ok && m.DnsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsStatus: %s. Supported values are: %s.", m.DnsStatus, strings.Join(GetEmailSubdomainDnsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailSubdomainCertificateManagementEnum(string(m.CertificateManagement)); !ok && m.CertificateManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateManagement: %s. Supported values are: %s.", m.CertificateManagement, strings.Join(GetEmailSubdomainCertificateManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailSubdomainCertificateStatusEnum(string(m.CertificateStatus)); !ok && m.CertificateStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateStatus: %s. Supported values are: %s.", m.CertificateStatus, strings.Join(GetEmailSubdomainCertificateStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEmailSubdomainLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetEmailSubdomainLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailSubdomainLifecycleStateEnum Enum with underlying type: string
type EmailSubdomainLifecycleStateEnum string

// Set of constants representing the allowable values for EmailSubdomainLifecycleStateEnum
const (
	EmailSubdomainLifecycleStateActive   EmailSubdomainLifecycleStateEnum = "ACTIVE"
	EmailSubdomainLifecycleStateInactive EmailSubdomainLifecycleStateEnum = "INACTIVE"
)

var mappingEmailSubdomainLifecycleStateEnum = map[string]EmailSubdomainLifecycleStateEnum{
	"ACTIVE":   EmailSubdomainLifecycleStateActive,
	"INACTIVE": EmailSubdomainLifecycleStateInactive,
}

var mappingEmailSubdomainLifecycleStateEnumLowerCase = map[string]EmailSubdomainLifecycleStateEnum{
	"active":   EmailSubdomainLifecycleStateActive,
	"inactive": EmailSubdomainLifecycleStateInactive,
}

// GetEmailSubdomainLifecycleStateEnumValues Enumerates the set of values for EmailSubdomainLifecycleStateEnum
func GetEmailSubdomainLifecycleStateEnumValues() []EmailSubdomainLifecycleStateEnum {
	values := make([]EmailSubdomainLifecycleStateEnum, 0)
	for _, v := range mappingEmailSubdomainLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailSubdomainLifecycleStateEnumStringValues Enumerates the set of values in String for EmailSubdomainLifecycleStateEnum
func GetEmailSubdomainLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingEmailSubdomainLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailSubdomainLifecycleStateEnum(val string) (EmailSubdomainLifecycleStateEnum, bool) {
	enum, ok := mappingEmailSubdomainLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailSubdomainLifecycleDetailsEnum Enum with underlying type: string
type EmailSubdomainLifecycleDetailsEnum string

// Set of constants representing the allowable values for EmailSubdomainLifecycleDetailsEnum
const (
	EmailSubdomainLifecycleDetailsPending                               EmailSubdomainLifecycleDetailsEnum = "PENDING"
	EmailSubdomainLifecycleDetailsNone                                  EmailSubdomainLifecycleDetailsEnum = "NONE"
	EmailSubdomainLifecycleDetailsCreatingTempProperty                  EmailSubdomainLifecycleDetailsEnum = "CREATING_TEMP_PROPERTY"
	EmailSubdomainLifecycleDetailsFailureTempPropertyCreation           EmailSubdomainLifecycleDetailsEnum = "FAILURE_TEMP_PROPERTY_CREATION"
	EmailSubdomainLifecycleDetailsActivatingTempProperty                EmailSubdomainLifecycleDetailsEnum = "ACTIVATING_TEMP_PROPERTY"
	EmailSubdomainLifecycleDetailsFailureTempPropertyActivation         EmailSubdomainLifecycleDetailsEnum = "FAILURE_TEMP_PROPERTY_ACTIVATION"
	EmailSubdomainLifecycleDetailsActivatingHostnameTempProperty        EmailSubdomainLifecycleDetailsEnum = "ACTIVATING_HOSTNAME_TEMP_PROPERTY"
	EmailSubdomainLifecycleDetailsFailureTempPropertyHostnameActivation EmailSubdomainLifecycleDetailsEnum = "FAILURE_TEMP_PROPERTY_HOSTNAME_ACTIVATION"
	EmailSubdomainLifecycleDetailsReadyForMainActivation                EmailSubdomainLifecycleDetailsEnum = "READY_FOR_MAIN_ACTIVATION"
	EmailSubdomainLifecycleDetailsActivatingHostname                    EmailSubdomainLifecycleDetailsEnum = "ACTIVATING_HOSTNAME"
	EmailSubdomainLifecycleDetailsFailureHostnameActivation             EmailSubdomainLifecycleDetailsEnum = "FAILURE_HOSTNAME_ACTIVATION"
	EmailSubdomainLifecycleDetailsSendingBrandEvent                     EmailSubdomainLifecycleDetailsEnum = "SENDING_BRAND_EVENT"
	EmailSubdomainLifecycleDetailsFailureBrandEvent                     EmailSubdomainLifecycleDetailsEnum = "FAILURE_BRAND_EVENT"
)

var mappingEmailSubdomainLifecycleDetailsEnum = map[string]EmailSubdomainLifecycleDetailsEnum{
	"PENDING":                                   EmailSubdomainLifecycleDetailsPending,
	"NONE":                                      EmailSubdomainLifecycleDetailsNone,
	"CREATING_TEMP_PROPERTY":                    EmailSubdomainLifecycleDetailsCreatingTempProperty,
	"FAILURE_TEMP_PROPERTY_CREATION":            EmailSubdomainLifecycleDetailsFailureTempPropertyCreation,
	"ACTIVATING_TEMP_PROPERTY":                  EmailSubdomainLifecycleDetailsActivatingTempProperty,
	"FAILURE_TEMP_PROPERTY_ACTIVATION":          EmailSubdomainLifecycleDetailsFailureTempPropertyActivation,
	"ACTIVATING_HOSTNAME_TEMP_PROPERTY":         EmailSubdomainLifecycleDetailsActivatingHostnameTempProperty,
	"FAILURE_TEMP_PROPERTY_HOSTNAME_ACTIVATION": EmailSubdomainLifecycleDetailsFailureTempPropertyHostnameActivation,
	"READY_FOR_MAIN_ACTIVATION":                 EmailSubdomainLifecycleDetailsReadyForMainActivation,
	"ACTIVATING_HOSTNAME":                       EmailSubdomainLifecycleDetailsActivatingHostname,
	"FAILURE_HOSTNAME_ACTIVATION":               EmailSubdomainLifecycleDetailsFailureHostnameActivation,
	"SENDING_BRAND_EVENT":                       EmailSubdomainLifecycleDetailsSendingBrandEvent,
	"FAILURE_BRAND_EVENT":                       EmailSubdomainLifecycleDetailsFailureBrandEvent,
}

var mappingEmailSubdomainLifecycleDetailsEnumLowerCase = map[string]EmailSubdomainLifecycleDetailsEnum{
	"pending":                                   EmailSubdomainLifecycleDetailsPending,
	"none":                                      EmailSubdomainLifecycleDetailsNone,
	"creating_temp_property":                    EmailSubdomainLifecycleDetailsCreatingTempProperty,
	"failure_temp_property_creation":            EmailSubdomainLifecycleDetailsFailureTempPropertyCreation,
	"activating_temp_property":                  EmailSubdomainLifecycleDetailsActivatingTempProperty,
	"failure_temp_property_activation":          EmailSubdomainLifecycleDetailsFailureTempPropertyActivation,
	"activating_hostname_temp_property":         EmailSubdomainLifecycleDetailsActivatingHostnameTempProperty,
	"failure_temp_property_hostname_activation": EmailSubdomainLifecycleDetailsFailureTempPropertyHostnameActivation,
	"ready_for_main_activation":                 EmailSubdomainLifecycleDetailsReadyForMainActivation,
	"activating_hostname":                       EmailSubdomainLifecycleDetailsActivatingHostname,
	"failure_hostname_activation":               EmailSubdomainLifecycleDetailsFailureHostnameActivation,
	"sending_brand_event":                       EmailSubdomainLifecycleDetailsSendingBrandEvent,
	"failure_brand_event":                       EmailSubdomainLifecycleDetailsFailureBrandEvent,
}

// GetEmailSubdomainLifecycleDetailsEnumValues Enumerates the set of values for EmailSubdomainLifecycleDetailsEnum
func GetEmailSubdomainLifecycleDetailsEnumValues() []EmailSubdomainLifecycleDetailsEnum {
	values := make([]EmailSubdomainLifecycleDetailsEnum, 0)
	for _, v := range mappingEmailSubdomainLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailSubdomainLifecycleDetailsEnumStringValues Enumerates the set of values in String for EmailSubdomainLifecycleDetailsEnum
func GetEmailSubdomainLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PENDING",
		"NONE",
		"CREATING_TEMP_PROPERTY",
		"FAILURE_TEMP_PROPERTY_CREATION",
		"ACTIVATING_TEMP_PROPERTY",
		"FAILURE_TEMP_PROPERTY_ACTIVATION",
		"ACTIVATING_HOSTNAME_TEMP_PROPERTY",
		"FAILURE_TEMP_PROPERTY_HOSTNAME_ACTIVATION",
		"READY_FOR_MAIN_ACTIVATION",
		"ACTIVATING_HOSTNAME",
		"FAILURE_HOSTNAME_ACTIVATION",
		"SENDING_BRAND_EVENT",
		"FAILURE_BRAND_EVENT",
	}
}

// GetMappingEmailSubdomainLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailSubdomainLifecycleDetailsEnum(val string) (EmailSubdomainLifecycleDetailsEnum, bool) {
	enum, ok := mappingEmailSubdomainLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailSubdomainDnsManagementEnum Enum with underlying type: string
type EmailSubdomainDnsManagementEnum string

// Set of constants representing the allowable values for EmailSubdomainDnsManagementEnum
const (
	EmailSubdomainDnsManagementOracleManaged   EmailSubdomainDnsManagementEnum = "ORACLE_MANAGED"
	EmailSubdomainDnsManagementCustomerManaged EmailSubdomainDnsManagementEnum = "CUSTOMER_MANAGED"
)

var mappingEmailSubdomainDnsManagementEnum = map[string]EmailSubdomainDnsManagementEnum{
	"ORACLE_MANAGED":   EmailSubdomainDnsManagementOracleManaged,
	"CUSTOMER_MANAGED": EmailSubdomainDnsManagementCustomerManaged,
}

var mappingEmailSubdomainDnsManagementEnumLowerCase = map[string]EmailSubdomainDnsManagementEnum{
	"oracle_managed":   EmailSubdomainDnsManagementOracleManaged,
	"customer_managed": EmailSubdomainDnsManagementCustomerManaged,
}

// GetEmailSubdomainDnsManagementEnumValues Enumerates the set of values for EmailSubdomainDnsManagementEnum
func GetEmailSubdomainDnsManagementEnumValues() []EmailSubdomainDnsManagementEnum {
	values := make([]EmailSubdomainDnsManagementEnum, 0)
	for _, v := range mappingEmailSubdomainDnsManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailSubdomainDnsManagementEnumStringValues Enumerates the set of values in String for EmailSubdomainDnsManagementEnum
func GetEmailSubdomainDnsManagementEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED",
	}
}

// GetMappingEmailSubdomainDnsManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailSubdomainDnsManagementEnum(val string) (EmailSubdomainDnsManagementEnum, bool) {
	enum, ok := mappingEmailSubdomainDnsManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailSubdomainDnsStatusEnum Enum with underlying type: string
type EmailSubdomainDnsStatusEnum string

// Set of constants representing the allowable values for EmailSubdomainDnsStatusEnum
const (
	EmailSubdomainDnsStatusActive       EmailSubdomainDnsStatusEnum = "ACTIVE"
	EmailSubdomainDnsStatusNotConfirmed EmailSubdomainDnsStatusEnum = "NOT_CONFIRMED"
	EmailSubdomainDnsStatusPending      EmailSubdomainDnsStatusEnum = "PENDING"
)

var mappingEmailSubdomainDnsStatusEnum = map[string]EmailSubdomainDnsStatusEnum{
	"ACTIVE":        EmailSubdomainDnsStatusActive,
	"NOT_CONFIRMED": EmailSubdomainDnsStatusNotConfirmed,
	"PENDING":       EmailSubdomainDnsStatusPending,
}

var mappingEmailSubdomainDnsStatusEnumLowerCase = map[string]EmailSubdomainDnsStatusEnum{
	"active":        EmailSubdomainDnsStatusActive,
	"not_confirmed": EmailSubdomainDnsStatusNotConfirmed,
	"pending":       EmailSubdomainDnsStatusPending,
}

// GetEmailSubdomainDnsStatusEnumValues Enumerates the set of values for EmailSubdomainDnsStatusEnum
func GetEmailSubdomainDnsStatusEnumValues() []EmailSubdomainDnsStatusEnum {
	values := make([]EmailSubdomainDnsStatusEnum, 0)
	for _, v := range mappingEmailSubdomainDnsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailSubdomainDnsStatusEnumStringValues Enumerates the set of values in String for EmailSubdomainDnsStatusEnum
func GetEmailSubdomainDnsStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NOT_CONFIRMED",
		"PENDING",
	}
}

// GetMappingEmailSubdomainDnsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailSubdomainDnsStatusEnum(val string) (EmailSubdomainDnsStatusEnum, bool) {
	enum, ok := mappingEmailSubdomainDnsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailSubdomainCertificateManagementEnum Enum with underlying type: string
type EmailSubdomainCertificateManagementEnum string

// Set of constants representing the allowable values for EmailSubdomainCertificateManagementEnum
const (
	EmailSubdomainCertificateManagementOracleManaged   EmailSubdomainCertificateManagementEnum = "ORACLE_MANAGED"
	EmailSubdomainCertificateManagementCustomerManaged EmailSubdomainCertificateManagementEnum = "CUSTOMER_MANAGED"
)

var mappingEmailSubdomainCertificateManagementEnum = map[string]EmailSubdomainCertificateManagementEnum{
	"ORACLE_MANAGED":   EmailSubdomainCertificateManagementOracleManaged,
	"CUSTOMER_MANAGED": EmailSubdomainCertificateManagementCustomerManaged,
}

var mappingEmailSubdomainCertificateManagementEnumLowerCase = map[string]EmailSubdomainCertificateManagementEnum{
	"oracle_managed":   EmailSubdomainCertificateManagementOracleManaged,
	"customer_managed": EmailSubdomainCertificateManagementCustomerManaged,
}

// GetEmailSubdomainCertificateManagementEnumValues Enumerates the set of values for EmailSubdomainCertificateManagementEnum
func GetEmailSubdomainCertificateManagementEnumValues() []EmailSubdomainCertificateManagementEnum {
	values := make([]EmailSubdomainCertificateManagementEnum, 0)
	for _, v := range mappingEmailSubdomainCertificateManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailSubdomainCertificateManagementEnumStringValues Enumerates the set of values in String for EmailSubdomainCertificateManagementEnum
func GetEmailSubdomainCertificateManagementEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
		"CUSTOMER_MANAGED",
	}
}

// GetMappingEmailSubdomainCertificateManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailSubdomainCertificateManagementEnum(val string) (EmailSubdomainCertificateManagementEnum, bool) {
	enum, ok := mappingEmailSubdomainCertificateManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailSubdomainCertificateStatusEnum Enum with underlying type: string
type EmailSubdomainCertificateStatusEnum string

// Set of constants representing the allowable values for EmailSubdomainCertificateStatusEnum
const (
	EmailSubdomainCertificateStatusActive       EmailSubdomainCertificateStatusEnum = "ACTIVE"
	EmailSubdomainCertificateStatusNotConfirmed EmailSubdomainCertificateStatusEnum = "NOT_CONFIRMED"
	EmailSubdomainCertificateStatusExpired      EmailSubdomainCertificateStatusEnum = "EXPIRED"
)

var mappingEmailSubdomainCertificateStatusEnum = map[string]EmailSubdomainCertificateStatusEnum{
	"ACTIVE":        EmailSubdomainCertificateStatusActive,
	"NOT_CONFIRMED": EmailSubdomainCertificateStatusNotConfirmed,
	"EXPIRED":       EmailSubdomainCertificateStatusExpired,
}

var mappingEmailSubdomainCertificateStatusEnumLowerCase = map[string]EmailSubdomainCertificateStatusEnum{
	"active":        EmailSubdomainCertificateStatusActive,
	"not_confirmed": EmailSubdomainCertificateStatusNotConfirmed,
	"expired":       EmailSubdomainCertificateStatusExpired,
}

// GetEmailSubdomainCertificateStatusEnumValues Enumerates the set of values for EmailSubdomainCertificateStatusEnum
func GetEmailSubdomainCertificateStatusEnumValues() []EmailSubdomainCertificateStatusEnum {
	values := make([]EmailSubdomainCertificateStatusEnum, 0)
	for _, v := range mappingEmailSubdomainCertificateStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailSubdomainCertificateStatusEnumStringValues Enumerates the set of values in String for EmailSubdomainCertificateStatusEnum
func GetEmailSubdomainCertificateStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NOT_CONFIRMED",
		"EXPIRED",
	}
}

// GetMappingEmailSubdomainCertificateStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailSubdomainCertificateStatusEnum(val string) (EmailSubdomainCertificateStatusEnum, bool) {
	enum, ok := mappingEmailSubdomainCertificateStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
