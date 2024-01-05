// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TrailFileSummary Summary of the TrailFiles.
type TrailFileSummary struct {

	// The TrailFile Id.
	TrailFileId *string `mandatory:"true" json:"trailFileId"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The size of the backup stored in object storage (in bytes)
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeLastUpdated *common.SDKTime `mandatory:"false" json:"timeLastUpdated"`

	// Number of sequences for a specific trail file
	NumberOfSequences *int `mandatory:"false" json:"numberOfSequences"`

	// Minimum sequence number
	MinSequenceNumber *string `mandatory:"false" json:"minSequenceNumber"`

	// Maximum sequence number
	MaxSequenceNumber *string `mandatory:"false" json:"maxSequenceNumber"`

	// Producer Process Name if any.
	Producer *string `mandatory:"false" json:"producer"`

	// array of consumer process names
	Consumers []string `mandatory:"false" json:"consumers"`
}

func (m TrailFileSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TrailFileSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
