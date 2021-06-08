package models

type Contact struct {
	ContactId       int
	FirstName       string
	LastName        string
	Mobile          string
	AlternateMobile string
	MailId          string
}

/*
create table contacts(
	contact_id int primary key,
	first_name text not null,
	last_name text not null,
	mobile char(10) not null,
	alternate_mobile char(10) not null,
	mail_id text not null
);

insert into contacts values
(1,'Vamsy Kiran','Aripaka','9052224753','9550204753','vamsy@gmail.com'),
(2,'Suseela','Aripaka','9152224753','9150204753','suseela@gmail.com'),
(3,'Sagar','Aripaka','3152224753','8150204753','sagar@gmail.com'),
(4,'Srinivas','Dachepalli','6752224753','8750204753','srinu@gmail.com'),
(5,'Deepa','Dachepalli','9752224753','8890204753','deepa@gmail.com');
*/
