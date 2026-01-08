// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuditLifecyclePolicy Lifecycle policy definition for audits.
type AuditLifecyclePolicy interface {
}

type auditlifecyclepolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *auditlifecyclepolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauditlifecyclepolicy auditlifecyclepolicy
	s := struct {
		Model Unmarshalerauditlifecyclepolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *auditlifecyclepolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RETAIN":
		mm := RetainAuditLifecyclePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USE_KB_LIFECYCLE_POLICY":
		mm := UseKbPolicyAuditLifecyclePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELETE_AFTER":
		mm := DeleteAfterAuditLifecyclePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AuditLifecyclePolicy: %s.", m.Type)
		return *m, nil
	}
}

func (m auditlifecyclepolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m auditlifecyclepolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuditLifecyclePolicyTypeEnum Enum with underlying type: string
type AuditLifecyclePolicyTypeEnum string

// Set of constants representing the allowable values for AuditLifecyclePolicyTypeEnum
const (
	AuditLifecyclePolicyTypeRetain               AuditLifecyclePolicyTypeEnum = "RETAIN"
	AuditLifecyclePolicyTypeDeleteAfter          AuditLifecyclePolicyTypeEnum = "DELETE_AFTER"
	AuditLifecyclePolicyTypeUseKbLifecyclePolicy AuditLifecyclePolicyTypeEnum = "USE_KB_LIFECYCLE_POLICY"
)

var mappingAuditLifecyclePolicyTypeEnum = map[string]AuditLifecyclePolicyTypeEnum{
	"RETAIN":                  AuditLifecyclePolicyTypeRetain,
	"DELETE_AFTER":            AuditLifecyclePolicyTypeDeleteAfter,
	"USE_KB_LIFECYCLE_POLICY": AuditLifecyclePolicyTypeUseKbLifecyclePolicy,
}

var mappingAuditLifecyclePolicyTypeEnumLowerCase = map[string]AuditLifecyclePolicyTypeEnum{
	"retain":                  AuditLifecyclePolicyTypeRetain,
	"delete_after":            AuditLifecyclePolicyTypeDeleteAfter,
	"use_kb_lifecycle_policy": AuditLifecyclePolicyTypeUseKbLifecyclePolicy,
}

// GetAuditLifecyclePolicyTypeEnumValues Enumerates the set of values for AuditLifecyclePolicyTypeEnum
func GetAuditLifecyclePolicyTypeEnumValues() []AuditLifecyclePolicyTypeEnum {
	values := make([]AuditLifecyclePolicyTypeEnum, 0)
	for _, v := range mappingAuditLifecyclePolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditLifecyclePolicyTypeEnumStringValues Enumerates the set of values in String for AuditLifecyclePolicyTypeEnum
func GetAuditLifecyclePolicyTypeEnumStringValues() []string {
	return []string{
		"RETAIN",
		"DELETE_AFTER",
		"USE_KB_LIFECYCLE_POLICY",
	}
}

// GetMappingAuditLifecyclePolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditLifecyclePolicyTypeEnum(val string) (AuditLifecyclePolicyTypeEnum, bool) {
	enum, ok := mappingAuditLifecyclePolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
