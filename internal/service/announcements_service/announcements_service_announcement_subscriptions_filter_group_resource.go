// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"
)

func AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAnnouncementsServiceAnnouncementSubscriptionsFilterGroup,
		Read:     readAnnouncementsServiceAnnouncementSubscriptionsFilterGroup,
		Update:   updateAnnouncementsServiceAnnouncementSubscriptionsFilterGroup,
		Delete:   deleteAnnouncementsServiceAnnouncementSubscriptionsFilterGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"announcement_subscription_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filters": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createAnnouncementsServiceAnnouncementSubscriptionsFilterGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.CreateResource(d, sync)
}

func readAnnouncementsServiceAnnouncementSubscriptionsFilterGroup(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateAnnouncementsServiceAnnouncementSubscriptionsFilterGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAnnouncementsServiceAnnouncementSubscriptionsFilterGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnnouncementSubscriptionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_announcements_service.AnnouncementSubscriptionClient
	Res                    *oci_announcements_service.FilterGroup
	DisableNotFoundRetries bool
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud) ID() string {
	return s.D.Get("announcement_subscription_id").(string)
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud) Create() error {
	request := oci_announcements_service.CreateFilterGroupRequest{}

	if announcementSubscriptionId, ok := s.D.GetOkExists("announcement_subscription_id"); ok {
		tmp := announcementSubscriptionId.(string)
		request.AnnouncementSubscriptionId = &tmp
	}

	if filters, ok := s.D.GetOkExists("filters"); ok {
		interfaces := filters.([]interface{})
		tmp := make([]oci_announcements_service.Filter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "filters", stateDataIndex)
			converted, err := s.mapToFilter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("filters") {
			request.Filters = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	response, err := s.Client.CreateFilterGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FilterGroup
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud) Update() error {
	request := oci_announcements_service.UpdateFilterGroupRequest{}

	if announcementSubscriptionId, ok := s.D.GetOkExists("announcement_subscription_id"); ok {
		tmp := announcementSubscriptionId.(string)
		request.AnnouncementSubscriptionId = &tmp
	}

	if filterGroupName, ok := s.D.GetOkExists("name"); ok {
		tmp := filterGroupName.(string)
		request.FilterGroupName = &tmp
	}

	if filters, ok := s.D.GetOkExists("filters"); ok {
		interfaces := filters.([]interface{})
		tmp := make([]oci_announcements_service.Filter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "filters", stateDataIndex)
			converted, err := s.mapToFilter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("filters") {
			request.Filters = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	response, err := s.Client.UpdateFilterGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FilterGroup
	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud) Delete() error {
	request := oci_announcements_service.DeleteFilterGroupRequest{}

	if announcementSubscriptionId, ok := s.D.GetOkExists("announcement_subscription_id"); ok {
		tmp := announcementSubscriptionId.(string)
		request.AnnouncementSubscriptionId = &tmp
	}

	if filterGroupName, ok := s.D.GetOkExists("name"); ok {
		tmp := filterGroupName.(string)
		request.FilterGroupName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "announcements_service")

	_, err := s.Client.DeleteFilterGroup(context.Background(), request)
	return err
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud) SetData() error {
	filters := []interface{}{}
	for _, item := range s.Res.Filters {
		filters = append(filters, FilterToMap(item))
	}
	s.D.Set("filters", filters)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return nil
}

func (s *AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResourceCrud) mapToFilter(fieldKeyFormat string) (oci_announcements_service.Filter, error) {
	result := oci_announcements_service.Filter{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_announcements_service.FilterTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func FilterToMap(obj oci_announcements_service.Filter) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
