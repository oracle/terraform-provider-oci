// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagedInstanceInstalledPackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagedInstanceInstalledPackages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_install_date_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_install_date_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"installed_package_collection": {
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
									"architecture": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"package_classification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_sources": {
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
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_mandatory_for_autonomous_linux": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"software_source_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"time_installed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_issued": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
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

func readOsManagementHubManagedInstanceInstalledPackages(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceInstalledPackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceInstalledPackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.ListManagedInstanceInstalledPackagesResponse
}

func (s *OsManagementHubManagedInstanceInstalledPackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceInstalledPackagesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagedInstanceInstalledPackagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if timeInstallDateEnd, ok := s.D.GetOkExists("time_install_date_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeInstallDateEnd.(string))
		if err != nil {
			return err
		}
		request.TimeInstallDateEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeInstallDateStart, ok := s.D.GetOkExists("time_install_date_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeInstallDateStart.(string))
		if err != nil {
			return err
		}
		request.TimeInstallDateStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstanceInstalledPackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstanceInstalledPackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagedInstanceInstalledPackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagedInstanceInstalledPackagesDataSource-", OsManagementHubManagedInstanceInstalledPackagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstanceInstalledPackage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InstalledPackageSummaryToMap(item))
	}
	managedInstanceInstalledPackage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagedInstanceInstalledPackagesDataSource().Schema["installed_package_collection"].Elem.(*schema.Resource).Schema)
		managedInstanceInstalledPackage["items"] = items
	}

	resources = append(resources, managedInstanceInstalledPackage)
	if err := s.D.Set("installed_package_collection", resources); err != nil {
		return err
	}

	return nil
}

func InstalledPackageSummaryToMap(obj oci_os_management_hub.InstalledPackageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["architecture"] = string(obj.Architecture)

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	// FIXME
	//result["package_classification"] = string(obj.PackageClassification)

	softwareSources := []interface{}{}
	for _, item := range obj.SoftwareSources {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
	}
	result["software_sources"] = softwareSources

	if obj.TimeInstalled != nil {
		result["time_installed"] = obj.TimeInstalled.String()
	}

	if obj.TimeIssued != nil {
		result["time_issued"] = obj.TimeIssued.String()
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
