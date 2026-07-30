// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSafeTargetSummary Summary of the targets.
type DataSafeTargetSummary struct {

	// The OCID of the target database.
	DbId *string `mandatory:"true" json:"dbId"`

	// The name of the target database.
	DbName *string `mandatory:"true" json:"dbName"`

	// The date and time the target was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Data Safe Private Endpoint OCID.
	DataSafePrivateEndpointId *string `mandatory:"false" json:"dataSafePrivateEndpointId"`

	// Data Safe Agent OCID.
	DataSafeAgentId *string `mandatory:"false" json:"dataSafeAgentId"`

	// The OCID of the on-premises connector.
	OnPremConnectorId *string `mandatory:"false" json:"onPremConnectorId"`
}

func (m DataSafeTargetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSafeTargetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
