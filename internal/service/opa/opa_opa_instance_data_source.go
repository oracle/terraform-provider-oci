// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opa "github.com/oracle/oci-go-sdk/v65/opa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpaOpaInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["opa_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpaOpaInstanceResource(), fieldMap, readSingularOpaOpaInstance)
}

func readSingularOpaOpaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OpaOpaInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpaInstanceClient()

	return tfresource.ReadResource(sync)
}

type OpaOpaInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opa.OpaInstanceClient
	Res    *oci_opa.GetOpaInstanceResponse
}

func (s *OpaOpaInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpaOpaInstanceDataSourceCrud) Get() error {
	request := oci_opa.GetOpaInstanceRequest{}

	if opaInstanceId, ok := s.D.GetOkExists("opa_instance_id"); ok {
		tmp := opaInstanceId.(string)
		request.OpaInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opa")

	response, err := s.Client.GetOpaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpaOpaInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	attachments := []interface{}{}
	for _, item := range s.Res.Attachments {
		attachments = append(attachments, AttachmentDetailsToMap(item))
	}
	s.D.Set("attachments", attachments)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

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

	if s.Res.IdentityAppDisplayName != nil {
		s.D.Set("identity_app_display_name", *s.Res.IdentityAppDisplayName)
	}

	if s.Res.IdentityAppGuid != nil {
		s.D.Set("identity_app_guid", *s.Res.IdentityAppGuid)
	}

	if s.Res.IdentityAppOpcServiceInstanceGuid != nil {
		s.D.Set("identity_app_opc_service_instance_guid", *s.Res.IdentityAppOpcServiceInstanceGuid)
	}

	if s.Res.IdentityDomainUrl != nil {
		s.D.Set("identity_domain_url", *s.Res.IdentityDomainUrl)
	}

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	if s.Res.IsBreakglassEnabled != nil {
		s.D.Set("is_breakglass_enabled", *s.Res.IsBreakglassEnabled)
	}

	s.D.Set("metering_type", s.Res.MeteringType)

	s.D.Set("shape_name", s.Res.ShapeName)

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

	return nil
}
