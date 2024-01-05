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

func CoreVirtualCircuitPublicPrefixesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVirtualCircuitPublicPrefixes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"verification_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_circuit_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_circuit_public_prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"verification_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreVirtualCircuitPublicPrefixes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitPublicPrefixesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVirtualCircuitPublicPrefixesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListVirtualCircuitPublicPrefixesResponse
}

func (s *CoreVirtualCircuitPublicPrefixesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVirtualCircuitPublicPrefixesDataSourceCrud) Get() error {
	request := oci_core.ListVirtualCircuitPublicPrefixesRequest{}

	if verificationState, ok := s.D.GetOkExists("verification_state"); ok {
		request.VerificationState = oci_core.VirtualCircuitPublicPrefixVerificationStateEnum(verificationState.(string))
	}

	if virtualCircuitId, ok := s.D.GetOkExists("virtual_circuit_id"); ok {
		tmp := virtualCircuitId.(string)
		request.VirtualCircuitId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListVirtualCircuitPublicPrefixes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVirtualCircuitPublicPrefixesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVirtualCircuitPublicPrefixesDataSource-", CoreVirtualCircuitPublicPrefixesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		virtualCircuitPublicPrefix := map[string]interface{}{}

		if r.CidrBlock != nil {
			virtualCircuitPublicPrefix["cidr_block"] = *r.CidrBlock
		}

		virtualCircuitPublicPrefix["verification_state"] = string(r.VerificationState)

		resources = append(resources, virtualCircuitPublicPrefix)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVirtualCircuitPublicPrefixesDataSource().Schema["virtual_circuit_public_prefixes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_circuit_public_prefixes", resources); err != nil {
		return err
	}

	return nil
}
