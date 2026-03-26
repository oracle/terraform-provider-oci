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

// MicrositeSummary microsite details for a marketing brand
type MicrositeSummary struct {

	// The unique identifier (OCID) of microsite. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// microsite name
	Name *string `mandatory:"true" json:"name"`

	// Marketing Brand Identifier
	MarketingBrandId *string `mandatory:"true" json:"marketingBrandId"`

	// Fusion Environment Identifier
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// microsite lifecyclestate
	LifecycleState MicrositeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// dns management type for microsite
	DnsManagement MicrositeDnsManagementEnum `mandatory:"true" json:"dnsManagement"`

	// dns status for microsite
	DnsStatus MicrositeDnsStatusEnum `mandatory:"true" json:"dnsStatus"`

	// The time the Microsite was created. An RFC3339 formatted datetime string
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

	// Intermediate state for microsite
	LifecycleDetails MicrositeLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// certificate type for microsite
	CertificateManagement MicrositeCertificateManagementEnum `mandatory:"false" json:"certificateManagement,omitempty"`
}

func (m MicrositeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MicrositeSummary) ValidateEnumValue() (bool, error) {
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

	if _, ok := GetMappingMicrositeLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMicrositeLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMicrositeCertificateManagementEnum(string(m.CertificateManagement)); !ok && m.CertificateManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateManagement: %s. Supported values are: %s.", m.CertificateManagement, strings.Join(GetMicrositeCertificateManagementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
