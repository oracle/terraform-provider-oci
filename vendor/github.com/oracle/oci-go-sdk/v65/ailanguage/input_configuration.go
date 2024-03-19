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

// InputConfiguration input documents configuration
// by default TXT files will be processed and this behaviour will not change in future after adding new types
type InputConfiguration struct {

	// Type of documents supported
	// for this release only TXT,CSV  and one element is allowed here.
	// for future scope this is marked as list
	DocumentTypes []string `mandatory:"false" json:"documentTypes"`

	// meta data about documents
	//  For CSV valid JSON format is {"CSV" :{inputColumn: "reviewDetails", rowId: "reviewId", copyColumnsToOutput: ["reviewId" "userId"] , delimiter: ","}
	// Note: In future if new file types added we will update here in documentation about input file meta data
	Configuration map[string]DocumentsConfiguration `mandatory:"false" json:"configuration"`
}

func (m InputConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InputConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
