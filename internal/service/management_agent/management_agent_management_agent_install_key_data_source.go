// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentInstallKeyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["management_agent_install_key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ManagementAgentManagementAgentInstallKeyResource(), fieldMap, readSingularManagementAgentManagementAgentInstallKey)
}

func readSingularManagementAgentManagementAgentInstallKey(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentInstallKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentInstallKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.GetManagementAgentInstallKeyResponse
}

func (s *ManagementAgentManagementAgentInstallKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentInstallKeyDataSourceCrud) Get() error {
	request := oci_management_agent.GetManagementAgentInstallKeyRequest{}

	if managementAgentInstallKeyId, ok := s.D.GetOkExists("management_agent_install_key_id"); ok {
		tmp := managementAgentInstallKeyId.(string)
		request.ManagementAgentInstallKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.GetManagementAgentInstallKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentInstallKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AllowedKeyInstallCount != nil {
		s.D.Set("allowed_key_install_count", *s.Res.AllowedKeyInstallCount)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedByPrincipalId != nil {
		s.D.Set("created_by_principal_id", *s.Res.CreatedByPrincipalId)
	}

	if s.Res.CurrentKeyInstallCount != nil {
		s.D.Set("current_key_install_count", *s.Res.CurrentKeyInstallCount)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
