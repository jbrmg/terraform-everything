terraform {
  required_providers {
    ikea = {
      source  = "qaware.com/terraform/ikea"
      version = "0.0.1"
    }
  }

  backend "local" {
    path = "terraform.tfstate"
  }
}

provider ikea {
}