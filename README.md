# Testing for Terraform in Terraform

Terraform is great but there's not much in the way of ability to test what you've written.
This project aims to address that problem.

This project was inspired by some of the [internal terraform code][1] and some
[experimental code][2] by [Martin Atkins][3]. Martin's code was a good starting
point for some design considerations, but unlike the code from Martin, this uses
the official Gotlang Terraform SDK.

> **Disclaimer**
>
> This is a living project and should be used with caution. No expections should be
> made about it's stability or ability to introduce breaking changes without notice.

## But what does it do?

Writing good Terraform code inevitably means you'll be writing some custom modules
in order to abstract away some repetative or otherwise complex logic into a nice
reusable package, which you can then share with your team or the wider community.
The problem with Terraform though, is it's hard to test what you've done. Short of
running your code constantly against some test environment where you can bring up
and tear down entire projects, it's hard to see if you've introduced any breaking
changes to your modules. Of course you can version them and do some manual testing
with each new version, but it's a lot of manual work to do this.

Now there are some tools out there that attempt to address this issue, such as
[Terratest][4] by Gruntwork, which is powerful and extensive, but it requires
knowing go, how Go's test suite works as well as a reasonable bit of Terraform's
internal workings. Long sotry short it's a lot of effort and it doesn't seem right
that you should have to learn another language to test another one - that's not
the norm!

Wouldn't it be nice if you could just write your tests in Terraform along with
the actual Terraform code you're writing? Well hopefully this project will sate
your appetite for such a thing.

## Example

Suppose you have the following module:
```hcl
variable "input" {
    type = number
    description = "The number to multiply by 10"
}

output "output" {
    value = var.input * 10
}
```
It's a simple module that takes a number in, and outputs that number multiplied by 10.

Within a `test` directory of your module, you could have the following test file:
```hcl
module "times_ten" {
    source = "../"

    input = 10
}

data "testing_unit" "calculator" {
    subject = "test the calculator"

    integer_equal {
        statement = "10 * 10 should be 100"

        input = module.times_ten.output
        expect = 100
    }
}
```
When running a `terraform apply` against this `test` directory, Terraform will instantiate
your calculator module, and the `testing_unit.calculator` data source will in this
instance run an integer equality test against the response given from the module and the
expected output as specified in the `integer_equal` block.

## Installing

You can grab a binary build from the [GitHub releases][5] page, or you can build it yourself
with `go build -v`.

Place the binary inside the `~/.terraform.d/plugins` directory (create it if it doesn't exist)
then you should be good to go!

## Running the Example

Checkout this repository and within the `example` directory, run the following inside your terminal:

```shell
$ terraform init -plugin-dir=../
```

```shell
$ terraform apply
```

You should see a list of tests that have passed and failed as outputs from the example Terraform code.

[1]: https://github.com/hashicorp/terraform/tree/master/builtin/providers/test
[2]: https://github.com/apparentlymart/terraform-provider-testing
[3]: https://github.com/apparentlymart
[4]: https://github.com/gruntwork-io/terratest
[5]: https://github.com/m13t/terraform-provider-testing/releases
