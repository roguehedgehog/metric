CREATE TABLE track (
    track_uri VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    artist VARCHAR(255),
    album VARCHAR(255)
);

CREATE TABLE play (
    play_id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    track_uri VARCHAR(255),
    ts DATETIME,
    play_time SMALLINT UNSIGNED,
    skipped BOOLEAN,
    shuffle BOOLEAN,

    INDEX(track_uri),
    INDEX(ts)
);