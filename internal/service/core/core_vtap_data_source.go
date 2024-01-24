// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreVtapDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vtap_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreVtapResource(), fieldMap, readSingularCoreVtap)
}

func readSingularCoreVtap(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVtapDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVtapDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVtapResponse
}

func (s *CoreVtapDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVtapDataSourceCrud) Get() error {
	request := oci_core.GetVtapRequest{}

	if vtapId, ok := s.D.GetOkExists("vtap_id"); ok {
		tmp := vtapId.(string)
		request.VtapId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVtap(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVtapDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CaptureFilterId != nil {
		s.D.Set("capture_filter_id", *s.Res.CaptureFilterId)
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

	s.D.Set("encapsulation_protocol", s.Res.EncapsulationProtocol)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsVtapEnabled != nil {
		s.D.Set("is_vtap_enabled", *s.Res.IsVtapEnabled)
	}

	s.D.Set("lifecycle_state_details", s.Res.LifecycleStateDetails)

	if s.Res.MaxPacketSize != nil {
		s.D.Set("max_packet_size", *s.Res.MaxPacketSize)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	if s.Res.SourcePrivateEndpointIp != nil {
		s.D.Set("source_private_endpoint_ip", *s.Res.SourcePrivateEndpointIp)
	}

	if s.Res.SourcePrivateEndpointSubnetId != nil {
		s.D.Set("source_private_endpoint_subnet_id", *s.Res.SourcePrivateEndpointSubnetId)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TargetIp != nil {
		s.D.Set("target_ip", *s.Res.TargetIp)
	}

	s.D.Set("target_type", s.Res.TargetType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("traffic_mode", s.Res.TrafficMode)

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VxlanNetworkIdentifier != nil {
		s.D.Set("vxlan_network_identifier", strconv.FormatInt(*s.Res.VxlanNetworkIdentifier, 10))
	}

	return nil
}
