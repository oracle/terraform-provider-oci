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

func CoreIpInventorySubnetDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreIpInventorySubnet,
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
			"ip_inventory_subnet_resource_summary": {
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
						"assigned_resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"assigned_resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"assigned_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"associated_public_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"associated_public_ip_pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address_lifetime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip_lifetime": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ip_inventory_subnet_count": {
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

func readSingularCoreIpInventorySubnet(d *schema.ResourceData, m interface{}) error {
	sync := &CoreIpInventorySubnetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreIpInventorySubnetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetSubnetIpInventoryResponse
}

func (s *CoreIpInventorySubnetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreIpInventorySubnetDataSourceCrud) Get() error {
	request := oci_core.GetSubnetIpInventoryRequest{}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetSubnetIpInventory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreIpInventorySubnetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreIpInventorySubnetDataSource-", CoreIpInventorySubnetDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	ipInventorySubnetResourceSummary := []interface{}{}
	for _, item := range s.Res.IpInventorySubnetResourceSummary {
		ipInventorySubnetResourceSummary = append(ipInventorySubnetResourceSummary, IpInventorySubnetResourceSummaryToMap(item))
	}
	s.D.Set("ip_inventory_subnet_resource_summary", ipInventorySubnetResourceSummary)

	if s.Res.Count != nil {
		s.D.Set("ip_inventory_subnet_count", *s.Res.Count)
	}

	if s.Res.LastUpdatedTimestamp != nil {
		s.D.Set("last_updated_timestamp", s.Res.LastUpdatedTimestamp.String())
	}

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	return nil
}

func IpInventorySubnetResourceSummaryToMap(obj oci_core.IpInventorySubnetResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddressType != nil {
		result["address_type"] = string(*obj.AddressType)
	}

	if obj.AssignedResourceName != nil {
		result["assigned_resource_name"] = string(*obj.AssignedResourceName)
	}

	result["assigned_resource_type"] = string(obj.AssignedResourceType)

	if obj.AssignedTime != nil {
		result["assigned_time"] = obj.AssignedTime.String()
	}

	if obj.AssociatedPublicIp != nil {
		result["associated_public_ip"] = string(*obj.AssociatedPublicIp)
	}

	result["associated_public_ip_pool"] = string(obj.AssociatedPublicIpPool)

	if obj.DnsHostName != nil {
		result["dns_host_name"] = string(*obj.DnsHostName)
	}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	result["ip_address_lifetime"] = string(obj.IpAddressLifetime)

	if obj.IpId != nil {
		result["ip_id"] = string(*obj.IpId)
	}

	if obj.ParentCidr != nil {
		result["parent_cidr"] = string(*obj.ParentCidr)
	}

	result["public_ip_lifetime"] = string(obj.PublicIpLifetime)

	return result
}
