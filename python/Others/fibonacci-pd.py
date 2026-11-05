def fibonacci_list(num):
	lista = [0,1]

	while len(lista) <= num:
		front = len(lista) - 1
		n = lista[front] + lista[front - 1]
		lista.append(n)

	return lista.pop()

def fibonice(num, value = 1, prev = 0):
	if num == 0:
		return prev
	if num == 1:
		return value

	return fibonice(num-1, value+prev, value)

def fibo_recur(num):

	if num <= 1:
		return num

	return fibo_recur(num-1) + fibo_recur(num-2)


n = 35
print "Calculating the n-th number of Fibonacci sequence:"
print "    1 - with lists:\t" + str( fibonacci_list(n) )
print "    2 - with recursion: " + str( fibonice(n,1,0) ) + " (faster)"
print "    3 - with recursion: " + str( fibo_recur(n) ) + " (slower)"