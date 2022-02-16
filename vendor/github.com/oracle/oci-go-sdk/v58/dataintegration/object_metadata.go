// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ObjectMetadata A summary type containing information about the object including its key, name and when/who created/updated it.
type ObjectMetadata struct {

	// The user that created the object.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The user that created the object.
	CreatedByName *string `mandatory:"false" json:"createdByName"`

	// The user that updated the object.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// The user that updated the object.
	UpdatedByName *string `mandatory:"false" json:"updatedByName"`

	// The date and time that the object was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time that the object was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The owning object key for this object.
	AggregatorKey *string `mandatory:"false" json:"aggregatorKey"`

	Aggregator *AggregatorSummary `mandatory:"false" json:"aggregator"`

	// The full path to identify this object.
	IdentifierPath *string `mandatory:"false" json:"identifierPath"`

	// Information property fields.
	InfoFields map[string]string `mandatory:"false" json:"infoFields"`

	// The registry version of the object.
	RegistryVersion *int `mandatory:"false" json:"registryVersion"`

	// Labels are keywords or tags that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	// Specifies whether this object is a favorite or not.
	IsFavorite *bool `mandatory:"false" json:"isFavorite"`

	CountStatistics *CountStatistic `mandatory:"false" json:"countStatistics"`
}

func (m ObjectMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
