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

func CoreDrgRouteDistributionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["drg_route_distribution_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreDrgRouteDistributionResource(), fieldMap, readSingularCoreDrgRouteDistribution)
}

func readSingularCoreDrgRouteDistribution(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreDrgRouteDistributionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetDrgRouteDistributionResponse
}

func (s *CoreDrgRouteDistributionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgRouteDistributionDataSourceCrud) Get() error {
	request := oci_core.GetDrgRouteDistributionRequest{}

	if drgRouteDistributionId, ok := s.D.GetOkExists("drg_route_distribution_id"); ok {
		tmp := drgRouteDistributionId.(string)
		request.DrgRouteDistributionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetDrgRouteDistribution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreDrgRouteDistributionDataSourceCrud) SetData() error {
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

	s.D.Set("distribution_type", s.Res.DistributionType)

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
