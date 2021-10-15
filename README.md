## Quiz

## Setup go environment and clone the project locally
1. Install golang - https://golang.org/doc/install
2. Git clone the repo into the src folder in the $GOPATH location (which is /Users/<username>/go) 
    i.e,. 
        a. $ cd /Users/<username>/go/src
        b. $ git clone git@github.com:ShreyasChikkabilathiHarisha/quiz.git
3. $ cd quiz
4. $ go get github.com/go-sql-driver/mysql

## Database: First run instructions:
Start mysql server and create the tables:
$brew services start mysql
$mysql -uroot

Now execute the commands in quiz.sql

Feel free to keep that terminal open to check the values in the DB and start the application in another terminal.

## Database: Subsequent runs:
The database is already set up, so we just need to start the mysql service. So run the following:
$brew services start mysql

## Running the service
In another terminal, from the project root directory, run this to start the application
$go run main.go

The service is interactive and in this current version, it uses command line for the quiz interaction.
With more time, web interface can be enabled to improve user experience.

## Database: Stopping the mysql server
Use the below command to stop the local mysql server:
$brew services stop mysql

## Features
1. Take the quiz
2. Add questions to the question bank/database
3. Check the previous answers stored in the database
4. Start the localhost with a static page - My goal was to complete the web user interface part of the application after setting up the functioning command line application. Now that I have that, I will iterate on this current version to enable Web UI interaction.