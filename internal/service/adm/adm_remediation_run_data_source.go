// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["remediation_run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AdmRemediationRunResource(), fieldMap, readSingularAdmRemediationRun)
}

func readSingularAdmRemediationRun(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmRemediationRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.GetRemediationRunResponse
}

func (s *AdmRemediationRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmRemediationRunDataSourceCrud) Get() error {
	request := oci_adm.GetRemediationRunRequest{}

	if remediationRunId, ok := s.D.GetOkExists("remediation_run_id"); ok {
		tmp := remediationRunId.(string)
		request.RemediationRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.GetRemediationRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AdmRemediationRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("current_stage_type", s.Res.CurrentStageType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RemediationRecipeId != nil {
		s.D.Set("remediation_recipe_id", *s.Res.RemediationRecipeId)
	}

	s.D.Set("remediation_run_source", s.Res.RemediationRunSource)

	stages := []interface{}{}
	for _, item := range s.Res.Stages {
		stages = append(stages, StageSummaryToMap(item))
	}
	s.D.Set("stages", stages)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
