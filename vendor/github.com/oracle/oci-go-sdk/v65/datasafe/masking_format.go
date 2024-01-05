// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MaskingFormat A masking format defines the logic to mask data in a database column. The condition attribute
// defines the condition that must be true for applying the masking format. It enables you to do
// <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>
// so that you can mask the column data values differently using different masking formats and
// the associated conditions. A masking format can have one or more format entries. A format entry
// can be a basic masking format such as Random Number, or it can be a library masking format.The
// combined output of all the format entries is used for masking. It provides the flexibility to
// define a masking format that can generate different parts of a data value separately and then
// combine them to get the final data value for masking.
type MaskingFormat struct {

	// An array of format entries. The combined output of all the format entries is
	// used for masking the column data values.
	FormatEntries []FormatEntry `mandatory:"true" json:"formatEntries"`

	// A condition that must be true for applying the masking format. It can be any valid
	// SQL construct that can be used in a SQL predicate. It enables you to do
	// <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>
	// so that you can mask the column data values differently using different masking
	// formats and the associated conditions.
	Condition *string `mandatory:"false" json:"condition"`

	// The description of the masking format.
	Description *string `mandatory:"false" json:"description"`
}

func (m MaskingFormat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingFormat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MaskingFormat) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Condition     *string       `json:"condition"`
		Description   *string       `json:"description"`
		FormatEntries []formatentry `json:"formatEntries"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Condition = model.Condition

	m.Description = model.Description

	m.FormatEntries = make([]FormatEntry, len(model.FormatEntries))
	for i, n := range model.FormatEntries {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.FormatEntries[i] = nn.(FormatEntry)
		} else {
			m.FormatEntries[i] = nil
		}
	}
	return
}
