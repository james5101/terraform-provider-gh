provider "gh" {
    token = ""
}

resource "gh_create_repo" "my_server" {
    name = "terraform-provider-gh"
}