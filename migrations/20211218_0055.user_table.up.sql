create table if not exists `users`(
  `id` int not null auto_increment primary key,
  `name` varchar(255) not null,
  `age` int,
  `birthdate` date,
  `password` varchar(255)  not null
);
