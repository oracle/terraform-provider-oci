// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
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
				// ShowConsoleHistoryData returns an error with length < 10240, though this is not documented in the API doc.
				ValidateFunc: func(i interface{}, k string) (s []string, es []error) {
					v, ok := i.(int)
					if !ok {
						es = append(es, fmt.Errorf("expected type of %s to be int", k))
						return
					}

					if v < 10240 {
						es = append(es, fmt.Errorf("expected %s to be less than %d, got %d", k, 10240, v))
						return
					}

					return
				},
			},
			"offset": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func readConsoleHistoryData(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	reader := &ConsoleHistoryDataDatasourceCrud{}
	reader.D = d
	reader.Client = client.client

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
