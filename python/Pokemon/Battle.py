import random as ran
import os
import sys
import time

from Player import Player

class Battle:

	def __init__(self, p1, p2):
		self.p1 = p1
		self.poke_p1 = p1.PokemonList[0]

		self.p2 = p2
		self.poke_p2 = p2.PokemonList[0]

		self.weather = 1

		self.turn = 0

	def round(self):
		os.system('clear')
		print('Turn: ' + str(self.turn) )
		print('==========\n')

		self.poke_p2.printBattleStats()

		print ('\n==========\n')

		self.poke_p1.printBattleStats(player=True)

		print ('\n==========\n')

	def playerSwitch(self, poke):
		self.poke_p1 = poke

	def redSwitch(self, poke):
		self.poke_p2 = poke
	
	def plusTurn(self):
		self.turn += 1