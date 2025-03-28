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

func OsManagementHubManagedInstanceDetachProfileManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagedInstanceDetachProfileManagement,
		Read:     readOsManagementHubManagedInstanceDetachProfileManagement,
		Delete:   deleteOsManagementHubManagedInstanceDetachProfileManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createOsManagementHubManagedInstanceDetachProfileManagement(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceDetachProfileManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubManagedInstanceDetachProfileManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteOsManagementHubManagedInstanceDetachProfileManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OsManagementHubManagedInstanceDetachProfileManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagedInstanceClient
	Res                    *string
	DisableNotFoundRetries bool
}

func (s *OsManagementHubManagedInstanceDetachProfileManagementResourceCrud) ID() string {
	return *s.Res
}

func (s *OsManagementHubManagedInstanceDetachProfileManagementResourceCrud) Create() error {
	request := oci_os_management_hub.DetachProfileFromManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
		s.Res = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DetachProfileFromManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	return nil
}

func (s *OsManagementHubManagedInstanceDetachProfileManagementResourceCrud) SetData() error {
	return nil
}
