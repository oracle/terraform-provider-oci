// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsAgentInstallersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsAgentInstallers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform_architecture": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_installer_collection": {
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
									"agent_installer_description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"agent_installer_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"agent_installer_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"agent_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"approximate_file_size_in_bytes": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"java_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_family": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"package_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"platform_architecture": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sha256": {
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

func readJmsAgentInstallers(d *schema.ResourceData, m interface{}) error {
	sync := &JmsAgentInstallersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsAgentInstallersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListAgentInstallersResponse
}

func (s *JmsAgentInstallersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsAgentInstallersDataSourceCrud) Get() error {
	request := oci_jms.ListAgentInstallersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_jms.ListAgentInstallersOsFamilyEnum(osFamily.(string))
	}

	if platformArchitecture, ok := s.D.GetOkExists("platform_architecture"); ok {
		request.PlatformArchitecture = oci_jms.ListAgentInstallersPlatformArchitectureEnum(platformArchitecture.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListAgentInstallers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAgentInstallers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsAgentInstallersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsAgentInstallersDataSource-", JmsAgentInstallersDataSource(), s.D))
	resources := []map[string]interface{}{}
	agentInstaller := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AgentInstallerSummaryToMap(item))
	}
	agentInstaller["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsAgentInstallersDataSource().Schema["agent_installer_collection"].Elem.(*schema.Resource).Schema)
		agentInstaller["items"] = items
	}

	resources = append(resources, agentInstaller)
	if err := s.D.Set("agent_installer_collection", resources); err != nil {
		return err
	}

	return nil
}

func AgentInstallerSummaryToMap(obj oci_jms.AgentInstallerSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentInstallerDescription != nil {
		result["agent_installer_description"] = string(*obj.AgentInstallerDescription)
	}

	if obj.AgentInstallerId != nil {
		result["agent_installer_id"] = strconv.FormatInt(*obj.AgentInstallerId, 10)
	}

	if obj.AgentInstallerVersion != nil {
		result["agent_installer_version"] = string(*obj.AgentInstallerVersion)
	}

	if obj.AgentVersion != nil {
		result["agent_version"] = string(*obj.AgentVersion)
	}

	if obj.ApproximateFileSizeInBytes != nil {
		result["approximate_file_size_in_bytes"] = strconv.FormatInt(*obj.ApproximateFileSizeInBytes, 10)
	}

	if obj.JavaVersion != nil {
		result["java_version"] = string(*obj.JavaVersion)
	}

	result["os_family"] = string(obj.OsFamily)

	result["package_type"] = string(obj.PackageType)

	result["platform_architecture"] = string(obj.PlatformArchitecture)

	if obj.Sha256 != nil {
		result["sha256"] = string(*obj.Sha256)
	}

	return result
}
