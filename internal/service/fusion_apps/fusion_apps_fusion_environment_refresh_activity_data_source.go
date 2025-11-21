// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentRefreshActivityDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fusion_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["refresh_activity_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FusionAppsFusionEnvironmentRefreshActivityResource(), fieldMap, readSingularFusionAppsFusionEnvironmentRefreshActivity)
}

func readSingularFusionAppsFusionEnvironmentRefreshActivity(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentRefreshActivityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentRefreshActivityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetRefreshActivityResponse
}

func (s *FusionAppsFusionEnvironmentRefreshActivityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentRefreshActivityDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetRefreshActivityRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	// the refresh_activity_id is a composite id
	if refreshActivityId, ok := s.D.GetOkExists("refresh_activity_id"); ok {
		_, refreshActivityIdString, err := parseFusionEnvironmentRefreshActivityCompositeId(refreshActivityId.(string))
		if err == nil {
			request.RefreshActivityId = &refreshActivityIdString
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetRefreshActivity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentRefreshActivityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsDataMaskingOpted != nil {
		s.D.Set("is_data_masking_opted", *s.Res.IsDataMaskingOpted)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	refreshIssueDetailsList := []interface{}{}
	for _, item := range s.Res.RefreshIssueDetailsList {
		refreshIssueDetailsList = append(refreshIssueDetailsList, RefreshIssueDetailsToMap(item))
	}
	s.D.Set("refresh_issue_details_list", refreshIssueDetailsList)

	s.D.Set("service_availability", s.Res.ServiceAvailability)

	if s.Res.SourceFusionEnvironmentId != nil {
		s.D.Set("source_fusion_environment_id", *s.Res.SourceFusionEnvironmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeExpectedFinish != nil {
		s.D.Set("time_expected_finish", s.Res.TimeExpectedFinish.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeOfRestorationPoint != nil {
		s.D.Set("time_of_restoration_point", s.Res.TimeOfRestorationPoint.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
