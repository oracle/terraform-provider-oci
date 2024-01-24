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

func JmsFleetInstallationSitesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetInstallationSites,
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
			"installation_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jre_distribution": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jre_security_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jre_vendor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jre_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"path_contains": {
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
			"installation_site_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
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
												"approximate_application_count": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"blocklist": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"operation": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"reason": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"installation_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"jre": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"distribution": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"jre_key": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"vendor": {
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
												"managed_instance_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"operating_system": {
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
															"family": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"managed_instance_count": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"name": {
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
												"path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"security_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_last_seen": {
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
				},
			},
		},
	}
}

func readJmsFleetInstallationSites(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetInstallationSitesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetInstallationSitesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListInstallationSitesResponse
}

func (s *JmsFleetInstallationSitesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetInstallationSitesDataSourceCrud) Get() error {
	request := oci_jms.ListInstallationSitesRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if installationPath, ok := s.D.GetOkExists("installation_path"); ok {
		tmp := installationPath.(string)
		request.InstallationPath = &tmp
	}

	if jreDistribution, ok := s.D.GetOkExists("jre_distribution"); ok {
		tmp := jreDistribution.(string)
		request.JreDistribution = &tmp
	}

	if jreSecurityStatus, ok := s.D.GetOkExists("jre_security_status"); ok {
		request.JreSecurityStatus = oci_jms.ListInstallationSitesJreSecurityStatusEnum(jreSecurityStatus.(string))
	}

	if jreVendor, ok := s.D.GetOkExists("jre_vendor"); ok {
		tmp := jreVendor.(string)
		request.JreVendor = &tmp
	}

	if jreVersion, ok := s.D.GetOkExists("jre_version"); ok {
		tmp := jreVersion.(string)
		request.JreVersion = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		interfaces := osFamily.([]interface{})
		tmp := make([]oci_jms.OsFamilyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_jms.OsFamilyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("os_family") {
			request.OsFamily = tmp
		}
	}

	if pathContains, ok := s.D.GetOkExists("path_contains"); ok {
		tmp := pathContains.(string)
		request.PathContains = &tmp
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

	response, err := s.Client.ListInstallationSites(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstallationSites(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetInstallationSitesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetInstallationSitesDataSource-", JmsFleetInstallationSitesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetInstallationSite := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InstallationSiteSummaryToMap(item))
	}
	fleetInstallationSite["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetInstallationSitesDataSource().Schema["installation_site_collection"].Elem.(*schema.Resource).Schema)
		fleetInstallationSite["items"] = items
	}

	resources = append(resources, fleetInstallationSite)
	if err := s.D.Set("installation_site_collection", resources); err != nil {
		return err
	}

	return nil
}

func BlocklistEntryToMap(obj oci_jms.BlocklistEntry) map[string]interface{} {
	result := map[string]interface{}{}

	result["operation"] = string(obj.Operation)

	if obj.Reason != nil {
		result["reason"] = string(*obj.Reason)
	}

	return result
}

func InstallationSiteSummaryToMap(obj oci_jms.InstallationSiteSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApproximateApplicationCount != nil {
		result["approximate_application_count"] = int(*obj.ApproximateApplicationCount)
	}

	blocklist := []interface{}{}
	for _, item := range obj.Blocklist {
		blocklist = append(blocklist, BlocklistEntryToMap(item))
	}
	result["blocklist"] = blocklist

	if obj.InstallationKey != nil {
		result["installation_key"] = string(*obj.InstallationKey)
	}

	if obj.Jre != nil {
		result["jre"] = []interface{}{JavaRuntimeIdToMap(obj.Jre)}
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.OperatingSystem != nil {
		result["operating_system"] = []interface{}{OperatingSystemToMap(obj.OperatingSystem)}
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	result["security_status"] = string(obj.SecurityStatus)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeLastSeen != nil {
		result["time_last_seen"] = obj.TimeLastSeen.String()
	}

	return result
}

func JavaRuntimeIdToMap(obj *oci_jms.JavaRuntimeId) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Distribution != nil {
		result["distribution"] = string(*obj.Distribution)
	}

	if obj.JreKey != nil {
		result["jre_key"] = string(*obj.JreKey)
	}

	if obj.Vendor != nil {
		result["vendor"] = string(*obj.Vendor)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func OperatingSystemToMap(obj *oci_jms.OperatingSystem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Architecture != nil {
		result["architecture"] = string(*obj.Architecture)
	}

	result["family"] = string(obj.Family)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
