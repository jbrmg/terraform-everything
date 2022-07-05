go build -o terraform-provider-everything_v0.0.1
mkdir -p /Users/j.bender/.terraform.d/plugins/qaware.com/terraform/everything/0.0.1/darwin_arm64/
cp -rf terraform-provider-everything_v0.0.1 /Users/j.bender/.terraform.d/plugins/qaware.com/terraform/everything/0.0.1/darwin_arm64/terraform-provider-everything_v0.0.1
rm "/Users/j.bender/Documents/qaware/talks/2022-07 EC22 - Terraform everything/codebase/iac/.terraform.lock.hcl"