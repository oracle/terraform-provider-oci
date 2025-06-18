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

func WlmsWlsDomainScanResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsWlsDomainScanResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"wls_domain_id": {
				Type:     schema.TypeString,
				Required: true,
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

func readWlmsWlsDomainScanResults(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsWlsDomainScanResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsWlsDomainScanResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListWlsDomainScanResultsResponse
}

func (s *WlmsWlsDomainScanResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsWlsDomainScanResultsDataSourceCrud) Get() error {
	request := oci_wlms.ListWlsDomainScanResultsRequest{}

	if serverName, ok := s.D.GetOkExists("server_name"); ok {
		tmp := serverName.(string)
		request.ServerName = &tmp
	}

	if wlsDomainId, ok := s.D.GetOkExists("wls_domain_id"); ok {
		tmp := wlsDomainId.(string)
		request.WlsDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListWlsDomainScanResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWlsDomainScanResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsWlsDomainScanResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsWlsDomainScanResultsDataSource-", WlmsWlsDomainScanResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	wlsDomainScanResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, WlsDomainScanResultSummaryToMap(item))
	}
	wlsDomainScanResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsWlsDomainScanResultsDataSource().Schema["scan_result_collection"].Elem.(*schema.Resource).Schema)
		wlsDomainScanResult["items"] = items
	}

	resources = append(resources, wlsDomainScanResult)
	if err := s.D.Set("scan_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func WlsDomainScanResultSummaryToMap(obj oci_wlms.ScanResultSummary) map[string]interface{} {
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
