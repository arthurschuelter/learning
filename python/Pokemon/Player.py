import random as ran
import copy
import os
#
from Pokemon import Pokemon

genderlist = ['BOY', 'GIRL', 'OTHER'] #Polemic

class Player:

	def __init__(self, name = 'Red', gender = 1 ):
		self.name = name
		self.trainer_id = ran.randint(0,65535)
		self.gender = gender - 1

		self.PokemonList= []
		self.PokemonBox = [] 
		self.Badges = []

		if name == 'Blue':
			self.trainer_id = 1
		if name == 'Red':
			self.trainer_id = 0

	def playerStats(self):
		os.system('clear')
		print('Player: ' + self.name)
		print('Trainer ID: ' + str(self.trainer_id).zfill(5) )
		print('')
		
		for i in range(0, len(self.PokemonList) ):
			print( str(i+1) + ')'),
			self.PokemonList[i].printSimpleStats()

		print('')

	def teamPokemon(self):
		os.system('clear')
		print(self.name + '\'s team:')
		
		for i in range(0, len(self.PokemonList) ):
			print( str(i+1) + ')'),
			self.PokemonList[i].printSimpleStats()

		print('')

	def addOT(self, pokemon):
		pokemon.trainer_ot = self.trainer_id

	def catchPokemon(self, pokemon):
		poke = copy.deepcopy(pokemon)
		poke.updateStats()
		self.addOT(poke)

		self.PokemonList.append(poke)