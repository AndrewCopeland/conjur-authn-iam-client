# conjur-authn-iam-client
Authenticate AWS resources (EC2, Lambda, ECS) with Conjur and easily retrieve secrets.

[![conjur-authn-iam-client CI](https://github.com/AndrewCopeland/conjur-authn-iam-client/workflows/conjur-authn-iam-client%20CI/badge.svg)](https://github.com/AndrewCopeland/conjur-authn-iam-client/actions?query=workflow%3A%22conjur-authn-iam-client+CI%22)

## Certification level
![](https://img.shields.io/badge/Certification%20Level-Community-28A745?link=https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md)

This repo is a **Community** level project. It's a community contributed project that **is not reviewed or supported
by CyberArk**. For more detailed information on our certification levels, see [our community guidelines](https://github.com/cyberark/community/blob/master/Conjur/conventions/certification-levels.md#community).

## Requirements

- +v5 of Conjur

## Usage instructions
- Write access token to a file. Use the `-token-path` flag.
- Retrieve a secret from Conjur. Use the `-secret` flag.
- Write Conjur access token to stdout that can be used in a bash script. Do not set the `-secret` or `-token-path` flags.

```
$ ./authenticator -h
Usage of ./authenticator:
  -account string
    	The account Conjur has been configued with
  -authn-url string
    	URL Conjur will be authenticating to. Environment variable equivalent 'CONJUR_AUTHN_URL'. e.g. https://conjur.com/authn-iam/global
  -aws-name string
    	AWS Resource type name. Environment variable equivalent 'CONJUR_AWS_TYPE'. e.g. ec2, lambda, ecs
  -ignore-ssl-verify
    	WARNING: Do not verify the SSL certificate provided by Conjur server. THIS SHOULD ONLY BE USED FOR POC
  -login string
    	Conjur login that will be used. Environment variable equivalent 'CONJUR_AUTHN_LOGIN'. e.g. host/6634674884744/iam-role-name
  -secret string
    	Retrieve a specific secret from Conjur. e.g. db/postgres/username
  -silence
    	Silence debug and info messages
  -token-path string
    	Write the access token to this file. Environment variable equivalent 'CONJUR_ACCESS_TOKEN_PATH'. e.g. /path/to/access-token.json
  -url string
    	The URL to the Conjur instance. e.g. https://conjur.com
```

### EC2
More information about [authenticating EC2 instances](docs/ec2/README.md)

### Lambda
More information about [authenticating Lambda functions](docs/lambda/README.md)

### ECS/Fargate
More information about [authenticating ECS/Fargate containers](docs/ecs/README.md)

## Contributing

We welcome contributions of all kinds to this repository. For instructions on how to get started and descriptions
of our development workflows, please see our [contributing guide](CONTRIBUTING.md).

## License

Copyright (c) 2020 CyberArk Software Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

For the full license text see [`LICENSE`](LICENSE).
