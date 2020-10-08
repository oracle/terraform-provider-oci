// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v27/core"
)

func init() {
	RegisterResource("oci_core_vcn", CoreVcnResource())
}

func CoreVcnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreVcn,
		Read:     readCoreVcn,
		Update:   updateCoreVcn,
		Delete:   deleteCoreVcn,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"cidr_block": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cidr_blocks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ipv6cidr_block": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: ipv6CompressionDiffSuppressFunction,
			},
			"is_ipv6enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"default_dhcp_options_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_route_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_security_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6public_cidr_block": {
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
			"vcn_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return CreateResource(d, sync)
}

func readCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

func updateCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return UpdateResource(d, sync)
}

func deleteCoreVcn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreVcnResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Vcn
	DisableNotFoundRetries bool
}

func (s *CoreVcnResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVcnResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VcnLifecycleStateProvisioning),
	}
}

func (s *CoreVcnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VcnLifecycleStateAvailable),
	}
}

func (s *CoreVcnResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VcnLifecycleStateTerminating),
	}
}

func (s *CoreVcnResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VcnLifecycleStateTerminated),
	}
}

func (s *CoreVcnResourceCrud) Create() error {
	request := oci_core.CreateVcnRequest{}

	if cidrBlock, ok := s.D.GetOkExists("cidr_block"); ok {
		tmp := cidrBlock.(string)
		request.CidrBlock = &tmp
	}

	if cidrBlocks, ok := s.D.GetOkExists("cidr_blocks"); ok {
		interfaces := cidrBlocks.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("cidr_blocks") {
			request.CidrBlocks = tmp
		}
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

	if ipv6CidrBlock, ok := s.D.GetOkExists("ipv6cidr_block"); ok {
		tmp := ipv6CidrBlock.(string)
		request.Ipv6CidrBlock = &tmp
	}

	if isIpv6Enabled, ok := s.D.GetOkExists("is_ipv6enabled"); ok {
		tmp := isIpv6Enabled.(bool)
		request.IsIpv6Enabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *CoreVcnResourceCrud) Get() error {
	request := oci_core.GetVcnRequest{}

	tmp := s.D.Id()
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *CoreVcnResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateVcnRequest{}

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
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *CoreVcnResourceCrud) Delete() error {
	request := oci_core.DeleteVcnRequest{}

	tmp := s.D.Id()
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVcn(context.Background(), request)
	return err
}

func (s *CoreVcnResourceCrud) SetData() error {
	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
	}

	if s.Res.CidrBlocks != nil {
		s.D.Set("cidr_blocks", s.Res.CidrBlocks)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultDhcpOptionsId != nil {
		s.D.Set("default_dhcp_options_id", *s.Res.DefaultDhcpOptionsId)
	}

	if s.Res.DefaultRouteTableId != nil {
		s.D.Set("default_route_table_id", *s.Res.DefaultRouteTableId)
	}

	if s.Res.DefaultSecurityListId != nil {
		s.D.Set("default_security_list_id", *s.Res.DefaultSecurityListId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsLabel != nil {
		s.D.Set("dns_label", *s.Res.DnsLabel)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Ipv6CidrBlock != nil {
		s.D.Set("ipv6cidr_block", *s.Res.Ipv6CidrBlock)
		s.D.Set("is_ipv6enabled", true)
	}

	if s.Res.Ipv6PublicCidrBlock != nil {
		s.D.Set("ipv6public_cidr_block", *s.Res.Ipv6PublicCidrBlock)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnDomainName != nil {
		s.D.Set("vcn_domain_name", *s.Res.VcnDomainName)
	}

	return nil
}

func (s *CoreVcnResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVcnCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VcnId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVcnCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
