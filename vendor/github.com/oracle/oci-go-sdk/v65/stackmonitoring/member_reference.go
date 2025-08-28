// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MemberReference The Member Reference details
type MemberReference struct {

	// The OCID of the resourceInstance/resourceType/resourceGroup
	Id *string `mandatory:"true" json:"id"`

	// Type of the member reference RESOURCE_INSTANCE, RESOURCE_TYPE, RESOURCE_GROUP
	Type MemberReferenceTypeEnum `mandatory:"true" json:"type"`

	// The OCID of the composite resource type like EBS or Peoplesoft.
	CompositeType *string `mandatory:"false" json:"compositeType"`
}

func (m MemberReference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MemberReference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMemberReferenceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMemberReferenceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
