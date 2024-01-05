// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_ai_private_endpoint", AiAnomalyDetectionAiPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_ai_private_endpoints", AiAnomalyDetectionAiPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_data_asset", AiAnomalyDetectionDataAssetDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_data_assets", AiAnomalyDetectionDataAssetsDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_detect_anomaly_job", AiAnomalyDetectionDetectAnomalyJobDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_detect_anomaly_jobs", AiAnomalyDetectionDetectAnomalyJobsDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_model", AiAnomalyDetectionModelDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_models", AiAnomalyDetectionModelsDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_project", AiAnomalyDetectionProjectDataSource())
	tfresource.RegisterDatasource("oci_ai_anomaly_detection_projects", AiAnomalyDetectionProjectsDataSource())
}
