// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_labeling_service

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v65/datalabelingservice"
)

func DataLabelingServiceAnnotationFormatDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataLabelingServiceAnnotationFormat,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_labeling_service_annotation_format", "oci_data_labeling_service_annotation_formats"),
	}
}

func readSingularDataLabelingServiceAnnotationFormat(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceAnnotationFormatDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()

	return tfresource.ReadResource(sync)
}

type DataLabelingServiceAnnotationFormatDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_labeling_service.DataLabelingManagementClient
	Res    *oci_data_labeling_service.ListAnnotationFormatsResponse
}

func (s *DataLabelingServiceAnnotationFormatDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataLabelingServiceAnnotationFormatDataSourceCrud) Get() error {
	request := oci_data_labeling_service.ListAnnotationFormatsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_labeling_service")

	response, err := s.Client.ListAnnotationFormats(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataLabelingServiceAnnotationFormatDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataLabelingServiceAnnotationFormatDataSource-", DataLabelingServiceAnnotationFormatDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AnnotationFormatSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
