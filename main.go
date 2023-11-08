package main

import (
	"fmt"
)

type automate struct {
	alphabet    map[string]bool
	etats       map[string]bool
	transition  map[string]map[string]string
	etatsFinaux map[string]bool
	etatInitial string
}

func (a automate) afficher() {
	fmt.Println("Alphabet : ")
	for lettre := range a.alphabet {
		fmt.Println(lettre)
	}

	fmt.Println("États : ")
	for etat := range a.etats {
		fmt.Println(etat)
	}

	fmt.Println("État Final : ")
	for etatFinal := range a.etatsFinaux {
		fmt.Println(etatFinal)
	}

	fmt.Println("État Initial : ")
	fmt.Println(a.etatInitial)

	fmt.Println("Transition : ")
	for etat, transitions := range a.transition {
		for lettre, etatSuivant := range transitions {
			fmt.Printf("%s %s %s\n", etat, lettre, etatSuivant)
		}
	}
}

func isWordValid(word string, alphabet map[string]bool) bool {
	for _, lettre := range word {
		if !alphabet[string(lettre)] {
			return false
		}
	}
	return true
}

func (a automate) peutLireMot(mot string) (bool, error) {
	if !isWordValid(mot, a.alphabet) {
		return false, fmt.Errorf("le mot '%s' contient des lettres invalides", mot)
	}
	etatCourant := a.etatInitial
	for _, lettre := range mot {
		lettreStr := string(lettre)
		prochainEtat, transitionValide := a.transition[etatCourant][lettreStr]
		if !transitionValide {
			return false, fmt.Errorf("aucune transition valide pour la lettre '%s' dans l'état '%s'", lettreStr, etatCourant)
		}
		etatCourant = prochainEtat
	}
	estEtatFinal := a.etatsFinaux[etatCourant]
	return estEtatFinal, nil
}

func main() {
	automate := automate{}
	fmt.Println("Entrez les informations de l'automate d’états finis simple déterministe <X, S, S0, F, II> :")
	var tailleAlphabet int
	fmt.Print("Entrez la taille de l'alphabet : ")
	_, err := fmt.Scanln(&tailleAlphabet)
	if err != nil {
		return
	}
	alphabet := make(map[string]bool)
	fmt.Println("Entrez l'alphabet :")
	for i := 0; i < tailleAlphabet; i++ {
		var lettre string
		_, err := fmt.Scan(&lettre)
		if err != nil {
			return
		}
		alphabet[lettre] = true
	}
	automate.alphabet = alphabet

	var tailleEtats int
	fmt.Print("Entrez le nombre d'états : ")
	_, err = fmt.Scanln(&tailleEtats)
	if err != nil {
		return
	}
	etats := make(map[string]bool)
	fmt.Println("Entrez les états :")
	for i := 0; i < tailleEtats; i++ {
		var etat string
		_, err := fmt.Scan(&etat)
		if err != nil {
			return
		}
		etats[etat] = true
	}
	automate.etats = etats

	transition := make(map[string]map[string]string)
	for etat := range etats {
		transition[etat] = make(map[string]string)
		fmt.Printf("Transitions à partir de l'état %s :\n", etat)
		for lettre := range alphabet {
			var etatSuivant string
			fmt.Printf("Entrez l'état suivant pour la lettre %s : ", lettre)
			_, err := fmt.Scan(&etatSuivant)
			if err != nil {
				return
			}
			transition[etat][lettre] = etatSuivant
		}
	}
	automate.transition = transition

	fmt.Print("Entrez l'état initial : ")
	_, err3 := fmt.Scan(&automate.etatInitial)
	if err3 != nil {
		return
	}
	var tailleEtatsFinaux int
	fmt.Print("Entrez le nombre d'états finaux : ")
	_, err = fmt.Scanln(&tailleEtatsFinaux)
	if err != nil {
		return
	}
	etatsFinaux := make(map[string]bool)
	fmt.Println("Entrez les états finaux :")
	for i := 0; i < tailleEtatsFinaux; i++ {
		var etatFinal string
		_, err := fmt.Scan(&etatFinal)
		if err != nil {
			return
		}
		etatsFinaux[etatFinal] = true
	}
	automate.etatsFinaux = etatsFinaux

	automate.afficher()

	var mot string
	for {
		fmt.Print("Entrez un mot à vérifier : ")
		_, err2 := fmt.Scan(&mot)
		if err2 != nil {
			return
		}
		estValide, err := automate.peutLireMot(mot)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Le mot '%s' est valide : %t\n", mot, estValide)
		}
		fmt.Println("voulez-vous continuer ? (y/n)")
		var reponse string
		_, err = fmt.Scan(&reponse)
		if err != nil {
			return
		}
		if reponse != "y" {
			break
		}
	}
}
