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

func JmsFleetLibraryApplicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetLibraryApplications,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"library_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_instance_id": {
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
			"library_application_usage_collection": {
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
									"application_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"first_seen_in_classpath": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_detected_dynamically": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_seen_in_classpath": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_count": {
										Type:     schema.TypeInt,
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

func readJmsFleetLibraryApplications(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetLibraryApplicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetLibraryApplicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListLibraryApplicationUsageResponse
}

func (s *JmsFleetLibraryApplicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetLibraryApplicationsDataSourceCrud) Get() error {
	request := oci_jms.ListLibraryApplicationUsageRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if applicationName, ok := s.D.GetOkExists("application_name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if applicationNameContains, ok := s.D.GetOkExists("application_name_contains"); ok {
		tmp := applicationNameContains.(string)
		request.ApplicationNameContains = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if libraryKey, ok := s.D.GetOkExists("library_key"); ok {
		tmp := libraryKey.(string)
		request.LibraryKey = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
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

	response, err := s.Client.ListLibraryApplicationUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLibraryApplicationUsage(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetLibraryApplicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetLibraryApplicationsDataSource-", JmsFleetLibraryApplicationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetLibraryApplication := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LibraryApplicationUsageSummaryToMap(item))
	}
	fleetLibraryApplication["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetLibraryApplicationsDataSource().Schema["library_application_usage_collection"].Elem.(*schema.Resource).Schema)
		fleetLibraryApplication["items"] = items
	}

	resources = append(resources, fleetLibraryApplication)
	if err := s.D.Set("library_application_usage_collection", resources); err != nil {
		return err
	}

	return nil
}

func LibraryApplicationUsageSummaryToMap(obj oci_jms.LibraryApplicationUsageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationKey != nil {
		result["application_key"] = string(*obj.ApplicationKey)
	}

	if obj.ApplicationName != nil {
		result["application_name"] = string(*obj.ApplicationName)
	}

	if obj.FirstSeenInClasspath != nil {
		result["first_seen_in_classpath"] = obj.FirstSeenInClasspath.String()
	}

	if obj.LastDetectedDynamically != nil {
		result["last_detected_dynamically"] = obj.LastDetectedDynamically.String()
	}

	if obj.LastSeenInClasspath != nil {
		result["last_seen_in_classpath"] = obj.LastSeenInClasspath.String()
	}

	if obj.ManagedInstanceCount != nil {
		result["managed_instance_count"] = int(*obj.ManagedInstanceCount)
	}

	return result
}
