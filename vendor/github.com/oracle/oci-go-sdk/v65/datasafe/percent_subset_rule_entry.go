// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PercentSubsetRuleEntry Defines a subsetting rule based on a percentage of the data to retain
type PercentSubsetRuleEntry struct {

	// The percentage of rows to retain in the subset (between 0 and 100)
	Percent *int `mandatory:"true" json:"percent"`
}

func (m PercentSubsetRuleEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PercentSubsetRuleEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PercentSubsetRuleEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePercentSubsetRuleEntry PercentSubsetRuleEntry
	s := struct {
		DiscriminatorParam string `json:"ruleType"`
		MarshalTypePercentSubsetRuleEntry
	}{
		"PERCENT",
		(MarshalTypePercentSubsetRuleEntry)(m),
	}

	return json.Marshal(&s)
}
