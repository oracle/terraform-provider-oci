// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsManagedInstanceScanResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsManagedInstanceScanResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_result_collection": {
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
									"server_check_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_check_result": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_check_result_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_check_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_of_server_check": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wls_domain_id": {
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

func readWlmsManagedInstanceScanResults(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsManagedInstanceScanResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsManagedInstanceScanResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListManagedInstanceScanResultsResponse
}

func (s *WlmsManagedInstanceScanResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsManagedInstanceScanResultsDataSourceCrud) Get() error {
	request := oci_wlms.ListManagedInstanceScanResultsRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if serverName, ok := s.D.GetOkExists("server_name"); ok {
		tmp := serverName.(string)
		request.ServerName = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListManagedInstanceScanResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceScanResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsManagedInstanceScanResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsManagedInstanceScanResultsDataSource-", WlmsManagedInstanceScanResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceScanResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScanResultSummaryToMap(item))
	}
	managedInstanceScanResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsManagedInstanceScanResultsDataSource().Schema["scan_result_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceScanResult["items"] = items
	}

	resources = append(resources, managedInstanceScanResult)
	if err := s.D.Set("scan_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func ScanResultSummaryToMap(obj oci_wlms.ScanResultSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ServerCheckName != nil {
		result["server_check_name"] = string(*obj.ServerCheckName)
	}

	if obj.ServerCheckResult != nil {
		result["server_check_result"] = string(*obj.ServerCheckResult)
	}

	if obj.ServerCheckResultId != nil {
		result["server_check_result_id"] = string(*obj.ServerCheckResultId)
	}

	result["server_check_status"] = string(obj.ServerCheckStatus)

	if obj.ServerName != nil {
		result["server_name"] = string(*obj.ServerName)
	}

	if obj.TimeOfServerCheck != nil {
		result["time_of_server_check"] = obj.TimeOfServerCheck.String()
	}

	if obj.WlsDomainId != nil {
		result["wls_domain_id"] = string(*obj.WlsDomainId)
	}

	return result
}
