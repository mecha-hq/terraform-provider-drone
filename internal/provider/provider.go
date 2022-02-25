package provider

import (
	"context"

	"github.com/drone/drone-go/drone"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/oauth2"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"server": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("DRONE_SERVER", nil),
				},
				"token": {
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("DRONE_TOKEN", nil),
				},
			},

			ResourcesMap: map[string]*schema.Resource{
				"drone_user": resourceUser(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		server := d.Get("server").(string)
		token := d.Get("token").(string)

		httpClient := new(oauth2.Config).Client(
			oauth2.NoContext,
			&oauth2.Token{
				AccessToken: token,
			},
		)

		client := drone.NewClient(server, httpClient)

		return client, nil
	}
}
