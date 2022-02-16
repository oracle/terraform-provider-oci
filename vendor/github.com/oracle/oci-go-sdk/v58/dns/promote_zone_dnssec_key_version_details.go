// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PromoteZoneDnssecKeyVersionDetails Details for promoting a DnssecKeyVersion on a zone.
type PromoteZoneDnssecKeyVersionDetails struct {

	// The UUID of the DnssecKeyVersion that is being promoted.
	DnssecKeyVersionUuid *string `mandatory:"true" json:"dnssecKeyVersionUuid"`

	// The minimum length of time, in seconds, from now until when the predecessor DnssecKeyVersion should be
	// unpublished. If that precedes scheduled removal of the predecessor, the removal will be accelerated.
	// For KskDnssecKeyVersion, to avoid service disruption, this delay should be no less than the TTL of the
	// replaced DS records in the parent zone, increased to account for slow propagation if new DS records in the
	// parent zone are not yet resolvable.
	// When providing a value for a ZskDnssecKeyVersion, in order to avoid disruption, more time has to pass
	// between the new key introduction and old key removal than: the key set TTL combined with the maximum
	// record TTL on the zone combined with buffer time for system automation to handle signing changes. The
	// provided value cannot result in a time past the current scheduled expiration for the predecessor
	// DnssecKeyVersion.
	// TODO: Link to docs with expanded explanation and examples on timing.
	PredecessorUnpublishDelayInSeconds *int `mandatory:"false" json:"predecessorUnpublishDelayInSeconds"`

	// Optional length of time, in seconds, from now until when the new ZskDnssecKeyVersion should start being
	// used for signing and when the old ZskDnssecKeyVersion should stop being used for signing. If this change has
	// already occurred then this will have no effect.
	// TODO: Link to docs with expanded explanation and examples on timing.
	ActivationDelayInSeconds *int `mandatory:"false" json:"activationDelayInSeconds"`
}

func (m PromoteZoneDnssecKeyVersionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PromoteZoneDnssecKeyVersionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
