// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/base64"
	"io/ioutil"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseAutonomousDatabaseWalletResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousDatabaseWallet,
		Read:     readDatabaseAutonomousDatabaseWallet,
		Delete:   deleteDatabaseAutonomousDatabaseWallet,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},

			// Optional
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"generate_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "SINGLE",
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},

			// Computed
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousDatabaseWallet(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseWalletResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseWallet(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseAutonomousDatabaseWallet(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousDatabaseWalletResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *[]byte
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDatabaseWalletResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseWalletResource-", DatabaseAutonomousDatabaseWalletResource(), s.D)
}

func (s *DatabaseAutonomousDatabaseWalletResourceCrud) Create() error {
	request := oci_database.GenerateAutonomousDatabaseWalletRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if generateType, ok := s.D.GetOkExists("generate_type"); ok {
		request.GenerateType = oci_database.GenerateAutonomousDatabaseWalletDetailsGenerateTypeEnum(generateType.(string))
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GenerateAutonomousDatabaseWallet(context.Background(), request)
	if err != nil {
		return err
	}

	if response.Content != nil {
		defer response.Content.Close()
		if contentBytes, err := ioutil.ReadAll(response.Content); err == nil {
			s.Res = &contentBytes
		} else {
			return err
		}
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseWalletResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseWalletResource-", DatabaseAutonomousDatabaseWalletResource(), s.D))

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	if base64EncodeContent {
		s.D.Set("content", base64.StdEncoding.EncodeToString(*s.Res))
	} else {
		s.D.Set("content", string(*s.Res))
	}

	return nil
}
