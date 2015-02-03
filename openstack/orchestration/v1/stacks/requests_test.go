package stacks

import (
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func TestCreateStack(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t, CreateOutput)

	createOpts := CreateOpts{
		Name:    "stackcreated",
		Timeout: 60,
		Template: `
    {
      "stack_name": "postman_stack",
      "template": {
        "heat_template_version": "2013-05-23",
        "description": "Simple template to test heat commands",
        "parameters": {
          "flavor": {
            "default": "m1.tiny",
            "type": "string"
          }
        },
        "resources": {
          "hello_world": {
            "type":"OS::Nova::Server",
            "properties": {
              "key_name": "heat_key",
              "flavor": {
                "get_param": "flavor"
              },
              "image": "ad091b52-742f-469e-8f3c-fd81cadf0743",
              "user_data": "#!/bin/bash -xv\necho \"hello world\" &gt; /root/hello-world.txt\n"
            }
          }
        }
      }
    }`,
		DisableRollback: Disable,
	}
	actual, err := Create(fake.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	expected := CreateExpected
	th.AssertDeepEquals(t, expected, actual)
}
