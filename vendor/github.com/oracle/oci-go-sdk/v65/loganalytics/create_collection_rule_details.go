// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCollectionRuleDetails Information required to create a collection rule. The rule could be created based on a collection template or
// by directly specifying the details, along with other required parameters.
type CreateCollectionRuleDetails struct {

	// Compartment Identifier OCID  (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The collection rule name.
	Name *string `mandatory:"true" json:"name"`

	// The collection rule type.
	Type CollectionRuleTypeEnum `mandatory:"true" json:"type"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	Context *CollectionRuleContext `mandatory:"false" json:"context"`

	// The OCID of the collection template that contains details from which information
	// required for setting up log collection can be obtained. For example, in case of agent based collections,
	// the collection template would specify the entity types and corresponding log sources to create associations.
	TemplateId *string `mandatory:"false" json:"templateId"`

	// The OCID of the log group which would contain the collected logs.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	AssociationDetails *CollectionRuleAssociationDetails `mandatory:"false" json:"associationDetails"`
}

func (m CreateCollectionRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCollectionRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCollectionRuleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCollectionRuleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
