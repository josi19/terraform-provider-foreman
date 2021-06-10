
# foreman_katello_product


Poducts are mostly operating systems to which repositories are assigned.


## Example Usage

```
# Autogenerated example with required keys
data "foreman_katello_product" "example" {
  name = "Debian 10"
}
```


## Argument Reference

The following arguments are supported:

- `name` - (Required) Product name.


## Attributes Reference

The following attributes are exported:

- `description` - Product description.
- `gpg_key_id` - Identifier of the GPG key.
- `label` - 
- `name` - Product name.
- `sync_plan_id` - Plan numeric identifier.
