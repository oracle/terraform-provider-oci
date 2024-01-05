// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
)

func DatascienceJobRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceJobRuns,
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
			"job_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"job_runs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceJobRunResource()),
			},
		},
	}
}

func readDatascienceJobRuns(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceJobRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListJobRunsResponse
}

func (s *DatascienceJobRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceJobRunsDataSourceCrud) Get() error {
	request := oci_datascience.ListJobRunsRequest{}

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

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListJobRunsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListJobRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJobRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceJobRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceJobRunsDataSource-", DatascienceJobRunsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		jobRun := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			jobRun["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			jobRun["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			jobRun["display_name"] = *r.DisplayName
		}

		jobRun["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			jobRun["id"] = *r.Id
		}

		if r.JobId != nil {
			jobRun["job_id"] = *r.JobId
		}

		if r.LifecycleDetails != nil {
			jobRun["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ProjectId != nil {
			jobRun["project_id"] = *r.ProjectId
		}

		jobRun["state"] = r.LifecycleState

		if r.TimeAccepted != nil {
			jobRun["time_accepted"] = r.TimeAccepted.String()
		}

		if r.TimeFinished != nil {
			jobRun["time_finished"] = r.TimeFinished.String()
		}

		if r.TimeStarted != nil {
			jobRun["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, jobRun)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceJobRunsDataSource().Schema["job_runs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("job_runs", resources); err != nil {
		return err
	}

	return nil
}
