// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v47/aianomalydetection"

	oci_common "github.com/oracle/oci-go-sdk/v47/common"
)

func init() {
	RegisterOracleClient("oci_ai_anomaly_detection.AnomalyDetectionClient", &OracleClient{initClientFn: initAianomalydetectionAnomalyDetectionClient})
}

func initAianomalydetectionAnomalyDetectionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ai_anomaly_detection.NewAnomalyDetectionClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) anomalyDetectionClient() *oci_ai_anomaly_detection.AnomalyDetectionClient {
	return m.GetClient("oci_ai_anomaly_detection.AnomalyDetectionClient").(*oci_ai_anomaly_detection.AnomalyDetectionClient)
}
