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

func JmsFleetInstallationSiteDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetInstallationSite,
		Schema: map[string]*schema.Schema{
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_jms_fleet_installation_site", "oci_jms_fleet_installation_sites"),
	}
}

func readSingularJmsFleetInstallationSite(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetInstallationSiteDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetInstallationSiteDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListInstallationSitesResponse
}

func (s *JmsFleetInstallationSiteDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetInstallationSiteDataSourceCrud) Get() error {
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
	return nil
}

func (s *JmsFleetInstallationSiteDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetInstallationSiteDataSource-", JmsFleetInstallationSiteDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InstallationSiteSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
