// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_ai_anomaly_detection_ai_private_endpoint", AiAnomalyDetectionAiPrivateEndpointResource())
	tfresource.RegisterResource("oci_ai_anomaly_detection_data_asset", AiAnomalyDetectionDataAssetResource())
	tfresource.RegisterResource("oci_ai_anomaly_detection_detect_anomaly_job", AiAnomalyDetectionDetectAnomalyJobResource())
	tfresource.RegisterResource("oci_ai_anomaly_detection_model", AiAnomalyDetectionModelResource())
	tfresource.RegisterResource("oci_ai_anomaly_detection_project", AiAnomalyDetectionProjectResource())
}
