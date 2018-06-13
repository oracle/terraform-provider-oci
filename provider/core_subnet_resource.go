// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func SubnetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSubnet,
		Read:     readSubnet,
		Update:   updateSubnet,
		Delete:   deleteSubnet,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dhcp_options_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_label": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"prohibit_public_ip_on_vnic": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"security_list_ids": {
				// @CODEGEN: The ordering of security_list_ids may change, but shouldn't result in a diff.
				// Change it to a TypeSet instead of TypeList (as generated).
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 5,
				MinItems: 0,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteSubnet(d *schema.ResourceData, m interface{}) error {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type SubnetResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Subnet
	DisableNotFoundRetries bool
}

func (s *SubnetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SubnetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.SubnetLifecycleStateProvisioning),
	}
}

func (s *SubnetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.SubnetLifecycleStateAvailable),
	}
}

func (s *SubnetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.SubnetLifecycleStateTerminating),
	}
}

func (s *SubnetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.SubnetLifecycleStateTerminated),
	}
}

func (s *SubnetResourceCrud) Create() error {
	request := oci_core.CreateSubnetRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if cidrBlock, ok := s.D.GetOkExists("cidr_block"); ok {
		tmp := cidrBlock.(string)
		request.CidrBlock = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if dhcpOptionsId, ok := s.D.GetOkExists("dhcp_options_id"); ok {
		tmp := dhcpOptionsId.(string)
		request.DhcpOptionsId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsLabel, ok := s.D.GetOkExists("dns_label"); ok {
		tmp := dnsLabel.(string)
		request.DnsLabel = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	// TODO: GetOk malfunction with this bool: 'ok' is always the value of the bool
	// newer versions of terraform support GetOkExists which should resolve this problem
	prohibitPublicIpOnVnic := s.D.Get("prohibit_public_ip_on_vnic").(bool)
	request.ProhibitPublicIpOnVnic = &prohibitPublicIpOnVnic

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	request.SecurityListIds = []string{}
	if securityListIds, ok := s.D.GetOkExists("security_list_ids"); ok {
		secListIdSet, assertOk := securityListIds.(*schema.Set)
		if !assertOk {
			return fmt.Errorf("Could not assert security_list_ids as type schema.Set")
		}

		interfaces := secListIdSet.List()
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.SecurityListIds = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subnet
	return nil
}

func (s *SubnetResourceCrud) Get() error {
	request := oci_core.GetSubnetRequest{}

	tmp := s.D.Id()
	request.SubnetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subnet
	return nil
}

func (s *SubnetResourceCrud) Update() error {
	request := oci_core.UpdateSubnetRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SubnetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateSubnet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subnet
	return nil
}

func (s *SubnetResourceCrud) Delete() error {
	request := oci_core.DeleteSubnetRequest{}

	tmp := s.D.Id()
	request.SubnetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteSubnet(context.Background(), request)
	return err
}

func (s *SubnetResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DhcpOptionsId != nil {
		s.D.Set("dhcp_options_id", *s.Res.DhcpOptionsId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsLabel != nil {
		s.D.Set("dns_label", *s.Res.DnsLabel)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.ProhibitPublicIpOnVnic != nil {
		s.D.Set("prohibit_public_ip_on_vnic", *s.Res.ProhibitPublicIpOnVnic)
	}

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	if err := s.D.Set("security_list_ids", crud.StringsToSet(s.Res.SecurityListIds)); err != nil {
		log.Printf("Unable to set security_list_ids. Error: %q", err)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetDomainName != nil {
		s.D.Set("subnet_domain_name", *s.Res.SubnetDomainName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	if s.Res.VirtualRouterIp != nil {
		s.D.Set("virtual_router_ip", *s.Res.VirtualRouterIp)
	}

	if s.Res.VirtualRouterMac != nil {
		s.D.Set("virtual_router_mac", *s.Res.VirtualRouterMac)
	}

}
