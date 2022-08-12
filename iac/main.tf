terraform {
  required_providers {
    everything = {
      source  = "qaware.com/terraform/everything"
      version = "0.0.1"
    }
  }

  backend "local" {
    path = "terraform.tfstate"
  }
}

provider everything {
}