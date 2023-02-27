create table if not exists `activity`
(
    id         int auto_increment
        primary key,
    creator_id int                                   null,
    created_ts timestamp   default CURRENT_TIMESTAMP not null,
    type       varchar(64) default 'INFO'            not null,
    level      varchar(16) default ''                not null,
    payload    text                                  not null
);

create table if not exists `idp`
(
    id                int auto_increment
        primary key,
    name              varchar(32)             not null,
    type              varchar(32)             not null,
    identifier_filter varchar(256) default '' not null,
    config            text                    not null
);

create table if not exists `memo`
(
    id         int auto_increment
        primary key,
    creator_id int                                   not null,
    created_ts timestamp   default CURRENT_TIMESTAMP not null,
    updated_ts timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    row_status varchar(32) default 'NORMAL'          null,
    content    text                                  not null,
    visibility varchar(32) default 'PRIVATE'         not null
);

create index idx_creator_id
    on memo (creator_id);

create table if not exists `memo_organizer`
(
    id      int auto_increment
        primary key,
    memo_id int               not null,
    user_id int               not null,
    pinned  tinyint default 0 not null,
    constraint unique_key_memo_id_user_id
        unique (memo_id, user_id)
);

create index idx_user_id
    on memo_organizer (user_id);

create table if not exists `memo_resource`
(
    memo_id     int                                 null,
    resource_id int                                 not null,
    created_ts  timestamp default CURRENT_TIMESTAMP not null,
    updated_ts  timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint unique_key_memo_id_resource_id
        unique (memo_id, resource_id)
);

create table if not exists `migration_history`
(
    version    varchar(32)                         not null
        primary key,
    created_ts timestamp default CURRENT_TIMESTAMP not null
);

create table if not exists `resource`
(
    id            int auto_increment
        primary key,
    creator_id    int                                    not null,
    created_ts    timestamp    default CURRENT_TIMESTAMP not null,
    updated_ts    timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    filename      varchar(128) default ''                not null,
    `blob`        blob                                   null,
    external_link text                                   null,
    type          varchar(16)  default ''                not null,
    size          int          default 0                 not null
);

create table if not exists `shortcut`
(
    id         int auto_increment
        primary key,
    creator_id int                                   not null,
    created_ts timestamp   default CURRENT_TIMESTAMP not null,
    updated_ts timestamp   default CURRENT_TIMESTAMP not null,
    row_status varchar(16) default 'NORMAL'          not null,
    title      varchar(32) default ''                not null,
    payload    text                                  not null
);

create table if not exists `storage`
(
    id     int auto_increment
        primary key,
    name   varchar(64) not null,
    type   varchar(16) not null,
    config text        not null
);

create table if not exists `system_setting`
(
    name        varchar(32)             not null,
    value       varchar(128)            not null,
    description varchar(256) default '' not null,
    constraint unique_key_name
        unique (name)
);

create table if not exists `tag`
(
    name       varchar(32) not null,
    creator_id int         null,
    constraint unique_key_name_creator_id
        unique (name, creator_id)
);

create index idx_creator_id
    on tag (creator_id);

create table if not exists `user`
(
    id            int auto_increment
        primary key,
    created_ts    timestamp   default CURRENT_TIMESTAMP not null,
    updated_ts    timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    row_status    varchar(16) default 'NORMAL'          not null,
    username      varchar(32)                           not null,
    role          varchar(16) default 'USER'            not null,
    email         varchar(32) default ''                not null,
    nickname      varchar(32) default ''                not null,
    password_hash varchar(128)                          not null,
    open_id       varchar(128)                          not null,
    avatar_url    text                                  not null,
    constraint unique_key_open_id
        unique (open_id),
    constraint unique_key_user_name
        unique (username)
);

create table if not exists `user_setting`
(
    user_id int          not null,
    `key`   varchar(128) not null,
    value   varchar(128) not null,
    constraint unique_key_user_id_key
        unique (user_id, `key`)
);
