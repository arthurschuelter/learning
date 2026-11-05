import random as ran
import os
import sys
import time
##
from Pokemon import Pokemon
from Battle import Battle
from Player import Player
from Move import Move

NORMAL = 0
FIGHT  = 1
FLYING = 2
POISON = 3
GROUND = 4
ROCK   = 5
BUG    = 6
GHOST  = 7
STEEL  = 8
FIRE   = 9
WATER  = 10
GRASS  = 11
ELECTR = 12
PSYCHIC= 13
ICE    = 14
DRAGON = 15
DARK   = 16
FAIRY  = 17
NULL   = 18

typelist = 	['NORMAL', 'FIGHTING', 'FLYING', 'POISON',  	# [0  ~  3]
			 'GROUND', 'ROCK', 'BUG', 'GHOST', 'STEEL', 	# [4  ~  8]
			 'FIRE', 'WATER', 'GRASS', 'ELECTRIC',			# [9  ~ 12]
			 'PSYCHIC', 'ICE', 'DRAGON', 'DARK', 'FAIRY',	# [13 ~ 17]
			 'NULL']										# [   18  ]

#(DEF ->)	 Nor, Fig, Fly, Poi, Gro, Roc, Bug, Gho, Ste, Fre, Wat, Gra, Ele, Psy, Ice, Dra, Dar, Fai, NULL
typeadv =[  [1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 0.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0], #NORMAL
			[2.0, 1.0, 0.5, 0.5, 1.0, 2.0, 0.5, 0.0, 2.0, 1.0, 1.0, 1.0, 1.0, 0.5, 2.0, 1.0, 2.0, 0.5, 1.0], #FIGHTING
			[1.0, 2.0, 1.0, 1.0, 1.0, 0.5, 2.0, 1.0, 0.5, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0], #FLYING
			[1.0, 1.0, 1.0, 0.5, 0.5, 0.5, 1.0, 0.5, 0.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0], #POISON
			[1.0, 1.0, 0.0, 2.0, 1.0, 2.0, 0.5, 1.0, 2.0, 2.0, 1.0, 0.5, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0], #GROUND
			[1.0, 0.5, 2.0, 1.0, 0.5, 1.0, 2.0, 1.0, 0.5, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0], #ROCK
			[1.0, 0.5, 0.5, 0.5, 1.0, 1.0, 1.0, 0.5, 0.5, 0.5, 1.0, 2.0, 1.0, 2.0, 1.0, 1.0, 2.0, 0.5, 1.0], #BUG
			[0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 0.5, 1.0, 1.0], #GHOST
			[1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 0.5, 0.5, 0.5, 1.0, 0.5, 1.0, 2.0, 1.0, 1.0, 2.0, 1.0], #STEEL
			[1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 2.0, 1.0, 2.0, 0.5, 0.5, 2.0, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0], #FIRE
			[1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 1.0, 1.0, 1.0, 2.0, 0.5, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0], #WATER
			[1.0, 1.0, 0.5, 0.5, 2.0, 2.0, 0.5, 1.0, 0.5, 0.5, 2.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0], #GRASS
			[1.0, 1.0, 2.0, 1.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 0.5, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0], #ELECTRIC
			[1.0, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 0.0, 1.0, 1.0], #PSYCHIC
			[1.0, 1.0, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 0.5, 0.5, 0.5, 2.0, 1.0, 1.0, 0.5, 2.0, 1.0, 1.0, 1.0], #ICE
			[1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.0, 1.0], #DRAGON
			[1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 0.5, 0.5, 1.0], #DARK
			[1.0, 2.0, 1.0, 0.5, 1.0, 1.0, 1.0, 2.0, 0.5, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 1.0, 1.0], #FAIRY
			[1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0]  #NULL
		]

category= ['PHYSICAL', 'SPECIAL', ' OTHER']

#GAIN |Lose:  attack | defense | spc atk | spc def | speed
attack =    [ 'Hardy', 'Lonely','Adamant','Naughty', 'Brave']
defense=    [ 'Bold' , 'Docile', 'Impish',  'Lax'  , 'Relaxed']
sp_atk =    ['Modest', 'Mild'  ,'Serious',  'Rash' , 'Quiet']
sp_def =    [ 'Calm' , 'Gentle','Careful','Bashful', 'Sassy']
speed  =    ['Timid' , 'Hasty' , 'Jolly' , 'Naive' , 'Quirky']

