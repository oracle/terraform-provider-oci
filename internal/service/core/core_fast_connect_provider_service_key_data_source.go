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

func CoreFastConnectProviderServiceKeyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreFastConnectProviderServiceKey,
		Schema: map[string]*schema.Schema{
			"provider_service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_service_key_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"bandwidth_shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peering_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreFastConnectProviderServiceKey(d *schema.ResourceData, m interface{}) error {
	sync := &CoreFastConnectProviderServiceKeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreFastConnectProviderServiceKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetFastConnectProviderServiceKeyResponse
}

func (s *CoreFastConnectProviderServiceKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreFastConnectProviderServiceKeyDataSourceCrud) Get() error {
	request := oci_core.GetFastConnectProviderServiceKeyRequest{}

	if providerServiceId, ok := s.D.GetOkExists("provider_service_id"); ok {
		tmp := providerServiceId.(string)
		request.ProviderServiceId = &tmp
	}

	if providerServiceKeyName, ok := s.D.GetOkExists("provider_service_key_name"); ok {
		tmp := providerServiceKeyName.(string)
		request.ProviderServiceKeyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetFastConnectProviderServiceKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreFastConnectProviderServiceKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreFastConnectProviderServiceKeyDataSource-", CoreFastConnectProviderServiceKeyDataSource(), s.D))

	if s.Res.BandwidthShapeName != nil {
		s.D.Set("bandwidth_shape_name", *s.Res.BandwidthShapeName)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PeeringLocation != nil {
		s.D.Set("peering_location", *s.Res.PeeringLocation)
	}

	return nil
}
