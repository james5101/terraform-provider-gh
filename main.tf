provider "gh" {
    token = "a27fe4c27f5404fe8fdf4215e4b435845e8aa1ed"
}

resource "gh_create_repo" "my_server" {
    name = "terraform-provider-gh"
}