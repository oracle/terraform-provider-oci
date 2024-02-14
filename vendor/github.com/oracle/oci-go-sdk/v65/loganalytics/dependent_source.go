// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DependentSource A source that uses the parser, either directly or indirectly.
type DependentSource struct {

	// The source name.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The source display name.
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The source type.
	SourceType *string `mandatory:"false" json:"sourceType"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// A flag indicating whether or not the source is marked for auto association.
	IsAutoAssociationEnabled *bool `mandatory:"false" json:"isAutoAssociationEnabled"`

	// The entity types.
	EntityTypes []LogAnalyticsSourceEntityType `mandatory:"false" json:"entityTypes"`

	// The list of dependencies defined by the source.
	Dependencies []Dependency `mandatory:"false" json:"dependencies"`
}

func (m DependentSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DependentSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
