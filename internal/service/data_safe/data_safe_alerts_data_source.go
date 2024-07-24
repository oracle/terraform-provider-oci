// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAlertsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAlerts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"field": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataSafeAlertResource()),
						},
					},
				},
			},
		},
	}
}

func readDataSafeAlerts(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAlertsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAlertsResponse
}

func (s *DataSafeAlertsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAlertsDataSourceCrud) Get() error {
	request := oci_data_safe.ListAlertsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAlertsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if field, ok := s.D.GetOkExists("field"); ok {
		interfaces := field.([]interface{})
		tmp := make([]oci_data_safe.ListAlertsFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListAlertsFieldEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("field") {
			request.Field = tmp
		}
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAlerts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlerts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAlertsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAlertsDataSource-", DataSafeAlertsDataSource(), s.D))
	resources := []map[string]interface{}{}
	alert := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AlertSummaryToMap(item))
	}
	alert["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAlertsDataSource().Schema["alert_collection"].Elem.(*schema.Resource).Schema)
		alert["items"] = items
	}

	resources = append(resources, alert)
	if err := s.D.Set("alert_collection", resources); err != nil {
		return err
	}

	return nil
}
