// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// StandardTagNamespaceTemplate The template of the standard tag namespace. This object includes necessary details to create the provided standard tag namespace.
type StandardTagNamespaceTemplate struct {

	// The default description of the tag namespace that users can use to create the tag namespace
	Description *string `mandatory:"true" json:"description"`

	// The reserved name of this standard tag namespace
	StandardTagNamespaceName *string `mandatory:"true" json:"standardTagNamespaceName"`

	// The template of the tag definition. This object includes necessary details to create the provided standard tag definition.
	TagDefinitionTemplates []StandardTagDefinitionTemplate `mandatory:"true" json:"tagDefinitionTemplates"`

	// The status of the standard tag namespace
	Status *string `mandatory:"true" json:"status"`
}

func (m StandardTagNamespaceTemplate) String() string {
	return common.PointerString(m)
}
