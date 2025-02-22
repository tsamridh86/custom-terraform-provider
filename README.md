# custom-terraform-provider
A custom terraform provider base example, this doesn't do anything, you'll probably require
some more APIs to get the current state and then update the state file along with it.


## Set up and how to use


### section 1 : setting up terraform to know that you've made a custom resource for your custom cloud.

1. Copy out `.terraformrc` from `/terraformrc` folder and move it to `~/.terraformrc` in unix based system, if you are using it in windows, then you should place it as `terraform.rc` in `<youruser>/AppData/Roaming/terraform.rc`. Keep in mind that windows is dramatic about extensions.
2. Update the `.terraformrc` file to where you intend to keep your own custom provider. This will be used in section 2. 


### section 2 : creating your own module.

1. `terraform-provider-<name>` : this should be the base directory, i've left it as `tuladhar` for now, rename it to the company.
2. it is smarter to use golang's features to create this folder, fire the commands in sequence:
3. `go mod init github.com/tsamridh86/terraform-provider-tuladhar`
4. `go get github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema`
5. `go get github.com/hashicorp/terraform-plugin-sdk/v2/plugin`
6. actually build your package: `go build -o terraform-provider-tuladhar`
7. you should see a file called : `terraform-provider-tuladhar` must have been created.
8. rename the file : `terraform-provider-tuladhar_v0.0.1_x5` add `.exe` at the end if you are in windows.
9. In the location that you have edited in Section 1.2, you need to add the following folders: `<base location>/tuladhars.providers.org/tsamridh86/tuladhar/0.0.1/windows_amd64`
10. Copy the built exe into this location.

### section 3 : actually operating it.

1. the `version` in `provider.tf` should match as `0.0.1` as that's the version that we have made.
2. `main.tf` is simple, doesn't do much.


## Pending work
1. Improve the documentation so that this can be access via gcs bucket or even `registry.terraform.io`
2. Test with remote state ( should work fine )
3. Actually call APIs and implement something lol