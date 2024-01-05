// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsAnnouncementsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsAnnouncements,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"summary_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"announcement_collection": {
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
									"key": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_released": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
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

func readJmsAnnouncements(d *schema.ResourceData, m interface{}) error {
	sync := &JmsAnnouncementsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsAnnouncementsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListAnnouncementsResponse
}

func (s *JmsAnnouncementsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsAnnouncementsDataSourceCrud) Get() error {
	request := oci_jms.ListAnnouncementsRequest{}

	if summaryContains, ok := s.D.GetOkExists("summary_contains"); ok {
		tmp := summaryContains.(string)
		request.SummaryContains = &tmp
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListAnnouncements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAnnouncements(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsAnnouncementsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsAnnouncementsDataSource-", JmsAnnouncementsDataSource(), s.D))
	resources := []map[string]interface{}{}
	announcement := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AnnouncementSummaryToMap(item))
	}
	announcement["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsAnnouncementsDataSource().Schema["announcement_collection"].Elem.(*schema.Resource).Schema)
		announcement["items"] = items
	}

	resources = append(resources, announcement)
	if err := s.D.Set("announcement_collection", resources); err != nil {
		return err
	}

	return nil
}

func AnnouncementSummaryToMap(obj oci_jms.AnnouncementSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = int(*obj.Key)
	}

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	if obj.TimeReleased != nil {
		result["time_released"] = obj.TimeReleased.String()
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}
