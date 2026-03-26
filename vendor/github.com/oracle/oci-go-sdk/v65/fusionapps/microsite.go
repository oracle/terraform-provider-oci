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

// Microsite microsite details for a marketing brand
type Microsite struct {

	// The unique identifier (OCID) of microsite. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// microsite sudomain name
	Name *string `mandatory:"true" json:"name"`

	// Marketing Brand Identifier
	MarketingBrandId *string `mandatory:"true" json:"marketingBrandId"`

	// Fusion Environment Identifier
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// microsite lifecycle state
	LifecycleState MicrositeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// dns management type for microsite
	DnsManagement MicrositeDnsManagementEnum `mandatory:"true" json:"dnsManagement"`

	// dns status for microsite
	DnsStatus MicrositeDnsStatusEnum `mandatory:"true" json:"dnsStatus"`

	// certificate management type for microsite
	CertificateManagement MicrositeCertificateManagementEnum `mandatory:"true" json:"certificateManagement"`

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

	// microsite intermediate states
	LifecycleDetails MicrositeLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The time the Microsite was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Microsite) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Microsite) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMicrositeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMicrositeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrositeDnsManagementEnum(string(m.DnsManagement)); !ok && m.DnsManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsManagement: %s. Supported values are: %s.", m.DnsManagement, strings.Join(GetMicrositeDnsManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrositeDnsStatusEnum(string(m.DnsStatus)); !ok && m.DnsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnsStatus: %s. Supported values are: %s.", m.DnsStatus, strings.Join(GetMicrositeDnsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrositeCertificateManagementEnum(string(m.CertificateManagement)); !ok && m.CertificateManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateManagement: %s. Supported values are: %s.", m.CertificateManagement, strings.Join(GetMicrositeCertificateManagementEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMicrositeLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMicrositeLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MicrositeLifecycleStateEnum Enum with underlying type: string
type MicrositeLifecycleStateEnum string

// Set of constants representing the allowable values for MicrositeLifecycleStateEnum
const (
	MicrositeLifecycleStateActive   MicrositeLifecycleStateEnum = "ACTIVE"
	MicrositeLifecycleStateInactive MicrositeLifecycleStateEnum = "INACTIVE"
)

var mappingMicrositeLifecycleStateEnum = map[string]MicrositeLifecycleStateEnum{
	"ACTIVE":   MicrositeLifecycleStateActive,
	"INACTIVE": MicrositeLifecycleStateInactive,
}

var mappingMicrositeLifecycleStateEnumLowerCase = map[string]MicrositeLifecycleStateEnum{
	"active":   MicrositeLifecycleStateActive,
	"inactive": MicrositeLifecycleStateInactive,
}

// GetMicrositeLifecycleStateEnumValues Enumerates the set of values for MicrositeLifecycleStateEnum
func GetMicrositeLifecycleStateEnumValues() []MicrositeLifecycleStateEnum {
	values := make([]MicrositeLifecycleStateEnum, 0)
	for _, v := range mappingMicrositeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrositeLifecycleStateEnumStringValues Enumerates the set of values in String for MicrositeLifecycleStateEnum
func GetMicrositeLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingMicrositeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrositeLifecycleStateEnum(val string) (MicrositeLifecycleStateEnum, bool) {
	enum, ok := mappingMicrositeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MicrositeLifecycleDetailsEnum Enum with underlying type: string
type MicrositeLifecycleDetailsEnum string

// Set of constants representing the allowable values for MicrositeLifecycleDetailsEnum
const (
	MicrositeLifecycleDetailsPending                               MicrositeLifecycleDetailsEnum = "PENDING"
	MicrositeLifecycleDetailsNone                                  MicrositeLifecycleDetailsEnum = "NONE"
	MicrositeLifecycleDetailsCreatingTempProperty                  MicrositeLifecycleDetailsEnum = "CREATING_TEMP_PROPERTY"
	MicrositeLifecycleDetailsFailureTempPropertyCreation           MicrositeLifecycleDetailsEnum = "FAILURE_TEMP_PROPERTY_CREATION"
	MicrositeLifecycleDetailsActivatingTempProperty                MicrositeLifecycleDetailsEnum = "ACTIVATING_TEMP_PROPERTY"
	MicrositeLifecycleDetailsFailureTempPropertyActivation         MicrositeLifecycleDetailsEnum = "FAILURE_TEMP_PROPERTY_ACTIVATION"
	MicrositeLifecycleDetailsActivatingHostnameTempProperty        MicrositeLifecycleDetailsEnum = "ACTIVATING_HOSTNAME_TEMP_PROPERTY"
	MicrositeLifecycleDetailsFailureTempPropertyHostnameActivation MicrositeLifecycleDetailsEnum = "FAILURE_TEMP_PROPERTY_HOSTNAME_ACTIVATION"
	MicrositeLifecycleDetailsReadyForMainActivation                MicrositeLifecycleDetailsEnum = "READY_FOR_MAIN_ACTIVATION"
	MicrositeLifecycleDetailsActivatingHostname                    MicrositeLifecycleDetailsEnum = "ACTIVATING_HOSTNAME"
	MicrositeLifecycleDetailsFailureHostnameActivation             MicrositeLifecycleDetailsEnum = "FAILURE_HOSTNAME_ACTIVATION"
	MicrositeLifecycleDetailsSendingBrandEvent                     MicrositeLifecycleDetailsEnum = "SENDING_BRAND_EVENT"
	MicrositeLifecycleDetailsFailureBrandEvent                     MicrositeLifecycleDetailsEnum = "FAILURE_BRAND_EVENT"
)

var mappingMicrositeLifecycleDetailsEnum = map[string]MicrositeLifecycleDetailsEnum{
	"PENDING":                                   MicrositeLifecycleDetailsPending,
	"NONE":                                      MicrositeLifecycleDetailsNone,
	"CREATING_TEMP_PROPERTY":                    MicrositeLifecycleDetailsCreatingTempProperty,
	"FAILURE_TEMP_PROPERTY_CREATION":            MicrositeLifecycleDetailsFailureTempPropertyCreation,
	"ACTIVATING_TEMP_PROPERTY":                  MicrositeLifecycleDetailsActivatingTempProperty,
	"FAILURE_TEMP_PROPERTY_ACTIVATION":          MicrositeLifecycleDetailsFailureTempPropertyActivation,
	"ACTIVATING_HOSTNAME_TEMP_PROPERTY":         MicrositeLifecycleDetailsActivatingHostnameTempProperty,
	"FAILURE_TEMP_PROPERTY_HOSTNAME_ACTIVATION": MicrositeLifecycleDetailsFailureTempPropertyHostnameActivation,
	"READY_FOR_MAIN_ACTIVATION":                 MicrositeLifecycleDetailsReadyForMainActivation,
	"ACTIVATING_HOSTNAME":                       MicrositeLifecycleDetailsActivatingHostname,
	"FAILURE_HOSTNAME_ACTIVATION":               MicrositeLifecycleDetailsFailureHostnameActivation,
	"SENDING_BRAND_EVENT":                       MicrositeLifecycleDetailsSendingBrandEvent,
	"FAILURE_BRAND_EVENT":                       MicrositeLifecycleDetailsFailureBrandEvent,
}

var mappingMicrositeLifecycleDetailsEnumLowerCase = map[string]MicrositeLifecycleDetailsEnum{
	"pending":                                   MicrositeLifecycleDetailsPending,
	"none":                                      MicrositeLifecycleDetailsNone,
	"creating_temp_property":                    MicrositeLifecycleDetailsCreatingTempProperty,
	"failure_temp_property_creation":            MicrositeLifecycleDetailsFailureTempPropertyCreation,
	"activating_temp_property":                  MicrositeLifecycleDetailsActivatingTempProperty,
	"failure_temp_property_activation":          MicrositeLifecycleDetailsFailureTempPropertyActivation,
	"activating_hostname_temp_property":         MicrositeLifecycleDetailsActivatingHostnameTempProperty,
	"failure_temp_property_hostname_activation": MicrositeLifecycleDetailsFailureTempPropertyHostnameActivation,
	"ready_for_main_activation":                 MicrositeLifecycleDetailsReadyForMainActivation,
	"activating_hostname":                       MicrositeLifecycleDetailsActivatingHostname,
	"failure_hostname_activation":               MicrositeLifecycleDetailsFailureHostnameActivation,
	"sending_brand_event":                       MicrositeLifecycleDetailsSendingBrandEvent,
	"failure_brand_event":                       MicrositeLifecycleDetailsFailureBrandEvent,
}

// GetMicrositeLifecycleDetailsEnumValues Enumerates the set of values for MicrositeLifecycleDetailsEnum
func GetMicrositeLifecycleDetailsEnumValues() []MicrositeLifecycleDetailsEnum {
	values := make([]MicrositeLifecycleDetailsEnum, 0)
	for _, v := range mappingMicrositeLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrositeLifecycleDetailsEnumStringValues Enumerates the set of values in String for MicrositeLifecycleDetailsEnum
func GetMicrositeLifecycleDetailsEnumStringValues() []string {
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

// GetMappingMicrositeLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrositeLifecycleDetailsEnum(val string) (MicrositeLifecycleDetailsEnum, bool) {
	enum, ok := mappingMicrositeLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MicrositeDnsManagementEnum Enum with underlying type: string
type MicrositeDnsManagementEnum string

// Set of constants representing the allowable values for MicrositeDnsManagementEnum
const (
	MicrositeDnsManagementCustomerManaged MicrositeDnsManagementEnum = "CUSTOMER_MANAGED"
)

var mappingMicrositeDnsManagementEnum = map[string]MicrositeDnsManagementEnum{
	"CUSTOMER_MANAGED": MicrositeDnsManagementCustomerManaged,
}

var mappingMicrositeDnsManagementEnumLowerCase = map[string]MicrositeDnsManagementEnum{
	"customer_managed": MicrositeDnsManagementCustomerManaged,
}

// GetMicrositeDnsManagementEnumValues Enumerates the set of values for MicrositeDnsManagementEnum
func GetMicrositeDnsManagementEnumValues() []MicrositeDnsManagementEnum {
	values := make([]MicrositeDnsManagementEnum, 0)
	for _, v := range mappingMicrositeDnsManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrositeDnsManagementEnumStringValues Enumerates the set of values in String for MicrositeDnsManagementEnum
func GetMicrositeDnsManagementEnumStringValues() []string {
	return []string{
		"CUSTOMER_MANAGED",
	}
}

// GetMappingMicrositeDnsManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrositeDnsManagementEnum(val string) (MicrositeDnsManagementEnum, bool) {
	enum, ok := mappingMicrositeDnsManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MicrositeDnsStatusEnum Enum with underlying type: string
type MicrositeDnsStatusEnum string

// Set of constants representing the allowable values for MicrositeDnsStatusEnum
const (
	MicrositeDnsStatusActive       MicrositeDnsStatusEnum = "ACTIVE"
	MicrositeDnsStatusNotConfirmed MicrositeDnsStatusEnum = "NOT_CONFIRMED"
)

var mappingMicrositeDnsStatusEnum = map[string]MicrositeDnsStatusEnum{
	"ACTIVE":        MicrositeDnsStatusActive,
	"NOT_CONFIRMED": MicrositeDnsStatusNotConfirmed,
}

var mappingMicrositeDnsStatusEnumLowerCase = map[string]MicrositeDnsStatusEnum{
	"active":        MicrositeDnsStatusActive,
	"not_confirmed": MicrositeDnsStatusNotConfirmed,
}

// GetMicrositeDnsStatusEnumValues Enumerates the set of values for MicrositeDnsStatusEnum
func GetMicrositeDnsStatusEnumValues() []MicrositeDnsStatusEnum {
	values := make([]MicrositeDnsStatusEnum, 0)
	for _, v := range mappingMicrositeDnsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrositeDnsStatusEnumStringValues Enumerates the set of values in String for MicrositeDnsStatusEnum
func GetMicrositeDnsStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NOT_CONFIRMED",
	}
}

// GetMappingMicrositeDnsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrositeDnsStatusEnum(val string) (MicrositeDnsStatusEnum, bool) {
	enum, ok := mappingMicrositeDnsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MicrositeCertificateManagementEnum Enum with underlying type: string
type MicrositeCertificateManagementEnum string

// Set of constants representing the allowable values for MicrositeCertificateManagementEnum
const (
	MicrositeCertificateManagementOracleManaged MicrositeCertificateManagementEnum = "ORACLE_MANAGED"
)

var mappingMicrositeCertificateManagementEnum = map[string]MicrositeCertificateManagementEnum{
	"ORACLE_MANAGED": MicrositeCertificateManagementOracleManaged,
}

var mappingMicrositeCertificateManagementEnumLowerCase = map[string]MicrositeCertificateManagementEnum{
	"oracle_managed": MicrositeCertificateManagementOracleManaged,
}

// GetMicrositeCertificateManagementEnumValues Enumerates the set of values for MicrositeCertificateManagementEnum
func GetMicrositeCertificateManagementEnumValues() []MicrositeCertificateManagementEnum {
	values := make([]MicrositeCertificateManagementEnum, 0)
	for _, v := range mappingMicrositeCertificateManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetMicrositeCertificateManagementEnumStringValues Enumerates the set of values in String for MicrositeCertificateManagementEnum
func GetMicrositeCertificateManagementEnumStringValues() []string {
	return []string{
		"ORACLE_MANAGED",
	}
}

// GetMappingMicrositeCertificateManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMicrositeCertificateManagementEnum(val string) (MicrositeCertificateManagementEnum, bool) {
	enum, ok := mappingMicrositeCertificateManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