natures = [attack, defense, sp_atk, sp_def, speed]
statName = ['ATTACK', 'DEFENSE', 'SP. ATTACK', 'SP. DEFENSE', 'SPEED']


def print_slow(string, speed = 0.08):
	for letter in string:
		sys.stdout.write(letter)
		sys.stdout.flush()
		time.sleep(speed)

def initializeGame():
	moves = []
	pokemon = []

	moves.append( Move('Tackle' , 'No additional effect.', 40, 100, 35, 0, 0, 1) )
	moves.append( Move('Scratch', 'No additional effect.', 40, 100, 35, 0, 0, 1) )
	moves.append( Move('Growl', 'Lowers opponent\'s Attack by one stage.', 0 , 100, 40, 2, 0, 0) )
	moves.append( Move('Tail Whip', 'Lowers opponent\'s Defense by one stage.', 0 , 100, 30, 2, 0, 0) )
	moves.append( Move('Ember', 'May induce burn.', 40, 100, 25, 1, 9, 1) )
	moves.append( Move('Water Gun', 'No additional effect.', 40, 100, 25, 1, 10, 1) )
	moves.append( Move('Vine Whip', 'No additional effect.', 45, 100, 25, 0, 11, 1) )

	pokemon.append( Pokemon("BULBASAUR", "Bulbasaur", 11, 3) )
	pokemon[0].setStats(21, 11, 11, 13, 13, 11)
	pokemon[0].learnMove(moves[0]) # Tackle
	pokemon[0].learnMove(moves[2]) # Growl

	pokemon.append( Pokemon("CHARMANDER", "Charmander", 9) )
	pokemon[1].setStats(20, 11, 10, 12, 11, 13)
	pokemon[1].learnMove(moves[1]) # Sratch
	pokemon[1].learnMove(moves[2]) # Growl

	pokemon.append( Pokemon("SQUIRTLE", "Squirtle", 10) )
	pokemon[2].setStats(20, 11, 13, 11, 12, 10)
	pokemon[2].learnMove(moves[0]) # Tackle
	pokemon[2].learnMove(moves[3]) # Tail Whip

	return pokemon, moves

###

def listAllMoves(Moves):

	os.system('clear')

	print('All moves:\n')

	for i in range(0, len(Moves) ):
		print ('  ' + str(i+1) + ') ' + Moves[i].getName() + '\t[' + typelist[Moves[i].getTyping()] +  ']  ')


	print('\nSelect a move:')
	opt = input()

	if (opt > 0 and opt <= len(Moves) ):
		moveInfo(Moves[opt-1])
	else:
		return

def listAllPokemon(Pokemon):
	os.system('clear')

	print ('All currently available Pokemon:\n')

	for i in range(0,len(Pokemon)):
		print ( '   #' + str(i+1).zfill(3)  + ' - ' + Pokemon[i].species)

	print ('---\n')
	print ('Choose a Pokemon:'),
	opt = input()
	printGenericStats(Pokemon[opt-1])

def listTeamPokemon(player):
	os.system('clear')
	player.teamPokemon()

	print('===\nSelect one:'),
	opt = input()

	if opt == 0:
		return 0

	poke = player.PokemonList[opt-1]
	printPokeStats(poke)
	return 1

###

def printGenericStats(Pokemon):
	os.system('clear')

	print(Pokemon.species)
	#print('Level:  ' + str(Pokemon.level))

	print('Type: '),
	for i in range(0, len(Pokemon.typing) ):
		if Pokemon.typing[i] != 18:
			print('| ' + typelist[ Pokemon.typing[i] ]),

	print('|')

	Pokemon.printGenericBaseStats()

	printMoves(Pokemon)

def printPokeStats(Pokemon):
	os.system('clear')
	Pokemon.printSimpleStats()
	print ('OT: ' + str(Pokemon.trainer_ot).zfill(5))
	print ('Nature: ' + natures[Pokemon.gainStat][Pokemon.loseStat]),

	if Pokemon.loseStat == Pokemon.gainStat:
		print('(NEUTRAL)')
	else:
		print ('(UP: ' + statName[Pokemon.gainStat] + ', DOWN: ' + statName[Pokemon.loseStat] + ')' )

	Pokemon.printPokeBaseStats()
	printMoves(Pokemon)

