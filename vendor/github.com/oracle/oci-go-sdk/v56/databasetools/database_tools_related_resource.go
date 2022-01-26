// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseToolsRelatedResource A related resource
type DatabaseToolsRelatedResource struct {

	// The resource entity type.
	EntityType RelatedResourceEntityTypeEnum `mandatory:"true" json:"entityType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related resource.
	Identifier *string `mandatory:"true" json:"identifier"`
}

func (m DatabaseToolsRelatedResource) String() string {
	return common.PointerString(m)
}
