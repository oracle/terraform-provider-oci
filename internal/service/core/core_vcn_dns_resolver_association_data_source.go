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

func CoreVcnDnsResolverAssociationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreVcnDnsResolverAssociation,
		Schema: map[string]*schema.Schema{
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wait_for_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"dns_resolver_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreVcnDnsResolverAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVcnDnsResolverAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVcnDnsResolverAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetVcnDnsResolverAssociationResponse
}

func (s *CoreVcnDnsResolverAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVcnDnsResolverAssociationDataSourceCrud) Get() error {
	request := oci_core.GetVcnDnsResolverAssociationRequest{}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVcnDnsResolverAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVcnDnsResolverAssociationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	waitForState := ""
	if v, ok := s.D.GetOkExists("wait_for_state"); ok {
		waitForState = v.(string)
	}

	if waitForState != "" {
		state := string(s.Res.LifecycleState)
		if state != waitForState {
			if err := tfresource.WaitForResourceCondition(
				s,
				func() bool {
					return s.Res != nil &&
						s.Res.LifecycleState == oci_core.VcnDnsResolverAssociationLifecycleStateEnum(waitForState)
				},
				s.D.Timeout(schema.TimeoutRead),
			); err != nil {
				return err
			}
		}
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVcnDnsResolverAssociationDataSource-", CoreVcnDnsResolverAssociationDataSource(), s.D))

	if s.Res.DnsResolverId != nil {
		s.D.Set("dns_resolver_id", *s.Res.DnsResolverId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