def printMoves(Pokemon):
	print ('Moves:')

	for move in range(0, len(Pokemon.moves) ):
		if Pokemon.moves[move].name != '-':
			print ('  ' + str(move+1) + ') ' + Pokemon.moves[move].name + '\t' + str(Pokemon.moves[move].current_PP) + '/' + str(Pokemon.moves[move].total_PP) + '\t[' + typelist[Pokemon.moves[move].typing] +  ']  ')
		else:
			print ('  ' + str(move+1) + ') ' + '  -')

def moveInfo(Move):

	os.system('clear')

	if(Move.name == '-'):
		print ('Invalid choice')
		return -1


	print (Move.name)
	print ('+---------------+---------------+')
	print ('|  TYPE:\t|     ' + 	typelist[Move.typing] + '\t|')
	print ('|  CATEGORY:\t|    ' + 	category[Move.category] + '\t|')
	print ('|  POWER:\t|\t' + str(Move.base_power) + '\t|')
	print ('|  ACCURACY:\t|      ' + str(Move.accuracy) + '\t|' )
	print ('|  PP:\t\t|     ' + str(Move.current_PP) + '/' + str(Move.total_PP) + '\t|' )
	print ('+---------------+---------------+')
	print ('Description: ')
	print (Move.effect)
	print ('')

###

def createPlayer():
	speed = 0.1
	os.system('clear')
	print_slow('Good morning!\n')
	time.sleep(1)
	print_slow('My name is Professor Oak.\n')
	time.sleep(1)
	print_slow('But everyone calls me the Pokemon Professor.\n')
	time.sleep(1)
	print_slow('This world is widely inhabited by creatures known as Pokemon.\n')
	time.sleep(1)
	print_slow('Now, ')
	time.sleep(1)
	print_slow('why don\'t you tell me a little about yourself?\n')
	time.sleep(1)
	print_slow('\nPlease,')
	time.sleep(1)
	print_slow(' tell me your name: '),
	name = raw_input()

	os.system('clear')
	print_slow('Oh, hello ' + name + '!')
	time.sleep(1)

	gender = 0
	while(gender != 1 and gender != 2 and gender != 3):
		os.system('clear')
		print_slow('Are you a boy?\n', speed)
		time.sleep(1)
		print_slow('Or are you a girl?\n', speed)

		print('  1) Boy')
		print('  2) Girl')
		print('  3) OTHER')
		gender = input('Choice: ')

		speed = 0.05

	player = Player(name, gender)

	os.system('clear')
	print_slow('Your very own tale of grand adventure is about to unfold.\n')
	time.sleep(1)
	print_slow('Dreams! ')
	time.sleep(1)
	print_slow('Adventure! ')
	time.sleep(1)
	print_slow('Let\'s go to the world of Pokemon!')
	time.sleep(4)

	return player

def getYourFirstPokemon(player, pokemon):
	os.system('clear')
	print_slow('You can choose between three Pokemon:\n')
	time.sleep(1)



	print_slow('\nBulbasaur, the Seed Pokemon,\n')
	time.sleep(1)
	print_slow('It likes to take naps on bright sunlight!\n')
	time.sleep(1)

	print_slow('\nCharmander, the Lizard Pokemon,\n')
	time.sleep(1)
	print_slow('The flame at the tip of its tail is an indication of its emotions!\n')
	time.sleep(1)

	print_slow('\nSquirtle, the Tiny Turtle Pokemon,\n')
	time.sleep(1)
	print_slow('The shells rounded shape helps minimize resistance in water, enabling it to swim at high speeds!\n')
	time.sleep(3)



	choice = 0
	speed = 0.08
	while(choice != 1 and choice != 2 and choice != 3):
		os.system('clear')
		print_slow('So who will you choose?\n', speed)
		print('  1) Bulbasaur, a GRASS-TYPE Pokemon')
		print('  2) Charmander, a FIRE-TYPE Pokemon')
		print('  3) Squirtle, a WATER-TYPE Pokemon')
		choice = input('Choice: ')

		speed = 0.05

	choice -= 1

	starter = pokemon[choice]
	player.catchPokemon(starter)

	os.system('clear')
	print_slow('I see, you chose ')
	print_slow(starter.species + '.\n')
	time.sleep(1)
	print_slow('A very good choice!\n')
	time.sleep(1)
	print_slow('Please, choose a nickname for your Pokemon: '),

	nickname = raw_input()
	player.PokemonList[0].changeNickname(nickname)

	os.system('clear')
	blue_speed = 0.03
	print('Blue: '),
	print_slow('So, you picked ', blue_speed)
	print_slow(starter.species + ',', blue_speed)

	time.sleep(0.5)
	print('')
	print_slow('I will pick the real best Pokemon!', blue_speed)

	time.sleep(1.5)
	os.system('clear')

	blue = Player('Blue', 1)
	blue.catchPokemon(pokemon[(choice+1)%3])


	print('Blue picked ' + blue.PokemonList[0].species)
	time.sleep(1)

	return blue

