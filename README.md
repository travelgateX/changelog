# ChangeLog API Graphql
"ChangeLog API Graphql" is an API interface implemented in Graphql that allow to save and categorize changes to available resources centrally.

One of the objectives of this project is to have a basic schema using `graphql + gorm` that serves as a guide for future projects.

We have been inspired by projects like [github.com/graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go), but we have not found a consensus majority on a pattern for projects that implement a Graphql API in golang projects. Also we are inspired in the ChangeLog manifest good practiques [keepachangelog.com](https://keepachangelog.com/). 

Any suggestions on how to improve the structure of the project, lines of code that are left over or missing, possible improvements, will be welcome.    

>Don’t let your friends dump git logs into changelogs

## Config file
The application load the configuration file in `config/config.toml`. This file is located outside the git repository, instead there is a file named `config.toml.example`.
Make a copy of this file in and rename it to `config.toml` with the configuration values of your environment.

```toml
[db]
host = "localhost"      # ENV["CHL_DBHOST"]
port = "5432"           # ENV["CHL_DBPORT"]
user = "default"        # ENV["CHL_DBUSER"]
password = "default"    # ENV["CHL_DBPASS"]
dbname = "changelog-db" # ENV["CHL_DBNAME"]
```

Configuration it's prepared to load the values from the environmet variables. If the environment variables are set, the application will take these values.

## Generating Graphql schema
The graphql schema is segmented by type for a better understanding. We use go generate command tool to join and generate the final schema file, the go-bindata package is required for this action. You can see more there [go-bindata](https://github.com/jteeuwen/go-bindata). After execute this command you will see `bindata.go` file in your schema folder.

```bash
$ go generate ./schema/
```

## Load package dependecy
Golang dependency management tool [dep](https://github.com/golang/dep) it's required for this project. Use the following command to fetch all packages that project needs inside `vendor` folder.

```bash
$ dep ensure
```
Testing packages need to be imported manually
```go
$ go get github.com/onsi/ginkgo github.com/onsi/gomega
```
## Database migrations
The data persistence used for this project is Postgresql DB. We use [gorm.io](http://gorm.io) to map data from the database, we also use [gormigrate](https://github.com/go-gormigrate/gormigrate) to maintain different versions of the database.

The full migrations functionality is not yet developed but now there are two operations available. 

```bash
# Migrate all migrations files
$ go run context/migration/*.go migrate

# Rollback last migration
$ go run context/migration/*.go rollback
```

To add a new migration follow the steps below:
1. Creates a new single file inside `context/migrations` folder and names it with format: `yyyymmddhhmm_MigrateShortDescription.go`
2. Follow the structure of existing migrations like 201806051541_CreateCommitsTable and generates new one. Remember to pay special attention to the generation of the format number, it must be sequential with the previous one.
3. Name the main function of your migration like: `myyyymmddhhmmMigrateShortDescription()` (add `m` before and remove underscore)
4. Tries to do atomic migrations, which can add a field/table but have an equivalent rollback function to go "backwards".
5. Call this funcion in `context/migration/migration.go` and add this one to allMigrations array.


## Launch the server
Once the above actions have been executed, you can run the server using the following:

```bash
$ go run server.go
```
And you can visit localhost with graphiql playground in root route. [http://localhost:3000/](http://localhost:3000/)

# Testing
We use the Ginkgo framework to run unity tests as well as integration tests. Ginkgo is an easy, human readable way of testing your code. We can autogenerate our testing files:
## Boostrapping a suite
We must first create a test suite. This file could be considered the "main.go" of the testing suite. To generate it just move to the directory where the package is and run the command. This will generarte a `packagename_suite_test.go` file.

```bash
$ cd path/to/package
$ ginkgo bootstrap
```

## Adding Specs to a Suite
You can start adding tests directly into the suite file, but you can also separate them into different files (especially for packages with multiple files). This will generarte a `filename_test.go` file.

```bash
$ ginkgo generate filename
```


>It is recommended that when the files are created you move them to a `/test` folder so the testing files don't pile up.

## Understanding the code

At first may be weir to look at a test file code, but it's very intuitive. At the suite file we have de the testing "main" and the data base initialization.

```go
func TestResolver(t *testing.T) {
	RegisterFailHandler(Fail)
    RunSpecs(t, "Resolver Suite")

    var _ = BeforeSuite(func() {
	c, err1 := config.LoadConfig("./../config")
	Expect(err1).NotTo(HaveOccurred())

	db, err2 := context.OpenDB(c)
	Expect(err2).NotTo(HaveOccurred())

	err3 := db.DB().Ping()
	Expect(err3).NotTo(HaveOccurred())

    })
}
```
Let's break this down:
- TestResovler() function will be called by the go test or ginkgo commands
- RegisterFailHandler(Fail) line is the failure signal in ginkgo. The test will stop if a failer signal is recived.
- RunSpecs() starts the test suite, it will automatically fail testing.T if any of the specs fail.
- The variable BeforeSuite() will run before all the tests, so it is very usefull to initialize global variables, like the data base connection.
 
 Now to write the test itself we follow this syntax:
 ```go
 var _ = Describe("Resolver", func() {

	BeforeEach(func() {
        ...
	})
	JustBeforeEach(func() {
        ...
	})

	// Function test
	Describe("Describe what the spec is going to do", func() {
		Context("In what state are we going to test our code", func() {
		   It("What result we spect", func() {
                ...
                Expect(result).To(Equal(spected))
		   })
        })
    })
 })  
 ```

 Just as before:
 - All the specs are declared inside the `var _ = Describe()`
 - `BeforeEach(func(){})` the code inside this block will excute before every `Describe` block.
 - Although our code can be inside all the blocks, the `Expect()` function is commonly only present inside the `It()` block.

 You can lear more about how to structure specs at http://onsi.github.io/ginkgo/#structuring-your-specs
 
 ## Running the test
To run your test you just need to be in the directory where the suite file is located and execute ginkgo.
```bash
$ cd /path/to/suitefile
$ ginkgo
```

## Remember
- If you change the schema.graphql you need to generate again bindata.go.
- If you receive an error during compilation with the message `undefined: AssetNames` or `undefined: MustAsset` you should to generate properly the bindata.go.

## Links
[http://gorm.io/](http://gorm.io/)

[https://keepachangelog.com](https://keepachangelog.com/)

[https://github.com/graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)

[https://github.com/graphql/graphiql](https://github.com/graphql/graphiql)

[https://github.com/stretchr/testify](https://github.com/stretchr/testify)



# TESTING FUNCTIONS (unfinished)

## Testing results
There are mainly two ways to see the output of a test:  
```
t.Error("The test has failed") 
```
and

```
t.Fatal("The test has failed")
```
The main difference between these two is that t.Fatal will stop the execution of the text. See this [link](https://blog.codeship.com/testing-in-go/) for an extended explanation of testing t. use.