// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentDataIngestionJobsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGenerativeAiAgentDataIngestionJobs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_ingestion_job_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(GenerativeAiAgentDataIngestionJobResource()),
						},
					},
				},
			},
		},
	}
}

func readGenerativeAiAgentDataIngestionJobs(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataIngestionJobsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentDataIngestionJobsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.ListDataIngestionJobsResponse
}

func (s *GenerativeAiAgentDataIngestionJobsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentDataIngestionJobsDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.ListDataIngestionJobsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSourceId, ok := s.D.GetOkExists("data_source_id"); ok {
		tmp := dataSourceId.(string)
		request.DataSourceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_generative_ai_agent.DataIngestionJobLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.ListDataIngestionJobs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataIngestionJobs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GenerativeAiAgentDataIngestionJobsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenerativeAiAgentDataIngestionJobsDataSource-", GenerativeAiAgentDataIngestionJobsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dataIngestionJob := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataIngestionJobSummaryToMap(item))
	}
	dataIngestionJob["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GenerativeAiAgentDataIngestionJobsDataSource().Schema["data_ingestion_job_collection"].Elem.(*schema.Resource).Schema)
		dataIngestionJob["items"] = items
	}

	resources = append(resources, dataIngestionJob)
	if err := s.D.Set("data_ingestion_job_collection", resources); err != nil {
		return err
	}

	return nil
}
