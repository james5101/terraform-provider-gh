package main

import (
	"context"
	"log"

	"github.com/google/go-github/v30/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRepo() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepoCreate,
		Read:   resourceRepoRead,
		Update: resourceRepoUpdate,
		Delete: resourceRepoDelete,
		Exists: resourceRepoExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceRepoExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

func resourceRepoCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	name1 := &name
	log.Printf(name)
	ctx := context.Background()
	client := meta.(*github.Client)
	r := &github.Repository{Name: name1}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		log.Print(repo)
	}
	return nil
}

func resourceRepoUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRepoRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRepoDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
