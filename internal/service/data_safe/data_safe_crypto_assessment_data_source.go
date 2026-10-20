// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["crypto_assessment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DataSafeCryptoAssessmentResource(), fieldMap, readSingularDataSafeCryptoAssessmentWithContext)
}

func readSingularDataSafeCryptoAssessmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetCryptoAssessmentResponse
}

func (s *DataSafeCryptoAssessmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetCryptoAssessmentRequest{}

	if cryptoAssessmentId, ok := s.D.GetOkExists("crypto_assessment_id"); ok {
		tmp := cryptoAssessmentId.(string)
		request.CryptoAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetCryptoAssessment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeCryptoAssessmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CryptoPosture != nil {
		s.D.Set("crypto_posture", []interface{}{CryptoPostureToMap(s.Res.CryptoPosture)})
	} else {
		s.D.Set("crypto_posture", nil)
	}

	if s.Res.CryptoProvider != nil {
		s.D.Set("crypto_provider", *s.Res.CryptoProvider)
	}

	if s.Res.DatabaseArchitecture != nil {
		s.D.Set("database_architecture", *s.Res.DatabaseArchitecture)
	}

	if s.Res.DatabaseName != nil {
		s.D.Set("database_name", *s.Res.DatabaseName)
	}

	if s.Res.DatabaseVersion != nil {
		s.D.Set("database_version", *s.Res.DatabaseVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAssessmentScheduled != nil {
		s.D.Set("is_assessment_scheduled", *s.Res.IsAssessmentScheduled)
	}

	if s.Res.IssueCount != nil {
		s.D.Set("issue_count", *s.Res.IssueCount)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("posture_category", s.Res.PostureCategory)

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDatabaseGroupId != nil {
		s.D.Set("target_database_group_id", *s.Res.TargetDatabaseGroupId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	s.D.Set("target_type", s.Res.TargetType)

	if s.Res.TargetsWithIssuesCount != nil {
		s.D.Set("targets_with_issues_count", *s.Res.TargetsWithIssuesCount)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastAssessed != nil {
		s.D.Set("time_last_assessed", s.Res.TimeLastAssessed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("triggered_by", s.Res.TriggeredBy)

	s.D.Set("type", s.Res.Type)

	return nil
}
