CREATE TABLE public.links (
    id uuid NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	deleted_at timestamptz NULL,
	originLink varchar NOT NULL,
	resultLink varchar NOT NULL,
	link_at timestamptz NOT NULL,
	rank integer,
	CONSTRAINT links_pk PRIMARY KEY (id)
);
