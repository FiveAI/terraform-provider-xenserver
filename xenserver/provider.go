package xenserver

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

// Returns the schema for the provider
// and all of its resources and datasources
func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["url"],
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["username"],
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["password"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"xenserver_pifs": dataSourceXenServerPifs(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"xenserver_vm":      resourceVM(),
			"xenserver_vdi":     resourceVDI(),
			"xenserver_network": resourceNetwork(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

// Initialised the descriptions map
func init() {
	descriptions = map[string]string{
		"url": "The URL to the XenAPI endpoint, typically \"https://<XenServer Management IP>\"",

		"username": "The username to use to authenticate to XenServer",

		"password": "The password to use to authenticate to XenServer",
	}
}

// Loads the provider's configuration
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL:      d.Get("url").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	return config.NewConnection()
}

// ignoreCaseDiffSuppressFunc is a DiffSuppressFunc from helper/schema that is
// used to ignore any case-changes in a return value.
func ignoreCaseDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	return strings.ToLower(old) == strings.ToLower(new)
}