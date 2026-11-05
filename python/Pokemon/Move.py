import os

class Move:

	def __init__(self, name = '-', effect = '-', base_power = 0, accuracy = 0, total_PP = 0, category = '-', typing = 18, priority = 0):
		self.name = name
		self.effect = effect
		self.base_power = base_power
		self.accuracy = accuracy

		self.total_PP = total_PP
		self.current_PP = total_PP

		self.category = category
		self.typing = typing

		self.priority = priority

	# -----------------------------------------------
	# Getters:
	def getName(self):
		return self.name

	def getEffect(self):
		return self.effect

	def getBasePower(self):
		return self.base_power

	def getAccuracy(self):
		return self.accuracy

	def getTotal_PP(self):
		return self.total_PP

	def getCategory(self):
		return self.category

	def getTyping(self):
		return self.typing

	# -----------------------------------------------

