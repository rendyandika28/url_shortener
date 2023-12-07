create table url_shorteners(
   id varchar(100) not null,
   slug varchar(100) character set utf8mb4 collate utf8mb4_bin not null unique ,
   url varchar(2803) not null,
   last_visited timestamp not null default current_timestamp,
   created_at timestamp not null default current_timestamp,
   updated_at timestamp not null default current_timestamp on update current_timestamp,
   deleted_at timestamp null,
   primary key (id)
)engine=innodb;