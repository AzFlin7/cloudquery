package compute

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func DedicatedVmHosts() *schema.Table {
	return &schema.Table{
		Name:      "oracle_compute_dedicated_vm_hosts",
		Resolver:  fetchDedicatedVmHosts,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&core.DedicatedVmHostSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}
