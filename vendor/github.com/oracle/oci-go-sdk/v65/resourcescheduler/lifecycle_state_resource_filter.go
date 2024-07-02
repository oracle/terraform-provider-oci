// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LifecycleStateResourceFilter This is a resource filter for filtering resources based on their lifecycle state.
type LifecycleStateResourceFilter struct {

	// This is a collection of resource lifecycle state values.
	Value []string `mandatory:"false" json:"value"`
}

func (m LifecycleStateResourceFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LifecycleStateResourceFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LifecycleStateResourceFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLifecycleStateResourceFilter LifecycleStateResourceFilter
	s := struct {
		DiscriminatorParam string `json:"attribute"`
		MarshalTypeLifecycleStateResourceFilter
	}{
		"LIFECYCLE_STATE",
		(MarshalTypeLifecycleStateResourceFilter)(m),
	}

	return json.Marshal(&s)
}
