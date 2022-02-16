// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InfluxDetailsV2v0 Influx details for V_2_0.
type InfluxDetailsV2v0 struct {

	// Bucket Name for influx connection
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Org name for the influx db
	OrganizationName *string `mandatory:"true" json:"organizationName"`
}

func (m InfluxDetailsV2v0) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfluxDetailsV2v0) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InfluxDetailsV2v0) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInfluxDetailsV2v0 InfluxDetailsV2v0
	s := struct {
		DiscriminatorParam string `json:"influxVersion"`
		MarshalTypeInfluxDetailsV2v0
	}{
		"V_2_0",
		(MarshalTypeInfluxDetailsV2v0)(m),
	}

	return json.Marshal(&s)
}
