

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    key TEXT UNIQUE NOT NULL,
    -- user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    active BOOLEAN DEFAULT TRUE,
    permissions TEXT[] -- opcional, por si deseas permisos por clave
);


CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    key TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    active BOOLEAN DEFAULT TRUE,
    permissions TEXT[]
);



CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    descripcion TEXT,
    precio NUMERIC(10,2),
    stock INTEGER,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
