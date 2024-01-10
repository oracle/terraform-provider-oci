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

// BitbucketCloudFilter The filter for Bitbucket Cloud events.
type BitbucketCloudFilter struct {
	Include *BitbucketCloudFilterAttributes `mandatory:"false" json:"include"`

	Exclude *BitbucketCloudFilterExclusionAttributes `mandatory:"false" json:"exclude"`

	// The events, for example, PUSH, PULL_REQUEST_MERGE.
	Events []BitbucketCloudFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m BitbucketCloudFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketCloudFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingBitbucketCloudFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetBitbucketCloudFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BitbucketCloudFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketCloudFilter BitbucketCloudFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeBitbucketCloudFilter
	}{
		"BITBUCKET_CLOUD",
		(MarshalTypeBitbucketCloudFilter)(m),
	}

	return json.Marshal(&s)
}

// BitbucketCloudFilterEventsEnum Enum with underlying type: string
type BitbucketCloudFilterEventsEnum string

// Set of constants representing the allowable values for BitbucketCloudFilterEventsEnum
const (
	BitbucketCloudFilterEventsPush               BitbucketCloudFilterEventsEnum = "PUSH"
	BitbucketCloudFilterEventsPullRequestCreated BitbucketCloudFilterEventsEnum = "PULL_REQUEST_CREATED"
	BitbucketCloudFilterEventsPullRequestUpdated BitbucketCloudFilterEventsEnum = "PULL_REQUEST_UPDATED"
	BitbucketCloudFilterEventsPullRequestMerged  BitbucketCloudFilterEventsEnum = "PULL_REQUEST_MERGED"
)

var mappingBitbucketCloudFilterEventsEnum = map[string]BitbucketCloudFilterEventsEnum{
	"PUSH":                 BitbucketCloudFilterEventsPush,
	"PULL_REQUEST_CREATED": BitbucketCloudFilterEventsPullRequestCreated,
	"PULL_REQUEST_UPDATED": BitbucketCloudFilterEventsPullRequestUpdated,
	"PULL_REQUEST_MERGED":  BitbucketCloudFilterEventsPullRequestMerged,
}

var mappingBitbucketCloudFilterEventsEnumLowerCase = map[string]BitbucketCloudFilterEventsEnum{
	"push":                 BitbucketCloudFilterEventsPush,
	"pull_request_created": BitbucketCloudFilterEventsPullRequestCreated,
	"pull_request_updated": BitbucketCloudFilterEventsPullRequestUpdated,
	"pull_request_merged":  BitbucketCloudFilterEventsPullRequestMerged,
}

// GetBitbucketCloudFilterEventsEnumValues Enumerates the set of values for BitbucketCloudFilterEventsEnum
func GetBitbucketCloudFilterEventsEnumValues() []BitbucketCloudFilterEventsEnum {
	values := make([]BitbucketCloudFilterEventsEnum, 0)
	for _, v := range mappingBitbucketCloudFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetBitbucketCloudFilterEventsEnumStringValues Enumerates the set of values in String for BitbucketCloudFilterEventsEnum
func GetBitbucketCloudFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
		"PULL_REQUEST_CREATED",
		"PULL_REQUEST_UPDATED",
		"PULL_REQUEST_MERGED",
	}
}

// GetMappingBitbucketCloudFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBitbucketCloudFilterEventsEnum(val string) (BitbucketCloudFilterEventsEnum, bool) {
	enum, ok := mappingBitbucketCloudFilterEventsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
