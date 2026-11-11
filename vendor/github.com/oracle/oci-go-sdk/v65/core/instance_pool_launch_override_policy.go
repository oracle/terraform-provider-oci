// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstancePoolLaunchOverridePolicy Optional pool-level launch override policy for future instance launches.
// The policy is evaluated only after placement is resolved for a launch.
// In V1, provide between 1 and 3 availability domain-scoped statements and
// no more than one statement for the same availability domain.
// For a matched availability domain, request order defines the canonical
// override candidate sequence. Preferred-candidate reuse may try one
// current option once as a fast path. If that fast path misses, canonical
// evaluation restarts from option 1. During canonical evaluation, the
// worker advances to the next option only for `out-of-host-capacity`,
// `LimitExceeded`, and `QuotaExceeded`, and only after same-option retries
// are exhausted. Other failures are terminal for that launch attempt.
// When used on create, omitting `overridePolicy` preserves the existing
// base-plus-placement behavior. When used on update, omitting
// `overridePolicy` leaves the stored policy unchanged, a non-null value
// replaces the stored policy, and `null` clears it. Read surfaces omit
// `overridePolicy` when no policy is stored on the pool.
type InstancePoolLaunchOverridePolicy struct {

	// The ordered set of availability domain-scoped override statements for this policy.
	OverrideStatements []InstancePoolLaunchOverrideStatement `mandatory:"true" json:"overrideStatements"`
}

func (m InstancePoolLaunchOverridePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolLaunchOverridePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
