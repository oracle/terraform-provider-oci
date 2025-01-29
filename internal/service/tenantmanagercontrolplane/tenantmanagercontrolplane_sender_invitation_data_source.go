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

func TenantmanagercontrolplaneSenderInvitationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneSenderInvitation,
		Schema: map[string]*schema.Schema{
			"sender_invitation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
			"recipient_email_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recipient_invitation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recipient_tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subjects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func readSingularTenantmanagercontrolplaneSenderInvitation(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSenderInvitationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SenderInvitationClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneSenderInvitationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SenderInvitationClient
	Res    *oci_tenantmanagercontrolplane.GetSenderInvitationResponse
}

func (s *TenantmanagercontrolplaneSenderInvitationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneSenderInvitationDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetSenderInvitationRequest{}

	if senderInvitationId, ok := s.D.GetOkExists("sender_invitation_id"); ok {
		tmp := senderInvitationId.(string)
		request.SenderInvitationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetSenderInvitation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneSenderInvitationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RecipientEmailAddress != nil {
		s.D.Set("recipient_email_address", *s.Res.RecipientEmailAddress)
	}

	if s.Res.RecipientInvitationId != nil {
		s.D.Set("recipient_invitation_id", *s.Res.RecipientInvitationId)
	}

	if s.Res.RecipientTenancyId != nil {
		s.D.Set("recipient_tenancy_id", *s.Res.RecipientTenancyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	s.D.Set("subjects", s.Res.Subjects)

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
