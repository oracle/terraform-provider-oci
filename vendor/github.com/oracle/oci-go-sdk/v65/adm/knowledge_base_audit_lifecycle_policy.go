// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// KnowledgeBaseAuditLifecyclePolicy A lifecycle policy that defines the default lifecycle policy for audits in the knowledge base.
type KnowledgeBaseAuditLifecyclePolicy interface {
}

type knowledgebaseauditlifecyclepolicy struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *knowledgebaseauditlifecyclepolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerknowledgebaseauditlifecyclepolicy knowledgebaseauditlifecyclepolicy
	s := struct {
		Model Unmarshalerknowledgebaseauditlifecyclepolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *knowledgebaseauditlifecyclepolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RETAIN":
		mm := RetainKnowledgeBaseAuditLifecyclePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELETE_AFTER":
		mm := DeleteAfterKnowledgeBaseAuditLifecyclePolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for KnowledgeBaseAuditLifecyclePolicy: %s.", m.Type)
		return *m, nil
	}
}

func (m knowledgebaseauditlifecyclepolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m knowledgebaseauditlifecyclepolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KnowledgeBaseAuditLifecyclePolicyTypeEnum Enum with underlying type: string
type KnowledgeBaseAuditLifecyclePolicyTypeEnum string

// Set of constants representing the allowable values for KnowledgeBaseAuditLifecyclePolicyTypeEnum
const (
	KnowledgeBaseAuditLifecyclePolicyTypeRetain      KnowledgeBaseAuditLifecyclePolicyTypeEnum = "RETAIN"
	KnowledgeBaseAuditLifecyclePolicyTypeDeleteAfter KnowledgeBaseAuditLifecyclePolicyTypeEnum = "DELETE_AFTER"
)

var mappingKnowledgeBaseAuditLifecyclePolicyTypeEnum = map[string]KnowledgeBaseAuditLifecyclePolicyTypeEnum{
	"RETAIN":       KnowledgeBaseAuditLifecyclePolicyTypeRetain,
	"DELETE_AFTER": KnowledgeBaseAuditLifecyclePolicyTypeDeleteAfter,
}

var mappingKnowledgeBaseAuditLifecyclePolicyTypeEnumLowerCase = map[string]KnowledgeBaseAuditLifecyclePolicyTypeEnum{
	"retain":       KnowledgeBaseAuditLifecyclePolicyTypeRetain,
	"delete_after": KnowledgeBaseAuditLifecyclePolicyTypeDeleteAfter,
}

// GetKnowledgeBaseAuditLifecyclePolicyTypeEnumValues Enumerates the set of values for KnowledgeBaseAuditLifecyclePolicyTypeEnum
func GetKnowledgeBaseAuditLifecyclePolicyTypeEnumValues() []KnowledgeBaseAuditLifecyclePolicyTypeEnum {
	values := make([]KnowledgeBaseAuditLifecyclePolicyTypeEnum, 0)
	for _, v := range mappingKnowledgeBaseAuditLifecyclePolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKnowledgeBaseAuditLifecyclePolicyTypeEnumStringValues Enumerates the set of values in String for KnowledgeBaseAuditLifecyclePolicyTypeEnum
func GetKnowledgeBaseAuditLifecyclePolicyTypeEnumStringValues() []string {
	return []string{
		"RETAIN",
		"DELETE_AFTER",
	}
}

// GetMappingKnowledgeBaseAuditLifecyclePolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKnowledgeBaseAuditLifecyclePolicyTypeEnum(val string) (KnowledgeBaseAuditLifecyclePolicyTypeEnum, bool) {
	enum, ok := mappingKnowledgeBaseAuditLifecyclePolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
