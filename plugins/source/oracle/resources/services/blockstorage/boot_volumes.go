package blockstorage

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func BootVolumes() *schema.Table {
	return &schema.Table{
		Name:      "oracle_blockstorage_boot_volumes",
		Resolver:  fetchBootVolumes,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.BootVolume{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
