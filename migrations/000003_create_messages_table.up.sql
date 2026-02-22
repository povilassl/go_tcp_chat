create table messages (
    id varchar(36) not null primary key,
    user_from_id varchar(36) not null,
    user_to_id varchar(36),
    channel_id varchar(36),
    content varchar(512) not null,
    created_at timestamp not null default current_timestamp,

    constraint fk_messages_channel_id
        foreign key (channel_id) 
        references channels(id)
        on delete cascade,

    constraint fk_messages_user_from_id
        foreign key (user_from_id) 
        references users(id)
        on delete cascade,

    constraint fk_messages_user_to_id
        foreign key (user_to_id)
        references users(id)
        on delete set null
)