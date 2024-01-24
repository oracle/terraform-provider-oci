// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_labeling_service

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_data_labeling_service_annotation_format", DataLabelingServiceAnnotationFormatDataSource())
	tfresource.RegisterDatasource("oci_data_labeling_service_annotation_formats", DataLabelingServiceAnnotationFormatsDataSource())
	tfresource.RegisterDatasource("oci_data_labeling_service_dataset", DataLabelingServiceDatasetDataSource())
	tfresource.RegisterDatasource("oci_data_labeling_service_datasets", DataLabelingServiceDatasetsDataSource())
}
