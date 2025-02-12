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

func TenantmanagercontrolplaneRecipientInvitationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneRecipientInvitations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sender_tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recipient_invitation_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recipient_email_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sender_invitation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sender_tenancy_id": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readTenantmanagercontrolplaneRecipientInvitations(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneRecipientInvitationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RecipientInvitationClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneRecipientInvitationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.RecipientInvitationClient
	Res    *oci_tenantmanagercontrolplane.ListRecipientInvitationsResponse
}

func (s *TenantmanagercontrolplaneRecipientInvitationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneRecipientInvitationsDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListRecipientInvitationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if senderTenancyId, ok := s.D.GetOkExists("sender_tenancy_id"); ok {
		tmp := senderTenancyId.(string)
		request.SenderTenancyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_tenantmanagercontrolplane.ListRecipientInvitationsLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_tenantmanagercontrolplane.ListRecipientInvitationsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListRecipientInvitations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRecipientInvitations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneRecipientInvitationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneRecipientInvitationsDataSource-", TenantmanagercontrolplaneRecipientInvitationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	recipientInvitation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecipientInvitationSummaryToMap(item))
	}
	recipientInvitation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneRecipientInvitationsDataSource().Schema["recipient_invitation_collection"].Elem.(*schema.Resource).Schema)
		recipientInvitation["items"] = items
	}

	resources = append(resources, recipientInvitation)
	if err := s.D.Set("recipient_invitation_collection", resources); err != nil {
		return err
	}

	return nil
}

func RecipientInvitationSummaryToMap(obj oci_tenantmanagercontrolplane.RecipientInvitationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.RecipientEmailAddress != nil {
		result["recipient_email_address"] = string(*obj.RecipientEmailAddress)
	}

	if obj.SenderInvitationId != nil {
		result["sender_invitation_id"] = string(*obj.SenderInvitationId)
	}

	if obj.SenderTenancyId != nil {
		result["sender_tenancy_id"] = string(*obj.SenderTenancyId)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	result["subjects"] = obj.Subjects

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
