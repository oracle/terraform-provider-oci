// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"
)

func LoggingUnifiedAgentConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["unified_agent_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LoggingUnifiedAgentConfigurationResource(), fieldMap, readSingularLoggingUnifiedAgentConfiguration)
}

func readSingularLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

type LoggingUnifiedAgentConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.GetUnifiedAgentConfigurationResponse
}

func (s *LoggingUnifiedAgentConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingUnifiedAgentConfigurationDataSourceCrud) Get() error {
	request := oci_logging.GetUnifiedAgentConfigurationRequest{}

	if unifiedAgentConfigurationId, ok := s.D.GetOkExists("unified_agent_configuration_id"); ok {
		tmp := unifiedAgentConfigurationId.(string)
		request.UnifiedAgentConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "logging")

	response, err := s.Client.GetUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoggingUnifiedAgentConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("configuration_state", s.Res.ConfigurationState)

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

	if s.Res.GroupAssociation != nil {
		s.D.Set("group_association", []interface{}{GroupAssociationDetailsToMap(s.Res.GroupAssociation)})
	} else {
		s.D.Set("group_association", nil)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.ServiceConfiguration != nil {
		serviceConfigurationArray := []interface{}{}
		if serviceConfigurationMap := UnifiedAgentServiceConfigurationDetailsToMap(&s.Res.ServiceConfiguration); serviceConfigurationMap != nil {
			serviceConfigurationArray = append(serviceConfigurationArray, serviceConfigurationMap)
		}
		s.D.Set("service_configuration", serviceConfigurationArray)
	} else {
		s.D.Set("service_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}
