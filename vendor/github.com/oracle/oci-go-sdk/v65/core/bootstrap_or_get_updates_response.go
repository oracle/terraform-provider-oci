// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BootstrapOrGetUpdatesResponse Response body for a successful `Bootstrap` or `GetUpdates` request.
type BootstrapOrGetUpdatesResponse struct {

	// Ordered list of data plane updates.
	Updates []PanamaDsUpdate `mandatory:"true" json:"updates"`

	// For a `Bootstrap` request, true iff there are more updates to return
	// for this bootstrap process, by passing the returned `sequenceToken`
	// to another `Bootstrap` request.  If false, then the bootstrap
	// process is complete, and the client has a consistent snapshot of NAT
	// gateway state that was true at some point during the bootstrap
	// process; it should then proceed to polling for updates with a
	// `GetUpdates` request, starting with the `sequenceToken` returned by
	// this terminal `Bootstrap` request.
	// For a `GetUpdates` request, true if there are known to be more
	// updates already available to return by passing the returned
	// `sequenceToken` to another `GetUpdates` request; otherwise, the
	// client should wait before polling again.
	MoreUpdates *bool `mandatory:"true" json:"moreUpdates"`

	// Opaque token to pass to another request of the same type to continue
	// a bootstrap process or poll for further updates.
	SequenceToken *string `mandatory:"true" json:"sequenceToken"`
}

func (m BootstrapOrGetUpdatesResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BootstrapOrGetUpdatesResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *BootstrapOrGetUpdatesResponse) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Updates       []panamadsupdate `json:"updates"`
		MoreUpdates   *bool            `json:"moreUpdates"`
		SequenceToken *string          `json:"sequenceToken"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Updates = make([]PanamaDsUpdate, len(model.Updates))
	for i, n := range model.Updates {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Updates[i] = nn.(PanamaDsUpdate)
		} else {
			m.Updates[i] = nil
		}
	}
	m.MoreUpdates = model.MoreUpdates

	m.SequenceToken = model.SequenceToken

	return
}
