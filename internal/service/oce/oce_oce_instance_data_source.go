// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oce

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_oce "github.com/oracle/oci-go-sdk/v65/oce"
)

func OceOceInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oce_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OceOceInstanceResource(), fieldMap, readSingularOceOceInstance)
}

func readSingularOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OceInstanceClient()

	return tfresource.ReadResource(sync)
}

type OceOceInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oce.OceInstanceClient
	Res    *oci_oce.GetOceInstanceResponse
}

func (s *OceOceInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OceOceInstanceDataSourceCrud) Get() error {
	request := oci_oce.GetOceInstanceRequest{}

	if oceInstanceId, ok := s.D.GetOkExists("oce_instance_id"); ok {
		tmp := oceInstanceId.(string)
		request.OceInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oce")

	response, err := s.Client.GetOceInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OceOceInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("add_on_features", s.Res.AddOnFeatures)

	if s.Res.AdminEmail != nil {
		s.D.Set("admin_email", *s.Res.AdminEmail)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DrRegion != nil {
		s.D.Set("dr_region", *s.Res.DrRegion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Guid != nil {
		s.D.Set("guid", *s.Res.Guid)
	}

	if s.Res.IdcsTenancy != nil {
		s.D.Set("idcs_tenancy", *s.Res.IdcsTenancy)
	}

	s.D.Set("instance_access_type", s.Res.InstanceAccessType)

	s.D.Set("instance_license_type", s.Res.InstanceLicenseType)

	s.D.Set("instance_usage_type", s.Res.InstanceUsageType)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStorageNamespace != nil {
		s.D.Set("object_storage_namespace", *s.Res.ObjectStorageNamespace)
	}

	s.D.Set("service", tfresource.GenericMapToJsonMap(s.Res.Service))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TenancyName != nil {
		s.D.Set("tenancy_name", *s.Res.TenancyName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("upgrade_schedule", s.Res.UpgradeSchedule)

	if s.Res.WafPrimaryDomain != nil {
		s.D.Set("waf_primary_domain", *s.Res.WafPrimaryDomain)
	}

	return nil
}
