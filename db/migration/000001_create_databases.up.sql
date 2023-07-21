CREATE TABLE fakultas (
    ID SERIAL PRIMARY KEY,
    Nama  VARCHAR(50),
    CONSTRAINT UC_NamaFakultas UNIQUE (Nama)
);

CREATE TABLE jurusan (
    ID SERIAL PRIMARY KEY,
    Nama  VARCHAR(50),
    Fakultas  VARCHAR(50),
    FOREIGN KEY (Fakultas) REFERENCES fakultas(Nama) ON DELETE CASCADE,
    CONSTRAINT UC_NamaJurusan UNIQUE (Nama)
);

CREATE TABLE mata_kuliah (
    ID SERIAL PRIMARY KEY,
    Nama VARCHAR(50),
    SKS INT,
    Jurusan VARCHAR(50),
    Fakultas  VARCHAR(50),
    SemMin INT,
    Prediksi VARCHAR(2),
    FOREIGN KEY (Jurusan) REFERENCES jurusan(Nama) ON DELETE CASCADE,
    FOREIGN KEY (Fakultas) REFERENCES fakultas(Nama) ON DELETE CASCADE,
    CONSTRAINT UC_NamaMatkul UNIQUE (Nama)
);