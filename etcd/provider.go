package etcd

import (
	"github.com/coreos/go-etcd/etcd"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns the ResourceProvider implemented by this package. Serve
// this with the Terraform plugin helper and you are golden.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"machines": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"etcd_key": resourceKey(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	ifMachines := d.Get("machines").([]interface{})
	strMachines := make([]string, len(ifMachines))
	for i, v := range ifMachines {
		strMachines[i] = v.(string)
	}

	api := etcd.NewClient(strMachines)
	return api, nil
}
