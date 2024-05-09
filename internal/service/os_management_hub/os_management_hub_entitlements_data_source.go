// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubEntitlementsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubEntitlements,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"csi": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entitlement_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"csi": {
							Type:     schema.TypeString,
							Required: true,
						},

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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csi": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vendor_name": {
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

func readOsManagementHubEntitlements(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEntitlementsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubEntitlementsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListEntitlementsResponse
}

func (s *OsManagementHubEntitlementsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubEntitlementsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListEntitlementsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if csi, ok := s.D.GetOkExists("csi"); ok {
		tmp := csi.(string)
		request.Csi = &tmp
	}

	if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
		request.VendorName = oci_os_management_hub.ListEntitlementsVendorNameEnum(vendorName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListEntitlements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEntitlements(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubEntitlementsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubEntitlementsDataSource-", OsManagementHubEntitlementsDataSource(), s.D))
	resources := []map[string]interface{}{}
	entitlement := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EntitlementSummaryToMap(item))
	}
	entitlement["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubEntitlementsDataSource().Schema["entitlement_collection"].Elem.(*schema.Resource).Schema)
		entitlement["items"] = items
	}

	resources = append(resources, entitlement)
	if err := s.D.Set("entitlement_collection", resources); err != nil {
		return err
	}

	return nil
}

func EntitlementSummaryToMap(obj oci_os_management_hub.EntitlementSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Csi != nil {
		result["csi"] = string(*obj.Csi)
	}

	if obj.VendorName != nil {
		result["vendor_name"] = string(*obj.VendorName)
	}

	return result
}
