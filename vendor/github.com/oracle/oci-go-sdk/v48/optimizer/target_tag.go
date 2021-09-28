// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v48/common"
)

// TargetTag A target tag with tag namespace, tag definition, tag value type, and tag values attached to the current profile override.
type TargetTag struct {

	// The name of the tag namespace.
	TagNamespaceName *string `mandatory:"true" json:"tagNamespaceName"`

	// The name of the tag definition.
	TagDefinitionName *string `mandatory:"true" json:"tagDefinitionName"`

	// The tag value type.
	TagValueType TagValueTypeEnum `mandatory:"true" json:"tagValueType"`

	// The list of tag values.
	TagValues []string `mandatory:"false" json:"tagValues"`
}

func (m TargetTag) String() string {
	return common.PointerString(m)
}
