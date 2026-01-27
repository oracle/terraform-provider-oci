// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceSnapsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readOsManagementHubManagedInstanceSnapsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snap_collection": {
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
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"publisher": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"revision": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"store_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_refreshed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tracking": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
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

func readOsManagementHubManagedInstanceSnapsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubManagedInstanceSnapsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OsManagementHubManagedInstanceSnapsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.ListManagedInstanceSnapsResponse
}

func (s *OsManagementHubManagedInstanceSnapsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceSnapsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_os_management_hub.ListManagedInstanceSnapsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceSnaps(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceSnaps(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceSnapsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceSnapsDataSource-", OsManagementHubManagedInstanceSnapsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceSnap := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SnapSummaryToMap(item))
	}
	managedInstanceSnap["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceSnapsDataSource().Schema["snap_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceSnap["items"] = items
	}

	resources = append(resources, managedInstanceSnap)
	if err := s.D.Set("snap_collection", resources); err != nil {
		return err
	}

	return nil
}

func SnapSummaryToMap(obj oci_os_management_hub.SnapSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Publisher != nil {
		result["publisher"] = string(*obj.Publisher)
	}

	if obj.Revision != nil {
		result["revision"] = string(*obj.Revision)
	}

	if obj.StoreUrl != nil {
		result["store_url"] = string(*obj.StoreUrl)
	}

	if obj.TimeRefreshed != nil {
		result["time_refreshed"] = obj.TimeRefreshed.String()
	}

	if obj.Tracking != nil {
		result["tracking"] = string(*obj.Tracking)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
