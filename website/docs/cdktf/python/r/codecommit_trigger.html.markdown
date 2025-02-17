---
subcategory: "CodeCommit"
layout: "aws"
page_title: "AWS: aws_codecommit_trigger"
description: |-
  Provides a CodeCommit Trigger Resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_codecommit_trigger

Provides a CodeCommit Trigger Resource.

~> **NOTE:** Terraform currently can create only one trigger per repository, even if multiple aws_codecommit_trigger resources are defined. Moreover, creating triggers with Terraform will delete all other triggers in the repository (also manually-created triggers).

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.codecommit_repository import CodecommitRepository
from imports.aws.codecommit_trigger import CodecommitTrigger
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        test = CodecommitRepository(self, "test",
            repository_name="test"
        )
        aws_codecommit_trigger_test = CodecommitTrigger(self, "test_1",
            repository_name=test.repository_name,
            trigger=[CodecommitTriggerTrigger(
                destination_arn=Token.as_string(aws_sns_topic_test.arn),
                events=["all"],
                name="all"
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_codecommit_trigger_test.override_logical_id("test")
```

## Argument Reference

This resource supports the following arguments:

* `repository_name` - (Required) The name for the repository. This needs to be less than 100 characters.
* `trigger` - (Required) The name of the trigger.
    * `name` - (Required) The name of the trigger.
    * `destination_arn` - (Required) The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
    * `events` - (Required) The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
    * `custom_data` - (Optional) Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
    * `branches` - (Optional) The branches that will be included in the trigger configuration. If no branches   are specified, the trigger will apply to all branches.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `configuration_id` - System-generated unique identifier.

<!-- cache-key: cdktf-0.20.8 input-64d0507dd9a3e74264f53d83bdf059ff3d9ad3e20d558761501e6dbf7e913d3f -->