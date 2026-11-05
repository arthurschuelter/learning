import random
import math

def calProb(vector, num, total):
	for i in range(0, num):
		vector.append(0)

	for i in range(0, total):
		vector[random.randint(1,num) - 1] += 1.0
	return vector

def calAverage(lista, total):

	average = 0.0

	for i in range(0, len(lista)):
		average	+= lista[i]*(i+1)

	return average/total

def stdDev(lista, avg):

	soma = 0

	for i in range(0, len(lista)):
		soma += ( i - avg + 1 )**2

	return soma / len(lista)

def _main():
	num = 10
	total = 100000
	prob = []

	prob = calProb(prob, num, total)
	media = calAverage(prob, total)
	desvio = stdDev(prob, media)

	print "The list: "
	for i in range(0, len(prob) ):
		print "\t(" + str(i+1) + "): " + str( (prob[i] * 100) / total ) + " %"

	print
	print "Average number: " + str(media)
	print "Stnd Deviation: " + str(desvio)


_main()