def battleStart(player, rival):
	opt = 0
	battle = Battle(player, rival)


	while( opt != -1):
		os.system('clear')
		battle.round()

		print('What will you do?')
		print('1 - Fight')
		print('2 - Run')

		opt = input()

		if opt == 1:
			mov = 0
			while(mov == 0):
				os.system('clear')

				battle.round()

				printMoves(battle.poke_p1)
				print('Select a move:'),
				mov = input()

				if mov <= 0 or mov > len(battle.poke_p1.moves):
					print('Invalid move!')
				else:
					move1 = battle.poke_p1.moves[mov-1]
					moveR = battle.poke_p2.moves[ran.randint(0, len(battle.poke_p2.moves) - 1) ]

					if battle.poke_p1.stats[5] > battle.poke_p2.stats[5]: # PLAYER IS FASTER
						battle.poke_p1.useMove(move1, battle.poke_p2, battle)

						time.sleep(3)
						if battle.poke_p2.current_hp <= 0:
							print_slow('\n' + battle.poke_p2.nickname + ' fainted...\n')
							return player

						battle.poke_p2.useMove(moveR, battle.poke_p1, battle)
						time.sleep(3)
						if battle.poke_p1.current_hp <= 0:
							print_slow('\n' + battle.poke_p1.nickname + ' fainted...\n')
							return rival

					elif battle.poke_p1.stats[5] < battle.poke_p2.stats[5]:
						battle.poke_p2.useMove(moveR, battle.poke_p1, battle)

						time.sleep(3)
						if battle.poke_p1.current_hp <= 0:
							print_slow('\n' + battle.poke_p1.nickname + ' fainted...\n')
							return rival


						battle.poke_p1.useMove(move1, battle.poke_p2, battle)

						time.sleep(3)
						if battle.poke_p2.current_hp <= 0:
							print_slow('\n' + battle.poke_p2.nickname + ' fainted...\n')
							return player

					else:
						luck = ran.randint(0,10)
						if (luck % 2) == 0: # PLAYER
							battle.poke_p1.useMove(move1, battle.poke_p2, battle)

							time.sleep(3)
							if battle.poke_p2.current_hp <= 0:
								print_slow('\n' + battle.poke_p2.nickname + ' fainted...\n')
								return player

							battle.poke_p2.useMove(moveR, battle.poke_p1, battle)
							time.sleep(3)
							if battle.poke_p1.current_hp <= 0:
								print_slow('\n' + battle.poke_p1.nickname + ' fainted...\n')
								return rival

						else: # RIVAL
							battle.poke_p2.useMove(moveR, battle.poke_p1, battle)

							time.sleep(3)
							if battle.poke_p1.current_hp <= 0:
								print_slow('\n' + battle.poke_p1.nickname + ' fainted...\n')
								return rival


							battle.poke_p1.useMove(move1, battle.poke_p2, battle)

							time.sleep(3)
							if battle.poke_p2.current_hp <= 0:
								print_slow('\n' + battle.poke_p2.nickname + ' fainted...\n')



					time.sleep(3)
					battle.plusTurn()

		elif opt == 2:
			print ('You can\'t run from a trainer battle...')
			time.sleep(1)

		else:
			opt = 0

	time.sleep(2)


def main():
	print('')

if __name__ == "__main__": main()
