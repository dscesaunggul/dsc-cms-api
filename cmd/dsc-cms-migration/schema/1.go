package schema

// First schema
const First = `
CREATE TABLE metas (
	"key" varchar(30) NOT NULL,
	value varchar(255) NOT NULL,
	CONSTRAINT metas_pk PRIMARY KEY ("key")
);
`
