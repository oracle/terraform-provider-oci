// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_labeling_service

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v56/datalabelingservice"
)

func DataLabelingServiceDatasetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataLabelingServiceDatasets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"annotation_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dataset_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataLabelingServiceDatasetResource()),
						},
					},
				},
			},
		},
	}
}

func readDataLabelingServiceDatasets(d *schema.ResourceData, m interface{}) error {
	sync := &DataLabelingServiceDatasetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataLabelingManagementClient()

	return tfresource.ReadResource(sync)
}

type DataLabelingServiceDatasetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_labeling_service.DataLabelingManagementClient
	Res    *oci_data_labeling_service.ListDatasetsResponse
}

func (s *DataLabelingServiceDatasetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataLabelingServiceDatasetsDataSourceCrud) Get() error {
	request := oci_data_labeling_service.ListDatasetsRequest{}

	if annotationFormat, ok := s.D.GetOkExists("annotation_format"); ok {
		tmp := annotationFormat.(string)
		request.AnnotationFormat = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_labeling_service.DatasetLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_labeling_service")

	response, err := s.Client.ListDatasets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatasets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataLabelingServiceDatasetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataLabelingServiceDatasetsDataSource-", DataLabelingServiceDatasetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dataset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatasetSummaryToMap(item))
	}
	dataset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataLabelingServiceDatasetsDataSource().Schema["dataset_collection"].Elem.(*schema.Resource).Schema)
		dataset["items"] = items
	}

	resources = append(resources, dataset)
	if err := s.D.Set("dataset_collection", resources); err != nil {
		return err
	}

	return nil
}
