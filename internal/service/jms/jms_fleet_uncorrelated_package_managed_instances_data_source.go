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

func JmsFleetUncorrelatedPackageManagedInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetUncorrelatedPackageManagedInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uncorrelated_package_managed_instance_usage_collection": {
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
									"application_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_detected_dynamically": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
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

func readJmsFleetUncorrelatedPackageManagedInstances(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetUncorrelatedPackageManagedInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetUncorrelatedPackageManagedInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListUncorrelatedPackageManagedInstanceUsageResponse
}

func (s *JmsFleetUncorrelatedPackageManagedInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetUncorrelatedPackageManagedInstancesDataSourceCrud) Get() error {
	request := oci_jms.ListUncorrelatedPackageManagedInstanceUsageRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if packageName, ok := s.D.GetOkExists("package_name"); ok {
		tmp := packageName.(string)
		request.PackageName = &tmp
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

	response, err := s.Client.ListUncorrelatedPackageManagedInstanceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListUncorrelatedPackageManagedInstanceUsage(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetUncorrelatedPackageManagedInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetUncorrelatedPackageManagedInstancesDataSource-", JmsFleetUncorrelatedPackageManagedInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetUncorrelatedPackageManagedInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, UncorrelatedPackageManagedInstanceUsageSummaryToMap(item))
	}
	fleetUncorrelatedPackageManagedInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetUncorrelatedPackageManagedInstancesDataSource().Schema["uncorrelated_package_managed_instance_usage_collection"].Elem.(*schema.Resource).Schema)
		fleetUncorrelatedPackageManagedInstance["items"] = items
	}

	resources = append(resources, fleetUncorrelatedPackageManagedInstance)
	if err := s.D.Set("uncorrelated_package_managed_instance_usage_collection", resources); err != nil {
		return err
	}

	return nil
}

func UncorrelatedPackageManagedInstanceUsageSummaryToMap(obj oci_jms.UncorrelatedPackageManagedInstanceUsageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationCount != nil {
		result["application_count"] = int(*obj.ApplicationCount)
	}

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.LastDetectedDynamically != nil {
		result["last_detected_dynamically"] = obj.LastDetectedDynamically.String()
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	return result
}
