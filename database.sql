-- Table Definition
CREATE TABLE "users" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "name" varchar,
    "username" varchar,
    "password_hash" varchar
);

-- Table Definition
CREATE TABLE "todo_list" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "title" varchar,
    "description" varchar
);

-- Table Definition
CREATE TABLE "todo_item" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "title" varchar,
    "description" varchar,
    "done" bool
);

-- Table Definition
CREATE TABLE "users_lists" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "user_id" INTEGER REFERENCES users (id),
    "list_id" INTEGER REFERENCES todo_list (id)
);

-- Table Definition
CREATE TABLE "lists_items" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "list_id" INTEGER REFERENCES todo_list (id),
    "item_id" INTEGER REFERENCES todo_item (id)
);