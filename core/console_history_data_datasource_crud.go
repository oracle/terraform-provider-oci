package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type ConsoleHistoryDataDatasourceCrud struct {
	D                  *schema.ResourceData
	Client             client.BareMetalClient
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
