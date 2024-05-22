-- Inport Extentions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Drop existing tables
DROP TABLE IF EXISTS replacement_lessees;
DROP TABLE IF EXISTS replacement_lessors;
DROP TABLE IF EXISTS split_lessees;
DROP TABLE IF EXISTS split_lessors;
DROP TABLE IF EXISTS replacements;
DROP TABLE IF EXISTS splits;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS groups;

-- DDL for tables
CREATE TABLE groups (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(300) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    group_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name VARCHAR(300) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE splits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount INTEGER DEFAULT 0 NOT NULL,
    group_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    title VARCHAR NOT NULL,
    done BOOLEAN DEFAULT FALSE NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE replacements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount INTEGER DEFAULT 0 NOT NULL,
    group_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    title VARCHAR NOT NULL,
    done BOOLEAN DEFAULT FALSE NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- DDL for split-lessors table
CREATE TABLE split_lessors (
    split_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    amount INTEGER NOT NULL,
    PRIMARY KEY (split_id, user_id),
    FOREIGN KEY (split_id) REFERENCES splits(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- DDL for split-lessees table
CREATE TABLE split_lessees (
    split_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (split_id, user_id),
    FOREIGN KEY (split_id) REFERENCES splits(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- DDL for replacement-lessors table
CREATE TABLE replacement_lessors (
    replacement_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    amount INTEGER NOT NULL,
    PRIMARY KEY (replacement_id, user_id),
    FOREIGN KEY (replacement_id) REFERENCES replacements(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- DDL for replacement-lessees table
CREATE TABLE replacement_lessees (
    replacement_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    amount INTEGER NOT NULL,
    PRIMARY KEY (replacement_id, user_id),
    FOREIGN KEY (replacement_id) REFERENCES replacements(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Insert test data into groups table
INSERT INTO groups (name, created_at, updated_at) VALUES
('Group A', NOW(), NOW()),
('Group B', NOW(), NOW()),
('Group C', NOW(), NOW());

-- Insert test data into users table
INSERT INTO users (group_id, created_at, updated_at, name) VALUES
((SELECT id FROM groups WHERE name = 'Group A'), NOW(), NOW(), 'Alice'),
((SELECT id FROM groups WHERE name = 'Group A'), NOW(), NOW(), 'Bob'),
((SELECT id FROM groups WHERE name = 'Group B'), NOW(), NOW(), 'Charlie'),
((SELECT id FROM groups WHERE name = 'Group B'), NOW(), NOW(), 'David'),
((SELECT id FROM groups WHERE name = 'Group C'), NOW(), NOW(), 'Eve'),
((SELECT id FROM groups WHERE name = 'Group C'), NOW(), NOW(), 'Frank');

-- Insert test data into splits table
INSERT INTO splits (amount, group_id, created_at, updated_at, title, done) VALUES
(100, (SELECT id FROM groups WHERE name = 'Group A'), NOW(), NOW(), 'Split A1', FALSE),
(200, (SELECT id FROM groups WHERE name = 'Group B'), NOW(), NOW(), 'Split B1', FALSE),
(150, (SELECT id FROM groups WHERE name = 'Group C'), NOW(), NOW(), 'Split C1', TRUE),
(120, (SELECT id FROM groups WHERE name = 'Group A'), NOW(), NOW(), 'Split A2', TRUE),
(220, (SELECT id FROM groups WHERE name = 'Group B'), NOW(), NOW(), 'Split B2', FALSE);

-- Insert test data into replacements table
INSERT INTO replacements (amount, group_id, created_at, updated_at, title, done) VALUES
(300, (SELECT id FROM groups WHERE name = 'Group A'), NOW(), NOW(), 'Replacement A1', TRUE),
(400, (SELECT id FROM groups WHERE name = 'Group B'), NOW(), NOW(), 'Replacement B1', FALSE),
(350, (SELECT id FROM groups WHERE name = 'Group C'), NOW(), NOW(), 'Replacement C1', TRUE),
(310, (SELECT id FROM groups WHERE name = 'Group A'), NOW(), NOW(), 'Replacement A2', FALSE),
(420, (SELECT id FROM groups WHERE name = 'Group B'), NOW(), NOW(), 'Replacement B2', TRUE);

-- Insert test data into split_lessors table
INSERT INTO split_lessors (split_id, user_id, created_at, updated_at, amount) VALUES
((SELECT id FROM splits WHERE title = 'Split A1'), (SELECT id FROM users WHERE name = 'Alice'), NOW(), NOW(), 50),
((SELECT id FROM splits WHERE title = 'Split B1'), (SELECT id FROM users WHERE name = 'Charlie'), NOW(), NOW(), 100),
((SELECT id FROM splits WHERE title = 'Split C1'), (SELECT id FROM users WHERE name = 'Eve'), NOW(), NOW(), 75),
((SELECT id FROM splits WHERE title = 'Split A2'), (SELECT id FROM users WHERE name = 'Bob'), NOW(), NOW(), 60),
((SELECT id FROM splits WHERE title = 'Split B2'), (SELECT id FROM users WHERE name = 'David'), NOW(), NOW(), 110);

-- Insert test data into split_lessees table
INSERT INTO split_lessees (split_id, user_id, created_at, updated_at) VALUES
((SELECT id FROM splits WHERE title = 'Split A1'), (SELECT id FROM users WHERE name = 'Bob'), NOW(), NOW()),
((SELECT id FROM splits WHERE title = 'Split B1'), (SELECT id FROM users WHERE name = 'David'), NOW(), NOW()),
((SELECT id FROM splits WHERE title = 'Split C1'), (SELECT id FROM users WHERE name = 'Frank'), NOW(), NOW()),
((SELECT id FROM splits WHERE title = 'Split A2'), (SELECT id FROM users WHERE name = 'Alice'), NOW(), NOW()),
((SELECT id FROM splits WHERE title = 'Split B2'), (SELECT id FROM users WHERE name = 'Charlie'), NOW(), NOW());

-- Insert test data into replacement_lessors table
INSERT INTO replacement_lessors (replacement_id, user_id, created_at, updated_at, amount) VALUES
((SELECT id FROM replacements WHERE title = 'Replacement A1'), (SELECT id FROM users WHERE name = 'Alice'), NOW(), NOW(), 150),
((SELECT id FROM replacements WHERE title = 'Replacement B1'), (SELECT id FROM users WHERE name = 'Charlie'), NOW(), NOW(), 200),
((SELECT id FROM replacements WHERE title = 'Replacement C1'), (SELECT id FROM users WHERE name = 'Eve'), NOW(), NOW(), 175),
((SELECT id FROM replacements WHERE title = 'Replacement A2'), (SELECT id FROM users WHERE name = 'Bob'), NOW(), NOW(), 160),
((SELECT id FROM replacements WHERE title = 'Replacement B2'), (SELECT id FROM users WHERE name = 'David'), NOW(), NOW(), 210);

-- Insert test data into replacement_lessees table
INSERT INTO replacement_lessees (replacement_id, user_id, created_at, updated_at, amount) VALUES
((SELECT id FROM replacements WHERE title = 'Replacement A1'), (SELECT id FROM users WHERE name = 'Bob'), NOW(), NOW(), 150),
((SELECT id FROM replacements WHERE title = 'Replacement B1'), (SELECT id FROM users WHERE name = 'David'), NOW(), NOW(), 200),
((SELECT id FROM replacements WHERE title = 'Replacement C1'), (SELECT id FROM users WHERE name = 'Frank'), NOW(), NOW(), 175),
((SELECT id FROM replacements WHERE title = 'Replacement A2'), (SELECT id FROM users WHERE name = 'Alice'), NOW(), NOW(), 160),
((SELECT id FROM replacements WHERE title = 'Replacement B2'), (SELECT id FROM users WHERE name = 'Charlie'), NOW(), NOW(), 210);
