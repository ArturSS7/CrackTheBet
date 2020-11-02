CREATE TABLE IF NOT EXISTS events (
        id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
            event_type varchar(255),
                league varchar(255),
                    player1 varchar(255),
                        player2 varchar(255),
                            odds1 float,
                                odds2 float,
                                    status varchar(255),
                                        time integer,
                                            flashscore_id varchar(255)
                                        
    );

    CREATE TABLE IF NOT EXISTS bets (
            id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                user_id integer,
                    flashscore_id varchar(255),
                        bet varchar(255),
                            odds float,
                                amount integer,
                                    prize integer
                                    
        );
