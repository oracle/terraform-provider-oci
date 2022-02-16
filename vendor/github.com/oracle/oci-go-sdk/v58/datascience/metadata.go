// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Metadata Defines properties of each model metadata.
type Metadata struct {

	// Key of the model Metadata. The key can either be user defined or OCI defined.
	//    List of OCI defined keys:
	//          * useCaseType
	//          * libraryName
	//          * libraryVersion
	//          * estimatorClass
	//          * hyperParameters
	//          * testartifactresults
	Key *string `mandatory:"false" json:"key"`

	// Allowed values for useCaseType:
	//              binary_classification, regression, multinomial_classification, clustering, recommender,
	//              dimensionality_reduction/representation, time_series_forecasting, anomaly_detection,
	//              topic_modeling, ner, sentiment_analysis, image_classification, object_localization, other
	// Allowed values for libraryName:
	//              scikit-learn, xgboost, tensorflow, pytorch, mxnet, keras, lightGBM, pymc3, pyOD, spacy,
	//              prophet, sktime, statsmodels, cuml, oracle_automl, h2o, transformers, nltk, emcee, pystan,
	//              bert, gensim, flair, word2vec, ensemble, other
	Value *string `mandatory:"false" json:"value"`

	// Description of model metadata
	Description *string `mandatory:"false" json:"description"`

	// Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	Category *string `mandatory:"false" json:"category"`
}

func (m Metadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Metadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
