Setting up a Go web server with PostgreSQL and serving an API for CRUD operations involves several steps. Below is a guide tailored for Arch Linux.

---

### 1. **Install Dependencies**
Make sure you have the following installed:
- **Go**: Install it using `sudo pacman -S go`.
- **PostgreSQL**: Install it using `sudo pacman -S postgresql`.

---

### 2. **Set Up PostgreSQL**
1. **Initialize the PostgreSQL Database**:
   ```bash
   sudo -iu postgres initdb --locale en_US.UTF-8 -D /var/lib/postgres/data
   ```

2. **Start and Enable the Service**:
   ```bash
   sudo systemctl start postgresql
   sudo systemctl enable postgresql
   ```

3. **Create a Database and User**:
   ```bash
   sudo -iu postgres psql
   ```
   Inside the PostgreSQL shell:
   ```sql
   CREATE DATABASE mydb;
   CREATE USER myuser WITH PASSWORD 'mypassword';
   GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
   \q
   ```

4. **Test Connection**:
   Test with `psql` or a Go driver later.

---

### 3. **Create a Go Project**
1. **Set Up Your Go Module**:
   ```bash
   mkdir Go_CRUD_server
   cd Go_CRUD_server
   go mod init Go_CRUD_server
   ```

2. **Install Required Libraries**:
   ```bash
   go get github.com/lib/pq
   go get github.com/gorilla/mux
   go get github.com/jmoiron/sqlx
   ```

---

### 4. **Set Up the Database Schema**
1. Log into the database:
   ```bash
   psql -U myuser -d mydb
   ```

2. Create the `post` table:
   ```sql
   CREATE SEQUENCE Posts_id_seq;
   
   CREATE TABLE public.posts 
    ( 
    	id integer NOT NULL DEFAULT nextval('Posts_id_seq'::regclass), 
    	title text COLLATE pg_catalog."default", 
    	body text COLLATE pg_catalog."default", 
    	CONSTRAINT "Posts_pkey" PRIMARY KEY (id)
    ) 
    WITH
    ( 
    	OIDS = FALSE
    ) 
    
    TABLESPACE pg_default;
    	
    ALTER TABLE public.posts OWNER to postgres;
    INSERT INTO posts(title, body) 
    VALUES('the first title only','and this is the body of the first post.');
    
    INSERT INTO posts(title, body) 
    VALUES('yet another title','and some more copy text for another body.');
   ```

---

### 6. **Run the Server**
1. Run your server:
   ```bash
   go run main.go
   ```

2. Test the API using tools like **Postman** or **curl**:
   ####Postman
   - `hrrp://localhost:8000/post`
   - `hrrp://localhost:8000/post/{id}`
   ####Curl
   - Create:  `curl -X POST 127.0.0.1:8000/posts -H "Content-Type: application/json" -d '{"title": "A new post headline for a new post", "body": "but also a new body text here will appear"}'`
   - Update:  `curl -X PUT 127.0.0.1:8000/posts/3 -H "Content-Type: application/json" -d '{"title": "and this is the new title for post number 1","body": "and the brand new edited body text of post number 1"}'`
   - Delete:  `curl -X DELETE 127.0.0.1:8000/posts/4 -H "Accept: application/json"`

---

### 7. **Optional Improvements**
- Use environment variables for sensitive data like credentials.
- Add validation for API inputs.
- Implement PUT for updating items.

