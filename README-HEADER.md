# PBS TF EFS module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs/terraform-aws-efs-module?ref=x.y.z
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

This creates an EFS file system.

Integrate this module like so:

```hcl
module "efs" {
  source = "github.com/pbs/terraform-aws-efs-module?ref=x.y.z"

  # Tagging Parameters
  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo

  # Optional Parameters
  is_encrypted = true
  kms_key_id = "123-456-789"
  lifecycle_policy = "AFTER_7_DAYS"
  performance_mode = "generalPurpose"
  provisioned_throughput_in_mibps = "10000"
  subnets = ["subnet-1", "subnet-2", "subnet-3", "subnet-4"]
  security_groups = ["sg-123"]
  vpc_id = "vpc-123"
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
