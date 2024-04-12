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

func CoreIpInventorySubnetCidrDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreIpInventorySubnetCidr,
		Schema: map[string]*schema.Schema{
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_inventory_cidr_utilization_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"address_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"utilization": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"ip_inventory_subnet_cidr_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreIpInventorySubnetCidr(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpInventorySubnetCidrDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpInventorySubnetCidrDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetSubnetCidrUtilizationResponse
}

func (s *CoreIpInventorySubnetCidrDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpInventorySubnetCidrDataSourceCrud) Get() error {
	request := oci_core.GetSubnetCidrUtilizationRequest{}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetSubnetCidrUtilization(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreIpInventorySubnetCidrDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpInventorySubnetCidrDataSource-", CoreIpInventorySubnetCidrDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	ipInventoryCidrUtilizationSummary := []interface{}{}
	for _, item := range s.Res.IpInventoryCidrUtilizationSummary {
		ipInventoryCidrUtilizationSummary = append(ipInventoryCidrUtilizationSummary, IpInventoryCidrUtilizationSummaryToMap(item))
	}
	s.D.Set("ip_inventory_cidr_utilization_summary", ipInventoryCidrUtilizationSummary)

	if s.Res.Count != nil {
		s.D.Set("ip_inventory_subnet_cidr_count", *s.Res.Count)
	}

	if s.Res.LastUpdatedTimestamp != nil {
		s.D.Set("last_updated_timestamp", s.Res.LastUpdatedTimestamp.String())
	}

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	return nil
}

func IpInventoryCidrUtilizationSummaryToMap(obj oci_core.IpInventoryCidrUtilizationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddressType != nil {
		result["address_type"] = string(*obj.AddressType)
	}

	if obj.Cidr != nil {
		result["cidr"] = string(*obj.Cidr)
	}

	if obj.Utilization != nil {
		result["utilization"] = float32(*obj.Utilization)
	}

	return result
}
