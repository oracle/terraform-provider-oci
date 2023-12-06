// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func KmsEkmsPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ekms_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(KmsEkmsPrivateEndpointResource(), fieldMap, readSingularKmsEkmsPrivateEndpoint)
}

func readSingularKmsEkmsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEkmsPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EkmClient()

	return tfresource.ReadResource(sync)
}

type KmsEkmsPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.EkmClient
	Res    *oci_kms.GetEkmsPrivateEndpointResponse
}

func (s *KmsEkmsPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsEkmsPrivateEndpointDataSourceCrud) Get() error {
	request := oci_kms.GetEkmsPrivateEndpointRequest{}

	if ekmsPrivateEndpointId, ok := s.D.GetOkExists("ekms_private_endpoint_id"); ok {
		tmp := ekmsPrivateEndpointId.(string)
		request.EkmsPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.GetEkmsPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsEkmsPrivateEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CaBundle != nil {
		s.D.Set("ca_bundle", *s.Res.CaBundle)
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

	if s.Res.ExternalKeyManagerIp != nil {
		s.D.Set("external_key_manager_ip", *s.Res.ExternalKeyManagerIp)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
