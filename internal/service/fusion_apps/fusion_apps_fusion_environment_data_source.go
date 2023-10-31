// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fusion_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FusionAppsFusionEnvironmentResource(), fieldMap, readSingularFusionAppsFusionEnvironment)
}

func readSingularFusionAppsFusionEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetFusionEnvironmentResponse
}

func (s *FusionAppsFusionEnvironmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentRequest{}

	if fusionEnvironmentId, ok := s.D.GetOkExists("fusion_environment_id"); ok {
		tmp := fusionEnvironmentId.(string)
		request.FusionEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetFusionEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_language_packs", s.Res.AdditionalLanguagePacks)

	s.D.Set("applied_patch_bundles", s.Res.AppliedPatchBundles)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsPrefix != nil {
		s.D.Set("dns_prefix", *s.Res.DnsPrefix)
	}

	if s.Res.DomainId != nil {
		s.D.Set("domain_id", *s.Res.DomainId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.FusionEnvironmentFamilyId != nil {
		s.D.Set("fusion_environment_family_id", *s.Res.FusionEnvironmentFamilyId)
	}

	s.D.Set("fusion_environment_type", s.Res.FusionEnvironmentType)

	if s.Res.IdcsDomainUrl != nil {
		s.D.Set("idcs_domain_url", *s.Res.IdcsDomainUrl)
	}

	if s.Res.IsBreakGlassEnabled != nil {
		s.D.Set("is_break_glass_enabled", *s.Res.IsBreakGlassEnabled)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyInfo != nil {
		s.D.Set("kms_key_info", []interface{}{objectToMap(s.Res.KmsKeyInfo)})
	} else {
		s.D.Set("kms_key_info", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LockboxId != nil {
		s.D.Set("lockbox_id", *s.Res.LockboxId)
	}

	if s.Res.MaintenancePolicy != nil {
		s.D.Set("maintenance_policy", []interface{}{GetMaintenancePolicyDetailsToMap(s.Res.MaintenancePolicy)})
	} else {
		s.D.Set("maintenance_policy", nil)
	}

	if s.Res.PublicUrl != nil {
		s.D.Set("public_url", *s.Res.PublicUrl)
	}

	if s.Res.Refresh != nil {
		s.D.Set("refresh", []interface{}{RefreshDetailsToMap(s.Res.Refresh)})
	} else {
		s.D.Set("refresh", nil)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, RuleToMap(item))
	}
	s.D.Set("rules", rules)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscription_ids", s.Res.SubscriptionIds)

	if s.Res.SystemName != nil {
		s.D.Set("system_name", *s.Res.SystemName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpcomingMaintenance != nil {
		s.D.Set("time_upcoming_maintenance", s.Res.TimeUpcomingMaintenance.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
