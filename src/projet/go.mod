module projet

go 1.21.6

replace projet/fichiers => ./fichiers

replace projet/dossiers => ./dossiers

require (
	projet/dossiers v0.0.0-00010101000000-000000000000
	projet/fichiers v0.0.0-00010101000000-000000000000
)
