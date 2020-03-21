provider "gh" {
    token = ""
}


resource "gh_create_repo" "repo1" {
    name = "terraform-provider-gh1"
    private = false
    description = "new repo from custom tf provider"
}

resource "gh_create_repo" "repo2" {
    name = "terraform-provider-gh2"
    private = false
    description = "new repo from custom tf provider"
}