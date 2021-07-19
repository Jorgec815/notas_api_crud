CREATE TABLE definitiva.estudiante(
	id serial NOT NULL,
	uq_codigo_estudiante numeric(10,0),
	codigo_materia numeric(5,0),
	nota_1 numeric(2,1),
	nota_2 numeric(2,1),
	nota_3 numeric(2,1),
	nota_def numeric(2,1),
	CONSTRAINT pk_estudiante PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE definitiva.estudiante IS 'Datos del estuidante con su notas respectivas en una materia';
-- ddl-end --
COMMENT ON COLUMN definitiva.estudiante.codigo_materia IS 'codigo de la materia que cursa el estudiante';
-- ddl-end --
COMMENT ON COLUMN definitiva.estudiante.nota_1 IS 'nota correspondiente al primer 35% de la materia';
-- ddl-end --
COMMENT ON COLUMN definitiva.estudiante.nota_2 IS 'nota correspondiente al segundo 35% de la materia';
-- ddl-end --
COMMENT ON COLUMN definitiva.estudiante.nota_3 IS 'nota correspondiente al 30% final de la materia';
-- ddl-end --
COMMENT ON COLUMN definitiva.estudiante.nota_def IS 'nota definitiva de la materia';
-- ddl-end --
ALTER TABLE definitiva.estudiante OWNER TO postgres;
-- ddl-end --