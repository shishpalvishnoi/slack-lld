create table users
(
    id           bigint serial primary key,
    name         varchar(255) not null,
    email        varchar(255) not null,
    phone_number varchar(20)  not null,
    user_type    varchar      not null,
    created_at   timestamp    not null,
    updated_at   timestamp,
    version_id   bigint       not null
);

create table workspace
(
    id          serial primary key,
    name        varchar(255) not null,
    description varchar(400) not null,
    created_by  bigint       not null,
    created_at  timestamp    not null,
    updated_at  timestamp    not null,
    version_id  bigint       not null
)

create table user_workspace
(
    id           serial primary key,
    user_id      bigint    not null,
    workspace_id bigint    not null,
    role_type    bigint    not null,
    created_at   timestamp not null,
)

create table channel
(
    id           serial primary key,
    workspace_id bigint    not null,
    name: varchar(255),
    created_at   timestamp not null,
    created_by   bigint    not null,
)

create table user_channel
(
    id serial primary key,
    user_id bigint not null,
    channel_id bigint not null,
    role_type varchar(255),
    created_at   timestamp not null,
)

