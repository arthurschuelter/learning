# python3 stat-roll.py -min 70
import random
import argparse

def array_sum(array):
    soma = 0
    for i in range(0, len(array)):
        soma += array[i]
    return soma

def rolls():
    roll = []
    for i in range(0, 6):
        dice = []
        lowest = 0
        for j in range(0, 4):
            dice.append( random.randint(1,6) )
            if( dice[lowest] > dice[j]):
                lowest = j
        
        del dice[lowest]
        for j in range(0, len(dice)):
            roll.append( dice[j] )
    
    return roll

def show_demo(characters, minimun):
    for i in range(0, 5):
        print(characters[i])
        dice = rolls()
        while( array_sum(dice) < minimun ):
            dice = rolls()
        dice.sort()
        print(dice)
        print("Sum: " + str(array_sum(dice)) )
        print("STR: \nDEX: \nCON: \nINT: \nWIS: \nCHA: \n")
        print("==================")

def just_rolls(minimun):
    dice = rolls()
    while( array_sum(dice) < minimun ):
        dice = rolls()
    dice.sort()
    print(dice)
    print("Sum: " + str(array_sum(dice)) )


def _main(minimun, option):

    characters=["Cutter (Kenku - Barbarian)",
                "Qiao (Fallen - Cleric)",
                "Thokk (Orc - Wizard)",
                "Mo-los (Triton - Rogue)",
                "Zarij (Loxodon - Fighter)" ]
    if(option == 1):
        show_demo(characters, minimun)
    elif(option == 2):
        just_rolls(minimun)


    return 0


parser = argparse.ArgumentParser(description="Teste")

parser.add_argument('--minimun', '-min',
					type = int,
					help = "Stat mininum",
					default = 70)

parser.add_argument('--option', '-op',
					type = int,
					help = "",
					default = 1)

args = parser.parse_args()

_main(args.minimun, args.option)