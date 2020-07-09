// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_integration "github.com/oracle/oci-go-sdk/v25/integration"
)

func init() {
	RegisterDatasource("oci_integration_integration_instance", IntegrationIntegrationInstanceDataSource())
}

func IntegrationIntegrationInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["integration_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(IntegrationIntegrationInstanceResource(), fieldMap, readSingularIntegrationIntegrationInstance)
}

func readSingularIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).integrationInstanceClient()

	return ReadResource(sync)
}

type IntegrationIntegrationInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_integration.IntegrationInstanceClient
	Res    *oci_integration.GetIntegrationInstanceResponse
}

func (s *IntegrationIntegrationInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IntegrationIntegrationInstanceDataSourceCrud) Get() error {
	request := oci_integration.GetIntegrationInstanceRequest{}

	if integrationInstanceId, ok := s.D.GetOkExists("integration_instance_id"); ok {
		tmp := integrationInstanceId.(string)
		request.IntegrationInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "integration")

	response, err := s.Client.GetIntegrationInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IntegrationIntegrationInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	s.D.Set("integration_instance_type", s.Res.IntegrationInstanceType)

	if s.Res.IsByol != nil {
		s.D.Set("is_byol", *s.Res.IsByol)
	}

	if s.Res.IsFileServerEnabled != nil {
		s.D.Set("is_file_server_enabled", *s.Res.IsFileServerEnabled)
	}

	if s.Res.MessagePacks != nil {
		s.D.Set("message_packs", *s.Res.MessagePacks)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
