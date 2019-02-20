// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreCrossConnectDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreCrossConnect,
		Schema: map[string]*schema.Schema{
			"cross_connect_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cross_connect_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_speed_shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreCrossConnect(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type CoreCrossConnectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectResponse
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnect(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LocationName != nil {
		s.D.Set("location_name", *s.Res.LocationName)
	}

	if s.Res.PortName != nil {
		s.D.Set("port_name", *s.Res.PortName)
	}

	if s.Res.PortSpeedShapeName != nil {
		s.D.Set("port_speed_shape_name", *s.Res.PortSpeedShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
