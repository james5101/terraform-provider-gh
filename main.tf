provider "gh" {
    token = ""
}


resource "gh_create_repo" "repo1" {
    name = "terraform-provider-gh9"
    private = false
    description = "new repo from custom tf provider"
}