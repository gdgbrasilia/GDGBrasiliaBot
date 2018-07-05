package comandos

func GetComandos() string {
	comandos := "Alguns dos comandos que eu entendo: \n \n"
	comandos = comandos + "/meetups - Lista dos proximos meetups do GDG Brasília \n"
	comandos = comandos + "/vagas - Lista de vagas cadastrados no GDG Brasília \n"
	comandos = comandos + "/livros - Um livro gratuto por dia \n"
	comandos = comandos + "/telegram - Link do Grupo do Telegram do GDG Brasília \n"
	comandos = comandos + "/github - Link da Organização GDG Brasília no GitHub \n"
	comandos = comandos + "/projeto - Link do Projeto do GDGBrasiliaBot no GitHub \n"
	comandos = comandos + "/midias - Links das Mídias Sociais do GDGBrasiliaBot \n"
	comandos = comandos + "/contato - Email de Contato do GDG Brasília \n"
	return comandos
}
