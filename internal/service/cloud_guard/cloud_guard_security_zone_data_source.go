// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardSecurityZoneDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["security_zone_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardSecurityZoneResource(), fieldMap, readSingularCloudGuardSecurityZone)
}

func readSingularCloudGuardSecurityZone(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardSecurityZoneDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardSecurityZoneDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetSecurityZoneResponse
}

func (s *CloudGuardSecurityZoneDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardSecurityZoneDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetSecurityZoneRequest{}

	if securityZoneId, ok := s.D.GetOkExists("security_zone_id"); ok {
		tmp := securityZoneId.(string)
		request.SecurityZoneId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetSecurityZone(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardSecurityZoneDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("inherited_by_compartments", s.Res.InheritedByCompartments)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecurityZoneRecipeId != nil {
		s.D.Set("security_zone_recipe_id", *s.Res.SecurityZoneRecipeId)
	}

	if s.Res.SecurityZoneTargetId != nil {
		s.D.Set("security_zone_target_id", *s.Res.SecurityZoneTargetId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
