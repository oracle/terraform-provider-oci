// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDataWarehouseWalletDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAutonomousDataWarehouseWallet,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_data_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},

			// Computed
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularAutonomousDataWarehouseWallet(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseWalletDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDataWarehouseWalletDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *[]byte
}

func (s *AutonomousDataWarehouseWalletDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDataWarehouseWalletDataSourceCrud) Get() error {
	request := oci_database.GenerateAutonomousDataWarehouseWalletRequest{}

	if autonomousDataWarehouseId, ok := s.D.GetOkExists("autonomous_data_warehouse_id"); ok {
		tmp := autonomousDataWarehouseId.(string)
		request.AutonomousDataWarehouseId = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GenerateAutonomousDataWarehouseWallet(context.Background(), request)
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

func (s *AutonomousDataWarehouseWalletDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("content", string(*s.Res))

	return nil
}
