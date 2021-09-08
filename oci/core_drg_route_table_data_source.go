// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v47/core"
)

func init() {
	RegisterDatasource("oci_core_drg_route_table", CoreDrgRouteTableDataSource())
}

func CoreDrgRouteTableDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["drg_route_table_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CoreDrgRouteTableResource(), fieldMap, readSingularCoreDrgRouteTable)
}

func readSingularCoreDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreDrgRouteTableDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetDrgRouteTableResponse
}

func (s *CoreDrgRouteTableDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgRouteTableDataSourceCrud) Get() error {
	request := oci_core.GetDrgRouteTableRequest{}

	if drgRouteTableId, ok := s.D.GetOkExists("drg_route_table_id"); ok {
		tmp := drgRouteTableId.(string)
		request.DrgRouteTableId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetDrgRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreDrgRouteTableDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImportDrgRouteDistributionId != nil {
		s.D.Set("import_drg_route_distribution_id", *s.Res.ImportDrgRouteDistributionId)
	}

	if s.Res.IsEcmpEnabled != nil {
		s.D.Set("is_ecmp_enabled", *s.Res.IsEcmpEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
