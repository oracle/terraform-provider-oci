// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatasciencePipelineRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatasciencePipelineRuns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pipeline_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pipeline_runs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatasciencePipelineRunResource()),
			},
		},
	}
}

func readDatasciencePipelineRuns(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatasciencePipelineRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListPipelineRunsResponse
}

func (s *DatasciencePipelineRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatasciencePipelineRunsDataSourceCrud) Get() error {
	request := oci_datascience.ListPipelineRunsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createdBy, ok := s.D.GetOkExists("created_by"); ok {
		tmp := createdBy.(string)
		request.CreatedBy = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListPipelineRunsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListPipelineRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPipelineRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatasciencePipelineRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatasciencePipelineRunsDataSource-", DatasciencePipelineRunsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pipelineRun := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			pipelineRun["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			pipelineRun["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			pipelineRun["display_name"] = *r.DisplayName
		}

		pipelineRun["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			pipelineRun["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			pipelineRun["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PipelineId != nil {
			pipelineRun["pipeline_id"] = *r.PipelineId
		}

		if r.ProjectId != nil {
			pipelineRun["project_id"] = *r.ProjectId
		}

		pipelineRun["state"] = r.LifecycleState

		if r.SystemTags != nil {
			pipelineRun["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeAccepted != nil {
			pipelineRun["time_accepted"] = r.TimeAccepted.String()
		}

		if r.TimeFinished != nil {
			pipelineRun["time_finished"] = r.TimeFinished.String()
		}

		if r.TimeStarted != nil {
			pipelineRun["time_started"] = r.TimeStarted.String()
		}

		if r.TimeUpdated != nil {
			pipelineRun["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, pipelineRun)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatasciencePipelineRunsDataSource().Schema["pipeline_runs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("pipeline_runs", resources); err != nil {
		return err
	}

	return nil
}
