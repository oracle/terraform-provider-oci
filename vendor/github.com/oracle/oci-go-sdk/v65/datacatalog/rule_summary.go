// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RuleSummary A list of rule resources. One or more rules can be defined for a data entity.
// Each rule can be defined on one or more attributes of the data entity.
type RuleSummary struct {

	// Immutable unique key of a rule.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of a rule.
	Description *string `mandatory:"false" json:"description"`

	// Type of a rule.
	RuleType RuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`

	// External URI that can be used to reference the object. Format will differ based on the type of object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Attributes associated with a rule.
	// A UNIQUEKEY rule would contain (at least) one attribute, for the local table column(s) on which uniqueness is defined.
	Attributes []RuleAttribute `mandatory:"false" json:"attributes"`

	// Folder key that represents the referenced folder, applicable only when rule type FOREIGNKEY.
	ReferencedFolderKey *string `mandatory:"false" json:"referencedFolderKey"`

	// Folder name that represents the referenced folder, applicable only when rule type FOREIGNKEY.
	ReferencedFolderName *string `mandatory:"false" json:"referencedFolderName"`

	// Entity key that represents the referenced entity, applicable only when rule type is FOREIGNKEY.
	ReferencedEntityKey *string `mandatory:"false" json:"referencedEntityKey"`

	// Entity name that represents the referenced entity, applicable only when rule type is FOREIGNKEY.
	ReferencedEntityName *string `mandatory:"false" json:"referencedEntityName"`

	// Rule key that represents the referenced rule, applicable only when rule type is FOREIGNKEY.
	ReferencedRuleKey *string `mandatory:"false" json:"referencedRuleKey"`

	// Rule name that represents the referenced rule, applicable only when rule type is FOREIGNKEY.
	ReferencedRuleName *string `mandatory:"false" json:"referencedRuleName"`

	// Attributes associated with referenced rule, applicable only when rule type is FOREIGNKEY.
	// A FOREIGNKEY rule would contain (at least) one attribute, for the local table column(s), and (at least) one referencedAttribute for referenced table column(s).
	ReferencedAttributes []RuleAttribute `mandatory:"false" json:"referencedAttributes"`

	// Origin type of the rule.
	OriginType RuleOriginTypeEnum `mandatory:"false" json:"originType,omitempty"`

	// URI to the rule instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// The date and time the rule was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// State of the rule.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m RuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetRuleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuleOriginTypeEnum(string(m.OriginType)); !ok && m.OriginType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OriginType: %s. Supported values are: %s.", m.OriginType, strings.Join(GetRuleOriginTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
