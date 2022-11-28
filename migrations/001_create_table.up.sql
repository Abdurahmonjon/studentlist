create table if not exists students (
    id uuid primary key not null ,
    first_name varchar(25) not null ,
    last_name varchar(25) not null,
    user_name varchar(25) unique not null
);