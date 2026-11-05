#python random-number.py -p1 'P1' -p2 'P2' -t 10000
import random
import argparse


def chooseRandomly(number, p1, p2):
	player1 = 0
	player2 = 0

	for i in range(0, number):
		if (random.randint(0,number) % 2) == 0:
			player1 += 1
		else:
			player2 += 1

	print "\n------------------"
	print "|  " + p1 + "\t|  " + str(player1) + "  |"
	print "|  " + p2 + "\t|  " + str(player2) + "  |"
	print "------------------\n"

	if(player1 > player2):
		return p1, player1

	return p2, player2

def _main(args):

	player, points = chooseRandomly(args.total, args.player1, args.player2)
	print player + " wins (score: " + str(points) + " / " + str(args.total) + ")!\n"

	return 0


# Getting input through terminal
parser = argparse.ArgumentParser(description="Teste")

parser.add_argument('--player1', '-p1',
					type=str,
					help="Choose the player 1 name",
					default="Player1")

parser.add_argument('--player2', '-p2',
					type=str,
					help="Choose the player 2 name",
					default="Player2")
parser.add_argument('--total', '-t',
					type=int,
					help="The total number of iterations",
					default=10000)

args = parser.parse_args()

#

_main(args)