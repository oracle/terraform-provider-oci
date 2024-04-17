// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResolvedEntity Resolved entity.
type ResolvedEntity struct {

	// id of the resolved entity in input
	Id *string `mandatory:"true" json:"id"`

	// offset of resolved entity in input
	Offset *int `mandatory:"true" json:"offset"`

	// length of resolved entity in input
	Length *int `mandatory:"true" json:"length"`

	// Entity text like name of person, location, and so on.
	Text *string `mandatory:"true" json:"text"`

	// Type of entity text like PER, LOC.
	Type *string `mandatory:"true" json:"type"`

	// key and value pair for resolved entities. keys can be specific for each type of resolved entity. Values can be instances of resolvedEntity, arrays of resolvedEntities, primitives, or custom JSON.
	Value map[string]string `mandatory:"true" json:"value"`
}

func (m ResolvedEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResolvedEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
