// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["monitored_resource_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMonitoredResourceResource(), fieldMap, readSingularStackMonitoringMonitoredResource)
}

func readSingularStackMonitoringMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoredResourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetMonitoredResourceResponse
}

func (s *StackMonitoringMonitoredResourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoredResourceDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoredResourceRequest{}

	if monitoredResourceId, ok := s.D.GetOkExists("monitored_resource_id"); ok {
		tmp := monitoredResourceId.(string)
		request.MonitoredResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetMonitoredResource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMonitoredResourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Aliases != nil {
		s.D.Set("aliases", []interface{}{MonitoredResourceAliasCredentialToMap(s.Res.Aliases)})
	} else {
		s.D.Set("aliases", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Credentials != nil {
		credentialsArray := []interface{}{}
		if credentialsMap := MonitoredResourceCredentialToMap(&s.Res.Credentials); credentialsMap != nil {
			credentialsArray = append(credentialsArray, credentialsMap)
		}
		s.D.Set("credentials", credentialsArray)
	} else {
		s.D.Set("credentials", nil)
	}

	if s.Res.DatabaseConnectionDetails != nil {
		s.D.Set("database_connection_details", []interface{}{ConnectionDetailsToMap(s.Res.DatabaseConnectionDetails)})
	} else {
		s.D.Set("database_connection_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	s.D.Set("license", s.Res.License)

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	properties := []interface{}{}
	for _, item := range s.Res.Properties {
		properties = append(properties, MonitoredResourcePropertyToMap(item))
	}
	s.D.Set("properties", properties)

	if s.Res.ResourceTimeZone != nil {
		s.D.Set("resource_time_zone", *s.Res.ResourceTimeZone)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}
