// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VcnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVcn,
		Read:     readVcn,
		Update:   updateVcn,
		Delete:   deleteVcn,
		Schema: map[string]*schema.Schema{
			// Required
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

			// Optional
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
				DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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
			"id": {
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

func createVcn(d *schema.ResourceData, m interface{}) error {
	sync := &VcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readVcn(d *schema.ResourceData, m interface{}) error {
	sync := &VcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateVcn(d *schema.ResourceData, m interface{}) error {
	sync := &VcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteVcn(d *schema.ResourceData, m interface{}) error {
	sync := &VcnResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VcnResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Vcn
	DisableNotFoundRetries bool
}

func (s *VcnResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VcnResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VcnLifecycleStateProvisioning),
	}
}

func (s *VcnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VcnLifecycleStateAvailable),
	}
}

func (s *VcnResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VcnLifecycleStateTerminating),
	}
}

func (s *VcnResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VcnLifecycleStateTerminated),
	}
}

func (s *VcnResourceCrud) Create() error {
	request := oci_core.CreateVcnRequest{}

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVcn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Vcn
	return nil
}

func (s *VcnResourceCrud) Get() error {
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

func (s *VcnResourceCrud) Update() error {
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

func (s *VcnResourceCrud) Delete() error {
	request := oci_core.DeleteVcnRequest{}

	tmp := s.D.Id()
	request.VcnId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVcn(context.Background(), request)
	return err
}

func (s *VcnResourceCrud) SetData() {
	if s.Res.CidrBlock != nil {
		s.D.Set("cidr_block", *s.Res.CidrBlock)
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

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnDomainName != nil {
		s.D.Set("vcn_domain_name", *s.Res.VcnDomainName)
	}

}
