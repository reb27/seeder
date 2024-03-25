-- Criação da View ##################################################################################################

CREATE OR REPLACE VIEW livraria.LIVRO_MAIOR_EDICAO AS
SELECT DISTINCT lv.codigo AS codigo_livro, lv.nome AS livro_nome, lv.lingua AS lingua, lv.ano AS ano, edit.nome AS editora_nome, ed.isbn,
(
SELECT SUM(ed2.qtdEstoque) FROM livraria.Edicao AS ed2
INNER JOIN livraria.EdicaoLivro edlv ON edlv.edicaoIsbn = ed2.isbn
WHERE edlv.codigoLivro = lv.codigo
) AS qtd_estoque
FROM livraria.Livro AS lv
INNER JOIN livraria.EdicaoLivro edlv ON edlv.codigoLivro = lv.codigo
INNER JOIN livraria.Edicao ed ON ed.isbn = edlv.edicaoIsbn
INNER JOIN livraria.Editora edit ON edit.codigo = ed.editoraCodigo
ORDER BY qtd_estoque DESC