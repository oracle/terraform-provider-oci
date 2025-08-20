// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityPolicyEntryStateSummary The resource represents the state of a specific entry type deployment on a target.
type SecurityPolicyEntryStateSummary struct {

	// Unique id of the security policy entry state.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the security policy entry associated.
	SecurityPolicyEntryId *string `mandatory:"true" json:"securityPolicyEntryId"`

	// The OCID of the target on which the security policy is deployed.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The security policy entry type. Allowed values:
	// - FIREWALL_POLICY - The SQL Firewall policy entry type.
	// - AUDIT_POLICY - The audit policy entry type.
	// - CONFIG - Config changes deployment.
	EntryType SecurityPolicyEntryStateSummaryEntryTypeEnum `mandatory:"true" json:"entryType"`

	// The current deployment status of the security policy deployment and the security policy entry associated.
	DeploymentStatus SecurityPolicyEntryStateDeploymentStatusEnum `mandatory:"true" json:"deploymentStatus"`

	// The OCID of the security policy deployment associated.
	SecurityPolicyDeploymentId *string `mandatory:"false" json:"securityPolicyDeploymentId"`

	// Details about the current deployment status.
	DeploymentStatusDetails *string `mandatory:"false" json:"deploymentStatusDetails"`
}

func (m SecurityPolicyEntryStateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityPolicyEntryStateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityPolicyEntryStateSummaryEntryTypeEnum(string(m.EntryType)); !ok && m.EntryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntryType: %s. Supported values are: %s.", m.EntryType, strings.Join(GetSecurityPolicyEntryStateSummaryEntryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityPolicyEntryStateDeploymentStatusEnum(string(m.DeploymentStatus)); !ok && m.DeploymentStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentStatus: %s. Supported values are: %s.", m.DeploymentStatus, strings.Join(GetSecurityPolicyEntryStateDeploymentStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityPolicyEntryStateSummaryEntryTypeEnum Enum with underlying type: string
type SecurityPolicyEntryStateSummaryEntryTypeEnum string

// Set of constants representing the allowable values for SecurityPolicyEntryStateSummaryEntryTypeEnum
const (
	SecurityPolicyEntryStateSummaryEntryTypeFirewallPolicy SecurityPolicyEntryStateSummaryEntryTypeEnum = "FIREWALL_POLICY"
	SecurityPolicyEntryStateSummaryEntryTypeAuditPolicy    SecurityPolicyEntryStateSummaryEntryTypeEnum = "AUDIT_POLICY"
	SecurityPolicyEntryStateSummaryEntryTypeConfig         SecurityPolicyEntryStateSummaryEntryTypeEnum = "CONFIG"
)

var mappingSecurityPolicyEntryStateSummaryEntryTypeEnum = map[string]SecurityPolicyEntryStateSummaryEntryTypeEnum{
	"FIREWALL_POLICY": SecurityPolicyEntryStateSummaryEntryTypeFirewallPolicy,
	"AUDIT_POLICY":    SecurityPolicyEntryStateSummaryEntryTypeAuditPolicy,
	"CONFIG":          SecurityPolicyEntryStateSummaryEntryTypeConfig,
}

var mappingSecurityPolicyEntryStateSummaryEntryTypeEnumLowerCase = map[string]SecurityPolicyEntryStateSummaryEntryTypeEnum{
	"firewall_policy": SecurityPolicyEntryStateSummaryEntryTypeFirewallPolicy,
	"audit_policy":    SecurityPolicyEntryStateSummaryEntryTypeAuditPolicy,
	"config":          SecurityPolicyEntryStateSummaryEntryTypeConfig,
}

// GetSecurityPolicyEntryStateSummaryEntryTypeEnumValues Enumerates the set of values for SecurityPolicyEntryStateSummaryEntryTypeEnum
func GetSecurityPolicyEntryStateSummaryEntryTypeEnumValues() []SecurityPolicyEntryStateSummaryEntryTypeEnum {
	values := make([]SecurityPolicyEntryStateSummaryEntryTypeEnum, 0)
	for _, v := range mappingSecurityPolicyEntryStateSummaryEntryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityPolicyEntryStateSummaryEntryTypeEnumStringValues Enumerates the set of values in String for SecurityPolicyEntryStateSummaryEntryTypeEnum
func GetSecurityPolicyEntryStateSummaryEntryTypeEnumStringValues() []string {
	return []string{
		"FIREWALL_POLICY",
		"AUDIT_POLICY",
		"CONFIG",
	}
}

// GetMappingSecurityPolicyEntryStateSummaryEntryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityPolicyEntryStateSummaryEntryTypeEnum(val string) (SecurityPolicyEntryStateSummaryEntryTypeEnum, bool) {
	enum, ok := mappingSecurityPolicyEntryStateSummaryEntryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
