
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