CREATE TABLE queue (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(64),
    location VARCHAR(30),
    question VARCHAR(512),
    googled VARCHAR(30),
    asked_student BOOLEAN,
    has_debugged BOOLEAN,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    contacted BOOLEAN,
    completed BOOLEAN,
    CONSTRAINT queue_pk PRIMARY KEY (id)
);
