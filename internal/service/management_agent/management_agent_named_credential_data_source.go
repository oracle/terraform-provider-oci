// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentNamedCredentialDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["named_credential_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ManagementAgentNamedCredentialResource(), fieldMap, readSingularManagementAgentNamedCredential)
}

func readSingularManagementAgentNamedCredential(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentNamedCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentNamedCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.GetNamedCredentialResponse
}

func (s *ManagementAgentNamedCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentNamedCredentialDataSourceCrud) Get() error {
	request := oci_management_agent.GetNamedCredentialRequest{}

	if namedCredentialId, ok := s.D.GetOkExists("named_credential_id"); ok {
		tmp := namedCredentialId.(string)
		request.NamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.GetNamedCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentNamedCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	properties := []interface{}{}
	for _, item := range s.Res.Properties {
		properties = append(properties, NamedCredentialPropertyToMap(item))
	}
	s.D.Set("properties", properties)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
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
