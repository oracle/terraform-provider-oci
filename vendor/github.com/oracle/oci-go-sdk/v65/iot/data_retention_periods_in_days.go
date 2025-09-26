// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataRetentionPeriodsInDays Data Retention periods
type DataRetentionPeriodsInDays struct {

	// Number of days for which any raw data sent to IoT devices would be retained for.
	RawData *int `mandatory:"true" json:"rawData"`

	// Number of days for which any data sent to IoT devices would be retained for.
	RejectedData *int `mandatory:"true" json:"rejectedData"`

	// Number of days for which any normalized data sent to IoT devices would be retained for.
	HistorizedData *int `mandatory:"true" json:"historizedData"`

	// Number of days for which any raw command data sent to IoT devices would be retained for.
	RawCommandData *int `mandatory:"true" json:"rawCommandData"`
}

func (m DataRetentionPeriodsInDays) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataRetentionPeriodsInDays) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
