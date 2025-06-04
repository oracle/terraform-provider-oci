// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedMySqlDatabaseBinaryLogInformation,
		Schema: map[string]*schema.Schema{
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"binary_log_compression": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"binary_log_compression_percent": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"binary_log_format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"binary_log_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"binary_log_position": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"binary_logging": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedMySqlDatabaseBinaryLogInformation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.GetBinaryLogInformationResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSourceCrud) Get() error {
	request := oci_database_management.GetBinaryLogInformationRequest{}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetBinaryLogInformation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSource-", DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSource(), s.D))

	if s.Res.BinaryLogCompression != nil {
		s.D.Set("binary_log_compression", *s.Res.BinaryLogCompression)
	}

	if s.Res.BinaryLogCompressionPercent != nil {
		s.D.Set("binary_log_compression_percent", *s.Res.BinaryLogCompressionPercent)
	}

	if s.Res.BinaryLogFormat != nil {
		s.D.Set("binary_log_format", *s.Res.BinaryLogFormat)
	}

	if s.Res.BinaryLogName != nil {
		s.D.Set("binary_log_name", *s.Res.BinaryLogName)
	}

	if s.Res.BinaryLogPosition != nil {
		s.D.Set("binary_log_position", strconv.FormatInt(*s.Res.BinaryLogPosition, 10))
	}

	if s.Res.BinaryLogging != nil {
		s.D.Set("binary_logging", *s.Res.BinaryLogging)
	}

	return nil
}
