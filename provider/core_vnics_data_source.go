// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VnicsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVnics,
		Schema: map[string]*schema.Schema{
			//@CODEGEN 1/2018: Generated code would have added a filter here. We remove it since this is a singular
			// data source.
			"vnic_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// @CODEGEN 01/2018: Code gen incorrectly assumes that all datasources support List operations and
			// will encapsulate the following fields in its own schema under a TypeList property.
			//
			// In the case of this data source, only Get operation is supported, so we only have one result and promote
			// all the properties to the top-level. This also avoids a breaking change. The SetData function also
			// differs from generated code in this respect.
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"hostname_label": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_primary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// @CODEGEN 1/2018: private_ip => private_ip_address
			"private_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// @CODEGEN 1/2018: public_ip => public_ip_address
			"public_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skip_source_dest_check": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readVnics(d *schema.ResourceData, m interface{}) error {
	sync := &VnicsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type VnicsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVnicResponse
}

func (s *VnicsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VnicsDataSourceCrud) Get() error {
	request := oci_core.GetVnicRequest{}

	if vnicId, ok := s.D.GetOkExists("vnic_id"); ok {
		tmp := vnicId.(string)
		request.VnicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetVnic(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VnicsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	// @CODEGEN 1/2018: In most generated data sources, the ID is set to the current time stamp.
	// In the case of this datasource, the existing provider sets it to the resource ID.
	// This happens because it only supports a Get operation that returns 1 item.
	// Let's keep this as is to avoid potential breaking changes.
	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	if s.Res.MacAddress != nil {
		s.D.Set("mac_address", *s.Res.MacAddress)
	}

	if s.Res.PrivateIp != nil {
		s.D.Set("private_ip_address", *s.Res.PrivateIp)
	}

	if s.Res.PublicIp != nil {
		s.D.Set("public_ip_address", *s.Res.PublicIp)
	}

	if s.Res.SkipSourceDestCheck != nil {
		s.D.Set("skip_source_dest_check", *s.Res.SkipSourceDestCheck)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return
}
