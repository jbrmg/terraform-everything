go build -o terraform-provider-ikea_v0.0.1
mkdir -p /Users/j.bender/.terraform.d/plugins/qaware.com/terraform/ikea/0.0.1/darwin_arm64/
cp -rf terraform-provider-ikea_v0.0.1 /Users/j.bender/.terraform.d/plugins/qaware.com/terraform/ikea/0.0.1/darwin_arm64/terraform-provider-ikea_v0.0.1
cd ../iac
echo "" > ./.terraform.lock.hcl
terraform init