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

func DataLabelingServiceAnnotationFormatsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataLabelingServiceAnnotationFormats,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"annotation_format_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}

func readDataLabelingServiceAnnotationFormats(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceAnnotationFormatsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()

	return tfresource.ReadResource(sync)
}

type DataLabelingServiceAnnotationFormatsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_labeling_service.DataLabelingManagementClient
	Res    *oci_data_labeling_service.ListAnnotationFormatsResponse
}

func (s *DataLabelingServiceAnnotationFormatsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataLabelingServiceAnnotationFormatsDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAnnotationFormats(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataLabelingServiceAnnotationFormatsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataLabelingServiceAnnotationFormatsDataSource-", DataLabelingServiceAnnotationFormatsDataSource(), s.D))
	resources := []map[string]interface{}{}
	annotationFormat := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AnnotationFormatsSummaryToMap(item))
	}
	annotationFormat["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataLabelingServiceAnnotationFormatsDataSource().Schema["annotation_format_collection"].Elem.(*schema.Resource).Schema)
		annotationFormat["items"] = items
	}

	resources = append(resources, annotationFormat)
	if err := s.D.Set("annotation_format_collection", resources); err != nil {
		return err
	}

	return nil
}

func AnnotationFormatsSummaryToMap(obj oci_data_labeling_service.AnnotationFormatSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func AnnotationFormatSummaryToMap(obj oci_data_labeling_service.AnnotationFormatSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
