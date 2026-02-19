create table channels (
    id varchar(36) not null primary key,
    channel_name varchar(255) not null unique,
    created_by_id varchar(36) not null,
    created_at timestamp not null default current_timestamp,

    constraint fk_channels_created_by_id
        foreign key (created_by_id) 
        references users(id)
        on delete cascade
)

/*TODO lenghts not aligning with validation*/