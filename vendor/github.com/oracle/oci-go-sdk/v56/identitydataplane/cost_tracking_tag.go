// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CostTrackingTag The representation of CostTrackingTag
type CostTrackingTag struct {

	// The tag namespace id.
	TagNamespaceId *string `mandatory:"true" json:"Tag_Namespace_Id"`

	// The tag namespace name.
	TagNamespaceName *string `mandatory:"true" json:"Tag_Namespace_Name"`

	// The tag definition id.
	TagDefinitionId *string `mandatory:"true" json:"Tag_Definition_Id"`

	// The tag definition name.
	TagDefinitionName *string `mandatory:"true" json:"Tag_Definition_Name"`
}

func (m CostTrackingTag) String() string {
	return common.PointerString(m)
}
