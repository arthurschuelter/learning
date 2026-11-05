import time
##
from Move import Move


def useGrowl(user, target):
	if target.statModifier[0] == -6:
		print('')
		print(user.nickname + ' used Growl!')
		time.sleep(0.5)
		print(user.nickname + '\' ATTACK won\'t go any lower!')

	else:
		target.statModifier[0] -= 1

		print(user.nickname + ' used Growl!')
		time.sleep(0.5)
		print(target.nickname + '\'s ATTACK fell!')

	time.sleep(2)

def useTailWhip(user, target):

	if target.statModifier[1] == -6:
		print(user.nickname + ' used Tail Whip!')
		time.sleep(0.5)
		print(user.nickname + '\' DEFENSE won\'t go any lower!')

	else:
		target.statModifier[1] -= 1

		print(user.nickname + ' Tail Whip!')
		time.sleep(0.5)
		print(target.nickname + '\'s DEFENSE fell!')
	
	time.sleep(2)

def useSwordsDance(user):
	if user.statModifier[0] == 6:
		print(user.nickname + ' used Swords Dance!')
		time.sleep(0.5)
		print(user.nickname + '\' ATTACK won\'t go any higher!')

	else:
		user.statModifier[0] += 2

		if user.statModifier[0] > 6:
			user.statModifier[0] = 6

		print(user.nickname + ' used Swords Dance!')
		time.sleep(0.5)
		print(user.nickname + '\'s ATTACK sharply rose!')

	time.sleep(2)



def main():
	print('')

if __name__ == '__main__': main()