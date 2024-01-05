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

// ShuffleFormatEntry The Shuffle masking format randomly shuffles values within a column. It
// can also be used to shuffle column data within discrete units, or groups,
// where there is a relationship among the members of each group. To learn more,
// check Shuffle in the Data Safe documentation. The Shuffle masking format
// randomly shuffles values within a column. It can also be used to shuffle
// column data within discrete units, or groups, where there is a relationship
// among the members of each group. To learn more, check Shuffle in the
// Data Safe documentation.
type ShuffleFormatEntry struct {

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`

	// One or more reference columns to be used to group column values so that
	// they can be shuffled within their own group. The grouping columns and
	// the column to be masked must belong to the same table.
	GroupingColumns []string `mandatory:"false" json:"groupingColumns"`
}

// GetDescription returns Description
func (m ShuffleFormatEntry) GetDescription() *string {
	return m.Description
}

func (m ShuffleFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShuffleFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ShuffleFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeShuffleFormatEntry ShuffleFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeShuffleFormatEntry
	}{
		"SHUFFLE",
		(MarshalTypeShuffleFormatEntry)(m),
	}

	return json.Marshal(&s)
}
