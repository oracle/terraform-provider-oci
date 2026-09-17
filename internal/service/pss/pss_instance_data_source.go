// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package pss

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_pss "github.com/oracle/oci-go-sdk/v65/pss"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PssInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(PssInstanceResource(), fieldMap, readSingularPssInstanceWithContext)
}

func readSingularPssInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &PssInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PssInstanceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type PssInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_pss.InstanceClient
	Res    *oci_pss.GetInstanceResponse
}

func (s *PssInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PssInstanceDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_pss.GetInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "pss")

	response, err := s.Client.GetInstance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PssInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdminIdentityDomainUserId != nil {
		s.D.Set("admin_identity_domain_user_id", *s.Res.AdminIdentityDomainUserId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdentityDomainUrl != nil {
		s.D.Set("identity_domain_url", *s.Res.IdentityDomainUrl)
	}

	if s.Res.IdentityDomainWebAuthClientId != nil {
		s.D.Set("identity_domain_web_auth_client_id", *s.Res.IdentityDomainWebAuthClientId)
	}

	if s.Res.IdentityDomainWebAuthClientSecretVaultId != nil {
		s.D.Set("identity_domain_web_auth_client_secret_vault_id", *s.Res.IdentityDomainWebAuthClientSecretVaultId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PublicIp != nil {
		s.D.Set("public_ip", *s.Res.PublicIp)
	}

	if s.Res.ServiceEndpoint != nil {
		s.D.Set("service_endpoint", *s.Res.ServiceEndpoint)
	}

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

	if s.Res.UiEndpoint != nil {
		s.D.Set("ui_endpoint", *s.Res.UiEndpoint)
	}

	return nil
}
