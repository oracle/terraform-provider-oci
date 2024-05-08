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

func CloudGuardWlpAgentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["wlp_agent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardWlpAgentResource(), fieldMap, readSingularCloudGuardWlpAgent)
}

func readSingularCloudGuardWlpAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardWlpAgentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardWlpAgentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetWlpAgentResponse
}

func (s *CloudGuardWlpAgentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardWlpAgentDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetWlpAgentRequest{}

	if wlpAgentId, ok := s.D.GetOkExists("wlp_agent_id"); ok {
		tmp := wlpAgentId.(string)
		request.WlpAgentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetWlpAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardWlpAgentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgentVersion != nil {
		s.D.Set("agent_version", *s.Res.AgentVersion)
	}

	if s.Res.CertificateId != nil {
		s.D.Set("certificate_id", *s.Res.CertificateId)
	}

	if s.Res.CertificateSignedRequest != nil {
		s.D.Set("certificate_signed_request", *s.Res.CertificateSignedRequest)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostId != nil {
		s.D.Set("host_id", *s.Res.HostId)
	}

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

	return nil
}
