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

func FusionAppsFusionEnvironmentFamilyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fusion_environment_family_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FusionAppsFusionEnvironmentFamilyResource(), fieldMap, readSingularFusionAppsFusionEnvironmentFamily)
}

func readSingularFusionAppsFusionEnvironmentFamily(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentFamilyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetFusionEnvironmentFamilyResponse
}

func (s *FusionAppsFusionEnvironmentFamilyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentFamilyDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentFamilyRequest{}

	if fusionEnvironmentFamilyId, ok := s.D.GetOkExists("fusion_environment_family_id"); ok {
		tmp := fusionEnvironmentFamilyId.(string)
		request.FusionEnvironmentFamilyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetFusionEnvironmentFamily(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentFamilyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FamilyMaintenancePolicy != nil {
		s.D.Set("family_maintenance_policy", []interface{}{FamilyMaintenancePolicyToMap(s.Res.FamilyMaintenancePolicy)})
	} else {
		s.D.Set("family_maintenance_policy", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSubscriptionUpdateNeeded != nil {
		s.D.Set("is_subscription_update_needed", *s.Res.IsSubscriptionUpdateNeeded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscription_ids", s.Res.SubscriptionIds)

	if s.Res.SystemName != nil {
		s.D.Set("system_name", *s.Res.SystemName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
