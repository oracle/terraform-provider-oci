// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_analytics "github.com/oracle/oci-go-sdk/v56/analytics"
)

func AnalyticsAnalyticsInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAnalyticsAnalyticsInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"capacity_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"feature_set": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"analytics_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(AnalyticsAnalyticsInstanceResource()),
			},
		},
	}
}

func readAnalyticsAnalyticsInstances(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.ReadResource(sync)
}

type AnalyticsAnalyticsInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_analytics.AnalyticsClient
	Res    *oci_analytics.ListAnalyticsInstancesResponse
}

func (s *AnalyticsAnalyticsInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnalyticsAnalyticsInstancesDataSourceCrud) Get() error {
	request := oci_analytics.ListAnalyticsInstancesRequest{}

	if capacityType, ok := s.D.GetOkExists("capacity_type"); ok {
		request.CapacityType = oci_analytics.ListAnalyticsInstancesCapacityTypeEnum(capacityType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if featureSet, ok := s.D.GetOkExists("feature_set"); ok {
		request.FeatureSet = oci_analytics.ListAnalyticsInstancesFeatureSetEnum(featureSet.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_analytics.ListAnalyticsInstancesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "analytics")

	response, err := s.Client.ListAnalyticsInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAnalyticsInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AnalyticsAnalyticsInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AnalyticsAnalyticsInstancesDataSource-", AnalyticsAnalyticsInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		analyticsInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Capacity != nil {
			analyticsInstance["capacity"] = []interface{}{AnalyticsCapacityToMap(r.Capacity)}
		} else {
			analyticsInstance["capacity"] = nil
		}

		if r.Description != nil {
			analyticsInstance["description"] = *r.Description
		}

		if r.EmailNotification != nil {
			analyticsInstance["email_notification"] = *r.EmailNotification
		}

		analyticsInstance["feature_set"] = r.FeatureSet

		if r.Id != nil {
			analyticsInstance["id"] = *r.Id
		}

		analyticsInstance["license_type"] = r.LicenseType

		if r.Name != nil {
			analyticsInstance["name"] = *r.Name
		}

		if r.NetworkEndpointDetails != nil {
			networkEndpointDetailsArray := []interface{}{}
			if networkEndpointDetailsMap := NetworkEndpointDetailsToMap(&r.NetworkEndpointDetails); networkEndpointDetailsMap != nil {
				networkEndpointDetailsArray = append(networkEndpointDetailsArray, networkEndpointDetailsMap)
			}
			analyticsInstance["network_endpoint_details"] = networkEndpointDetailsArray
		} else {
			analyticsInstance["network_endpoint_details"] = nil
		}

		if r.ServiceUrl != nil {
			analyticsInstance["service_url"] = *r.ServiceUrl
		}

		analyticsInstance["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			analyticsInstance["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			analyticsInstance["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, analyticsInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, AnalyticsAnalyticsInstancesDataSource().Schema["analytics_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("analytics_instances", resources); err != nil {
		return err
	}

	return nil
}
