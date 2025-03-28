// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberDatabase The properties for a Base Database or Exadata Database member of a DR protection group.
type DrProtectionGroupMemberDatabase struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// The OCID of the vault secret where the database SYSDBA password is stored.
	// This password is used for performing database DR operations.
	// Example: `ocid1.vaultsecret.oc1..uniqueID`
	PasswordVaultSecretId *string `mandatory:"false" json:"passwordVaultSecretId"`
}

// GetMemberId returns MemberId
func (m DrProtectionGroupMemberDatabase) GetMemberId() *string {
	return m.MemberId
}

func (m DrProtectionGroupMemberDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroupMemberDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrProtectionGroupMemberDatabase) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrProtectionGroupMemberDatabase DrProtectionGroupMemberDatabase
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeDrProtectionGroupMemberDatabase
	}{
		"DATABASE",
		(MarshalTypeDrProtectionGroupMemberDatabase)(m),
	}

	return json.Marshal(&s)
}
