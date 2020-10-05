# Authenticating ECS/Fargate with Side Car Container
To Setup the Conjur `authn-iam` authenticator with ECS or Fargate follow the following instructions.

1. Retrieve the latest authenticator image from [docker hub](https://hub.docker.com/r/andrewcopeland/authenticator)
2. When defining your `Task Definition` use the image in the above link.
3. Define a Volume, this is where the Conjur access token can be written to and shared with other containers that require secrets from Conjur.
4. Define the following environment variables in your `Task Definition` (These environment variables can be overridden however these values typicallu tend to be static)
  - CONJUR_ACCESS_TOKEN_PATH: <path to volume mount>/access.json
  - CONJUR_ACCOUNT: company
  - CONJUR_APPLIANCE_URL: https://conjur.company.local
  - CONJUR_AUTHN_URL: https://conjur.company.local/authn-iam/prod
  - CONJUR_AWS_TYPE: ecs
5. Then `Run a New Task` and select the `Task Definition` we defined in the previous steps.
6. Select `Advanced Option` -> `Add Environment Variables` and add environment variable called `CONJUR_AUTHN_LOGIN` this is the ID of the application. e.g. `host/7363673737363/iam-role-name`
9. Share the volume mount with the containers that require a Conjur Access Token to retrieve secrets.


## Task Definition JSON
Here is a sample JSON that represents my Task Definition. You will see I defined a volume called `test` and I mount this volume to `/tmp/conjur`. I then write the Conjur access token to `/tmp/conjur/access.json`. This access.json file can be used by other containers to retrieve secrets from conjur.
```json
{
    "ipcMode": null,
    "executionRoleArn": "arn:aws:iam::[redacted]:role/ecsTaskExecutionRole",
    "containerDefinitions": [
        {
            "dnsSearchDomains": null,
            "environmentFiles": null,
            "logConfiguration": {
                "logDriver": "awslogs",
                "secretOptions": null,
                "options": {
                    "awslogs-group": "/ecs/conjur-authn-iam-client",
                    "awslogs-region": "us-east-1",
                    "awslogs-stream-prefix": "ecs"
                }
            },
            "entryPoint": null,
            "portMappings": [],
            "command": null,
            "linuxParameters": null,
            "cpu": 0,
            "environment": [
                {
                    "name": "CONJUR_ACCOUNT",
                    "value": "v1"
                },
                {
                    "name": "CONJUR_APPLIANCE_URL",
                    "value": "https://ec2-54-236-56-209.compute-1.amazonaws.com"
                },
                {
                    "name": "CONJUR_AWS_TYPE",
                    "value": "ecs"
                },
                {
                    "name": "CONJUR_AUTHN_URL",
                    "value": "https://ec2-54-236-56-209.compute-1.amazonaws.com/authn-iam/global"
                },
                {
                    "name": "CONJUR_ACCESS_TOKEN_PATH",
                    "value": "/tmp/conjur/access.json"
                }
            ],
            "resourceRequirements": null,
            "ulimits": null,
            "dnsServers": null,
            "mountPoints": [
                {
                    "readOnly": null,
                    "containerPath": "/tmp/conjur",
                    "sourceVolume": "test"
                }
            ],
            "workingDirectory": null,
            "secrets": null,
            "dockerSecurityOptions": null,
            "memory": null,
            "memoryReservation": null,
            "volumesFrom": [],
            "stopTimeout": null,
            "image": "andrewcopeland/authenticator:dev-20201005_150903",
            "startTimeout": null,
            "firelensConfiguration": null,
            "dependsOn": null,
            "disableNetworking": null,
            "interactive": null,
            "healthCheck": null,
            "essential": true,
            "links": null,
            "hostname": null,
            "extraHosts": null,
            "pseudoTerminal": null,
            "user": null,
            "readonlyRootFilesystem": null,
            "dockerLabels": null,
            "systemControls": null,
            "privileged": null,
            "name": "conjur-authn-iam-client"
        }
    ],
    "memory": "512",
    "taskRoleArn": "arn:aws:iam::[redacted]:role/ecsTaskExecutionRole",
    "family": "conjur-authn-iam-client",
    "pidMode": null,
    "requiresCompatibilities": [
        "FARGATE"
    ],
    "networkMode": "awsvpc",
    "cpu": "256",
    "inferenceAccelerators": [],
    "proxyConfiguration": null,
    "volumes": [
        {
            "efsVolumeConfiguration": null,
            "name": "test",
            "host": {
                "sourcePath": null
            },
            "dockerVolumeConfiguration": null
        }
    ],
    "tags": []
}
```