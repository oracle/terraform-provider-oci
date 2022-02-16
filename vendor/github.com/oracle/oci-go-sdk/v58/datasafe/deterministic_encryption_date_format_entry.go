// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DeterministicEncryptionDateFormatEntry The Deterministic Encryption (Date) masking format encrypts column data using a cryptographic
// key and Advanced Encryption Standard (AES 128). It can be used to encrypt date columns only.
// It requires a range of dates as input defined by the startDate and endDate attributes. The
// start date must be less than or equal to the end date.
// The original column values in all the rows must be within the specified date range. The
// encrypted values are also within the specified range. Therefore, to ensure uniqueness, the
// total number of dates in the range must be greater than or equal to the number of distinct
// original values in the column. If an original value is not in the specified date range, it
// might not produce a one-to-one mapping. All non-confirming values are mapped to a single
// encrypted value, thereby producing a many-to-one mapping.
// Deterministic Encryption (Date) is a format-preserving, deterministic and reversible masking
// format, which requires a seed value while submitting a masking work request. Passing the
// same seed value when masking multiple times or masking different databases ensures that
// the data is masked deterministically. To learn more, check Deterministic Encryption in the
// Data Safe documentation.
type DeterministicEncryptionDateFormatEntry struct {

	// The lower bound of the range within which all the original column values fall.
	// The start date must be less than or equal to the end date.
	StartDate *common.SDKTime `mandatory:"true" json:"startDate"`

	// The upper bound of the range within which all the original column values fall.
	// The end date must be greater than or equal to the start date.
	EndDate *common.SDKTime `mandatory:"true" json:"endDate"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m DeterministicEncryptionDateFormatEntry) GetDescription() *string {
	return m.Description
}

func (m DeterministicEncryptionDateFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeterministicEncryptionDateFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeterministicEncryptionDateFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeterministicEncryptionDateFormatEntry DeterministicEncryptionDateFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDeterministicEncryptionDateFormatEntry
	}{
		"DETERMINISTIC_ENCRYPTION_DATE",
		(MarshalTypeDeterministicEncryptionDateFormatEntry)(m),
	}

	return json.Marshal(&s)
}
