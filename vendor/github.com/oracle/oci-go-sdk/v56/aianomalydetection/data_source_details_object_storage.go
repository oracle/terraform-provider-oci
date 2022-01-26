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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DataSourceDetailsObjectStorage Data Source details for object storage
type DataSourceDetailsObjectStorage struct {

	// Object storage namespace
	Namespace *string `mandatory:"false" json:"namespace"`

	// Object storage bucket name
	BucketName *string `mandatory:"false" json:"bucketName"`

	// File name
	ObjectName *string `mandatory:"false" json:"objectName"`
}

func (m DataSourceDetailsObjectStorage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataSourceDetailsObjectStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataSourceDetailsObjectStorage DataSourceDetailsObjectStorage
	s := struct {
		DiscriminatorParam string `json:"dataSourceType"`
		MarshalTypeDataSourceDetailsObjectStorage
	}{
		"ORACLE_OBJECT_STORAGE",
		(MarshalTypeDataSourceDetailsObjectStorage)(m),
	}

	return json.Marshal(&s)
}
