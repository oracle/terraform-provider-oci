// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccCustomerGroupOccCustomerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCapacityManagementOccCustomerGroupOccCustomer,
		Read:     readCapacityManagementOccCustomerGroupOccCustomer,
		Update:   updateCapacityManagementOccCustomerGroupOccCustomer,
		Delete:   deleteCapacityManagementOccCustomerGroupOccCustomer,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"occ_customer_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createCapacityManagementOccCustomerGroupOccCustomer(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupOccCustomerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCapacityManagementOccCustomerGroupOccCustomer(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateCapacityManagementOccCustomerGroupOccCustomer(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupOccCustomerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCapacityManagementOccCustomerGroupOccCustomer(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccCustomerGroupOccCustomerResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CapacityManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CapacityManagementOccCustomerGroupOccCustomerResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_capacity_management.CapacityManagementClient
	Res                    *oci_capacity_management.OccCustomer
	DisableNotFoundRetries bool
}

func (s *CapacityManagementOccCustomerGroupOccCustomerResourceCrud) ID() string {
	return *s.Res.TenancyId
}

func (s *CapacityManagementOccCustomerGroupOccCustomerResourceCrud) Create() error {
	request := oci_capacity_management.CreateOccCustomerRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_capacity_management.CreateOccCustomerDetailsStatusEnum(status.(string))
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.CreateOccCustomer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCustomer
	return nil
}

func (s *CapacityManagementOccCustomerGroupOccCustomerResourceCrud) Update() error {
	request := oci_capacity_management.UpdateOccCustomerRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if occCustomerId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := occCustomerId.(string)
		request.OccCustomerId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_capacity_management.UpdateOccCustomerDetailsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	response, err := s.Client.UpdateOccCustomer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OccCustomer
	return nil
}

func (s *CapacityManagementOccCustomerGroupOccCustomerResourceCrud) Delete() error {
	request := oci_capacity_management.DeleteOccCustomerRequest{}

	if occCustomerGroupId, ok := s.D.GetOkExists("occ_customer_group_id"); ok {
		tmp := occCustomerGroupId.(string)
		request.OccCustomerGroupId = &tmp
	}

	if occCustomerId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := occCustomerId.(string)
		request.OccCustomerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "capacity_management")

	_, err := s.Client.DeleteOccCustomer(context.Background(), request)
	return err
}

func (s *CapacityManagementOccCustomerGroupOccCustomerResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.OccCustomerGroupId != nil {
		s.D.Set("occ_customer_group_id", *s.Res.OccCustomerGroupId)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	return nil
}
