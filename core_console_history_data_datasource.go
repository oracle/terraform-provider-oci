// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func ConsoleHistoryDataDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readConsoleHistoryData,
		Schema: map[string]*schema.Schema{
			"console_history_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"offset": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func readConsoleHistoryData(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &ConsoleHistoryDataDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type ConsoleHistoryDataDatasourceCrud struct {
	crud.BaseCrud
	ConsoleHistoryData *baremetal.ConsoleHistoryData
}

func (res *ConsoleHistoryDataDatasourceCrud) Get() (e error) {
	id := res.D.Get("console_history_id").(string)

	opts := &baremetal.ConsoleHistoryDataOptions{}
	opts.Length = uint64(res.D.Get("length").(int))
	opts.Offset = uint64(res.D.Get("offset").(int))

	res.ConsoleHistoryData, e = res.Client.ShowConsoleHistoryData(id, opts)

	return
}

func (res *ConsoleHistoryDataDatasourceCrud) SetData() {
	res.D.SetId(time.Now().UTC().String())
	res.D.Set("data", res.ConsoleHistoryData.Data)
}
