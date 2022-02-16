// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreationSource Details for auto-created entity.
type CreationSource struct {

	// Source that auto-created the entity.
	Type CreationSourceTypeEnum `mandatory:"false" json:"type,omitempty"`

	// This will provide additional details for source of auto-creation. For example, if entity is auto-created
	// by enterprise manager bridge, this field provides additional detail on enterprise manager that contributed
	// to the entity auto-creation.
	Details *string `mandatory:"false" json:"details"`
}

func (m CreationSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreationSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreationSourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreationSourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
