// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FirewallPolicyEntryDetails SQL Firewall policy details.
type FirewallPolicyEntryDetails struct {

	// The time the the SQL Firewall policy was generated on the target database, in the format defined by RFC3339.
	TimeGenerated *common.SDKTime `mandatory:"true" json:"timeGenerated"`

	// The last date and time the status of the SQL Firewall policy was updated on the target database, in the format defined by RFC3339.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" json:"timeStatusUpdated"`
}

func (m FirewallPolicyEntryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FirewallPolicyEntryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FirewallPolicyEntryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFirewallPolicyEntryDetails FirewallPolicyEntryDetails
	s := struct {
		DiscriminatorParam string `json:"entryType"`
		MarshalTypeFirewallPolicyEntryDetails
	}{
		"FIREWALL_POLICY",
		(MarshalTypeFirewallPolicyEntryDetails)(m),
	}

	return json.Marshal(&s)
}
