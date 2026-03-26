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

// EmailSubdomainSummary email subdomain details for a marketing brand
type EmailSubdomainSummary struct {

	// The unique identifier (OCID) of emailsubdomain. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// email subdomain name for a brand
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

	// certification expiration date
	TimeCertificateExpiration *common.SDKTime `mandatory:"true" json:"timeCertificateExpiration"`

	// The time the Email Subdomain was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

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

	// Intermediate state for Email subdomain
	LifecycleDetails EmailSubdomainLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m EmailSubdomainSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailSubdomainSummary) ValidateEnumValue() (bool, error) {
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
