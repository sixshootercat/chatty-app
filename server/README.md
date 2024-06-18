## Todos

1. Postgres image from Docker Hub to create a database
2. Database setup
   - make a connection to the database
   - add a db migration file to create users table
3. `/signup` endpoint route to create a new user
   - handler -> service -> repository -> database
4. `/login` and `/logout` endpoint routes to authenticate and logout users
   - jwt with http-only cookie
