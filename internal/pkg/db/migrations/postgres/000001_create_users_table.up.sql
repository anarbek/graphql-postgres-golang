CREATE TABLE public."Users"
(
    "ID" serial,
    "Username" character varying(255),
    "Password" character varying(255),
    CONSTRAINT "Primary" PRIMARY KEY ("ID"),
    CONSTRAINT "UsernameUnique" UNIQUE ("Username")
);

ALTER TABLE public."Users"
    OWNER to postgres;