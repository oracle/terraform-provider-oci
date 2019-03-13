// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityTagDefaultDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityTagDefault,
		Schema: map[string]*schema.Schema{
			"tag_default_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_definition_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_definition_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_namespace_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularIdentityTagDefault(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityTagDefaultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetTagDefaultResponse
}

func (s *IdentityTagDefaultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagDefaultDataSourceCrud) Get() error {
	request := oci_identity.GetTagDefaultRequest{}

	if tagDefaultId, ok := s.D.GetOkExists("tag_default_id"); ok {
		tmp := tagDefaultId.(string)
		request.TagDefaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.GetTagDefault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityTagDefaultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TagDefinitionId != nil {
		s.D.Set("tag_definition_id", *s.Res.TagDefinitionId)
	}

	if s.Res.TagDefinitionName != nil {
		s.D.Set("tag_definition_name", *s.Res.TagDefinitionName)
	}

	if s.Res.TagNamespaceId != nil {
		s.D.Set("tag_namespace_id", *s.Res.TagNamespaceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	return nil
}
