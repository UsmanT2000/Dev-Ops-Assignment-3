CREATE TABLE user_info_db (
                             id             SERIAL      PRIMARY KEY,
                             created_at     TIMESTAMP   WITH TIME ZONE NOT NULL,
                             updated_at     TIMESTAMP   WITH TIME ZONE NOT NULL,
                             deleted_at     TIMESTAMP   WITH TIME ZONE,
                             name           VARCHAR(255),
                             age            SMALLINT,
                             email          VARCHAR(255)
);
