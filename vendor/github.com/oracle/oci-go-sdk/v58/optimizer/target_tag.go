// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// TargetTag A tag key definition used in the current profile override, including the tag namespace, tag key, tag value type, and tag values.
// Only defined tags are supported.
// For more information about tagging, see Tagging Overview (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm)
type TargetTag struct {

	// The name of the tag namespace.
	TagNamespaceName *string `mandatory:"true" json:"tagNamespaceName"`

	// The name you use to refer to the tag, also known as the tag key.
	TagDefinitionName *string `mandatory:"true" json:"tagDefinitionName"`

	// Specifies which tag value types in the `tagValues` field result in overrides of the recommendation criteria.
	// When the value for this field is `ANY`, the `tagValues` field should be empty, which enforces overrides to the recommendation
	// for resources with any tag values attached to them.
	// When the value for this field value is `VALUE`, the `tagValues` field must include a specific value or list of values.
	// Overrides to the recommendation criteria only occur for resources that match the values in the `tagValues` fields.
	TagValueType TagValueTypeEnum `mandatory:"true" json:"tagValueType"`

	// The list of tag values. The tag value is the value that the user applying the tag adds to the tag key.
	TagValues []string `mandatory:"false" json:"tagValues"`
}

func (m TargetTag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetTag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTagValueTypeEnum(string(m.TagValueType)); !ok && m.TagValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TagValueType: %s. Supported values are: %s.", m.TagValueType, strings.Join(GetTagValueTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
