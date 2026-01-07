// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpManagementApplianceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["management_appliance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OcvpManagementApplianceResource(), fieldMap, readSingularOcvpManagementAppliance)
}

func readSingularOcvpManagementAppliance(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpManagementApplianceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementApplianceClient()

	return tfresource.ReadResource(sync)
}

type OcvpManagementApplianceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.ManagementApplianceClient
	Res    *oci_ocvp.GetManagementApplianceResponse
}

func (s *OcvpManagementApplianceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpManagementApplianceDataSourceCrud) Get() error {
	request := oci_ocvp.GetManagementApplianceRequest{}

	if managementApplianceId, ok := s.D.GetOkExists("management_appliance_id"); ok {
		tmp := managementApplianceId.(string)
		request.ManagementApplianceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GetManagementAppliance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpManagementApplianceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeInstanceId != nil {
		s.D.Set("compute_instance_id", *s.Res.ComputeInstanceId)
	}

	if s.Res.Configuration != nil {
		s.D.Set("configuration", []interface{}{ManagementApplianceConfigurationToMap(s.Res.Configuration)})
	} else {
		s.D.Set("configuration", nil)
	}

	connections := []interface{}{}
	for _, item := range s.Res.Connections {
		connections = append(connections, ManagementApplianceConnectionToMap(item))
	}
	s.D.Set("connections", connections)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	heartbeatConnectionStates := []interface{}{}
	for _, item := range s.Res.HeartbeatConnectionStates {
		heartbeatConnectionStates = append(heartbeatConnectionStates, ManagementApplianceConnectionStatusToMap(item))
	}
	s.D.Set("heartbeat_connection_states", heartbeatConnectionStates)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeConfigurationUpdated != nil {
		s.D.Set("time_configuration_updated", s.Res.TimeConfigurationUpdated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastHeartbeat != nil {
		s.D.Set("time_last_heartbeat", s.Res.TimeLastHeartbeat.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
