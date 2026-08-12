// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

// CoreDrgDefaultRouteTableManagementResource manages a DRG's per-attachment-type default DRG
// route table assignment (UpdateDrgDetails.DefaultDrgRouteTables) as its own resource, distinct
// from oci_core_drg itself. This lets a config point one of these four fields at a DRG route
// table created in the same apply without oci_core_drg and oci_core_drg_route_table forming a
// module/resource reference cycle -- same shape of problem oci_core_drg_attachment_management
// solves for drg_route_table_id/export_drg_route_distribution_id on attachments OCI auto-creates.
func CoreDrgDefaultRouteTableManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgDefaultRouteTableManagement,
		Read:     readCoreDrgDefaultRouteTableManagement,
		Update:   updateCoreDrgDefaultRouteTableManagement,
		Delete:   deleteCoreDrgDefaultRouteTableManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ipsec_tunnel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_peering_connection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_circuit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCoreDrgDefaultRouteTableManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgDefaultRouteTableManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreDrgDefaultRouteTableManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgDefaultRouteTableManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDrgDefaultRouteTableManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgDefaultRouteTableManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDrgDefaultRouteTableManagement(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgDefaultRouteTableManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreDrgDefaultRouteTableManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Drg
	DisableNotFoundRetries bool
}

func (s *CoreDrgDefaultRouteTableManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgDefaultRouteTableManagementResourceCrud) Get() error {
	request := oci_core.GetDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgDefaultRouteTableManagementResourceCrud) Create() error {
	return s.Update()
}

// A DRG's default route table assignment can't actually be deleted -- every attachment type
// always has SOME default route table (see oci_core_default_drg_route_table's LIFECYCLE.md).
// Same "management, not ownership" posture as oci_core_drg_attachment_management's Delete:
// Terraform just stops tracking the assignment, OCI keeps whatever it was last set to.
func (s *CoreDrgDefaultRouteTableManagementResourceCrud) Delete() error {
	return nil
}

func (s *CoreDrgDefaultRouteTableManagementResourceCrud) Update() error {
	request := oci_core.UpdateDrgRequest{}

	drgId := s.D.Get("drg_id").(string)
	request.DrgId = &drgId

	defaultDrgRouteTables := oci_core.DefaultDrgRouteTables{}

	if ipsecTunnel, ok := s.D.GetOkExists("ipsec_tunnel"); ok {
		tmp := ipsecTunnel.(string)
		defaultDrgRouteTables.IpsecTunnel = &tmp
	}

	if remotePeeringConnection, ok := s.D.GetOkExists("remote_peering_connection"); ok {
		tmp := remotePeeringConnection.(string)
		defaultDrgRouteTables.RemotePeeringConnection = &tmp
	}

	if vcn, ok := s.D.GetOkExists("vcn"); ok {
		tmp := vcn.(string)
		defaultDrgRouteTables.Vcn = &tmp
	}

	if virtualCircuit, ok := s.D.GetOkExists("virtual_circuit"); ok {
		tmp := virtualCircuit.(string)
		defaultDrgRouteTables.VirtualCircuit = &tmp
	}

	request.DefaultDrgRouteTables = &defaultDrgRouteTables

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *CoreDrgDefaultRouteTableManagementResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.Set("drg_id", *s.Res.Id)

	if s.Res.DefaultDrgRouteTables != nil {
		if s.Res.DefaultDrgRouteTables.IpsecTunnel != nil {
			s.D.Set("ipsec_tunnel", *s.Res.DefaultDrgRouteTables.IpsecTunnel)
		}
		if s.Res.DefaultDrgRouteTables.RemotePeeringConnection != nil {
			s.D.Set("remote_peering_connection", *s.Res.DefaultDrgRouteTables.RemotePeeringConnection)
		}
		if s.Res.DefaultDrgRouteTables.Vcn != nil {
			s.D.Set("vcn", *s.Res.DefaultDrgRouteTables.Vcn)
		}
		if s.Res.DefaultDrgRouteTables.VirtualCircuit != nil {
			s.D.Set("virtual_circuit", *s.Res.DefaultDrgRouteTables.VirtualCircuit)
		}
	}

	return nil
}
