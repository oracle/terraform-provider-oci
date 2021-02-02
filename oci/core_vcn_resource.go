// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v35/core"
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

	if _, ok := s.D.GetOkExists("cidr_blocks"); ok && s.D.HasChange("cidr_blocks") {
		oldRaw, newRaw := s.D.GetChange("cidr_blocks")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCidrBlocks(oldRaw, newRaw)
			if err != nil {
				return err
			}
		}
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

	s.D.Set("cidr_blocks", s.Res.CidrBlocks)

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

func (s *CoreVcnResourceCrud) updateCidrBlocks(oldRaw interface{}, newRaw interface{}) error {
	interfaces := oldRaw.([]interface{})
	oldBlocks := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			oldBlocks[i] = interfaces[i].(string)
		}
	}
	interfaces = newRaw.([]interface{})
	newBlocks := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			newBlocks[i] = interfaces[i].(string)
		}
	}
	canEdit, operation, oldCidr, newCidr := oneEditAway(oldBlocks, newBlocks)
	if !canEdit {
		return fmt.Errorf("only one add/remove or modification is allowed at once, new cidr_block must be added at the end of list")
	}
	if operation == "add" {
		addVcnCidrRequest := oci_core.AddVcnCidrRequest{}
		idTmp := s.D.Id()
		addVcnCidrRequest.VcnId = &idTmp
		addVcnCidrRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")
		addVcnCidrRequest.CidrBlock = &newCidr
		_, err := s.Client.AddVcnCidr(context.Background(), addVcnCidrRequest)
		if err != nil {
			return err
		}
	}
	if operation == "remove" {
		removeVcnCidrRequest := oci_core.RemoveVcnCidrRequest{}
		idTmp := s.D.Id()
		removeVcnCidrRequest.VcnId = &idTmp
		removeVcnCidrRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")
		removeVcnCidrRequest.CidrBlock = &oldCidr
		_, err := s.Client.RemoveVcnCidr(context.Background(), removeVcnCidrRequest)
		if err != nil {
			return err
		}
	}
	if operation == "modify" {
		modifyVcnCidrRequest := oci_core.ModifyVcnCidrRequest{}
		idTmp := s.D.Id()
		modifyVcnCidrRequest.VcnId = &idTmp
		modifyVcnCidrRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")
		modifyVcnCidrRequest.OriginalCidrBlock = &oldCidr
		modifyVcnCidrRequest.NewCidrBlock = &newCidr
		_, err := s.Client.ModifyVcnCidr(context.Background(), modifyVcnCidrRequest)
		if err != nil {
			return err
		}
	}
	return nil
}

func oneEditAway(oldBlocks []string, newBlocks []string) (bool, string, string, string) {
	if Abs(len(newBlocks)-len(oldBlocks)) > 1 {
		return false, "", "", ""
	}
	if len(newBlocks) == len(oldBlocks) {
		for i := 0; i < len(oldBlocks); i++ {
			if oldBlocks[i] != newBlocks[i] {
				for j := i + 1; j < len(oldBlocks); j++ {
					if oldBlocks[j] != newBlocks[j] {
						return false, "", "", ""
					}
				}
				return true, "modify", oldBlocks[i], newBlocks[i]
			}
		}
	}
	if len(newBlocks) > len(oldBlocks) {
		for i := 0; i < len(oldBlocks); i++ {
			if oldBlocks[i] != newBlocks[i] {
				return false, "", "", ""
			}
		}
		return true, "add", "", newBlocks[len(newBlocks)-1]
	}
	for i := 0; i < len(newBlocks); i++ {
		if oldBlocks[i] != newBlocks[i] {
			for j := i + 1; j < len(newBlocks); j++ {
				if oldBlocks[j] != newBlocks[j-1] {
					return false, "", "", ""
				}
			}
			return true, "remove", oldBlocks[i], ""
		}
	}
	return true, "remove", oldBlocks[len(oldBlocks)-1], ""
}
