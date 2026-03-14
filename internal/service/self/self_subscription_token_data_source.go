// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package self

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_self "github.com/oracle/oci-go-sdk/v65/self"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SelfSubscriptionTokenDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularSelfSubscriptionTokenWithContext,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularSelfSubscriptionTokenWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &SelfSubscriptionTokenDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SelfSubscriptionClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type SelfSubscriptionTokenDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_self.SubscriptionClient
	Res    *oci_self.GetSubscriptionTokenResponse
}

func (s *SelfSubscriptionTokenDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SelfSubscriptionTokenDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_self.GetSubscriptionTokenRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "self")

	response, err := s.Client.GetSubscriptionToken(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SelfSubscriptionTokenDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SelfSubscriptionTokenDataSource-", SelfSubscriptionTokenDataSource(), s.D))

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.Token != nil {
		s.D.Set("token", *s.Res.Token)
	}

	return nil
}
