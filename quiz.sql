create database quiz;

use quiz;

create table questions(
   id INT NOT NULL AUTO_INCREMENT,
   question VARCHAR(200) NOT NULL,
   answer VARCHAR(200) NOT NULL,
   latestanswer VARCHAR(200) NOT NULL,
   PRIMARY KEY ( id )
);