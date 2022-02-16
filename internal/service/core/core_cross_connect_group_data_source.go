// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreCrossConnectGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cross_connect_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreCrossConnectGroupResource(), fieldMap, readSingularCoreCrossConnectGroup)
}

func readSingularCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCrossConnectGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectGroupResponse
}

func (s *CoreCrossConnectGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCrossConnectGroupDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectGroupRequest{}

	if crossConnectGroupId, ok := s.D.GetOkExists("cross_connect_group_id"); ok {
		tmp := crossConnectGroupId.(string)
		request.CrossConnectGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreCrossConnectGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomerReferenceName != nil {
		s.D.Set("customer_reference_name", *s.Res.CustomerReferenceName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MacsecProperties != nil {
		s.D.Set("macsec_properties", []interface{}{MacsecPropertiesToMap(s.Res.MacsecProperties)})
	} else {
		s.D.Set("macsec_properties", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
