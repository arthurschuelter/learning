import time
import sys
import os
##
from Battle import Battle
from Player import Player
from Move import Move
import Mechanics
import Red

def devMenu(Pokemon, Moves):
	option = 1
	os.system('clear')

	while(option != 0):
		print ('===\nDeveloper Menu:')

		print('\t1) List all available Pokemon')
		print('\t2) List all available moves')
		print('\t0) To exit')

		print ('Choice:'),
		option = input()

		if option == 1:
			Mechanics.listAllPokemon(Pokemon)
		elif option == 2:
			Mechanics.listAllMoves(Moves)

		else:
			option=0

def gameMenu(Pokemon, Moves):
	player = Mechanics.createPlayer()

	os.system('clear')
	option = 1
	win = 0
	starter = 0

	while (option != 0):
		print ('===\nGame Menu:')

		if len(player.PokemonList) != 0:
			print('\t1) Catch a pokemon!')
		else:
			print('\t1) Get your first pokemon!')
		
		print('\t2) See your info')

		if len(player.PokemonList) != 0:
			print('\t3) See the stats of your Pokemon!')

		if win == 1:
			print('\t4) Challenge a strong opponent')
		
		print('\t0) Exit')

		print ('Choice:'),
		option = input()

		if option == 1:
			if starter == 0:
				blue = Mechanics.getYourFirstPokemon(player, Pokemon)
				starter += 1
			else:
				winner = Mechanics.battleStart(player, blue)

				if winner.trainer_id == player.trainer_id:
					Mechanics.print_slow('Congratulations, you beat the game!!')
					win = 1
				else:
					os.system('clear')
					Mechanics.print_slow(player.name + ' is out of usable Pokemon\n')
					Mechanics.print_slow(player.name + ' lost all hope\n')
					time.sleep(3)
					os.system('clear')
					Mechanics.print_slow('Don\'t give up! You can beat Blue!')
					time.sleep(2)
					os.system('clear')

				time.sleep(1)
					
				option = 0



		elif option == 2:
			player.playerStats()

		elif option == 3:

			if len(player.PokemonList) == 0:
				print('You have no Pokemon :/')
				time.sleep(1)
				os.system('clear')

			else:
				Mechanics.listTeamPokemon(player)

		elif option == 4:
			Red.RedMenu(player)

		else:
			option = 0

	return win



def mainMenu():
	pokemon, moves = Mechanics.initializeGame()
	option = 1
	win = 0

	while(option != 0):
		os.system('clear')
		print ("===\nMain Menu:")

		print ('\t1) Go to dev menu')
		print ('\t2) Play the game')
		print ('\t0) To exit')
		
		print ('Choice:'),
		option = input()

		if option == 1:
			devMenu(pokemon, moves)
		elif option == 2:
			win = gameMenu(pokemon, moves)

		else:
			option = 0


	os.system('clear')
	print ('Go in the shadow')

if __name__ == "__main__": mainMenu()