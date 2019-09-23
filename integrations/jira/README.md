# Jira

## For admin
- [Sign-in](https://www.atlassian.com/software/jira/try)
- choose a next-gen template  - scrum/kanban
    select a project name   - // site-name.atlassian.net/jira/software/projects/PROJECT_NAME/boards/1/backlog
Create API token - https://id.atlassian.com/manage/api-tokens

## For users
They will receive an email invitation, skip questions
- verify
- send invitations            - can skip
- set up (answer questions)   - can skip

## Env
```console
foo@bar:~$ env
SITE_API_URL=https://<site-name>.atlassian.net/rest/api/3
```

## Endpoints

### Issue
- [`POST /issue`](#POST%20issue)
### Project
- [`GET /project`](#GET%20projects)
- [`POST /project (join users to the project)`](#POST%20project)
### Role
- [`GET /role`](#GET%20roles)
### User
- [`GET /users`](#GET%20users)
- [`POST /user`](#user)

## REST calls
---

#### POST issue:
```console
foo@bar:~$ curl --request POST --url $SITE_API_URL/issue --header 'Content-Type: application/json' --header 'Accept: application/json' --data '{"fields":{"project":{"key":"project-name"},"summary":"something is wrong","description":{"type":"doc","version":1,"content":[{"type":"paragraph","content":[{"text":"This is the description","type":"text"}]}]},"issuetype":{"name":"Bug"},"reporter":{"id":"<account-id>"},"assignee":{"id":"<account-id>"}}}' --user <email>:<api-token>
```
```json
{
    "id": "10000",
    "key": "key-1",
    "self": "https://<site-name>.atlassian.net/rest/api/3/issue/10000"
}
```

#### GET projects:
```console
foo@bar:~$ curl --request GET --url $SITE_API_URL/project/search --header 'Accept: application/json' --user <email>:<api-token>
```
```json
{
    "self": "https://<site-name>.atlassian.net/rest/api/3/project/search?null=&maxResults=50&startAt=0",
    "maxResults": 50,
    "startAt": 0,
    "total": 1,
    "isLast": true,
    "values": [
        {
            "expand": "description,lead,issueTypes,url,projectKeys,permissions",
            "self": "https://<site-name>.atlassian.net/rest/api/3/project/10000",
            "id": "10000",
            "key": "key",
            "name": "<project-name>",
            "avatarUrls": { },
            "projectTypeKey": "software",
            "simplified": true,
            "style": "next-gen",
            "isPrivate": false,
            "properties": {},
            "entityId": "<id>",
            "uuid": "<id>"
        }
    ]
}
```

#### POST project:
```console
foo@bar:~$ curl --request POST --url $SITE_API_URL/project/<project-name>/role/10101 --header 'Content-Type: application/json' --header 'Accept: application/json' --data '{"user":["<account-id>"]}' --user <email>:<api-token>
```
```json
{
    "self": "https://<site-name>.atlassian.net/rest/api/3/project/10000/role/10101",
    "name": "Member",
    "id": 10101,
    "description": "Members are part of the team, and can add, edit, and collaborate on all work.",
    "actors": [
        {
            "id": 10011,
            "displayName": "<display-name>",
            "type": "atlassian-user-role-actor",
            "name": "<username>",
            "avatarUrl": "...",
            "actorUser": {
                "accountId": "<account-id>"
            }
        }
    ],
    "scope": {
        "type": "PROJECT",
        "project": {
            "id": "10000"
        }
    }
}
```

#### GET roles:
Get id where name="Member"
```console
foo@bar:~$ curl --request GET --url $SITE_API_URL/role --header 'Accept: application/json' --user <email>:<api-token>
```
```json
[
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/role/10100",
        "name": "Administrator",
        "id": 10100,
        "description": "Admins can do most things, like update settings and add other admins.",
        "scope": {
            "type": "PROJECT",
            "project": {
                "id": "10000"
            }
        }
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/role/10103",
        "name": "atlassian-addons-project-access",
        "id": 10103,
        "description": "A project role that represents Connect add-ons declaring a scope that requires more than read issue permissions",
        "scope": {
            "type": "PROJECT",
            "project": {
                "id": "10000"
            }
        }
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/role/10101",
        "name": "Member",
        "id": 10101,
        "description": "Members are part of the team, and can add, edit, and collaborate on all work.",
        "scope": {
            "type": "PROJECT",
            "project": {
                "id": "10000"
            }
        }
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/role/10102",
        "name": "Viewer",
        "id": 10102,
        "description": "Viewers can search through, view, and comment on your team's work, but not much else.",
        "scope": {
            "type": "PROJECT",
            "project": {
                "id": "10000"
            }
        }
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/role/10002",
        "name": "Administrators",
        "id": 10002,
        "description": "A project role that represents administrators in a project"
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/role/10003",
        "name": "atlassian-addons-project-access",
        "id": 10003,
        "description": "A project role that represents Connect add-ons declaring a scope that requires more than read issue permissions",
        "actors": [
            {
                "id": 10001,
                "displayName": "Jira Service Desk Widget",
                "type": "atlassian-user-role-actor",
                "name": "addon_com.atlassian.servicedesk.embedded",
                "actorUser": {
                    "accountId": "<account-id>"
                }
            },
            {
                "id": 10003,
                "displayName": "Statuspage for Jira",
                "type": "atlassian-user-role-actor",
                "name": "addon_stspg-jira-ops",
                "actorUser": {
                    "accountId": "<account-id>"
                }
            },
            {
                "id": 10004,
                "displayName": "Jira Cloud for Workplace",
                "type": "atlassian-user-role-actor",
                "name": "addon_jira-workplace-integration",
                "actorUser": {
                    "accountId": "<account-id>"
                }
            },
            {
                "id": 10000,
                "displayName": "Slack",
                "type": "atlassian-user-role-actor",
                "name": "addon_jira-slack-integration",
                "actorUser": {
                    "accountId": "<account-id>"
                }
            },
            {
                "id": 10002,
                "displayName": "Trello",
                "type": "atlassian-user-role-actor",
                "name": "addon_jira-trello-integration",
                "actorUser": {
                    "accountId": "<account-id>"
                }
            }
        ]
    }
]
```

#### GET users:
Get * where accountType="atlassian"
```console
foo@bar:~$ curl --request GET --url $SITE_API_URL/users/search --header 'Accept: application/json' --user <email>:<api-token>
```
```json
[
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "app",
        "avatarUrls": { },
        "displayName": "Slack",
        "active": true
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "atlassian",
        "avatarUrls": { },
        "displayName": "<display-name>",
        "active": true
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "app",
        "avatarUrls": { },
        "displayName": "Jira Service Desk Widget",
        "active": true
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "atlassian",
        "emailAddress": "<email>",
        "avatarUrls": { },
        "displayName": "<display-name>",
        "active": true,
        "locale": "en_US"
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "app",
        "avatarUrls": { },
        "displayName": "Jira Spreadsheets",
        "active": true
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "app",
        "avatarUrls": { },
        "displayName": "Jira Cloud for Workplace",
        "active": true
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "atlassian",
        "avatarUrls": { },
        "displayName": "<display-name>",
        "active": true,
        "locale": "en_US"
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "app",
        "avatarUrls": { },
        "displayName": "Trello",
        "active": true
    },
    {
        "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
        "accountId": "<account-id>",
        "accountType": "app",
        "avatarUrls": { },
        "displayName": "Statuspage for Jira",
        "active": true
    }
]
```

#### POST user:
```console
foo@bar:~$ curl --request POST --url $SITE_API_URL/user --header 'Content-Type: application/json' --header 'Accept: application/json' --data '{"password":"123", "emailAddress":"<email>", "displayName":"<display-name>", "name":"" }' --user <email>:<api-token>
```
```json
{
    "self": "https://<site-name>.atlassian.net/rest/api/3/user?accountId=<account-id>",
    "key": "<username>",
    "accountId": "<account-id>",
    "accountType": "atlassian",
    "name": "<username>",
    "emailAddress": "<email-address>",
    "avatarUrls": { },    
    "displayName": "<display-name>",    
    "active": true, 
    "timeZone": "America/Chicago",   
    "locale": "en_US", 
    "groups": {
        "size": 0,
        "items": []
    },
    "applicationRoles": {
        "size": 0,
        "items": []
    },
    "expand": "groups,applicationRoles"
}
```
