// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v37/core"
)

func init() {
	RegisterDatasource("oci_core_vlan", CoreVlanDataSource())
}

func CoreVlanDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vlan_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CoreVlanResource(), fieldMap, readSingularCoreVlan)
}

func readSingularCoreVlan(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreVlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVlanResponse
}

func (s *CoreVlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVlanDataSourceCrud) Get() error {
	request := oci_core.GetVlanRequest{}

	if vlanId, ok := s.D.GetOkExists("vlan_id"); ok {
		tmp := vlanId.(string)
		request.VlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetVlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VlanTag != nil {
		s.D.Set("vlan_tag", *s.Res.VlanTag)
	}

	return nil
}
