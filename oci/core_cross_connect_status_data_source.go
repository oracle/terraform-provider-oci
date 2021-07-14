// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v44/core"
)

func init() {
	RegisterDatasource("oci_core_cross_connect_status", CoreCrossConnectStatusDataSource())
}

func CoreCrossConnectStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreCrossConnectStatus,
		Schema: map[string]*schema.Schema{
			"cross_connect_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"interface_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"light_level_ind_bm": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"light_level_indicator": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreCrossConnectStatus(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreCrossConnectStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCrossConnectStatusResponse
}

func (s *CoreCrossConnectStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCrossConnectStatusDataSourceCrud) Get() error {
	request := oci_core.GetCrossConnectStatusRequest{}

	if crossConnectId, ok := s.D.GetOkExists("cross_connect_id"); ok {
		tmp := crossConnectId.(string)
		request.CrossConnectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetCrossConnectStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreCrossConnectStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreCrossConnectStatusDataSource-", CoreCrossConnectStatusDataSource(), s.D))

	s.D.Set("interface_state", s.Res.InterfaceState)

	if s.Res.LightLevelIndBm != nil {
		s.D.Set("light_level_ind_bm", *s.Res.LightLevelIndBm)
	}

	s.D.Set("light_level_indicator", s.Res.LightLevelIndicator)

	return nil
}
