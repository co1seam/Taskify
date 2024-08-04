CREATE TABLE users
(
    id serial primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    username varchar(255) not null unique,
    password_hash text not null,
    email varchar(255) not null unique,
    telegram_id bigint not null
);

CREATE TABLE birthdays
(
    id serial primary key,
    user_id int not null,
    foreign key (user_id) references users(id) on delete cascade,
    is_public boolean not null default false,
    birthdate date
);

CREATE TABLE subscriptions
(
    id serial primary key,
    subscriber_id int not null,
    subscribed_to_id int not null,
    foreign key (subscriber_id) references users(id) on delete cascade,
    foreign key (subscribed_to_id) references users(id) on delete cascade
);