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
  -authn-url string
    	URL Conjur will be authenticating to. Environment variable equivalent 'CONJUR_AUTHN_URL'. e.g. https://conjur.com/authn-iam/global
  -aws-name string
    	AWS Resource type name. Environment variable equivalent 'CONJUR_AWS_TYPE'. e.g. ec2, lambda, ecs
  -login string
    	Conjur login that will be used. Environment variable equivalent 'CONJUR_AUTHN_LOGIN'. e.g. host/6634674884744/iam-role-name
  -secret string
    	Retrieve a specific secret from Conjur. e.g. db/postgres/username
  -silence
    	Silence debug and info messages
  -token-path string
    	Write the access token to this file. Environment variable equivalent 'CONJUR_ACCESS_TOKEN_PATH'. e.g. /path/to/access-token.json
```

### EC2
Download the binary onto your EC2 instance from [here](path/to/Releases) and execute the following command:
```
$ ./authenticator -aws-name ec2 -login host/622705945757/ubuntu-client-conjur-identity -authn-url https://conjur-master/authn-iam/global 
INFO:  2020/09/30 18:44:11.738289 main.go:41: CAIC001I Retrieving AWS IAM credential for resource type 'ec2'
INFO:  2020/09/30 18:44:11.742063 main.go:47: CAIC002I Successfully retrieved AWS IAM credential
INFO:  2020/09/30 18:44:11.742257 main.go:55: CAIC003I Successfully created the Conjur authentication request
INFO:  2020/09/30 18:44:11.742364 main.go:58: CAIC004I Attempting to authenticate to conjur: authnUrl=https://conjur-master/authn-iam/global, login=host/team1/622705945757/ubuntu-client-conjur-identity, account=conjur
INFO:  2020/09/30 18:44:11.742742 authenticate.go:29: CAIC005I Attempting to authenticate to conjur 'https://conjur-master/authn-iam/global/conjur/host%2Fteam1%2F622705945757%2Fubuntu-client-conjur-identity/authenticate'
INFO:  2020/09/30 18:44:11.786632 main.go:64: CAIC006I Successfully authenticated to conjur with login 'host/team1/622705945757/ubuntu-client-conjur-identity'
{"protected":"eyJhbGciOiJjb25qdXIub3JnL3Nsb3NpbG8vdjIiLCJraWQiOiIzYmQ0ZTNkZmE3NmRhMzhkMjVlM2VjNjZlZDkyODcwNCJ9","payload":"eyJzdWIiOiJob3N0L3RlYW0xLzYyMjcwNTk0NTc1Ny91YnVudHUtY2xpZW50LWNvbmp1ci1pZGVudGl0eSIsImlhdCI6MTYwMTQ5MTQ1MX0=","signature":"QnrpYl32ddwh8EEuPR9EDQiI2VD8szKamjjWiholWkJFb9RmeZhDVnlGBpMZWeFNn-4G3jl4a2HSPwI2DROQjEJ2oaSpA49HgY5jGyt58vipZ47Dxmi8ECYXq4Js_NLaylwBDbx6lKFqizrF2-rDKoLDFZMBbbpfk6OzuPs0vvnrDukkGw3_eA3xpi6d2v_F_BtcXmrSlr5PSjnornL5aqcsY1rMTcuT7E4Yja8uZUP_hZEXHozI2KFmCbHUCS3EnR9-XifiAfdKPczmxSILWYOCArGaor7ZKBpxndjn1qA5L7358M8I-Y2TwCF-MRz2bS6j-0YXeLdO4v6rMeEyccGYu5Nk4YMYKz2L-FShpbfjtQzfEoSc3hpsnMDymb5L"}
```

To reduce this noise in the terminal use the `-silence` flag. This will return the Conjur access token.
```
$ ./authenticator -silence -aws-name ec2 -login host/622705945757/ubuntu-client-conjur-identity -authn-url https://conjur-master/authn-iam/global 
{"protected":"eyJhbGciOiJjb25qdXIub3JnL3Nsb3NpbG8vdjIiLCJraWQiOiIzYmQ0ZTNkZmE3NmRhMzhkMjVlM2VjNjZlZDkyODcwNCJ9","payload":"eyJzdWIiOiJob3N0L3RlYW0xLzYyMjcwNTk0NTc1Ny91YnVudHUtY2xpZW50LWNvbmp1ci1pZGVudGl0eSIsImlhdCI6MTYwMTQ5MTU5N30=","signature":"qAefwgQ_VVranWWoK39ZelAoFPuMfWa1zdaJPbk5ff31l5kKfTIIhG97ikVkp3MaK227ikL-2KSv42S2aCgn2BkklNot_f1Jn8CUPavKf9hP2vuhuL55TRFp_dtVgyScYx6n9nXftjhXHbYHHRugqQpT5cbeO_PVrb_Q9UKWWA0erY_-JSpjw25EhOwcMTgg3tgqrsrjztSHOSaKMY0_E2TOKWCGN02_xWGxJPBNp4qOwE0LwmJBkrEqQgJN13GKcrXGwRKSRvhtUiiEkx58-aCzye8dMbCY9cuD4fxpztFfFKpBw8n7tKGmQd6KipkdxHzSE1jgF9-mYuOLos85wekl8A-w28pgCHbwStBkjdFT1QseJ7ywWuNRAYHgCUa7"}
```

To write the Conjur access token to a file we can use the `-token-path` flag.
```
$ ./authenticator -silence -token-path ~/.conjur.json -aws-name ec2 -login host/team1/622705945757/ubuntu-client-conjur-identity -authn-url https://conjur-master/authn-iam/global
$ cat ~/.conjur.json
{"protected":"eyJhbGciOiJjb25qdXIub3JnL3Nsb3NpbG8vdjIiLCJraWQiOiIzYmQ0ZTNkZmE3NmRhMzhkMjVlM2VjNjZlZDkyODcwNCJ9","payload":"eyJzdWIiOiJob3N0L3RlYW0xLzYyMjcwNTk0NTc1Ny91YnVudHUtY2xpZW50LWNvbmp1ci1pZGVudGl0eSIsImlhdCI6MTYwMTQ5MTczNH0=","signature":"bhnDgM0u1aa1m85WXMGIIsrJlbDqMwpefNIwY9d5-ij5ZO0LrGlMds8X-660_qC8nwwxnV_mGCxvGXYPWhoqG5dNmB2KEp2-cE6IQv6HI0HcNuXh_yDFuggHmkWJwGagB5YiqhnIRvAvyG-IPsEacIENMmR97vzpj-b5K0Yp1IXvyDP0cRrCTYNwzLciJV57iHNHp1VMqYgpbxvfxyxiGRC1MFRkVhSbDxucl5jabUs7L3ZptQk1uaCN3w-TKEAd0C36EDJcAHl4tfr-TXhDP-hxDN8FFvjNscUmYYwy26rXJy40nSHDeEjVE35gipiJQFOZW1bdVrUt2yyS2LcOKBBOKNW3nadEZ8VNRWLgZbVhntL5WOjEN8MdTfVY9xZq"}
```

To retrieve a secret from Conjur use the `-secret` flag.
```
$ ./authenticator -silence -secret path/to/secret -aws-name ec2 -login host/team1/622705945757/ubuntu-client-conjur-identity -authn-url https://conjur-master/authn-iam/global
secretValueHere123!@#
```


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
