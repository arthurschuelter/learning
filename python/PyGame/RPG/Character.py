class Character:

	def __init__(self):

		self.Name = ""

		self.total_HP = 0
		self.current_HP = 0
		self.total_Mana = 0
		self.current_Mana = 0

		self.Str = 0
		self.Mag = 0
		self.Spd = 0
		self.Luk = 0
		
		self.Atk = 0
		self.Def = 0
		self.M_Def = 0
		self.Eva = 0

	def __init__(self, name, hp, mana, stg, mag, spd, luk):

		self.Name = name

		self.total_HP = hp
		self.current_HP = hp
		self.total_Mana = mana
		self.current_Mana = mana

		self.Str = stg
		self.Mag = mag
		self.Spd = spd
		self.Luk = luk
		
		self.Atk = stg
		self.Def = 0
		self.M_Def = 0
		self.Eva = spd + luk // 2

	def status(self):
		print( self.Name )
		print("---")
		print("HP: " +  str(self.current_HP) + "/" +  str(self.total_HP) )
		print("Mana: " +  str(self.current_Mana) + "/" +  str(self.total_Mana) )
		print("---")
		print("Strength: " +  str(self.Str) )
		print("Magic: " +  str(self.Mag) )
		print("Speed: " +  str(self.Spd) )
		print("Luck: " +  str(self.Luk) )
		print("---")
		print("Attack: " +  str(self.Atk) )
		print("Defense: " +  str(self.Def) )
		print("Magic Defense: " +  str(self.M_Def) )
		print("Evasion: " +  str(self.Eva) )





