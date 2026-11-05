import numpy as np
import random as rnd

alfabeto = "abcdefghijklmnopqrstuvwxyz "
goal = "to be or not to be"

def fitness(word):
    max = len(goal)
    fit = 0.0

    for i in range(0, len(word) - 1):
        if( word[i] == goal[i] ):
            fit += 1

    return fit/max


def start(pop_size=10):
    population = []

    for i in range(0, pop_size):
        word = ""
        for j in range(0, len(goal)-1):
            word = word + alfabeto[ rnd.randint(0, len(alfabeto) - 1) ]

        population.append( ( word, fitness(word) ) )

    return population

def main():
    population = start(pop_size=10)
    for ind in population:
        print( str(ind[0]) + " --> " + str(ind[1]))


if __name__ == "__main__": main()