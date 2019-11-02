# DSC CMS API

Content management system for our internal tools

# Main App

- API server at cmd/dsc-cms-api
- Migration schema for database at cmd/dsc-cms-migration

# How To Update Schema Database

1. Whenever you want to mutate your database structure, the query you must put in schema folder at cmd/dsc-cms-migration/schema (create a new one),
   why is this important ? when you update the database structure, some programmer who collaborate with you get a new database structure
2. After you put schema, running the migration with the command below :
   <pre>
   go run cmd/dsc-cms-migration/*.go
   </pre>
3. Then you can run your app WOHO !!!

# Folder structure responbility

<pre>
cmd
internal
|-- config
   |-- env (put your environment variable here)
   |-- postgres (connection db to postgres, if you have other connection db to another provider please create one inside config folder)
|-- entity
   |-- {entity name}
       |-- model (database model and response model)
       |-- repository (your query put inside here)
       |-- service (your logical put here)
|-- httpserver
    |-- container (injector for DI, if you create a new service register your service inside this code)
    |-- middleware (put your http middleware inside here)
    |-- route (versioning your http server api)
|-- pkg (list our helper function and can be reuse at another project)
</pre>

Any bug or feature please create an issue and then solve it
