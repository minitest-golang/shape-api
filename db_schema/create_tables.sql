CREATE  TABLE user_tbl (
	username             varchar(64)  NOT NULL  ,
	login_key            varchar(256)  NOT NULL  ,
	CONSTRAINT pk_user_tbl_username PRIMARY KEY ( username )
 );

CREATE  TABLE shape_tbl (
	shape_id             bigint  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	username             varchar(64)  NOT NULL  ,
	shape                varchar(32)  NOT NULL  ,
	edges                text[]  NOT NULL  ,
	CONSTRAINT pk_shape_tbl PRIMARY KEY ( shape_id ),
	CONSTRAINT fk_shape_tbl_user_tbl FOREIGN KEY ( username ) REFERENCES user_tbl( username ) ON DELETE CASCADE
 );