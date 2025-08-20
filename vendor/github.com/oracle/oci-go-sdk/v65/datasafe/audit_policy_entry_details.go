// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuditPolicyEntryDetails Audit policy details.
type AuditPolicyEntryDetails struct {

	// Specifies why exclusion of the Data Safe user did not succeed.
	ExcludeDatasafeUserFailureMsg *string `mandatory:"false" json:"excludeDatasafeUserFailureMsg"`

	// The status of Data Safe user exclusion in the audit policy.
	DatasafeUserExclusionStatus AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum `mandatory:"true" json:"datasafeUserExclusionStatus"`
}

func (m AuditPolicyEntryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditPolicyEntryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum(string(m.DatasafeUserExclusionStatus)); !ok && m.DatasafeUserExclusionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatasafeUserExclusionStatus: %s. Supported values are: %s.", m.DatasafeUserExclusionStatus, strings.Join(GetAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AuditPolicyEntryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAuditPolicyEntryDetails AuditPolicyEntryDetails
	s := struct {
		DiscriminatorParam string `json:"entryType"`
		MarshalTypeAuditPolicyEntryDetails
	}{
		"AUDIT_POLICY",
		(MarshalTypeAuditPolicyEntryDetails)(m),
	}

	return json.Marshal(&s)
}

// AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum Enum with underlying type: string
type AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum string

// Set of constants representing the allowable values for AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum
const (
	AuditPolicyEntryDetailsDatasafeUserExclusionStatusExcludedSuccess AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum = "EXCLUDED_SUCCESS"
	AuditPolicyEntryDetailsDatasafeUserExclusionStatusExcludedFailed  AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum = "EXCLUDED_FAILED"
	AuditPolicyEntryDetailsDatasafeUserExclusionStatusNotExcluded     AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum = "NOT_EXCLUDED"
)

var mappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum = map[string]AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum{
	"EXCLUDED_SUCCESS": AuditPolicyEntryDetailsDatasafeUserExclusionStatusExcludedSuccess,
	"EXCLUDED_FAILED":  AuditPolicyEntryDetailsDatasafeUserExclusionStatusExcludedFailed,
	"NOT_EXCLUDED":     AuditPolicyEntryDetailsDatasafeUserExclusionStatusNotExcluded,
}

var mappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumLowerCase = map[string]AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum{
	"excluded_success": AuditPolicyEntryDetailsDatasafeUserExclusionStatusExcludedSuccess,
	"excluded_failed":  AuditPolicyEntryDetailsDatasafeUserExclusionStatusExcludedFailed,
	"not_excluded":     AuditPolicyEntryDetailsDatasafeUserExclusionStatusNotExcluded,
}

// GetAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumValues Enumerates the set of values for AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum
func GetAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumValues() []AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum {
	values := make([]AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum, 0)
	for _, v := range mappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumStringValues Enumerates the set of values in String for AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum
func GetAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumStringValues() []string {
	return []string{
		"EXCLUDED_SUCCESS",
		"EXCLUDED_FAILED",
		"NOT_EXCLUDED",
	}
}

// GetMappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum(val string) (AuditPolicyEntryDetailsDatasafeUserExclusionStatusEnum, bool) {
	enum, ok := mappingAuditPolicyEntryDetailsDatasafeUserExclusionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
