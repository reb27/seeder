CREATE TABLE Autor (
    codigo VARCHAR(36) NOT NULL PRIMARY KEY,
    nome VARCHAR(255) NULL,
    dtNascimento DATETIME NULL,
    paisNascimento VARCHAR(255) NULL,
    biografia VARCHAR(500) NULL
);

CREATE TABLE Editora (
    codigo VARCHAR(36) NOT NULL PRIMARY KEY,
    nome VARCHAR(255) NULL,
    endereco VARCHAR(255) NULL,
    telefone VARCHAR(20) NULL
);

CREATE TABLE Edicao (
    isbn VARCHAR(36) NOT NULL PRIMARY KEY,
    preco FLOAT NULL,
    ano DATETIME NULL,
    numeroPaginas INT NULL,
    qtdEstoque INT NULL,
    editoraCodigo VARCHAR(36) NULL,
    CONSTRAINT Edicao_Editora_codigo_fk FOREIGN KEY (editoraCodigo) REFERENCES Editora (codigo)
);

CREATE TABLE Livro (
    codigo VARCHAR(36) NOT NULL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    lingua VARCHAR(255) NOT NULL,
    ano DATETIME NOT NULL
);

CREATE TABLE EdicaoLivro (
    codigoLivro VARCHAR(36) NOT NULL,
    edicaoIsbn VARCHAR(36) NOT NULL,
    CONSTRAINT EdicaoLivro_Edicao_isbn_fk FOREIGN KEY (edicaoIsbn) REFERENCES Edicao (isbn),
    CONSTRAINT EdicaoLivro_Livro_codigo_fk FOREIGN KEY (codigoLivro) REFERENCES Livro (codigo)
);

CREATE TABLE LivroAutor (
    livroCodigo VARCHAR(36) NOT NULL,
    codigoAutor VARCHAR(36) NOT NULL,
    CONSTRAINT LivroAutor_Autor_codigo_fk FOREIGN KEY (codigoAutor) REFERENCES Autor (codigo),
    CONSTRAINT LivroAutor_Livro_codigo_fk FOREIGN KEY (livroCodigo) REFERENCES Livro (codigo)
);
-- Criação dos Index ##################################################################################################

CREATE INDEX idx_autor_codigo ON Autor (codigo);

CREATE INDEX idx_editora_codigo ON Editora (codigo);

CREATE INDEX idx_edicao_isbn ON Edicao (isbn);

CREATE INDEX idx_livro_codigo ON Livro (codigo);

CREATE INDEX idx_edicaolivro_codigoLivro ON EdicaoLivro (codigoLivro);
CREATE INDEX idx_edicaolivro_edicaoIsbn ON EdicaoLivro (edicaoIsbn);

CREATE INDEX idx_livroautor_codigoAutor ON LivroAutor (codigoAutor);
CREATE INDEX idx_livroautor_livroCodigo ON LivroAutor (livroCodigo);

