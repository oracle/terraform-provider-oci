// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CandidateResponderRule Candidate Responder Rule list in Detector rule
type CandidateResponderRule struct {

	// The unique identifier of the Responder rule
	Id *string `mandatory:"false" json:"id"`

	// The display name of the Responder rule
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Preferred state
	IsPreferred *bool `mandatory:"false" json:"isPreferred"`
}

func (m CandidateResponderRule) String() string {
	return common.PointerString(m)
}
