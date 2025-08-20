// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgImportPolicyUpdate Data plane information about Drg Attachment import policy on
// a route table.
type DrgImportPolicyUpdate struct {

	// Unique identifier for the primary resource affected by this update,
	// such as its OCID.
	Id *string `mandatory:"false" json:"id"`

	// True iff this update signals deletion of the identified resource.
	// If true, the type-specific fields of this object may be null.
	IsDelete *bool `mandatory:"false" json:"isDelete"`

	// The date and time that the API call was made that led to this Update.
	// The date and time format is defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The label given to the drg attachment
	DrgAttachmentLabel *int `mandatory:"false" json:"drgAttachmentLabel"`

	// Unique identifier for the Drg Route Table that is affected by this update.
	VrfId *string `mandatory:"false" json:"vrfId"`

	// The label given to the Route Table.
	VrfLabel *int `mandatory:"false" json:"vrfLabel"`

	// The preference given to the attachment.
	DrgAttachmentPreference *int `mandatory:"false" json:"drgAttachmentPreference"`

	// The DP ID of the DRG Attachment's DRG
	DrgDpId *int `mandatory:"false" json:"drgDpId"`

	// Boolean flag indicating whether the import policy is for
	// a route table with ecmp enabled
	IsEcmpEnabled *bool `mandatory:"false" json:"isEcmpEnabled"`

	// Boolean flag indicating whether the tenancy needs to import VCN_CIDR instead of SUBNET_CIDR
	DoImportVcnCidrs *bool `mandatory:"false" json:"doImportVcnCidrs"`
}

// GetId returns Id
func (m DrgImportPolicyUpdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m DrgImportPolicyUpdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m DrgImportPolicyUpdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DrgImportPolicyUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgImportPolicyUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrgImportPolicyUpdate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrgImportPolicyUpdate DrgImportPolicyUpdate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDrgImportPolicyUpdate
	}{
		"DrgImportPolicyUpdate",
		(MarshalTypeDrgImportPolicyUpdate)(m),
	}

	return json.Marshal(&s)
}
