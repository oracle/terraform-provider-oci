// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataAsmEntity ASM entitie for an exadata
type ExadataAsmEntity struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// Instance name of ASM
	InstanceName *string `mandatory:"false" json:"instanceName"`

	// Cluster name of ASM
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// Software version
	SoftwareVersion *string `mandatory:"false" json:"softwareVersion"`
}

// GetTimeCollected returns TimeCollected
func (m ExadataAsmEntity) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m ExadataAsmEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataAsmEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExadataAsmEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataAsmEntity ExadataAsmEntity
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeExadataAsmEntity
	}{
		"ASM_ENTITY",
		(MarshalTypeExadataAsmEntity)(m),
	}

	return json.Marshal(&s)
}
