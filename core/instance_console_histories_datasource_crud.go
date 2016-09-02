package core

import (
	"time"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type InstanceConsoleHistoriesDatasourceCrud struct {
	D              *schema.ResourceData
	Client         client.BareMetalClient
	ConsoleHistory *baremetal.ShowConsoleHistoryMetadataResponse
}

func (i *InstanceConsoleHistoriesDatasourceCrud) Get() (e error) {
	id := i.D.Get("instance_console_history_id").(string)

	i.ConsoleHistory = &baremetal.ShowConsoleHistoryMetadataResponse{}

	mostWeCanFetch := i.D.Get("limit").(int)
	snapshotStartPosition := 0

	for {
		opts := baremetal.Options{
			Length: mostWeCanFetch,
			Offset: snapshotStartPosition,
		}

		var res *baremetal.ShowConsoleHistoryMetadataResponse
		if res, e = i.Client.ShowConsoleHistoryData(id, opts); e != nil {
			break
		}

		i.ConsoleHistory.ConsoleHistoryData += res.ConsoleHistoryData
		// no more bytes in snapshot we're done
		if res.BytesRemaining == 0 {
			break
		}

		mostWeCanFetch -= len(res.ConsoleHistoryData)
		snapshotStartPosition += len(res.ConsoleHistoryData)

		// Can't fetch anymore so we're done
		if mostWeCanFetch <= 0 {
			break
		}
	}

	return
}

func (i *InstanceConsoleHistoriesDatasourceCrud) SetData() {
	i.D.SetId(time.Now().UTC().String())
	i.D.Set("console_history", i.ConsoleHistory.ConsoleHistoryData)
}
