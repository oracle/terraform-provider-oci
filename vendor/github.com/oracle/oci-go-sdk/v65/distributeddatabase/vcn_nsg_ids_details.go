// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VcnNsgIdsDetails Specifies the Virtual Cloud Network (VCN) and the list of associated Network Security Group (NSG) OCIDs.
// This model is shared across create (POST), update (PUT/PATCH), and retrieve (GET) operations.
// - During create, `nsgIds` (if provided) specifies the NSGs to associate with the new resource.
// - During update or patch, `nsgIds` modifies the associations: provide an empty array (`[]`) to clear all NSGs, or provide a new list to replace them.
// - GET responses return the current associations (which may be empty).
type VcnNsgIdsDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer Virtual Cloud Network.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The list of OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) associated with resources in this VCN.
	// For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"true" json:"nsgIds"`
}

func (m VcnNsgIdsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VcnNsgIdsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
