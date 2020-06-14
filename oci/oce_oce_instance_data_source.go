// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_oce "github.com/oracle/oci-go-sdk/oce"
)

func init() {
	RegisterDatasource("oci_oce_oce_instance", OceOceInstanceDataSource())
}

func OceOceInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oce_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OceOceInstanceResource(), fieldMap, readSingularOceOceInstance)
}

func readSingularOceOceInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).oceInstanceClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "oce")

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

	if s.Res.AdminEmail != nil {
		s.D.Set("admin_email", *s.Res.AdminEmail)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStorageNamespace != nil {
		s.D.Set("object_storage_namespace", *s.Res.ObjectStorageNamespace)
	}

	s.D.Set("service", genericMapToJsonMap(s.Res.Service))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
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
