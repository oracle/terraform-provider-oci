// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bastion "github.com/oracle/oci-go-sdk/v56/bastion"
)

func BastionBastionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bastion_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BastionBastionResource(), fieldMap, readSingularBastionBastion)
}

func readSingularBastionBastion(d *schema.ResourceData, m interface{}) error {
	sync := &BastionBastionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.ReadResource(sync)
}

type BastionBastionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bastion.BastionClient
	Res    *oci_bastion.GetBastionResponse
}

func (s *BastionBastionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BastionBastionDataSourceCrud) Get() error {
	request := oci_bastion.GetBastionRequest{}

	if bastionId, ok := s.D.GetOkExists("bastion_id"); ok {
		tmp := bastionId.(string)
		request.BastionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bastion")

	response, err := s.Client.GetBastion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BastionBastionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BastionType != nil {
		s.D.Set("bastion_type", *s.Res.BastionType)
	}

	s.D.Set("client_cidr_block_allow_list", s.Res.ClientCidrBlockAllowList)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxSessionTtlInSeconds != nil {
		s.D.Set("max_session_ttl_in_seconds", *s.Res.MaxSessionTtlInSeconds)
	}

	if s.Res.MaxSessionsAllowed != nil {
		s.D.Set("max_sessions_allowed", *s.Res.MaxSessionsAllowed)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PhoneBookEntry != nil {
		s.D.Set("phone_book_entry", *s.Res.PhoneBookEntry)
	}

	if s.Res.PrivateEndpointIpAddress != nil {
		s.D.Set("private_endpoint_ip_address", *s.Res.PrivateEndpointIpAddress)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("static_jump_host_ip_addresses", s.Res.StaticJumpHostIpAddresses)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetSubnetId != nil {
		s.D.Set("target_subnet_id", *s.Res.TargetSubnetId)
	}

	if s.Res.TargetVcnId != nil {
		s.D.Set("target_vcn_id", *s.Res.TargetVcnId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
