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
			"private": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRepoExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	name := d.Get("name").(string)
	ctx := context.Background()
	client := meta.(*github.Client)
	repo, _, err := client.Repositories.Get(ctx, "", name)
	if err != nil {
		d.SetId(name)
		log.Print(repo)
		return true, nil
	}
	d.SetId("")
	return false, nil
}

func resourceRepoCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	private := d.Get("private").(bool)
	description := d.Get("description").(string)
	ctx := context.Background()
	client := meta.(*github.Client)
	r := &github.Repository{Name: &name, Private: &private, Description: &description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		log.Print(repo)
	}
	d.SetId(name)
	return nil
}

func resourceRepoUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRepoRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	ctx := context.Background()
	client := meta.(*github.Client)
	repo, _, err := client.Repositories.Get(ctx, "", name)
	if err != nil {
		d.SetId(name)
		log.Print(repo)
		return nil
	}
	d.SetId("")
	return nil
}

func resourceRepoDelete(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	ctx := context.Background()
	client := meta.(*github.Client)
	client.Repositories.Delete(ctx, "", name)
	log.Fatal("something happened")
	return nil
}
