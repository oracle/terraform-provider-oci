// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v58/jms"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func JmsFleetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsFleetResource(), fieldMap, readSingularJmsFleet)
}

func readSingularJmsFleet(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetFleetResponse
}

func (s *JmsFleetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetDataSourceCrud) Get() error {
	request := oci_jms.GetFleetRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetFleet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApproximateApplicationCount != nil {
		s.D.Set("approximate_application_count", *s.Res.ApproximateApplicationCount)
	}

	if s.Res.ApproximateInstallationCount != nil {
		s.D.Set("approximate_installation_count", *s.Res.ApproximateInstallationCount)
	}

	if s.Res.ApproximateJreCount != nil {
		s.D.Set("approximate_jre_count", *s.Res.ApproximateJreCount)
	}

	if s.Res.ApproximateManagedInstanceCount != nil {
		s.D.Set("approximate_managed_instance_count", *s.Res.ApproximateManagedInstanceCount)
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
