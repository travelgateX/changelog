# ChangeLog API Graphql
"ChangeLog API Graphql" is an API interface implemented in Graphql that allow to save and categorize changes to available resources centrally.

## Config file
TODO: complete info

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
## Remember
- If you change the schema.graphql you need to generate again bindata.go.

## Citas
*Don’t let your friends dump git logs into changelogs*