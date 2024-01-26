module projet

go 1.21.6

replace projet/fichiers => ./fichiers

replace projet/dossiers => ./dossiers

replace projet/sql => ./database

require (
	projet/dossiers v0.0.0-00010101000000-000000000000
	projet/fichiers v0.0.0-00010101000000-000000000000
	projet/sql v0.0.0-00010101000000-000000000000
)

require github.com/go-sql-driver/mysql v1.7.1 // indirect
