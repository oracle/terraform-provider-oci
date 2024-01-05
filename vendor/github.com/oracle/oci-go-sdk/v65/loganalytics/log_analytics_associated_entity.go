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

// LogAnalyticsAssociatedEntity LogAnalyticsAssociatedEntity
type LogAnalyticsAssociatedEntity struct {

	// The entity unique identifier.
	EntityId *string `mandatory:"false" json:"entityId"`

	// The entity name.
	EntityName *string `mandatory:"false" json:"entityName"`

	// The entity type.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The entity type display name.
	EntityTypeDisplayName *string `mandatory:"false" json:"entityTypeDisplayName"`

	// The host associated with the entity.
	OnHost *string `mandatory:"false" json:"onHost"`

	// The association count for the entity.
	AssociationCount *int64 `mandatory:"false" json:"associationCount"`
}

func (m LogAnalyticsAssociatedEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsAssociatedEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
