// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BitbucketServerFilter The filter for Bitbucket Server events.
type BitbucketServerFilter struct {
	Include *BitbucketServerFilterAttributes `mandatory:"false" json:"include"`

	// The events, for example, PUSH, PULL_REQUEST_MERGE.
	Events []BitbucketServerFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m BitbucketServerFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketServerFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingBitbucketServerFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetBitbucketServerFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BitbucketServerFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketServerFilter BitbucketServerFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeBitbucketServerFilter
	}{
		"BITBUCKET_SERVER",
		(MarshalTypeBitbucketServerFilter)(m),
	}

	return json.Marshal(&s)
}

// BitbucketServerFilterEventsEnum Enum with underlying type: string
type BitbucketServerFilterEventsEnum string

// Set of constants representing the allowable values for BitbucketServerFilterEventsEnum
const (
	BitbucketServerFilterEventsPush                BitbucketServerFilterEventsEnum = "PUSH"
	BitbucketServerFilterEventsPullRequestOpened   BitbucketServerFilterEventsEnum = "PULL_REQUEST_OPENED"
	BitbucketServerFilterEventsPullRequestModified BitbucketServerFilterEventsEnum = "PULL_REQUEST_MODIFIED"
	BitbucketServerFilterEventsPullRequestMerged   BitbucketServerFilterEventsEnum = "PULL_REQUEST_MERGED"
)

var mappingBitbucketServerFilterEventsEnum = map[string]BitbucketServerFilterEventsEnum{
	"PUSH":                  BitbucketServerFilterEventsPush,
	"PULL_REQUEST_OPENED":   BitbucketServerFilterEventsPullRequestOpened,
	"PULL_REQUEST_MODIFIED": BitbucketServerFilterEventsPullRequestModified,
	"PULL_REQUEST_MERGED":   BitbucketServerFilterEventsPullRequestMerged,
}

var mappingBitbucketServerFilterEventsEnumLowerCase = map[string]BitbucketServerFilterEventsEnum{
	"push":                  BitbucketServerFilterEventsPush,
	"pull_request_opened":   BitbucketServerFilterEventsPullRequestOpened,
	"pull_request_modified": BitbucketServerFilterEventsPullRequestModified,
	"pull_request_merged":   BitbucketServerFilterEventsPullRequestMerged,
}

// GetBitbucketServerFilterEventsEnumValues Enumerates the set of values for BitbucketServerFilterEventsEnum
func GetBitbucketServerFilterEventsEnumValues() []BitbucketServerFilterEventsEnum {
	values := make([]BitbucketServerFilterEventsEnum, 0)
	for _, v := range mappingBitbucketServerFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetBitbucketServerFilterEventsEnumStringValues Enumerates the set of values in String for BitbucketServerFilterEventsEnum
func GetBitbucketServerFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
		"PULL_REQUEST_OPENED",
		"PULL_REQUEST_MODIFIED",
		"PULL_REQUEST_MERGED",
	}
}

// GetMappingBitbucketServerFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBitbucketServerFilterEventsEnum(val string) (BitbucketServerFilterEventsEnum, bool) {
	enum, ok := mappingBitbucketServerFilterEventsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
