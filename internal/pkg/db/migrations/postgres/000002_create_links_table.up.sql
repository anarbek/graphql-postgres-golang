CREATE TABLE public."Links"
(
    "ID" serial,
    "Title" character varying(255),
    "Address" character varying(255),
    "UserID" integer,
    CONSTRAINT "PrimKeyLinks" PRIMARY KEY ("ID"),
    FOREIGN KEY ("UserID")
        REFERENCES public."Users" ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE public."Links"
    OWNER to postgres;