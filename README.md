# dbq-blog

## About Project
A quick project hosting all the codes I Wrote while writting the Blog about Dbq.

**To Run the Project as is**:
- Setup up a Database (call it whatever you want)
- Create a table called *store* in your database
```
CREATE TABLE store ( id INT PRIMARY KEY AUTO_INCREMENT, product VARCHAR(120) NOT NULL, quantity INT(10) NOT NULL, price DECIMAL NOT NULL, available BOOL, timing DATETIME NOT NULL);
```
- copy the [config-example.json](./config-example.json) file to a **config.json** file in the same directory and fill in your db connection details. 
- You can now Run codes in the project

## Structure of the Project

### Main
[*main.go*](./main.go) file is the Entry point of the project and it runs all the functions created in different files.
Before running `go run main.go` in you terminal you would have to uncomment any of the function you don't want to execute.

### Others
The file naming of the other files are self explanatory.
#### Insert
- [**insert_withdbq.go**](./insert_withdbq.go): Contains Sql Database Insert Code snippet that use Dbq for db insert.
- [**insert_withoutdbq.go**](./insert_withoutdbq.go): Contains Sql Database Insert Code snippet that does not use Dbq for db insert.

#### Query
- [**query_withdbq.go**](./query_withdbq.go): Contains Sql Database Query Code snippets that use Dbq for db query.
- [**query_withoutdbq.go**](./query_withoutdbq.go): Contains Sql Database Query Code snippets that does not use Dbq for db query.

#### Script
For the sake of Benchmarking to compare the differences between when Query uses Dbq and does not use it,
I wrote a [DbScript](./dbScript/main.go) which populates a database with up *100000* random name data (ofcourse you can increase or decrease the number if you want to investigate further).

**To Run the script**:
- Setup up a Database (call it whatever you want)
- Create a table called *benchmark* in your database
```
CREATE TABLE benchmark(id int primary key auto_increment not null, name varchar(255) not null, timestamp timestamp not null);
```
- *cd* into the [*dbScript*](./dbScript) directory
- make sure your db connection details are already set in a *config.json* file
- run the `go run main.go` command
Just wait and your table will be populated with up to 100000 rows at 50hops of 2000 data insert per hop using **DBQ**

### Benchmark Test
In the [*dbq-blog_test*](./dbq-blog_test.go) file, I wrote about five benchmark tests

**To Run the Benchmark Test**:
- make sure you are in the root directory
- make sure you have already set up *benchmark* Table and populated it with the script as directed above.
- make sure your db connection details are already set in a *config.json* file
- Run ```go test -bench .``` (This runs all the benchmark tests).
  Go to [Go's Documentation](https://golang.org/pkg/testing/) to see other options available for the benchmark command (e.g ```go test -bench . -benchtime=10x```)
 
