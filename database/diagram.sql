Table users {
  id char(36) [pk, not null]
  name varchar(100) [not null]
  email varchar(150) [unique, not null]
  password varchar(255) [not null]
  creation_date datetime [not null]
  update_date datetime [not null]
}