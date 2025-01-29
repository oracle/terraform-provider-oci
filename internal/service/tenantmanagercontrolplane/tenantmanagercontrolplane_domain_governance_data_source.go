// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneDomainGovernanceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneDomainGovernance,
		Schema: map[string]*schema.Schema{
			"domain_governance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_governance_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ons_subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ons_topic_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularTenantmanagercontrolplaneDomainGovernance(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneDomainGovernanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DomainGovernanceClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneDomainGovernanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.DomainGovernanceClient
	Res    *oci_tenantmanagercontrolplane.GetDomainGovernanceResponse
}

func (s *TenantmanagercontrolplaneDomainGovernanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneDomainGovernanceDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetDomainGovernanceRequest{}

	if domainGovernanceId, ok := s.D.GetOkExists("domain_governance_id"); ok {
		tmp := domainGovernanceId.(string)
		request.DomainGovernanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetDomainGovernance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneDomainGovernanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DomainId != nil {
		s.D.Set("domain_id", *s.Res.DomainId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsGovernanceEnabled != nil {
		s.D.Set("is_governance_enabled", *s.Res.IsGovernanceEnabled)
	}

	if s.Res.OnsSubscriptionId != nil {
		s.D.Set("ons_subscription_id", *s.Res.OnsSubscriptionId)
	}

	if s.Res.OnsTopicId != nil {
		s.D.Set("ons_topic_id", *s.Res.OnsTopicId)
	}

	if s.Res.OwnerId != nil {
		s.D.Set("owner_id", *s.Res.OwnerId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubscriptionEmail != nil {
		s.D.Set("subscription_email", *s.Res.SubscriptionEmail)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
