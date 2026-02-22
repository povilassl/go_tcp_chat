create table channel_members (
    user_id varchar(36) not null,
    channel_id varchar(36) not null,

    primary key (user_id, channel_id),

    constraint fk_channel_members_user_id
        foreign key (user_id)
        references users(id)
        on delete cascade,

    constraint fk_channel_members_channel_id
        foreign key (channel_id)
        references channels(id)
        on delete cascade
)
