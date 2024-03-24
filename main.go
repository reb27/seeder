package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/brianvoe/gofakeit/v6"
)

// minimo 400 livros
func seedLivros(db *sql.DB) {
	for i := 0; i < 400; i++ {
		codigo := gofakeit.UUID()
		nome := gofakeit.BookTitle()
		lingua := gofakeit.Language()
		ano := gofakeit.Date()

		_, err := db.Exec("INSERT INTO Livro (codigo, nome, lingua, ano) VALUES (?, ?, ?, ?)", codigo, nome, lingua, ano)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("____________Tabela de Livros populada com sucesso!_________________")
}

// Minimo 50 editoras
func seedEditoras(db *sql.DB) {
	for i := 0; i < 50; i++ {
		codigo := gofakeit.UUID()
		nome := gofakeit.Company()
		endereco := gofakeit.StreetName()
		telefone := gofakeit.Phone()

		_, err := db.Exec("INSERT INTO Editora (codigo, nome, endereco, telefone) VALUES (?, ?, ?, ?)", codigo, nome, endereco, telefone)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("____________Tabela de Editoras populada com sucesso!_________________")
}

// Minimo 100 autores
func seedAutores(db *sql.DB) {
	for i := 0; i < 100; i++ {
		codigo := gofakeit.UUID()
		nome := gofakeit.Name()
		dtNascimento := gofakeit.Date()
		paisNascimento := gofakeit.Country()
		biografia := gofakeit.Paragraph(1, 1, 10, " ")

		_, err := db.Exec("INSERT INTO Autor (codigo, nome, dtNascimento, paisNascimento, biografia) VALUES (?, ?, ?, ?, ?)", codigo, nome, dtNascimento, paisNascimento, biografia)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("____________Tabela de Autores populada com sucesso!_________________")
}

// Minimo 1000 edições
func seedEdicoes(db *sql.DB) {
	// Select para buscar todos os codigos pre existentes de editoras
	rowsEditora, err := db.Query("SELECT codigo FROM Editora")
	if err != nil {
		panic(err)
	}
	defer rowsEditora.Close()

	var codigosEditoras []string

	for rowsEditora.Next() {
		var cod string
		if err := rowsEditora.Scan(&cod); err != nil {
			panic(err)
		}
		codigosEditoras = append(codigosEditoras, cod)
	}

	for i := 0; i < 1000; i++ {
		isbn := geradorIsbn()
		preco := gofakeit.Price(10, 100)
		ano := gofakeit.Date()
		numeroPaginas := gofakeit.Number(100, 500)
		qtdEstoque := gofakeit.Number(10, 100)
		//Esse trecho pega um indice aletorio com um codigo de uma editora pre existente
		editoraCodigo := codigosEditoras[gofakeit.IntRange(0, len(codigosEditoras)-1)]

		_, err := db.Exec("INSERT INTO Edicao (isbn, preco, ano, numeroPaginas, qtdEstoque, editoraCodigo) VALUES (?, ?, ?, ?, ?, ?)",
			isbn, preco, ano, numeroPaginas, qtdEstoque, editoraCodigo)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("____________Tabela de Edições populada com sucesso!_________________")
}

// Função para gerar um ISBN aleatório
func geradorIsbn() string {
	ean := gofakeit.Number(100, 999)
	group := gofakeit.Number(10, 99)
	editor := gofakeit.Number(100, 999)
	titulo := gofakeit.Number(1000, 9999)
	digito := gofakeit.Number(0, 9)

	return fmt.Sprintf("%d-%d-%d-%d-%d", ean, group, editor, titulo, digito)
}

func seedEdicaoLivro(db *sql.DB) {
	var (
		isbns         []string
		codigosLivros []string
	)

	rowsLivros, err := db.Query("SELECT codigo FROM Livro")
	if err != nil {
		panic(err)
	}
	defer rowsLivros.Close()

	for rowsLivros.Next() {
		var cod string
		if err := rowsLivros.Scan(&cod); err != nil {
			panic(err)
		}
		codigosLivros = append(codigosLivros, cod)
	}

	isbnEdicoes, err := db.Query("SELECT isbn FROM Edicao")
	if err != nil {
		panic(err)
	}
	defer isbnEdicoes.Close()

	for isbnEdicoes.Next() {
		var cod string
		if err := isbnEdicoes.Scan(&cod); err != nil {
			panic(err)
		}
		isbns = append(isbns, cod)
	}

	for i := 0; i < 1000; i++ {
		edicaoIsbn := isbns[gofakeit.IntRange(0, len(isbns)-1)]
		codigoLivro := codigosLivros[gofakeit.IntRange(0, len(codigosLivros)-1)]

		_, err := db.Exec("INSERT INTO EdicaoLivro (codigo, codigoLivro, edicaoIsbn) VALUES (?, ?, ?)", gofakeit.UUID(), codigoLivro, edicaoIsbn)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("____________Tabela de Relaçao de Livros e Edições populada com sucesso!_________________")

}

func seedLivroAutor(db *sql.DB) {
	var (
		codigosLivros []string
		codigoAutores []string
	)

	rowsLivros, err := db.Query("SELECT codigo FROM Livro")
	if err != nil {
		panic(err)
	}
	defer rowsLivros.Close()

	for rowsLivros.Next() {
		var cod string
		if err := rowsLivros.Scan(&cod); err != nil {
			panic(err)
		}
		codigosLivros = append(codigosLivros, cod)
	}

	rowsAutores, err := db.Query("SELECT codigo FROM Autor")
	if err != nil {
		panic(err)
	}
	defer rowsAutores.Close()

	for rowsAutores.Next() {
		var cod string
		if err := rowsAutores.Scan(&cod); err != nil {
			panic(err)
		}
		codigoAutores = append(codigoAutores, cod)
	}

	for i := 0; i < 400; i++ {
		codigoAutor := codigoAutores[gofakeit.IntRange(0, len(codigoAutores)-1)]
		codigoLivro := codigosLivros[gofakeit.IntRange(0, len(codigosLivros)-1)]

		_, err := db.Exec("INSERT INTO LivroAutor (livroCodigo, codigoAutor) VALUES (?, ?)", codigoLivro, codigoAutor)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("____________Tabela de Relaçao de Livros e Autores populada com sucesso!_________________")
}
func main() {
	db, err := sql.Open("mysql", "root:livraria2024@tcp(127.0.0.1:3306)/livraria")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	seedLivros(db)
	seedEditoras(db)
	seedAutores(db)
	seedEdicoes(db)
	seedEdicaoLivro(db)
	seedLivroAutor(db)

	fmt.Println("################# Todas as tabelas forão populadas com sucesso! ___________FIM___________")
}
