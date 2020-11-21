CREATE TABLE IF NOT EXISTS events (
        id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
            event_type varchar(32),
                league varchar(255),
                    player1 varchar(255),
                        player2 varchar(255),
                            odds1 float,
                                odds2 float,
                                    status varchar(32),
                                        time bigint,
                                            flashscore_id varchar(32),
                                                winner int default -1
    );

CREATE TABLE IF NOT EXISTS bets (
        id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
            user_id integer,
                odds float,
                    amount integer,
                        player1 varchar(255),
                            player2 varchar(255),
                                time bigint,
                                    bet_time bigint,
                                        event_type varchar(32),
                                            bet_player int,
                                                flashscore_id varchar(32),
                                                    prize float default -1,
                                                        status varchar(32) default 'not_processed'
        );

CREATE TABLE IF NOT EXISTS users(
        id serial,
            username varchar(50) unique,
                email varchar(50) unique,
                    password varchar(100),
                        verified boolean default false,
                            balance float default 228
        );

CREATE TABLE IF NOT EXISTS verification(
        id int unique,
            token varchar(36) unique
        );

CREATE TABLE IF NOT EXISTS password_recovery(
        id int unique,
            token varchar(36) unique
        );