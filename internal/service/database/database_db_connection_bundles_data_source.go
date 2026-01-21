// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbConnectionBundlesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseDbConnectionBundlesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"associated_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_connection_bundle_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_connection_bundles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"associated_resource_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"resource_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_connection_bundle_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_protected": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_refreshed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbConnectionBundlesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDbConnectionBundlesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseDbConnectionBundlesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbConnectionBundlesResponse
}

func (s *DatabaseDbConnectionBundlesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbConnectionBundlesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.ListDbConnectionBundlesRequest{}

	if associatedResourceId, ok := s.D.GetOkExists("associated_resource_id"); ok {
		tmp := associatedResourceId.(string)
		request.AssociatedResourceId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbConnectionBundleType, ok := s.D.GetOkExists("db_connection_bundle_type"); ok {
		request.DbConnectionBundleType = oci_database.ListDbConnectionBundlesDbConnectionBundleTypeEnum(dbConnectionBundleType.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ListDbConnectionBundlesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbConnectionBundles(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbConnectionBundles(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbConnectionBundlesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbConnectionBundlesDataSource-", DatabaseDbConnectionBundlesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbConnectionBundle := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		associatedResourceDetails := []interface{}{}
		for _, item := range r.AssociatedResourceDetails {
			associatedResourceDetails = append(associatedResourceDetails, AssociatedResourceDetailsToMap(item))
		}
		dbConnectionBundle["associated_resource_details"] = associatedResourceDetails

		dbConnectionBundle["db_connection_bundle_type"] = r.DbConnectionBundleType

		if r.DefinedTags != nil {
			dbConnectionBundle["defined_tags"] = r.DefinedTags
		}

		if r.DisplayName != nil {
			dbConnectionBundle["display_name"] = *r.DisplayName
		}

		dbConnectionBundle["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			dbConnectionBundle["id"] = *r.Id
		}

		if r.IsProtected != nil {
			dbConnectionBundle["is_protected"] = *r.IsProtected
		}

		dbConnectionBundle["state"] = r.LifecycleState

		if r.SystemTags != nil {
			dbConnectionBundle["system_tags"] = r.SystemTags
		}

		if r.TimeCreated != nil {
			dbConnectionBundle["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastRefreshed != nil {
			dbConnectionBundle["time_last_refreshed"] = r.TimeLastRefreshed.String()
		}

		if r.TimeUpdated != nil {
			dbConnectionBundle["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, dbConnectionBundle)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbConnectionBundlesDataSource().Schema["db_connection_bundles"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_connection_bundles", resources); err != nil {
		return err
	}

	return nil
}

func AssociatedResourceDetailsToMap(obj oci_database.AssociatedResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["resource_ids"] = obj.ResourceIds

	return result
}
