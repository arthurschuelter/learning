import random as ran
import time
import os
##
from Move import Move
import MoveList

class Pokemon:

	def __init__(self, nickname, species, typing1, typing2 = 18):
		self.nickname = nickname
		self.species = species
		self.level = 5

		self.typing = []
		self.typing.append(typing1)
		self.typing.append(typing2)

		#self.nature = ran.randint(0,24)
		self.gainStat = ran.randint(0,4)
		self.loseStat = ran.randint(0,4)

		self.shiny = (ran.randint(1,8192) / 8192)

		self.moves = []

		self.statModifier = [0,0,0,0,0,0,0]
		#statModifier[0] = ATK
		#statModifier[1] = DEF
		#statModifier[2] = SP. ATK
		#statModifier[3] = SP. DEF
		#statModifier[4] = SPEED
		#statModifier[5] = EVASION
		#statModifier[6] = ACCURACCY

	def changeNickname(self, new_nickname):
		self.nickname = new_nickname

	def setLevel(self, level):
		self.level = level

	def setStats(self, hp, attack, defense, sp_attack, sp_defense, speed):
		self.current_hp = hp

		self.stats = []
		self.stats.append(hp)
		self.stats.append(attack)
		self.stats.append(defense)
		self.stats.append(sp_attack)
		self.stats.append(sp_defense)
		self.stats.append(speed)

	def updateStats(self):

		gain = self.gainStat
		loss = self.loseStat

		if gain == loss:
			return

		self.stats[gain+1] += (self.stats[gain+1])//10

		self.stats[loss+1] = int( (self.stats[loss+1]) - (self.stats[loss+1])/10.0 ) 

	# -----------------------------------------------
	# Getters:	
	def getMove(self, position):
		if position > 0 and position <= len(self.moves):
			return self.moves[position-1]
		else:
			print ('Invalid move!')
			return Move()
			
	# -----------------------------------------------

	
	def printSimpleStats(self):
		if self.shiny == 0:
			print (self.nickname  + ' (' + self.species + ')'), 
		else:
			print (self.nickname  + ' (' + self.species + ') *'), 

		print ('Level:  ' + str(self.level) )

	def printBattleStats(self, player=False):
		if self.shiny == 0:
			print (self.nickname  + ' (' + self.species + ')')
		else:
			print (self.nickname  + ' (' + self.species + ') *')

		print ('Level:  ' + str(self.level) )

		self.printHPbar()
		if (player == True):
			print(' ' + str(self.current_hp) + '/' + str(self.stats[0]))
		else:
			print('\n')

	def printGenericBaseStats(self):
		print('\nBase Stats:')
		print('HP:\t' + str(self.stats[0]))
		print('ATK:\t' + str(self.stats[1]))
		print('DEF:\t' + str(self.stats[2]))
		print('SP.ATK:\t' + str(self.stats[3]))
		print('SP.DEF:\t' + str(self.stats[4]))
		print('SPEED:\t' + str(self.stats[5]))

		print('--\n')

	def printPokeBaseStats(self):
		print('\nBase Stats:')
		print('HP:\t' + str(self.current_hp) + '/' + str(self.stats[0]))
		print('ATK:\t' + str(self.stats[1]))
		print('DEF:\t' + str(self.stats[2]))
		print('SP.ATK:\t' + str(self.stats[3]))
		print('SP.DEF:\t' + str(self.stats[4]))		
		print('SPEED:\t' + str(self.stats[5]))
		

		print('--\n')
	

	def learnMove(self, move):
		if len(self.moves) < 4:
			self.moves.append(move)

		
		else:
			print ('Your Pokemon can\'t learn more than four moves...')
			print ('Choose one to forget:')

			self.printMoves()

			option = input()

			self.moves[option-1] = move

	# BATTLE RELATED: ----------

	def useMove(self, move, target, battle):
		
		if move.category == 2:			
			battle.round()
			self.useSupportMove(move, target, battle)
			return

		crit = 0
		damage = 0
		crit = 0

		damage, crit = self.calculateDamage(move, target, battle)

		if damage < 1:
			damage = 1

		accuracy = move.accuracy

		luck = ran.randint(0,100)
		move.current_PP -= 1
		miss = 0


		if luck <= accuracy:
			miss = 0
			target.current_hp -= damage
			if target.current_hp < 0:
				target.current_hp = 0 
		else:
			miss = 1

		battle.round()

		print(self.nickname + ' used ' + move.name)
		if crit == 1 and move.category != 2:
			print('It was a critical hit!')

		#print('Luck: ' + str(luck))
		#print('Accuracy: ' + str(accuracy))
		#print('Damage: ' + str(damage))
		print('')

		if miss == 1:			
			print('But, it missed...')


		if move.category == 2:
			print(move.effect)

		time.sleep(1)

	def calculateModifier(self, move, target, battle):
		weather = battle.weather
		stab = 1
		typing = 1
		rdm = ran.randint(85,100) / 100.0
		crit = 0

		if self.typing[0] == move.typing or self.typing[1] == move.typing:
			stab = 1.5

		#typing = typing * effect(target.typing[0], move.typing) * effect(target.typing[1], move.typing)
		
		#crit, random

		if (ran.randint(1,96) % 16) == 0:
			crit = 1 

		mod = weather * stab * typing * rdm * (crit + 1)

		return mod, crit

	def calculateStatisticMod(self, target, category):

		a = 1
		d = 1

		if self.statModifier[category*2] > 1:
			a = 1 + self.statModifier[category*2] * 0.5

		if self.statModifier[category*2] < 1:

			self.statModifier[category*2] *= -1

			if self.statModifier[category*2]%2 == 0:
				a = (1.0 / ( (self.statModifier[category*2] / 2) + 1))
			
			else:
				a = (2.0 / (self.statModifier[category*2] + 2))

		if self.statModifier[category*2] == 0:
			a = 1

		###

		if target.statModifier[category*2 + 1] > 1:
			d = 1 + target.statModifier[category*2 + 1] * 0.5

		if target.statModifier[category*2 + 1] < 1:

			target.statModifier[category*2 + 1] *= -1

			if target.statModifier[category*2 + 1]%2 == 0:
				d = (1.0 / ( (target.statModifier[category*2 + 1] / 2) + 1))
			
			else:
				d = (2.0 / ( target.statModifier[category*2 + 1] + 2))

		if target.statModifier[category*2 + 1] == 0:
			d = 1

		return a,d

	def calculateDamage(self, move, target, battle):
		
		a_mod, d_mod = self.calculateStatisticMod(target, move.category)

		# Physical
		if move.category == 0: 

			a = self.stats[1] * a_mod 	# ATK
			d = target.stats[2] * d_mod	# DEF
		
		# Special
		elif move.category == 1:
			a = self.stats[3] * a_mod	#SP ATK
			d = target.stats[4] * d_mod	#SP DEF

		#Other
		else:
			return 0, 0

		if a <= 0:
			a = 1

		if d <= 0:
			d = 1
		

		modifier, crit = self.calculateModifier(move, target, battle)

		# CRITICAL ignores stat changes
		if crit == 1:

			if target.statModifier[1] > 0:
				d = target.stats[4]

			if self.statModifier[0] < 0:
				a = self.stats[1]


		damage = ( ( ( ( ( (2 * self.level)/5.0) + 2) * move.base_power * (a/ (d * 1.0) ) )/ 50.0 ) + 2 ) * modifier
		damage = int(damage)

		'''
		print('Base power: ' + move.base_power)
		print('A: ' + a)
		print('D: ' + d)
		print('Mod: ' + modifier)		
		time.sleep(1.0)
		'''

		return damage, crit
			
	def useSupportMove(self, move, target, battle):
		name = move.name
		move.current_PP -= 1

		if name == 'Growl':
			MoveList.useGrowl(self, target)
			
		if name == 'Tail Whip':
			MoveList.useTailWhip(self, target)

		if name == 'Swords Dance':
			MoveList.useSwordsDance(self, target)

	def printHPbar(self, length = 20, fill = '*'):
		current = self.current_hp
		total = self.stats[0]
		percent = ("{0:.1f}").format(100 * (current / float(total)))
		filledLength = int(length * current // total)
		bar = fill * filledLength + '-' * (length - filledLength)
		print('HP: |%s|' % bar),