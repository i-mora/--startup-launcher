# Slack



## Requirements

- A legacy token can be enough, you can get it here: https://api.slack.com/custom-integrations/legacy-tokens

## Get

### Users

You can fetch the users for a given space/group and store it a json file with the command below:
```shell
make slack get=users export=true out-dir=<output-directory> token=<legacy-token>
```
