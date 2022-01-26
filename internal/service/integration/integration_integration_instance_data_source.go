// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integration

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_integration "github.com/oracle/oci-go-sdk/v56/integration"
)

func IntegrationIntegrationInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["integration_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(IntegrationIntegrationInstanceResource(), fieldMap, readSingularIntegrationIntegrationInstance)
}

func readSingularIntegrationIntegrationInstance(d *schema.ResourceData, m interface{}) error {
	sync := &IntegrationIntegrationInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IntegrationInstanceClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "integration")

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

	alternateCustomEndpoints := []interface{}{}
	for _, item := range s.Res.AlternateCustomEndpoints {
		alternateCustomEndpoints = append(alternateCustomEndpoints, CustomEndpointDetailsToMap(&item))
	}
	s.D.Set("alternate_custom_endpoints", alternateCustomEndpoints)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

	if s.Res.CustomEndpoint != nil {
		s.D.Set("custom_endpoint", []interface{}{CustomEndpointDetailsToMap(s.Res.CustomEndpoint)})
	} else {
		s.D.Set("custom_endpoint", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.IsVisualBuilderEnabled != nil {
		s.D.Set("is_visual_builder_enabled", *s.Res.IsVisualBuilderEnabled)
	}

	if s.Res.MessagePacks != nil {
		s.D.Set("message_packs", *s.Res.MessagePacks)
	}

	if s.Res.NetworkEndpointDetails != nil {
		networkEndpointDetailsArray := []interface{}{}
		if networkEndpointDetailsMap := IntegNetworkEndpointDetailsToMap(&s.Res.NetworkEndpointDetails, true); networkEndpointDetailsMap != nil {
			networkEndpointDetailsArray = append(networkEndpointDetailsArray, networkEndpointDetailsMap)
		}
		s.D.Set("network_endpoint_details", networkEndpointDetailsArray)
	} else {
		s.D.Set("network_endpoint_details", nil)
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
