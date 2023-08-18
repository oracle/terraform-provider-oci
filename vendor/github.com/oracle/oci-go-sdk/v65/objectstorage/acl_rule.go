// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AclRule Specifies whether certain types of network traffic should be ALLOWED or DENIED. ACLs contain a list of ACL Rules
// to restrict the networks from which a request can originate.
type AclRule struct {

	// The action to perform if the incoming request matches all of this rule's restrictions. FORCE DENY will cause
	// all matching requests to fail, overriding any ALLOW rules at any level (tenancy, compartment, or bucket).
	Action AclRuleActionEnum `mandatory:"true" json:"action"`

	NetworkSource NetworkSource `mandatory:"true" json:"networkSource"`

	// Specifies whether this rule applies to ALL operations, or only READ or WRITE operations.
	Operation AclRuleOperationEnum `mandatory:"true" json:"operation"`

	// This rule only applies to objects in buckets with the specified tag(s).
	BucketTags []string `mandatory:"false" json:"bucketTags"`

	BucketNameFilter *BucketNameFilter `mandatory:"false" json:"bucketNameFilter"`

	ObjectNameFilter *ObjectNameFilter `mandatory:"false" json:"objectNameFilter"`
}

func (m AclRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AclRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAclRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetAclRuleActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAclRuleOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetAclRuleOperationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *AclRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BucketTags       []string             `json:"bucketTags"`
		BucketNameFilter *BucketNameFilter    `json:"bucketNameFilter"`
		ObjectNameFilter *ObjectNameFilter    `json:"objectNameFilter"`
		Action           AclRuleActionEnum    `json:"action"`
		NetworkSource    networksource        `json:"networkSource"`
		Operation        AclRuleOperationEnum `json:"operation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.BucketTags = make([]string, len(model.BucketTags))
	copy(model.BucketTags, m.BucketTags)
	m.BucketNameFilter = model.BucketNameFilter

	m.ObjectNameFilter = model.ObjectNameFilter

	m.Action = model.Action

	nn, e = model.NetworkSource.UnmarshalPolymorphicJSON(model.NetworkSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkSource = nn.(NetworkSource)
	} else {
		m.NetworkSource = nil
	}

	m.Operation = model.Operation

	return
}

// AclRuleActionEnum Enum with underlying type: string
type AclRuleActionEnum string

// Set of constants representing the allowable values for AclRuleActionEnum
const (
	AclRuleActionAllow     AclRuleActionEnum = "ALLOW"
	AclRuleActionDeny      AclRuleActionEnum = "DENY"
	AclRuleActionForceDeny AclRuleActionEnum = "FORCE_DENY"
)

var mappingAclRuleActionEnum = map[string]AclRuleActionEnum{
	"ALLOW":      AclRuleActionAllow,
	"DENY":       AclRuleActionDeny,
	"FORCE_DENY": AclRuleActionForceDeny,
}

var mappingAclRuleActionEnumLowerCase = map[string]AclRuleActionEnum{
	"allow":      AclRuleActionAllow,
	"deny":       AclRuleActionDeny,
	"force_deny": AclRuleActionForceDeny,
}

// GetAclRuleActionEnumValues Enumerates the set of values for AclRuleActionEnum
func GetAclRuleActionEnumValues() []AclRuleActionEnum {
	values := make([]AclRuleActionEnum, 0)
	for _, v := range mappingAclRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetAclRuleActionEnumStringValues Enumerates the set of values in String for AclRuleActionEnum
func GetAclRuleActionEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DENY",
		"FORCE_DENY",
	}
}

// GetMappingAclRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclRuleActionEnum(val string) (AclRuleActionEnum, bool) {
	enum, ok := mappingAclRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AclRuleOperationEnum Enum with underlying type: string
type AclRuleOperationEnum string

// Set of constants representing the allowable values for AclRuleOperationEnum
const (
	AclRuleOperationRead  AclRuleOperationEnum = "READ"
	AclRuleOperationWrite AclRuleOperationEnum = "WRITE"
	AclRuleOperationAll   AclRuleOperationEnum = "ALL"
)

var mappingAclRuleOperationEnum = map[string]AclRuleOperationEnum{
	"READ":  AclRuleOperationRead,
	"WRITE": AclRuleOperationWrite,
	"ALL":   AclRuleOperationAll,
}

var mappingAclRuleOperationEnumLowerCase = map[string]AclRuleOperationEnum{
	"read":  AclRuleOperationRead,
	"write": AclRuleOperationWrite,
	"all":   AclRuleOperationAll,
}

// GetAclRuleOperationEnumValues Enumerates the set of values for AclRuleOperationEnum
func GetAclRuleOperationEnumValues() []AclRuleOperationEnum {
	values := make([]AclRuleOperationEnum, 0)
	for _, v := range mappingAclRuleOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetAclRuleOperationEnumStringValues Enumerates the set of values in String for AclRuleOperationEnum
func GetAclRuleOperationEnumStringValues() []string {
	return []string{
		"READ",
		"WRITE",
		"ALL",
	}
}

// GetMappingAclRuleOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclRuleOperationEnum(val string) (AclRuleOperationEnum, bool) {
	enum, ok := mappingAclRuleOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
