// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v33/core"
)

func init() {
	RegisterDatasource("oci_core_vcn_dns_resolver_association", CoreVcnDnsResolverAssociationDataSource())
}

func CoreVcnDnsResolverAssociationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreVcnDnsResolverAssociation,
		Schema: map[string]*schema.Schema{
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
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
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

	s.D.SetId(GenerateDataSourceHashID("CoreVcnDnsResolverAssociationDataSource-", CoreVcnDnsResolverAssociationDataSource(), s.D))

	if s.Res.DnsResolverId != nil {
		s.D.Set("dns_resolver_id", *s.Res.DnsResolverId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
