// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeProtectedDatabaseSubscriptionDetails Subscription details of the new cloud service environment, such as Microsoft Azure, where the protected database will be provisioned.
type ChangeProtectedDatabaseSubscriptionDetails struct {

	// The OCID of the new cloud service subscription to which you want to link the protected database.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Indicates whether it is a Universal Credit Model (UCM) subscription.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m ChangeProtectedDatabaseSubscriptionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeProtectedDatabaseSubscriptionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
