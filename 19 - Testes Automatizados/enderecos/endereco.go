package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]
	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
