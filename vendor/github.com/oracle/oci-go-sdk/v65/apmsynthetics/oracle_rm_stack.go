// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleRmStack Details of the Oracle Resource Manager stack, which is a subtype of the Dedicated Vantage Point stack.
type OracleRmStack struct {

	// Version of the dedicated vantage point.
	DvpVersion *string `mandatory:"true" json:"dvpVersion"`

	// Stack OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Resource Manager stack for dedicated vantage point.
	DvpStackId *string `mandatory:"true" json:"dvpStackId"`

	// Stream OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Resource Manager stack for dedicated vantage point.
	DvpStreamId *string `mandatory:"true" json:"dvpStreamId"`
}

// GetDvpVersion returns DvpVersion
func (m OracleRmStack) GetDvpVersion() *string {
	return m.DvpVersion
}

func (m OracleRmStack) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleRmStack) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleRmStack) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleRmStack OracleRmStack
	s := struct {
		DiscriminatorParam string `json:"dvpStackType"`
		MarshalTypeOracleRmStack
	}{
		"ORACLE_RM_STACK",
		(MarshalTypeOracleRmStack)(m),
	}

	return json.Marshal(&s)
}
