// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CodeSearchResultCollection Results of a codesearch. Contains codesearch result items.
type CodeSearchResultCollection struct {

	// A page of codesearch results as a list returned by applying the search query.
	Items []CodeSearchResultSummary `mandatory:"true" json:"items"`

	// Total number of results returned by the query.
	TotalItems *int `mandatory:"true" json:"totalItems"`

	// Specifies if the search is performed using a modified query or not.
	// UNMODIFIED - Search performed with the query provided by user, which is in proper syntax.
	// MODIFIED - Search performed with a modified query, since user provided query is not in proper syntax, but its determined by the service that its modifiable.
	// INVALID - Unable to search due to incorrect query syntax and is not modifiable.
	QueryStatus CodeSearchResultCollectionQueryStatusEnum `mandatory:"false" json:"queryStatus,omitempty"`

	// Optional message from server.
	Message *string `mandatory:"false" json:"message"`
}

func (m CodeSearchResultCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CodeSearchResultCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCodeSearchResultCollectionQueryStatusEnum(string(m.QueryStatus)); !ok && m.QueryStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QueryStatus: %s. Supported values are: %s.", m.QueryStatus, strings.Join(GetCodeSearchResultCollectionQueryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CodeSearchResultCollectionQueryStatusEnum Enum with underlying type: string
type CodeSearchResultCollectionQueryStatusEnum string

// Set of constants representing the allowable values for CodeSearchResultCollectionQueryStatusEnum
const (
	CodeSearchResultCollectionQueryStatusUnmodified CodeSearchResultCollectionQueryStatusEnum = "UNMODIFIED"
	CodeSearchResultCollectionQueryStatusModified   CodeSearchResultCollectionQueryStatusEnum = "MODIFIED"
	CodeSearchResultCollectionQueryStatusInvalid    CodeSearchResultCollectionQueryStatusEnum = "INVALID"
)

var mappingCodeSearchResultCollectionQueryStatusEnum = map[string]CodeSearchResultCollectionQueryStatusEnum{
	"UNMODIFIED": CodeSearchResultCollectionQueryStatusUnmodified,
	"MODIFIED":   CodeSearchResultCollectionQueryStatusModified,
	"INVALID":    CodeSearchResultCollectionQueryStatusInvalid,
}

var mappingCodeSearchResultCollectionQueryStatusEnumLowerCase = map[string]CodeSearchResultCollectionQueryStatusEnum{
	"unmodified": CodeSearchResultCollectionQueryStatusUnmodified,
	"modified":   CodeSearchResultCollectionQueryStatusModified,
	"invalid":    CodeSearchResultCollectionQueryStatusInvalid,
}

// GetCodeSearchResultCollectionQueryStatusEnumValues Enumerates the set of values for CodeSearchResultCollectionQueryStatusEnum
func GetCodeSearchResultCollectionQueryStatusEnumValues() []CodeSearchResultCollectionQueryStatusEnum {
	values := make([]CodeSearchResultCollectionQueryStatusEnum, 0)
	for _, v := range mappingCodeSearchResultCollectionQueryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCodeSearchResultCollectionQueryStatusEnumStringValues Enumerates the set of values in String for CodeSearchResultCollectionQueryStatusEnum
func GetCodeSearchResultCollectionQueryStatusEnumStringValues() []string {
	return []string{
		"UNMODIFIED",
		"MODIFIED",
		"INVALID",
	}
}

// GetMappingCodeSearchResultCollectionQueryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCodeSearchResultCollectionQueryStatusEnum(val string) (CodeSearchResultCollectionQueryStatusEnum, bool) {
	enum, ok := mappingCodeSearchResultCollectionQueryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
