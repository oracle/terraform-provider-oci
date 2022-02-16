// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// EnableMacsManagedExternalHostInsightDetails The information about the MACS-managed external host to be analyzed.
type EnableMacsManagedExternalHostInsightDetails struct {
}

func (m EnableMacsManagedExternalHostInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableMacsManagedExternalHostInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EnableMacsManagedExternalHostInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEnableMacsManagedExternalHostInsightDetails EnableMacsManagedExternalHostInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEnableMacsManagedExternalHostInsightDetails
	}{
		"MACS_MANAGED_EXTERNAL_HOST",
		(MarshalTypeEnableMacsManagedExternalHostInsightDetails)(m),
	}

	return json.Marshal(&s)
}
