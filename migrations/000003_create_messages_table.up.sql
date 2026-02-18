create table messages (
    id varchar(36) not null primary key,
    channel_id varchar(36) not null,
    user_id varchar(36) not null,
    content varchar(512) not null unique,
    created_at timestamp not null default current_timestamp,

    constraint fk_messages_channel_id
        foreign key (channel_id) 
        references channels(id)
        on delete cascade,

    constraint fk_messages_user_id
        foreign key (user_id) 
        references users(id)
        on delete cascade
)