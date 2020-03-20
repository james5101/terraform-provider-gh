package main

import (
	"context"

	"github.com/google/go-github/v30/github"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"golang.org/x/oauth2"
	// "golang.org/x/oauth2/github"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
				Description: "API Token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gh_create_repo": resourceRepo(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	token := d.Get("token").(string)
	c, error := auth(token)
	if error != nil {
		return nil, error
	}
	return c, nil
}

func auth(token string) (x interface{}, err error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client, error := github.NewClient(tc), err
	if error != nil {
		return nil, err
	}

	return client, nil
}
