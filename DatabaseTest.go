package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Persona struct{
	c_f string
	nome string
	cognome string
}

type Telefono struct{
	intestatario string
	numero string
}

func main(){
	fmt.Println("Programma di test di un database con GO")
	
	//connessione al database
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/comunity")
	if err != nil {
        panic(err.Error())
    }
	fmt.Println("Connessione OK")
	
    //chiusura deferenziata della connessione
    defer db.Close()
	
	//eseguo la query
	results, err := db.Query("SELECT * FROM persona")
    if err != nil {
        panic(err.Error()) 
    }

	//leggo e stampo i valori ottenuti dalla select
	fmt.Println("\nIl contenuto della tabella persona:")
    for results.Next() {
        var persona Persona
		
        err = results.Scan(&persona.c_f, &persona.nome, &persona.cognome)
        if err != nil {
            panic(err.Error()) 
        }
        
        fmt.Println("nome: ", persona.nome, ", cognome: ", persona.cognome, ", Codice Fiscale: ", persona.c_f)
    }
	
	results, err = db.Query("SELECT * FROM telefono")
	if err != nil {
        panic(err.Error()) 
    }
	fmt.Println("\nIl contenuto della tabella telefono:")
    for results.Next() {
        var telefono Telefono
		
        err = results.Scan(&telefono.numero, &telefono.intestatario)
        if err != nil {
            panic(err.Error()) 
        }
        
        fmt.Println("numero di telefono: ", telefono.numero, ", Codice Fiscale Intestatario: ", telefono.intestatario)
    }
	
	//inserisco una nuova persona
	results, err = db.Query("INSERT INTO persona (c_f, nome, cognome) VALUES ('LMBGVN37D45L111J','Giovanni','Lamberti')")
	if err != nil {
        panic(err.Error()) 
    }
		
	//modifico il nome della persona appena aggiunta
	results, err = db.Query("UPDATE persona SET nome = 'Giovanna' WHERE c_f = 'LMBGVN37D45L111J'")
	if err != nil {
        panic(err.Error()) 
    }
		
	results, err = db.Query("SELECT * FROM persona")
    if err != nil {
        panic(err.Error()) 
    }
	
	fmt.Println("\nIl contenuto finale della tabella persona dopo le modifiche:")
    for results.Next() {
        var persona Persona
		
        err = results.Scan(&persona.c_f, &persona.nome, &persona.cognome)
        if err != nil {
            panic(err.Error()) 
        }
        
        fmt.Println("nome: ", persona.nome, ", cognome: ", persona.cognome, ", Codice Fiscale: ", persona.c_f)
    }
	
	//elimino la persona aggiunta
	results, err = db.Query("DELETE FROM persona WHERE c_f = 'LMBGVN37D45L111J'")
	if err != nil {
        panic(err.Error()) 
    }
}