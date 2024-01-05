// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
)

func AiAnomalyDetectionProjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiAnomalyDetectionProjects,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(AiAnomalyDetectionProjectResource()),
						},
					},
				},
			},
		},
	}
}

func readAiAnomalyDetectionProjects(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionProjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionProjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.ListProjectsResponse
}

func (s *AiAnomalyDetectionProjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionProjectsDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.ListProjectsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_ai_anomaly_detection.ProjectLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.ListProjects(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiAnomalyDetectionProjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiAnomalyDetectionProjectsDataSource-", AiAnomalyDetectionProjectsDataSource(), s.D))
	resources := []map[string]interface{}{}
	project := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, aiAnomalyDetectionProjectSummaryToMap(item))
	}
	project["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiAnomalyDetectionProjectsDataSource().Schema["project_collection"].Elem.(*schema.Resource).Schema)
		project["items"] = items
	}

	resources = append(resources, project)
	if err := s.D.Set("project_collection", resources); err != nil {
		return err
	}

	return nil
}
