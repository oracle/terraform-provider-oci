// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ZprTagNamespaceTagSummary A Zpr Tag definition that belongs to a specific Zpr Tag namespace.
type ZprTagNamespaceTagSummary struct {

	// The OCID of the compartment that contains the Zprtag definition.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the namespace that contains the Zpr tag definition.
	ZprTagNamespaceId *string `mandatory:"false" json:"zprTagNamespaceId"`

	// The name of the Zpr tag namespace that contains the Zpr tag definition.
	ZprTagNamespaceName *string `mandatory:"false" json:"zprTagNamespaceName"`

	// The OCID of the Zprtag definition.
	Id *string `mandatory:"false" json:"id"`

	// The name assigned to the Zpr tag during creation. This is the tag key definition.
	// The name must be unique within the ZPR tag namespace and cannot be changed.
	Name *string `mandatory:"false" json:"name"`

	// The description you assign to the ZPR tag.
	Description *string `mandatory:"false" json:"description"`

	// The data type of the ZPR tag.
	Type *string `mandatory:"false" json:"type"`

	// Whether the ZPR tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm#retiringkeys).
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// The ZPR tag's current state. After creating a ZPR tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a ZPR ag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a ZPR tag, you cannot delete another ZPR tag until the deleted ZPR tag's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState ZprTagNamespaceTagLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Date and time the ZPR tag was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ZprTagNamespaceTagSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZprTagNamespaceTagSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingZprTagNamespaceTagLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZprTagNamespaceTagLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
