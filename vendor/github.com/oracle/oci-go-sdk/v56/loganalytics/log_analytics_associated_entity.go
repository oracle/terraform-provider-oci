// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
