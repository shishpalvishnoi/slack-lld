# slack-lld

```
Slack -> workspace -> channels -> messages
Users

UseCase:
Create workspace
User registration -> roles
User can create channel in workspace
User can add another users to workspace or channel
User can send message to channel or another user


users:
Id: bigint (p_key)
Name: varchar
Email: varchar
phoneNumber: varchar
User_type

workspace:
Id: bigint (p_key)
Name: varchar
Description: varchar
Created_by: bigint (f_key) -> users table
Created_date: timestamp

user_workspace:
Id: bigint
User_id: bigint (f_key)
Workspace_id: bigint (f_key)
role_type: varchar (can be admin, member….)

channel:
Id: bigint (p_key)
workspace_id: bigint (f_key)
Name: varchar
created_by: bigint (f_key)

user_channel:
Id: bigint(p_key)
User_id: bigint (f_key)
Channel_id: bigint (f_key)
Role_type: varchar






Message:
id: bigint (p_key)
data: varchar
attachment: varchar
user_id: bigint (f_key) // posted_by
workspace_id

channel_message:
channel_id
messages: <message_ids>



Class Users {
addUser();
}


Workspace {
createWorkspace();
addUserToWorkspace();

}

Channel {
createChannel();
addUserInChannel();
}

Message {

}
```
