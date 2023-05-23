#The-Knapsack-Problem

Ce projet, intitulé "The-Knapsack-Problem", explore le problème du sac à dos en cryptographie. Il propose une implémentation de l'attaque de cryptanalyse sur le protocole Merkle-Hellman, l'algorithme de réduction de réseau LLL (Lenstra–Lenstra–Lovász), ainsi que des attaques sur d'autres cryptosystèmes.
Compilation et exécution du projet

##Pour compiler le projet, utilisez la commande suivante :

go build

##Pour exécuter le projet, utilisez la commande suivante :

./Kna...

##Pour lancer les tests, utilisez la commande suivante :


go test -bench=.

##Fonctionnalités

Le programme principal (main) du projet offre les fonctionnalités suivantes :

    Génération de données : Le programme génère des données aléatoires pour le problème du sac à dos à l'aide de la fonction create_data.GenerateData(). Cela crée un fichier JSON contenant les objets avec leurs valeurs et poids correspondants.

    Benchmark du problème du sac à dos : Le programme effectue un benchmark du problème du sac à dos en utilisant trois approches différentes : l'algorithme glouton, la programmation dynamique et la recherche exhaustive. Le benchmark mesure le temps d'exécution de chaque algorithme pour résoudre le problème du sac à dos avec les données générées précédemment. Les résultats du benchmark sont affichés à l'écran.

    Génération de clés : Le programme génère une paire de clés publiques et privées aléatoires pour le protocole Merkle-Hellman à l'aide de la fonction tools.GenerateKeys(). Les clés générées sont utilisées pour chiffrer et déchiffrer un message.

    Réduction de réseau : Le programme génère un réseau initial de Lagarias-Odlyzko et un réseau initial de Joux-Stern à l'aide des fonctions algo_reduc_reseau.GenerateLagariasOdlyzkoNetwork() et algo_reduc_reseau.GenerateJouxSternNetwork(). Ensuite, il applique l'algorithme LLL aux réseaux respectifs en utilisant les fonctions algo_reduc_reseau.LLL(LONetwork, big.NewRat(3, 4), 1000) et algo_reduc_reseau.LLL(JSNetwork, big.NewRat(3, 4), 1000). Les résultats de la réduction des réseaux sont affichés à l'écran, et une vérification est effectuée pour s'assurer de la validité des résultats.

    Attaque (en cours de développement) : Le programme inclut du code commenté pour une attaque basée sur la réduction de réseau sur le protocole Merkle-Hellman. Cependant, cette partie de l'attaque est en cours de développement et est actuellement désactivée.

##Notes

Veuillez noter que l'attaque basée sur la réduction de réseau est en cours de développement et n'est pas fonctionnelle dans l'état actuel du programme. Les autres fonctionnalités, telles que le benchmark du problème du sac à dos, la génération de clés et la réduction de réseau, sont pleinement fonctionnelles.

N'hésitez pas à explorer le code source pour une compréhension plus détaillée de chaque fonctionnalité et de son implémentation.

2022/2023 Tutored Projects - Knapsack Problem in Cryptography. Implement cryptanalysis of Merkle-Hellman protocol, LLL lattice reduction algorithm, and explore attacks on other cryptosystems.
