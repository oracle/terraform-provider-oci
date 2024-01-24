// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreNetworkSecurityGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_security_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreNetworkSecurityGroupResource(), fieldMap, readSingularCoreNetworkSecurityGroup)
}

func readSingularCoreNetworkSecurityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreNetworkSecurityGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreNetworkSecurityGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetNetworkSecurityGroupResponse
}

func (s *CoreNetworkSecurityGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreNetworkSecurityGroupDataSourceCrud) Get() error {
	request := oci_core.GetNetworkSecurityGroupRequest{}

	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetNetworkSecurityGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreNetworkSecurityGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
