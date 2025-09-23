// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscriptionAcknowledgmentConfiguration The configuration for subscription acknowledgment.
type SubscriptionAcknowledgmentConfiguration struct {

	// Flag to determine whether the subscription was acknowledged or not.
	IsAcknowledged *bool `mandatory:"true" json:"isAcknowledged"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the principal that ackwnoledged the subscription.
	AcknowledgedBy *string `mandatory:"false" json:"acknowledgedBy"`

	// The date and time the subscription was acknowledged (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeAcknowledged *common.SDKTime `mandatory:"false" json:"timeAcknowledged"`
}

func (m SubscriptionAcknowledgmentConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionAcknowledgmentConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
