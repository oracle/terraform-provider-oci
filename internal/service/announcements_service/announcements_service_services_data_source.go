// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AnnouncementsServiceServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAnnouncementsServiceServices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"comms_manager_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"platform_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"services_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"comms_manager_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"excluded_realms": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"platform_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"previous_service_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"short_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"team_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
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

func readAnnouncementsServiceServices(d *schema.ResourceData, m interface{}) error {
	sync := &AnnouncementsServiceServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceClient()

	return tfresource.ReadResource(sync)
}

type AnnouncementsServiceServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_announcements_service.ServiceClient
	Res    *oci_announcements_service.ListServicesResponse
}

func (s *AnnouncementsServiceServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnnouncementsServiceServicesDataSourceCrud) Get() error {
	request := oci_announcements_service.ListServicesRequest{}

	if commsManagerName, ok := s.D.GetOkExists("comms_manager_name"); ok {
		request.CommsManagerName = oci_announcements_service.ListServicesCommsManagerNameEnum(commsManagerName.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if platformType, ok := s.D.GetOkExists("platform_type"); ok {
		request.PlatformType = oci_announcements_service.ListServicesPlatformTypeEnum(platformType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "announcements_service")

	response, err := s.Client.ListServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AnnouncementsServiceServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AnnouncementsServiceServicesDataSource-", AnnouncementsServiceServicesDataSource(), s.D))
	resources := []map[string]interface{}{}
	service := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceSummaryToMap(item))
	}
	service["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AnnouncementsServiceServicesDataSource().Schema["services_collection"].Elem.(*schema.Resource).Schema)
		service["items"] = items
	}

	resources = append(resources, service)
	if err := s.D.Set("services_collection", resources); err != nil {
		return err
	}

	return nil
}

func ServiceSummaryToMap(obj oci_announcements_service.ServiceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["comms_manager_name"] = string(obj.CommsManagerName)

	result["excluded_realms"] = obj.ExcludedRealms

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["platform_type"] = string(obj.PlatformType)

	result["previous_service_names"] = obj.PreviousServiceNames

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	if obj.ShortName != nil {
		result["short_name"] = string(*obj.ShortName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TeamName != nil {
		result["team_name"] = string(*obj.TeamName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
