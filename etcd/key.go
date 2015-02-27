package etcd

import (
	"fmt"

	etcdErrors "github.com/coreos/etcd/error"
	"github.com/coreos/go-etcd/etcd"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKey() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},

		Create: resourceKeySet,
		Read:   resourceKeyRead,
		Delete: resourceKeyDelete,
		Update: resourceKeySet,
	}
}

func resourceKeySet(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*etcd.Client)

	key := d.Get("key").(string)
	value := d.Get("value").(string)

	_, err := api.Set(key, value, 0)
	if err != nil {
		return fmt.Errorf("could not set key %s: %s", key, err)
	}

	d.SetId(key)
	if err := d.Set("key", key); err != nil {
		return err
	}
	if err := d.Set("value", value); err != nil {
		return err
	}
	return nil
}

func resourceKeyRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*etcd.Client)

	response, err := api.Get(d.Id(), false, false)
	if err != nil {
		if etcdErr, ok := err.(*etcd.EtcdError); ok &&
			etcdErr.ErrorCode == etcdErrors.EcodeKeyNotFound {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("could not read key %s: %s", d.Id(), err)
	}

	if err := d.Set("key", response.Node.Key); err != nil {
		return err
	}
	if err := d.Set("value", response.Node.Value); err != nil {
		return err
	}

	return nil
}

func resourceKeyDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*etcd.Client)

	_, err := api.Delete(d.Id(), false)
	if err != nil {
		return fmt.Errorf("could not delete key %s: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}
