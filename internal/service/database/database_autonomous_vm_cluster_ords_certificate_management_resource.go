// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousVmClusterOrdsCertificateManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousVmClusterOrdsCertificateManagement,
		Read:     readDatabaseAutonomousVmClusterOrdsCertificateManagement,
		Delete:   deleteDatabaseAutonomousVmClusterOrdsCertificateManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_generation_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ca_bundle_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certificate_authority_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDatabaseAutonomousVmClusterOrdsCertificateManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterOrdsCertificateManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousVmClusterOrdsCertificateManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseAutonomousVmClusterOrdsCertificateManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousVmClusterOrdsCertificateManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.RotateAutonomousVmClusterOrdsCertsResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousVmClusterOrdsCertificateManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseAutonomousVmClusterOrdsCertificateManagementResource-", DatabaseAutonomousVmClusterOrdsCertificateManagementResource(), s.D)
}

func (s *DatabaseAutonomousVmClusterOrdsCertificateManagementResourceCrud) Create() error {
	request := oci_database.RotateAutonomousVmClusterOrdsCertsRequest{}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	if caBundleId, ok := s.D.GetOkExists("ca_bundle_id"); ok {
		tmp := caBundleId.(string)
		request.CaBundleId = &tmp
	}

	if certificateAuthorityId, ok := s.D.GetOkExists("certificate_authority_id"); ok {
		tmp := certificateAuthorityId.(string)
		request.CertificateAuthorityId = &tmp
	}

	if certificateGenerationType, ok := s.D.GetOkExists("certificate_generation_type"); ok {
		request.CertificateGenerationType = oci_database.RotateAutonomousVmClusterOrdsCertsDetailsCertificateGenerationTypeEnum(certificateGenerationType.(string))
	}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.RotateAutonomousVmClusterOrdsCerts(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousvmcluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseAutonomousVmClusterOrdsCertificateManagementResourceCrud) SetData() error {
	return nil
}
