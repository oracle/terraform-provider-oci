// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsMigrationPlanDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["migration_plan_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudMigrationsMigrationPlanResource(), fieldMap, readSingularCloudMigrationsMigrationPlan)
}

func readSingularCloudMigrationsMigrationPlan(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsMigrationPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.GetMigrationPlanResponse
}

func (s *CloudMigrationsMigrationPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsMigrationPlanDataSourceCrud) Get() error {
	request := oci_cloud_migrations.GetMigrationPlanRequest{}

	if migrationPlanId, ok := s.D.GetOkExists("migration_plan_id"); ok {
		tmp := migrationPlanId.(string)
		request.MigrationPlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.GetMigrationPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudMigrationsMigrationPlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("calculated_limits", s.Res.CalculatedLimits)
	s.D.Set("calculated_limits", s.Res.CalculatedLimits)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MigrationId != nil {
		s.D.Set("migration_id", *s.Res.MigrationId)
	}

	if s.Res.MigrationPlanStats != nil {
		s.D.Set("migration_plan_stats", []interface{}{MigrationPlanStatsToMap(s.Res.MigrationPlanStats)})
	} else {
		s.D.Set("migration_plan_stats", nil)
	}

	if s.Res.ReferenceToRmsStack != nil {
		s.D.Set("reference_to_rms_stack", *s.Res.ReferenceToRmsStack)
	}

	if s.Res.SourceMigrationPlanId != nil {
		s.D.Set("source_migration_plan_id", *s.Res.SourceMigrationPlanId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	strategies := []interface{}{}
	for _, item := range s.Res.Strategies {
		strategies = append(strategies, ResourceAssessmentStrategyToMap(item))
	}
	s.D.Set("strategies", strategies)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	targetEnvironments := []interface{}{}
	for _, item := range s.Res.TargetEnvironments {
		targetEnvironments = append(targetEnvironments, TargetEnvironmentToMap(item))
	}
	s.D.Set("target_environments", targetEnvironments)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
