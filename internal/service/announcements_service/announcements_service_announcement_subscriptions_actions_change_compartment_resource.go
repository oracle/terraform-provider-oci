// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"
)

func AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartment,
		Read:     readAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartment,
		Delete:   deleteAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartment,
		Schema: map[string]*schema.Schema{
			// Required
			"announcement_subscription_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartment(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.CreateResource(d, sync)
}

func readAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartment(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartment(d *schema.ResourceData, m interface{}) error {
	return nil
}

type AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_announcements_service.AnnouncementSubscriptionClient
	Res                    *oci_announcements_service.ChangeAnnouncementSubscriptionCompartmentResponse
	DisableNotFoundRetries bool
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceCrud) ID() string {
	return s.D.Get("announcement_subscription_id").(string)
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceCrud) Create() error {
	request := oci_announcements_service.ChangeAnnouncementSubscriptionCompartmentRequest{}

	if announcementSubscriptionId, ok := s.D.GetOkExists("announcement_subscription_id"); ok {
		tmp := announcementSubscriptionId.(string)
		request.AnnouncementSubscriptionId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	response, err := s.Client.ChangeAnnouncementSubscriptionCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResourceCrud) SetData() error {
	return nil
}
