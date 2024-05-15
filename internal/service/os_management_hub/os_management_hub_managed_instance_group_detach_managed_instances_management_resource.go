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

func OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceGroupDetachManagedInstancesManagement,
		Read:     readOsManagementHubManagedInstanceGroupDetachManagedInstancesManagement,
		Delete:   deleteOsManagementHubManagedInstanceGroupDetachManagedInstancesManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"managed_instances": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional

			// Computed
		},
	}
}

func createOsManagementHubManagedInstanceGroupDetachManagedInstancesManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceGroupDetachManagedInstancesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceGroupDetachManagedInstancesManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceGroupClient
	Res                    *oci_os_management_hub.GetManagedInstanceGroupResponse
	DisableNotFoundRetries bool
}

func (s *OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceCrud) Create() error {
	request := oci_os_management_hub.DetachManagedInstancesFromManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	if managedInstances, ok := s.D.GetOkExists("managed_instances"); ok {
		interfaces := managedInstances.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("managed_instances") {
			request.ManagedInstances = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DetachManagedInstancesFromManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceCrud) SetData() error {
	return nil
}
