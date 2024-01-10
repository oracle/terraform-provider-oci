// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DocumentClassificationModelMetrics Metrics for Document Classification Model.
type DocumentClassificationModelMetrics struct {

	// List of metrics entries per label.
	LabelMetricsReport []DocumentClassificationLabelMetricsReport `mandatory:"true" json:"labelMetricsReport"`

	OverallMetricsReport *DocumentClassificationOverallMetricsReport `mandatory:"true" json:"overallMetricsReport"`

	DatasetSummary *DatasetSummary `mandatory:"false" json:"datasetSummary"`
}

// GetDatasetSummary returns DatasetSummary
func (m DocumentClassificationModelMetrics) GetDatasetSummary() *DatasetSummary {
	return m.DatasetSummary
}

func (m DocumentClassificationModelMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentClassificationModelMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DocumentClassificationModelMetrics) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDocumentClassificationModelMetrics DocumentClassificationModelMetrics
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDocumentClassificationModelMetrics
	}{
		"DOCUMENT_CLASSIFICATION",
		(MarshalTypeDocumentClassificationModelMetrics)(m),
	}

	return json.Marshal(&s)
}
