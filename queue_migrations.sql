CREATE TABLE queue (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(30),
    location VARCHAR(30),
    question VARCHAR(30),
    googled VARCHAR(30),
    asked_student BOOLEAN,
    has_debugged BOOLEAN,
    contacted BOOLEAN,
    completed BOOLEAN,
    CONSTRAINT queue_pk PRIMARY KEY (id)
);
