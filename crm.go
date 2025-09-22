package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

// bdd map
var contacts = make(map[int]Contact)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// Flags
	idFlag := flag.Int("id", -1, "ID du contact")
	nomFlag := flag.String("nom", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *idFlag != -1 && *nomFlag != "" && *emailFlag != "" {
		ajouterContact(*idFlag, *nomFlag, *emailFlag)
		fmt.Println("Contact ajouté via flags !")
		listerContacts()
		return
	}

	// For pour menu
	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")

		fmt.Print("Choix: ")
		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			ajouterContactCLI()
		case "2":
			listerContacts()
		case "3":
			supprimerContact()
		case "4":
			mettreAJourContact()
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

// CRUD functions
func ajouterContact(id int, nom string, email string) {
	contacts[id] = Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}
}
func ajouterContactCLI() {
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	fmt.Print("Nom: ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	ajouterContact(id, nom, email)
	fmt.Println("Contact ajouté avec succès !")
}
func listerContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé.")
		return
	}
	fmt.Println("\n--- Liste des contacts ---")
	for _, c := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Nom, c.Email)
	}
}
func supprimerContact() {
	fmt.Print("ID du contact à supprimer: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println("Contact supprimé !")
	} else {
		fmt.Println("Contact non trouvé.")
	}
}
func mettreAJourContact() {
	fmt.Print("ID du contact à mettre à jour: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	contact, ok := contacts[id]
	if !ok {
		fmt.Println("Contact non trouvé.")
		return
	}

	fmt.Printf("Nouveau nom (%s): ", contact.Nom)
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom != "" {
		contact.Nom = nom
	}

	fmt.Printf("Nouvel email (%s): ", contact.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		contact.Email = email
	}

	contacts[id] = contact
	fmt.Println("Contact mis à jour avec succès !")
}
