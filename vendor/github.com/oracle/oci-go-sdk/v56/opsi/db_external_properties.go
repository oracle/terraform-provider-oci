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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DbExternalProperties Configuration parameters defined for external databases.
type DbExternalProperties struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// Name of the database.
	Name *string `mandatory:"false" json:"name"`

	// Archive log mode.
	LogMode *string `mandatory:"false" json:"logMode"`

	// Indicates if it is a CDB or not. This would be 'yes' or 'no'.
	Cdb *string `mandatory:"false" json:"cdb"`

	// Open mode information.
	OpenMode *string `mandatory:"false" json:"openMode"`

	// Current role of the database.
	DatabaseRole *string `mandatory:"false" json:"databaseRole"`

	// Data protection policy.
	GuardStatus *string `mandatory:"false" json:"guardStatus"`

	// Platform name of the database, OS with architecture.
	PlatformName *string `mandatory:"false" json:"platformName"`

	// Type of control file.
	ControlFileType *string `mandatory:"false" json:"controlFileType"`

	// Indicates whether switchover is allowed.
	SwitchoverStatus *string `mandatory:"false" json:"switchoverStatus"`

	// Creation time.
	Created *common.SDKTime `mandatory:"false" json:"created"`
}

//GetTimeCollected returns TimeCollected
func (m DbExternalProperties) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m DbExternalProperties) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DbExternalProperties) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbExternalProperties DbExternalProperties
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeDbExternalProperties
	}{
		"DB_EXTERNAL_PROPERTIES",
		(MarshalTypeDbExternalProperties)(m),
	}

	return json.Marshal(&s)
}
