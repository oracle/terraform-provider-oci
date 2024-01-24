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

func CoreVtapsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVtaps,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_vtap_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtaps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreVtapResource()),
			},
		},
	}
}

func readCoreVtaps(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVtapsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVtapsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVtapsResponse
}

func (s *CoreVtapsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVtapsDataSourceCrud) Get() error {
	request := oci_core.ListVtapsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isVtapEnabled, ok := s.D.GetOkExists("is_vtap_enabled"); ok {
		tmp := isVtapEnabled.(bool)
		request.IsVtapEnabled = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		tmp := source.(string)
		request.Source = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VtapLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if targetIp, ok := s.D.GetOkExists("target_ip"); ok {
		tmp := targetIp.(string)
		request.TargetIp = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListVtaps(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVtaps(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVtapsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVtapsDataSource-", CoreVtapsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vtap := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CaptureFilterId != nil {
			vtap["capture_filter_id"] = *r.CaptureFilterId
		}

		if r.DefinedTags != nil {
			vtap["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			vtap["display_name"] = *r.DisplayName
		}

		vtap["encapsulation_protocol"] = r.EncapsulationProtocol

		vtap["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			vtap["id"] = *r.Id
		}

		if r.IsVtapEnabled != nil {
			vtap["is_vtap_enabled"] = *r.IsVtapEnabled
		}

		vtap["lifecycle_state_details"] = r.LifecycleStateDetails

		if r.MaxPacketSize != nil {
			vtap["max_packet_size"] = *r.MaxPacketSize
		}

		if r.SourceId != nil {
			vtap["source_id"] = *r.SourceId
		}

		if r.SourcePrivateEndpointIp != nil {
			vtap["source_private_endpoint_ip"] = *r.SourcePrivateEndpointIp
		}

		if r.SourcePrivateEndpointSubnetId != nil {
			vtap["source_private_endpoint_subnet_id"] = *r.SourcePrivateEndpointSubnetId
		}

		vtap["source_type"] = r.SourceType

		vtap["state"] = r.LifecycleState

		if r.TargetId != nil {
			vtap["target_id"] = *r.TargetId
		}

		if r.TargetIp != nil {
			vtap["target_ip"] = *r.TargetIp
		}

		vtap["target_type"] = r.TargetType

		if r.TimeCreated != nil {
			vtap["time_created"] = r.TimeCreated.String()
		}

		vtap["traffic_mode"] = r.TrafficMode

		if r.VcnId != nil {
			vtap["vcn_id"] = *r.VcnId
		}

		if r.VxlanNetworkIdentifier != nil {
			vtap["vxlan_network_identifier"] = strconv.FormatInt(*r.VxlanNetworkIdentifier, 10)
		}

		resources = append(resources, vtap)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVtapsDataSource().Schema["vtaps"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vtaps", resources); err != nil {
		return err
	}

	return nil
}
