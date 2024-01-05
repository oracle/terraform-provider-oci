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

func DatascienceJobRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["job_run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatascienceJobRunResource(), fieldMap, readSingularDatascienceJobRun)
}

func readSingularDatascienceJobRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceJobRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetJobRunResponse
}

func (s *DatascienceJobRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceJobRunDataSourceCrud) Get() error {
	request := oci_datascience.GetJobRunRequest{}

	if jobRunId, ok := s.D.GetOkExists("job_run_id"); ok {
		tmp := jobRunId.(string)
		request.JobRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetJobRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceJobRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.JobConfigurationOverrideDetails != nil {
		jobConfigurationOverrideDetailsArray := []interface{}{}
		if jobConfigurationOverrideDetailsMap := JobConfigurationDetailsToMap(&s.Res.JobConfigurationOverrideDetails); jobConfigurationOverrideDetailsMap != nil {
			jobConfigurationOverrideDetailsArray = append(jobConfigurationOverrideDetailsArray, jobConfigurationOverrideDetailsMap)
		}
		s.D.Set("job_configuration_override_details", jobConfigurationOverrideDetailsArray)
	} else {
		s.D.Set("job_configuration_override_details", nil)
	}

	if s.Res.JobId != nil {
		s.D.Set("job_id", *s.Res.JobId)
	}

	if s.Res.JobInfrastructureConfigurationDetails != nil {
		jobInfrastructureConfigurationDetailsArray := []interface{}{}
		if jobInfrastructureConfigurationDetailsMap := JobInfrastructureConfigurationDetailsToMap(&s.Res.JobInfrastructureConfigurationDetails); jobInfrastructureConfigurationDetailsMap != nil {
			jobInfrastructureConfigurationDetailsArray = append(jobInfrastructureConfigurationDetailsArray, jobInfrastructureConfigurationDetailsMap)
		}
		s.D.Set("job_infrastructure_configuration_details", jobInfrastructureConfigurationDetailsArray)
	} else {
		s.D.Set("job_infrastructure_configuration_details", nil)
	}

	if s.Res.JobLogConfigurationOverrideDetails != nil {
		s.D.Set("job_log_configuration_override_details", []interface{}{JobLogConfigurationDetailsToMap(s.Res.JobLogConfigurationOverrideDetails)})
	} else {
		s.D.Set("job_log_configuration_override_details", nil)
	}

	jobStorageMountConfigurationDetailsList := []interface{}{}
	for _, item := range s.Res.JobStorageMountConfigurationDetailsList {
		jobStorageMountConfigurationDetailsList = append(jobStorageMountConfigurationDetailsList, StorageMountConfigurationDetailsToMap(item))
	}
	s.D.Set("job_storage_mount_configuration_details_list", jobStorageMountConfigurationDetailsList)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogDetails != nil {
		s.D.Set("log_details", []interface{}{JobRunLogDetailsToMap(s.Res.LogDetails)})
	} else {
		s.D.Set("log_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
