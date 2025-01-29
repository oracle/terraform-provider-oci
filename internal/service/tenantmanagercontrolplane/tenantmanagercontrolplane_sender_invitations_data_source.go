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

func TenantmanagercontrolplaneSenderInvitationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneSenderInvitations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recipient_tenancy_id": {
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
			"sender_invitation_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"recipient_tenancy_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"recipient_email_address": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"subjects": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recipient_invitation_id": {
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

func readTenantmanagercontrolplaneSenderInvitations(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSenderInvitationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SenderInvitationClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneSenderInvitationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SenderInvitationClient
	Res    *oci_tenantmanagercontrolplane.ListSenderInvitationsResponse
}

func (s *TenantmanagercontrolplaneSenderInvitationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneSenderInvitationsDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListSenderInvitationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if recipientTenancyId, ok := s.D.GetOkExists("recipient_tenancy_id"); ok {
		tmp := recipientTenancyId.(string)
		request.RecipientTenancyId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_tenantmanagercontrolplane.ListSenderInvitationsLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_tenantmanagercontrolplane.ListSenderInvitationsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListSenderInvitations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSenderInvitations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneSenderInvitationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneSenderInvitationsDataSource-", TenantmanagercontrolplaneSenderInvitationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	senderInvitation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SenderInvitationSummaryToMap(item))
	}
	senderInvitation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneSenderInvitationsDataSource().Schema["sender_invitation_collection"].Elem.(*schema.Resource).Schema)
		senderInvitation["items"] = items
	}

	resources = append(resources, senderInvitation)
	if err := s.D.Set("sender_invitation_collection", resources); err != nil {
		return err
	}

	return nil
}

func SenderInvitationSummaryToMap(obj oci_tenantmanagercontrolplane.SenderInvitationSummary) map[string]interface{} {
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

	if obj.RecipientInvitationId != nil {
		result["recipient_invitation_id"] = string(*obj.RecipientInvitationId)
	}

	if obj.RecipientTenancyId != nil {
		result["recipient_tenancy_id"] = string(*obj.RecipientTenancyId)
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
