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

func CoreCrossConnectDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cross_connect_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreCrossConnectResource(), fieldMap, readSingularCoreCrossConnect)
}

func readSingularCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCrossConnectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectResponse
	Loa    *oci_core.LetterOfAuthority
}

func (s *CoreCrossConnectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCrossConnectDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectRequest{}

	if crossConnectId, ok := s.D.GetOkExists("cross_connect_id"); ok {
		tmp := crossConnectId.(string)
		request.CrossConnectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	loa, err := getCrossConnectLetterOfAuthority(s.Client, *request.CrossConnectId, false)
	if err != nil {
		return err
	}
	s.Loa = loa
	return nil
}

func (s *CoreCrossConnectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CrossConnectGroupId != nil {
		s.D.Set("cross_connect_group_id", *s.Res.CrossConnectGroupId)
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

	if s.Res.InterfaceDownTimerValueInMilliseconds != nil {
		s.D.Set("interface_down_timer_value_in_milliseconds", *s.Res.InterfaceDownTimerValueInMilliseconds)
	}

	if s.Res.InterfaceName != nil {
		s.D.Set("interface_name", *s.Res.InterfaceName)
	}

	if s.Res.IsInterfaceHoldTimerEnabled != nil {
		s.D.Set("is_interface_hold_timer_enabled", *s.Res.IsInterfaceHoldTimerEnabled)
	}

	if s.Res.IsQosEnabled != nil {
		s.D.Set("is_qos_enabled", *s.Res.IsQosEnabled)
	}

	if s.Res.LocationName != nil {
		s.D.Set("location_name", *s.Res.LocationName)
	}

	if s.Res.MacsecProperties != nil {
		s.D.Set("macsec_properties", []interface{}{MacsecPropertiesToMap(s.Res.MacsecProperties)})
	} else {
		s.D.Set("macsec_properties", nil)
	}

	if s.Res.OciLogicalDeviceName != nil {
		s.D.Set("oci_logical_device_name", *s.Res.OciLogicalDeviceName)
	}

	if s.Res.OciPhysicalDeviceName != nil {
		s.D.Set("oci_physical_device_name", *s.Res.OciPhysicalDeviceName)
	}

	if s.Res.InterfaceName != nil {
		s.D.Set("interface_name", *s.Res.InterfaceName)
	}

	if s.Res.PortName != nil {
		s.D.Set("port_name", *s.Res.PortName)
	}

	if s.Res.PortSpeedShapeName != nil {
		s.D.Set("port_speed_shape_name", *s.Res.PortSpeedShapeName)
	}

	if hasMeaningfulLetterOfAuthorityProperties(s.Loa) {
		s.D.Set("loa_properties", []interface{}{LetterOfAuthorityPropertiesToMap(s.Loa)})
	} else {
		s.D.Set("loa_properties", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func getCrossConnectLetterOfAuthority(client *oci_core.VirtualNetworkClient, crossConnectId string, disableNotFoundRetries bool) (*oci_core.LetterOfAuthority, error) {
	request := oci_core.GetCrossConnectLetterOfAuthorityRequest{}
	request.CrossConnectId = &crossConnectId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(disableNotFoundRetries, "core")

	response, err := client.GetCrossConnectLetterOfAuthority(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return &response.LetterOfAuthority, nil
}
