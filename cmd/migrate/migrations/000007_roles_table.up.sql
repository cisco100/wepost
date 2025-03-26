CREATE TABLE IF NOT EXISTS roles(
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    level INT NOT NULL DEFAULT 1,
    description TEXT NOT NULL
);

INSERT INTO roles(id,name,level,description) VALUES('0094efbe-836e-40e3-vf34-dab75d7eadc7','user',1,'A user can create posts and comments');
INSERT INTO roles(id,name,level,description) VALUES('91bdc182-3f94-6y78-a797-23e9064a6c15','moderator',2,'A moderator can update other users posts');
INSERT INTO roles(id,name,level,description) VALUES('63c1ebd9-0455-7ea3-b87d-7bdde804b87d','admin',3,'An admin can update and delete other users posts');